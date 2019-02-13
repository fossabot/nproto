package nproto

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// MsgPublisher is used to publish messages reliably, e.g. at least once delivery.
type MsgPublisher interface {
	// Publish publishes a message to the given subject. It returns nil when succeeded.
	Publish(ctx context.Context, subject string, msg proto.Message) error
}

// MsgSubscriber is used to consume messages.
type MsgSubscriber interface {
	// Subscribe subscribes to a given subject. One subject can have many queues.
	// In normal case (excpet message redelivery) each message will be delivered to
	// one member of each queue.
	Subscribe(subject, queue string, newMsg func() proto.Message, handler MsgHandler, opts ...interface{}) error
}

// MsgHandler handles the message. The message should be redelivered if it returns an error.
type MsgHandler func(context.Context, proto.Message) error

// MsgAsyncPublisher extends MsgPublisher with PublishAsync for higher performance.
type MsgAsyncPublisher interface {
	// It's trivial to implement Publish if it supports PublishAsync. See MsgAsyncPublisherFunc.
	MsgPublisher

	// PublishAsync publishes a message to a given subject asynchronously.
	// NOTE: If PublishAsync returns nil, then the final result must be returned by calling cb
	// exactly once even context is done. Otherwise, cb must not called.
	PublishAsync(ctx context.Context, subject string, msg proto.Message, cb func(error)) error
}

// MsgAsyncPublisherFunc is an adapter to allow the use of ordinary functions as MsgAsyncPublisher.
type MsgAsyncPublisherFunc func(context.Context, string, proto.Message, func(error)) error

var (
	_ MsgAsyncPublisher = (MsgAsyncPublisherFunc)(nil)
)

// Publish implements MsgAsyncPublisher interface.
func (fn MsgAsyncPublisherFunc) Publish(ctx context.Context, subject string, msg proto.Message) error {
	var (
		err  error
		errc = make(chan struct{})
	)
	if err1 := fn(ctx, subject, msg, func(err2 error) {
		err = err2
		close(errc)
	}); err1 != nil {
		return err1
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-errc:
		return err
	}
}

// PublishAsync implements MsgAsyncPublisher interface.
func (fn MsgAsyncPublisherFunc) PublishAsync(ctx context.Context, subject string, msg proto.Message, cb func(error)) error {
	return fn(ctx, subject, msg, cb)
}
