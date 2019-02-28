// Code generated by protoc-gen-nproto. DO NOT EDIT.
// source: runner.proto

package runnerapi // import "github.com/huangjunwen/nproto/tests/runner/api"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
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
type Runner interface {
	// Sleep sleeps a while and returns.
	Sleep(ctx context.Context, input *duration.Duration) (output *empty.Empty, err error)
}

// ServeRunner serves Runner service using a RPC server.
func ServeRunner(server nproto.RPCServer, svcName string, svc Runner) error {
	return server.RegistSvc(svcName, map[*nproto.RPCMethod]nproto.RPCHandler{
		methodRunner__Sleep: func(ctx context.Context, input proto.Message) (proto.Message, error) {
			return svc.Sleep(ctx, input.(*duration.Duration))
		},
	})
}

// InvokeRunner invokes Runner service using a RPC client.
func InvokeRunner(client nproto.RPCClient, svcName string) Runner {
	return &clientRunner{
		handlerSleep: client.MakeHandler(svcName, methodRunner__Sleep),
	}
}

var methodRunner__Sleep = &nproto.RPCMethod{
	Name:      "Sleep",
	NewInput:  func() proto.Message { return &duration.Duration{} },
	NewOutput: func() proto.Message { return &empty.Empty{} },
}

type clientRunner struct {
	handlerSleep nproto.RPCHandler
}

// Sleep implements Runner interface.
func (svc *clientRunner) Sleep(ctx context.Context, input *duration.Duration) (*empty.Empty, error) {
	output, err := svc.handlerSleep(ctx, input)
	if err != nil {
		return nil, err
	}
	return output.(*empty.Empty), nil
}

// @@protoc_insertion_point(nprpc)

// @@protoc_insertion_point(npmsg)
