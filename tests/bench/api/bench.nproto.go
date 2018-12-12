// Code generated by protoc-gen-nproto. DO NOT EDIT.
// source: bench.proto

package benchapi // import "github.com/huangjunwen/nproto/tests/bench/api"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	nproto "github.com/huangjunwen/nproto/nproto"
)

// Avoid import and not used errors.
var (
	_ = context.Background
	_ = proto.Int
	_ = nproto.NewRPCCtx
)

// @@protoc_insertion_point(imports)

// @@nprpc@@
type Bench interface {
	// Nop receives empty and returns empty.
	Nop(ctx context.Context, input *empty.Empty) (output *empty.Empty, err error)

	// Echo echos string.
	Echo(ctx context.Context, input *EchoMsg) (output *EchoMsg, err error)
}

// ServeBench serves Bench service using a RPC server.
func ServeBench(server nproto.RPCServer, svcName string, svc Bench) error {
	return server.RegistSvc(svcName, map[*nproto.RPCMethod]nproto.RPCHandler{
		methodBench__Nop: func(ctx context.Context, input proto.Message) (proto.Message, error) {
			return svc.Nop(ctx, input.(*empty.Empty))
		},
		methodBench__Echo: func(ctx context.Context, input proto.Message) (proto.Message, error) {
			return svc.Echo(ctx, input.(*EchoMsg))
		},
	})
}

// InvokeBench invokes Bench service using a RPC client.
func InvokeBench(client nproto.RPCClient, svcName string) Bench {
	return &clientBench{
		handlerNop:  client.MakeHandler(svcName, methodBench__Nop),
		handlerEcho: client.MakeHandler(svcName, methodBench__Echo),
	}
}

var methodBench__Nop = &nproto.RPCMethod{
	Name:      "Nop",
	NewInput:  func() proto.Message { return &empty.Empty{} },
	NewOutput: func() proto.Message { return &empty.Empty{} },
}
var methodBench__Echo = &nproto.RPCMethod{
	Name:      "Echo",
	NewInput:  func() proto.Message { return &EchoMsg{} },
	NewOutput: func() proto.Message { return &EchoMsg{} },
}

type clientBench struct {
	handlerNop  nproto.RPCHandler
	handlerEcho nproto.RPCHandler
}

// Nop implements Bench interface.
func (svc *clientBench) Nop(ctx context.Context, input *empty.Empty) (*empty.Empty, error) {
	output, err := svc.handlerNop(ctx, input)
	if err != nil {
		return nil, err
	}
	return output.(*empty.Empty), nil
}

// Echo implements Bench interface.
func (svc *clientBench) Echo(ctx context.Context, input *EchoMsg) (*EchoMsg, error) {
	output, err := svc.handlerEcho(ctx, input)
	if err != nil {
		return nil, err
	}
	return output.(*EchoMsg), nil
}

// @@protoc_insertion_point(nprpc)

// @@protoc_insertion_point(npmsg)