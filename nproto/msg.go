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

// MsgAsyncPublisher is similar to MsgPublisher but in async manner.
type MsgAsyncPublisher interface {
	// PublishAsync publishes a message to the given subject asynchronously.
	// The final result is returned by cb.
	// NOTE: This method must be non-blocking.
	// And cb must be called exactly once (even after context done) if PublishAsync returns nil.
	PublishAsync(ctx context.Context, subject string, msg proto.Message, cb func(error)) error
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
