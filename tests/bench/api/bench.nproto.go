// Code generated by protoc-gen-nproto. DO NOT EDIT.
// source: bench.proto

package benchapi // import "github.com/huangjunwen/nproto/tests/bench/api"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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
	// Echo echos string.
	Echo(ctx context.Context, input *EchoMsg) (output *EchoMsg, err error)
}

// ServeBench serves Bench service using a RPC server.
func ServeBench(server nproto.RPCServer, svcName string, svc Bench) error {
	return server.RegistSvc(svcName, map[*nproto.RPCMethod]nproto.RPCHandler{
		methodBench__Echo: func(ctx context.Context, input proto.Message) (proto.Message, error) {
			return svc.Echo(ctx, input.(*EchoMsg))
		},
	})
}

// InvokeBench invokes Bench service using a RPC client.
func InvokeBench(client nproto.RPCClient, svcName string) Bench {
	return &clientBench{
		handlerEcho: client.MakeHandler(svcName, methodBench__Echo),
	}
}

var methodBench__Echo = &nproto.RPCMethod{
	Name:      "Echo",
	NewInput:  func() proto.Message { return &EchoMsg{} },
	NewOutput: func() proto.Message { return &EchoMsg{} },
}

type clientBench struct {
	handlerEcho nproto.RPCHandler
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
