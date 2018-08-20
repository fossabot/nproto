package durconn

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/nuid"

	"github.com/huangjunwen/nproto/util"
)

const (
	pkgName = "nproto.msg.durconn"
)

var (
	ErrNotConnected   = errors.New(pkgName + ".DurConn: not yet connected to streaming server")
	ErrEmptyGroupName = errors.New(pkgName + ".DurConn: empty group name")
)

// Can be mocked.
var (
	stanConnect                      = stan.Connect
	cfh         util.ControlFlowHook = util.ProdControlFlowHook{}
)

// DurConn provides re-connection/re-subscription functions on top of stan.DurConn.
// It supports Publish/PublishAsync and durable QueueSubscribe (not support Unsubscribe).
type DurConn struct {
	id        string
	options   Options
	connectMu sync.Mutex // sync between connect(s)/Close

	mu       sync.RWMutex
	sc       stan.Conn                   // nil if not connected
	scStaleC chan struct{}               // scStaleC is closed when sc is stale
	subs     map[[2]string]*subscription // (subject, group)->subscription
	closed   bool                        // closed or not
}

// subscription is a single subscription. NOTE: Not support Unsubscribe.
type subscription struct {
	conn *DurConn

	subject string
	group   string
	cb      stan.MsgHandler

	options SubscriptionOptions
}

// NewDurConn creates a new DurConn. nc should have MaxReconnects < 0 set (e.g. Always reconnect).
func NewDurConn(nc *nats.Conn, clusterID string, opts ...Option) *DurConn {

	c := &DurConn{
		id:      nuid.Next(),
		options: NewOptions(),
		subs:    make(map[[2]string]*subscription),
	}
	for _, opt := range opts {
		opt(&c.options)
	}

	// Add id.
	c.options.logger = c.options.logger.With().Str("client_id", c.id).Logger()

	var (
		connect func(bool)
		logger  = c.options.logger.With().Str("comp", pkgName+".NewDurConn.connect").Logger()
	)

	// This function is used to close old connection (if any), release resouces used by old connection.
	// Re-connect and re subscribe.
	connect = func(wait bool) {
		// Wait a while.
		if wait {
			time.Sleep(c.options.reconnectWait)
		}

		// At most one connect is allowed to run at any time.
		c.connectMu.Lock()
		defer c.connectMu.Unlock()

		// Set fields to blank and release old resources.
		{
			c.mu.Lock()
			sc := c.sc
			scStaleC := c.scStaleC
			closed := c.closed
			c.sc = nil
			c.scStaleC = nil
			c.mu.Unlock()

			if sc != nil {
				sc.Close()
			}
			if scStaleC != nil {
				// Notify stale of sc.
				close(scStaleC)
			}

			// DurConn has been closed.
			if closed {
				logger.Info().Msg("DurConn closed. connect aborted.")
				return
			}

		}

		// Start to connect.
		opts := []stan.Option{}
		opts = append(opts, c.options.stanOptions...)
		opts = append(opts, stan.NatsConn(nc))
		opts = append(opts, stan.SetConnectionLostHandler(func(_ stan.Conn, _ error) {
			// Reconnect when connection lost.
			go connect(true)
		}))

		sc, err := stanConnect(clusterID, c.id, opts...)
		if err != nil {
			// Connect failed. Retry.
			logger.Error().Err(err).Msg("Failed to connect.")
			go connect(true)
			return
		}
		logger.Info().Msg("Connected.")

		// Connected. Update fields and start to re-subscribe.
		c.mu.Lock()
		defer c.mu.Unlock()

		c.sc = sc
		c.scStaleC = make(chan struct{})

		for _, sub := range c.subs {
			go sub.queueSubscribeTo(c.sc, c.scStaleC)
		}

		return

	}

	go connect(false)
	return c
}

// Close DurConn.
func (c *DurConn) Close() {
	// Guarantee no connect is running.
	c.connectMu.Lock()
	defer c.connectMu.Unlock()

	// Set fields to blank. Set closed to true.
	c.mu.Lock()
	sc := c.sc
	scStaleC := c.scStaleC
	c.sc = nil
	c.scStaleC = nil
	c.closed = true
	c.mu.Unlock()

	if sc != nil {
		// Close old sc.
		// NOTE: The callback (SetConnectionLostHandler) will not be invoked on normal Conn.Close().
		sc.Close()
	}
	if scStaleC != nil {
		// Close scStaleC.
		close(scStaleC)
	}

	return
}

// Publish == stan.Conn.Publish
func (c *DurConn) Publish(subject string, data []byte) error {
	c.mu.RLock()
	sc := c.sc
	c.mu.RUnlock()
	if sc == nil {
		return ErrNotConnected
	}
	return sc.Publish(subject, data)
}

// PublishAsync == stan.Conn.PublishAsync
func (c *DurConn) PublishAsync(subject string, data []byte, ah stan.AckHandler) (string, error) {
	c.mu.RLock()
	sc := c.sc
	c.mu.RUnlock()
	if sc == nil {
		return "", ErrNotConnected
	}
	return sc.PublishAsync(subject, data, ah)
}

// QueueSubscribe creates a durable queue subscription. Will be re-subscription after reconnection.
// Can't unsubscribe. This function returns error only when parameter error.
func (c *DurConn) QueueSubscribe(subject, group string, cb stan.MsgHandler, opts ...SubscriptionOption) error {
	if group == "" {
		return ErrEmptyGroupName
	}

	sub := &subscription{
		conn:    c,
		subject: subject,
		group:   group,
		cb:      cb,
		options: NewSubscriptionOptions(),
	}
	for _, opt := range opts {
		opt(&sub.options)
	}

	// Check duplication.
	key := [2]string{subject, group}
	c.mu.Lock()
	if c.subs[key] != nil {
		c.mu.Unlock()
		return fmt.Errorf("%s.DurConn: subject=%+q group=%+q has already subscribed", pkgName, subject, group)
	}
	c.subs[key] = sub
	sc := c.sc
	scStaleC := c.scStaleC
	c.mu.Unlock()

	// If sc is not nil, start to subscribe.
	if sc != nil {
		go sub.queueSubscribeTo(sc, scStaleC)
	}
	return nil

}

func (sub *subscription) queueSubscribeTo(sc stan.Conn, scStaleC chan struct{}) {
	logger := sub.conn.options.logger.With().
		Str("comp", pkgName+".subscription.queueSubscribeTo").
		Str("subj", sub.subject).
		Str("grp", sub.group).Logger()

	// Group as DurableName
	opts := []stan.SubscriptionOption{}
	opts = append(opts, sub.options.stanOptions...)
	opts = append(opts, stan.DurableName(sub.group))

	// Keep re-subscribing util stale become true if error.
	stale := false
	for !stale {
		_, err := sc.QueueSubscribe(sub.subject, sub.group, sub.cb, opts...)
		if err == nil {
			// Normal case.
			logger.Info().Msg("Subscribed.")
			return
		}

		// Error case.
		logger.Error().Err(err).Msg("Subscribe error.")

		// Wait a while.
		t := time.NewTimer(sub.options.resubscribeWait)
		select {
		case <-scStaleC:
			stale = true
		case <-t.C:
		}
		t.Stop()
	}
}
