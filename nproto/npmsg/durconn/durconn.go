package durconn

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/huangjunwen/nproto/nproto/npmsg"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/rs/xid"
)

var (
	DefaultConnectWait   = 5 * time.Second
	DefaultSubRetryWait  = 5 * time.Second
	DefaultSubjectPreifx = "npmsg"
)

var (
	ErrNCMaxReconnect     = errors.New("nproto.npmsg.durconn.NewDurConn: nats.Conn should have MaxReconnects < 0")
	ErrClosed             = errors.New("nproto.npmsg.durconn.DurConn: DurConn closed.")
	ErrNotConnected       = errors.New("nproto.npmsg.durconn.DurConn: Not connected.")
	ErrBadSubscriptionOpt = errors.New("nproto.npmsg.durconn.DurConn: Expect SubscriptionOption.")
	ErrDupSubscription    = func(subject, queue string) error {
		return fmt.Errorf(
			"nproto.npmsg.durconn.DurConn: Duplicated subscription (%q, %q)",
			subject,
			queue,
		)
	}
)

var (
	stanConnect = stan.Connect
)

var (
	_ npmsg.RawMsgPublisher      = (*DurConn)(nil)
	_ npmsg.RawBatchMsgPublisher = (*DurConn)(nil)
	_ npmsg.RawMsgSubscriber     = (*DurConn)(nil)
)

type DurConn struct {
	// Options.
	stanOptions   []stan.Option
	reconnectWait time.Duration
	subjectPrefix string

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

	// Immutable fields.
	dc      *DurConn
	subject string
	queue   string
	handler npmsg.RawMsgHandler
}

type DurConnOption func(*DurConn) error

type SubscriptionOption func(*subscription) error

func NewDurConn(nc *nats.Conn, clusterID string, opts ...DurConnOption) (*DurConn, error) {

	if nc.Opts.MaxReconnect >= 0 {
		return nil, ErrNCMaxReconnect
	}

	dc := &DurConn{
		reconnectWait: DefaultConnectWait,
		subjectPrefix: DefaultSubjectPreifx,
		clusterID:     clusterID,
		nc:            nc,
		subNames:      make(map[[2]string]int),
	}
	for _, opt := range opts {
		if err := opt(dc); err != nil {
			return nil, err
		}
	}

	dc.connect(false)
	return dc, nil

}

func (dc *DurConn) Close() error {
	oldSc, oldStalec, err := dc.close_()
	if oldSc != nil {
		oldSc.Close()
		close(oldStalec)
	}
	return err
}

func (dc *DurConn) Subscribe(subject, queue string, handler npmsg.RawMsgHandler, opts ...interface{}) error {

	sub := &subscription{
		retryWait: DefaultSubRetryWait,
		dc:        dc,
		subject:   subject,
		queue:     queue,
		handler:   handler,
	}
	for _, opt := range opts {
		option, ok := opt.(SubscriptionOption)
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

func (dc *DurConn) Publish(ctx context.Context, subject string, data []byte) error {
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

	resultc := make(chan error)

	// Use a seperated go routine to call Publish since we need to listen to ctx.
	go func() {
		err := sc.Publish(dc.makeSubject(subject), data)
		select {
		case resultc <- err:
		case <-ctx.Done():
		}
	}()

	// Wait result or context done.
	select {
	case r := <-resultc:
		return r
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (dc *DurConn) PublishBatch(ctx context.Context, subjects []string, datas [][]byte) (errors []error) {

	dc.mu.RLock()
	closed := dc.closed
	sc := dc.sc
	dc.mu.RUnlock()

	var err error = notSetErr{}
	if closed {
		err = ErrClosed
	} else if sc == nil {
		err = ErrNotConnected
	}

	// Initialize error slice.
	n := len(subjects)
	errors = make([]error, n)
	for i, _ := range errors {
		errors[i] = err
	}

	// If real error.
	if err != (notSetErr{}) {
		return
	}

	// Channel to send/recv result.
	type result struct {
		I int
		E error
	}
	resultc := make(chan result)

	// Use a seperated go routine to call PublishAsync since we need to listen to ctx.
	go func() {
		for i, subject := range subjects {
			i := i
			h := func(_ string, err error) {
				select {
				case resultc <- result{I: i, E: err}:
				case <-ctx.Done():
					return
				}
			}

			_, err := sc.PublishAsync(dc.makeSubject(subject), datas[i], h)
			if err != nil {
				select {
				case resultc <- result{I: i, E: err}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	// Wait all results or context done.
	for n > 0 {
		select {
		case r := <-resultc:
			errors[r.I] = r.E
			n -= 1
		case <-ctx.Done():
			break
		}
	}

	// Context done before getting all results.
	if n > 0 {
		err = ctx.Err()
		for i, _ := range errors {
			if errors[i] == (notSetErr{}) {
				errors[i] = err
			}
		}
	}

	return
}

// connect is used to make a new stan connection.
func (dc *DurConn) connect(wait bool) {

	go func() {
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
					return
				}
				panic(err)
			}
		}

		// Connect.
		opts := []stan.Option{}
		opts = append(opts, dc.stanOptions...)
		opts = append(opts, stan.NatsConn(dc.nc))
		opts = append(opts, stan.SetConnectionLostHandler(func(_ stan.Conn, _ error) {
			// Reconnect when connection lost.
			// NOTE: This callback will no be invoked in normal close.
			dc.connect(true)
		}))

		// NOTE: Use a UUID-like id as client id sine we only use durable queue subscription.
		// See: https://groups.google.com/d/msg/natsio/SkWAdSU1AgU/tCX9f3ONBQAJ
		sc, err := stanConnect(dc.clusterID, xid.New().String(), opts...)
		if err != nil {
			// Reconnect immediately.
			dc.connect(true)
			return
		}

		// Update to the new connection.
		{
			stalec := make(chan struct{})
			_, _, subs, err := dc.reset(sc, stalec)
			if err != nil {
				sc.Close() // Release the new one.
				close(stalec)
				if err == ErrClosed {
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
		// Wrap sub.handler to stan.MsgHandler.
		handler := func(m *stan.Msg) {
			// Ack when no error.
			if err := sub.handler(context.Background(), m.Subject, m.Data); err == nil {
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
			_, err := sc.QueueSubscribe(dc.makeSubject(sub.subject), sub.queue, handler, opts...)
			if err == nil {
				return
			}
			select {
			case <-stalec:
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

func (dc *DurConn) makeSubject(subject string) string {
	return fmt.Sprintf("%s.%s", dc.subjectPrefix, subject)
}

type notSetErr struct{}

func (notSetErr) Error() string {
	return "Not set"
}

var _ error = notSetErr{}
