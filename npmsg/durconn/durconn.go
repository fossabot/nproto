package durconn

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/huangjunwen/nproto/npmsg"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

var (
	DefaultReconnectWait = 2 * time.Second
	DefaultSubRetryWait  = 2 * time.Second
	DefaultPubAckWait    = 2 * time.Second
	DefaultSubjectPrefix = "npmsg"
)

var (
	ErrNCMaxReconnect     = errors.New("nproto.npmsg.durconn.NewDurConn: nats.Conn should have MaxReconnects < 0")
	ErrClosed             = errors.New("nproto.npmsg.durconn.DurConn: Closed.")
	ErrNotConnected       = errors.New("nproto.npmsg.durconn.DurConn: Not connected.")
	ErrBadSubscriptionOpt = errors.New("nproto.npmsg.durconn.DurConn: Expect durconn.SubOption.")
	ErrDupSubscription    = func(subject, queue string) error {
		return fmt.Errorf(
			"nproto.npmsg.durconn.DurConn: Duplicated subscription (%q, %q)",
			subject,
			queue,
		)
	}
)

// For mocking.
var (
	stanConnect = stan.Connect
)

var (
	_ npmsg.RawMsgAsyncPublisher = (*DurConn)(nil)
	_ npmsg.RawMsgSubscriber     = (*DurConn)(nil)
)

// DurConn implements a "durable" connection to nats-streaming server: auto reconnection/auto resubscription.
// Implements RawMsgPublisher/RawBatchMsgPublisher/RawMsgSubscriber interfaces.
type DurConn struct {
	// Options.
	logger        zerolog.Logger
	reconnectWait time.Duration
	subjectPrefix string
	connectCb     func(stan.Conn)
	disconnectCb  func(stan.Conn)
	stanOptions   []stan.Option

	// Immutable fields.
	clusterID string
	nc        *nats.Conn

	// At most one connect can be run at any time.
	connectMu sync.Mutex

	// Mutable fields.
	// mu is used to protect the following fields.
	mu       sync.RWMutex
	closed   bool              // true if the DurConn is closed.
	sc       stan.Conn         // nil if not connected/reconnecting.
	stalec   chan struct{}     // stalec is pair with sc, it is closed when sc is stale.
	subs     []*subscription   // list of subscription, append-only (since no Unsubscribe)
	subNames map[[2]string]int // (subject, queue) -> subscription index
}

type subscription struct {
	// Options.
	stanOptions []stan.SubscriptionOption
	retryWait   time.Duration
	subscribeCb func(sc stan.Conn, subject, queue string)

	// Immutable fields.
	dc      *DurConn
	subject string
	queue   string
	handler npmsg.RawMsgHandler
}

// Option is option in creating DurConn.
type Option func(*DurConn) error

// SubOption is option in subscription.
type SubOption func(*subscription) error

// NewDurConn creates a new DurConn. `nc` must have MaxReconnect < 0 set (e.g. Always reconnect).
// `clusterID` is nats-streaming server's cluster id.
func NewDurConn(nc *nats.Conn, clusterID string, opts ...Option) (*DurConn, error) {

	if nc.Opts.MaxReconnect >= 0 {
		return nil, ErrNCMaxReconnect
	}

	dc := &DurConn{
		logger:        zerolog.Nop(),
		reconnectWait: DefaultReconnectWait,
		connectCb:     func(_ stan.Conn) {},
		disconnectCb:  func(_ stan.Conn) {},
		stanOptions: []stan.Option{
			stan.PubAckWait(DefaultPubAckWait),
		},
		clusterID: clusterID,
		nc:        nc,
		subNames:  make(map[[2]string]int),
	}
	dc.setSubjectPrefix(DefaultSubjectPrefix)
	for _, opt := range opts {
		if err := opt(dc); err != nil {
			return nil, err
		}
	}

	dc.connect(false)
	return dc, nil

}

// Close closes the DurConn.
func (dc *DurConn) Close() error {
	oldSc, oldStalec, err := dc.close_()
	if oldSc != nil {
		oldSc.Close()
		close(oldStalec)
	}
	return err
}

// Subscribe implements npmsg.RawMsgSubscriber interface. `opts` must be SubOption.
func (dc *DurConn) Subscribe(subject, queue string, handler npmsg.RawMsgHandler, opts ...interface{}) error {

	sub := &subscription{
		retryWait:   DefaultSubRetryWait,
		subscribeCb: func(_ stan.Conn, _, _ string) {},
		dc:          dc,
		subject:     subject,
		queue:       queue,
		handler:     handler,
	}
	for _, opt := range opts {
		option, ok := opt.(SubOption)
		if !ok {
			return ErrBadSubscriptionOpt
		}
		if err := option(sub); err != nil {
			return err
		}
	}

	// Add to subscription list and subscribe.
	sc, stalec, err := dc.addSub(sub)
	if err != nil {
		return err
	}

	if sc != nil {
		dc.subscribe(sub, sc, stalec)
	}
	return nil
}

// PublishAsync implements npmsg.RawMsgAsyncPublisher interface.
func (dc *DurConn) PublishAsync(ctx context.Context, subject string, data []byte, cb func(error)) error {
	dc.mu.RLock()
	closed := dc.closed
	sc := dc.sc
	dc.mu.RUnlock()

	if closed {
		return ErrClosed
	}
	if sc == nil {
		return ErrNotConnected
	}
	_, err := sc.PublishAsync(dc.addSubjectPrefix(subject), data, func(_ string, err error) { cb(err) })
	return err
}

// Publish implements npmsg.RawMsgPublisher interface.
func (dc *DurConn) Publish(ctx context.Context, subject string, data []byte) error {
	var (
		err  error
		errc = make(chan struct{})
	)
	if e1 := dc.PublishAsync(ctx, subject, data, func(e2 error) {
		err = e2
		close(errc)
	}); e1 != nil {
		return e1
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-errc:
		return err
	}
}

// connect is used to make a new stan connection.
func (dc *DurConn) connect(wait bool) {

	go func() {
		logger := dc.logger.With().Str("fn", "connect").Logger()

		if wait {
			time.Sleep(dc.reconnectWait)
		}

		dc.connectMu.Lock()
		defer dc.connectMu.Unlock()

		// Release old sc and reset to nil.
		{
			oldSc, oldStalec, _, err := dc.reset(nil, nil)
			if oldSc != nil {
				oldSc.Close()
				close(oldStalec)
			}
			if err != nil {
				if err == ErrClosed {
					logger.Info().Msg("closed")
					return
				}
				panic(err)
			}
		}

		// Connect.
		opts := []stan.Option{}
		opts = append(opts, dc.stanOptions...)
		opts = append(opts, stan.NatsConn(dc.nc))
		opts = append(opts, stan.SetConnectionLostHandler(func(sc stan.Conn, _ error) {
			// Reconnect when connection lost.
			// NOTE: This callback will no be invoked in normal close.
			dc.disconnectCb(sc)
			dc.connect(true)
		}))

		// NOTE: Use a UUID-like id as client id sine we only use durable queue subscription.
		// See: https://groups.google.com/d/msg/natsio/SkWAdSU1AgU/tCX9f3ONBQAJ
		sc, err := stanConnect(dc.clusterID, xid.New().String(), opts...)
		if err != nil {
			// Reconnect immediately.
			logger.Error().Err(err).Msg("connect failed")
			dc.connect(true)
			return
		}
		logger.Info().Msg("connected")
		dc.connectCb(sc)

		// Update to the new connection.
		{
			stalec := make(chan struct{})
			_, _, subs, err := dc.reset(sc, stalec)
			if err != nil {
				sc.Close() // Release the new one.
				close(stalec)
				if err == ErrClosed {
					logger.Info().Msg("closed")
					return
				}
				panic(err)
			}

			// Re-subscribe.
			for _, sub := range subs {
				dc.subscribe(sub, sc, stalec)
			}
		}

		return
	}()

}

// subscribe keeps subscribing unless success or the stan connection is stale
// (e.g. disconnected or closed).
func (dc *DurConn) subscribe(sub *subscription, sc stan.Conn, stalec chan struct{}) {
	// Use a seperated go routine.
	go func() {
		logger := dc.logger.With().
			Str("fn", "subscribe").
			Str("subject", sub.subject).
			Str("queue", sub.queue).
			Logger()

		// Wrap sub.handler to stan.MsgHandler.
		handler := func(m *stan.Msg) {
			// Ack when no error.
			if err := sub.handler(context.Background(), dc.trimSubjectPrefix(m.Subject), m.Data); err == nil {
				m.Ack()
			}
		}

		opts := []stan.SubscriptionOption{}
		opts = append(opts, sub.stanOptions...)
		opts = append(opts, stan.SetManualAckMode())     // Use manual ack mode. See above handler.
		opts = append(opts, stan.DurableName(sub.queue)) // Queue as durable name.

		// Loop until success or the stan connection is stale (e.g. scStaleC is closed).
		for {
			// We don't need the returned stan.Subscription object since no Unsubscribe.
			_, err := sc.QueueSubscribe(dc.addSubjectPrefix(sub.subject), sub.queue, handler, opts...)
			if err == nil {
				logger.Info().Msg("subscribed")
				sub.subscribeCb(sc, sub.subject, sub.queue)
				return
			}
			logger.Error().Err(err).Msg("subscribe failed")

			select {
			case <-stalec:
				logger.Info().Msg("stale")
				return
			case <-time.After(sub.retryWait):
			}
		}

	}()
}

// reset gets lock then reset stan connection.
// It returns the old stan connection to be released (if not nil) and current subscriptions.
// It returns an error only when closed.
func (dc *DurConn) reset(sc stan.Conn, stalec chan struct{}) (
	oldSc stan.Conn,
	oldStalec chan struct{},
	subs []*subscription,
	err error) {

	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.closed {
		return nil, nil, nil, ErrClosed
	}

	oldSc = dc.sc
	oldStalec = dc.stalec
	dc.sc = sc
	dc.stalec = stalec

	// NOTE: Since subs is append-only, we can simply get the full slice as a snapshot for resubscribing.
	subs = dc.subs[:]
	return
}

// addSub gets lock then adds subscription to DurConn.
// It returns current stan connection.
func (dc *DurConn) addSub(sub *subscription) (sc stan.Conn, stalec chan struct{}, err error) {
	key := [2]string{sub.subject, sub.queue}
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.closed {
		return nil, nil, ErrClosed
	}

	if _, found := dc.subNames[key]; found {
		return nil, nil, ErrDupSubscription(sub.subject, sub.queue)
	}

	dc.subs = append(dc.subs, sub)
	dc.subNames[key] = len(dc.subs) - 1
	sc = dc.sc
	stalec = dc.stalec
	return
}

// close_ gets lock then sets closed to true.
// It returns the old stan connection to be released (if not nil).
func (dc *DurConn) close_() (oldSc stan.Conn, oldStalec chan struct{}, err error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.closed {
		return nil, nil, ErrClosed
	}

	oldSc = dc.sc
	oldStalec = dc.stalec
	dc.sc = nil
	dc.stalec = nil

	dc.closed = true
	return
}

func (dc *DurConn) setSubjectPrefix(prefix string) {
	dc.subjectPrefix = prefix + "."
}

func (dc *DurConn) addSubjectPrefix(subject string) string {
	return fmt.Sprintf("%s%s", dc.subjectPrefix, subject)
}

func (dc *DurConn) trimSubjectPrefix(subject string) string {
	return strings.TrimPrefix(subject, dc.subjectPrefix)
}
