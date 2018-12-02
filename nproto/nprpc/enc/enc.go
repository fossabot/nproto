package enc

import (
	"time"

	"github.com/huangjunwen/nproto/nproto"

	"github.com/golang/protobuf/proto"
)

// RPCServerEncoder is the server-side encoder.
type RPCServerEncoder interface {
	// DecodeRequest decodes request from data.
	DecodeRequest(data []byte, req *RPCRequest) error
	// EncodeReply encodes reply to data.
	EncodeReply(reply *RPCReply) ([]byte, error)
}

// RPCClientEncoder is the client-side encoder.
type RPCClientEncoder interface {
	// EncodeRequest encodes request to data.
	EncodeRequest(req *RPCRequest) ([]byte, error)
	// DecodeReply decodes reply from data.
	DecodeReply(data []byte, reply *RPCReply) error
}

// RPCRequest is the request of an rpc.
type RPCRequest struct {
	// Param is the parameter of this rpc. Must be filled with an empty object before decoding.
	Param proto.Message
	// Timeout is an optional timeout of this rpc.
	Timeout *time.Duration
	// MetaData is an optional dict.
	MetaData nproto.MetaData
}

// RPCReply is the reply of an rpc.
type RPCReply struct {
	// Result is the normal result of this rpc. Must set to nil if there is an error.
	// Must be filled with an empty object before decoding.
	Result proto.Message
	// Error is the error result of this rpc. Must set to nil if there is no error.
	Error error
}
