// Code generated by protoc-gen-nproto. DO NOT EDIT.
// source: trace.proto

package traceapi // import "github.com/huangjunwen/nproto/tests/trace/api"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	nproto "github.com/huangjunwen/nproto/nproto"
)

// Avoid import and not used errors.
var (
	_ = context.Background
	_ = proto.Int
	_ = nproto.NewMetaDataPairs
)

// @@protoc_insertion_point(imports)

// @@nprpc@@
type Trace interface {
	// Recursive calls this method recursively.
	Recursive(ctx context.Context, input *RecursiveRequest) (output *RecursiveReply, err error)
}

// ServeTrace serves Trace service using a RPC server.
func ServeTrace(server nproto.RPCServer, svcName string, svc Trace) error {
	return server.RegistSvc(svcName, map[*nproto.RPCMethod]nproto.RPCHandler{
		methodTrace__Recursive: func(ctx context.Context, input proto.Message) (proto.Message, error) {
			return svc.Recursive(ctx, input.(*RecursiveRequest))
		},
	})
}

// InvokeTrace invokes Trace service using a RPC client.
func InvokeTrace(client nproto.RPCClient, svcName string) Trace {
	return &clientTrace{
		handlerRecursive: client.MakeHandler(svcName, methodTrace__Recursive),
	}
}

var methodTrace__Recursive = &nproto.RPCMethod{
	Name:      "Recursive",
	NewInput:  func() proto.Message { return &RecursiveRequest{} },
	NewOutput: func() proto.Message { return &RecursiveReply{} },
}

type clientTrace struct {
	handlerRecursive nproto.RPCHandler
}

// Recursive implements Trace interface.
func (svc *clientTrace) Recursive(ctx context.Context, input *RecursiveRequest) (*RecursiveReply, error) {
	output, err := svc.handlerRecursive(ctx, input)
	if err != nil {
		return nil, err
	}
	return output.(*RecursiveReply), nil
}

// @@protoc_insertion_point(nprpc)

// SubscribeRecursiveDepthNegative subscribes to the specified message channel.
func SubscribeRecursiveDepthNegative(subscriber nproto.MsgSubscriber, subject, queue string, handler func(context.Context, *RecursiveDepthNegative) error, opts ...interface{}) error {
	return subscriber.Subscribe(
		subject,
		queue,
		func(ctx context.Context, msgData []byte) error {
			msg := &RecursiveDepthNegative{}
			if err := proto.Unmarshal(msgData, msg); err != nil {
				return err
			}
			return handler(ctx, msg)
		},
		opts...,
	)
}

// @@protoc_insertion_point(npmsg)
