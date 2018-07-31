package enc

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestPBRequest(t *testing.T) {

	assert := assert.New(t)

	param := *ptypes.TimestampNow()
	timeout := 10 * time.Nanosecond
	passthru := map[string]string{"a": "z"}

	data := []byte{}
	err := error(nil)

	// Encode.
	{
		req := &RPCRequest{
			Param:    &param,
			Timeout:  &timeout,
			Passthru: passthru,
		}

		data, err = PBClientEncoder{}.EncodeRequest(req)
		assert.NoError(err)
	}

	// Normal decode.
	{
		p := timestamp.Timestamp{}
		req := &RPCRequest{
			Param: &p,
		}
		err = PBServerEncoder{}.DecodeRequest(data, req)
		assert.NoError(err)

		assert.Equal(param.Seconds, p.Seconds)
		assert.Equal(param.Nanos, p.Nanos)
		assert.Equal(timeout, *req.Timeout)
		assert.Equal(passthru, req.Passthru)
	}

	// Panic if Param not set
	{
		req := &RPCRequest{}
		assert.Panics(func() {
			PBServerEncoder{}.DecodeRequest(data, req)
		})
	}

}

func TestPBReply(t *testing.T) {

	assert := assert.New(t)

	// Normal result.
	{
		result := *ptypes.TimestampNow()

		data := []byte{}
		err := error(nil)

		// Encode.
		{
			reply := &RPCReply{
				Result: &result,
			}
			data, err = PBServerEncoder{}.EncodeReply(reply)
			assert.NoError(err)
		}

		// Normal decode.
		{
			r := timestamp.Timestamp{}
			reply := &RPCReply{
				Result: &r,
			}
			err = PBClientEncoder{}.DecodeReply(data, reply)
			assert.NoError(err)

			assert.Equal(r.Seconds, result.Seconds)
			assert.Equal(r.Nanos, result.Nanos)
			assert.Nil(reply.Error)
		}

	}

	// Error result.
	{
		errResult := errors.New("Some error")

		data := []byte{}
		err := error(nil)

		// Encode.
		{
			reply := &RPCReply{
				Error: errResult,
			}
			data, err = PBServerEncoder{}.EncodeReply(reply)
			assert.NoError(err)
		}

		// Normal decode.
		{
			r := timestamp.Timestamp{}
			reply := &RPCReply{
				Result: &r,
			}
			err = PBClientEncoder{}.DecodeReply(data, reply)
			assert.NoError(err)

			assert.Equal(reply.Error.Error(), errResult.Error())
			assert.Nil(reply.Result)

		}

	}

}

func BenchmarkPBEncode(b *testing.B) {

	param := *ptypes.TimestampNow()
	timeout := 10 * time.Nanosecond
	passthru := map[string]string{"a": "z"}
	req := &RPCRequest{
		Param:    &param,
		Timeout:  &timeout,
		Passthru: passthru,
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PBClientEncoder{}.EncodeRequest(req)
	}

}

func BenchmarkPBDecode(b *testing.B) {

	param := *ptypes.TimestampNow()
	timeout := 10 * time.Nanosecond
	passthru := map[string]string{"a": "z"}
	data, _ := PBClientEncoder{}.EncodeRequest(&RPCRequest{
		Param:    &param,
		Timeout:  &timeout,
		Passthru: passthru,
	})

	req := &RPCRequest{
		Param: &timestamp.Timestamp{},
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PBServerEncoder{}.DecodeRequest(data, req)
	}

}
