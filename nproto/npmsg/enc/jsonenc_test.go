package enc

import (
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	"github.com/huangjunwen/nproto/nproto"
)

func TestJSONEncodeDecode(t *testing.T) {

	assert := assert.New(t)

	msg := ptypes.TimestampNow()
	md := nproto.NewMetaDataPairs("a", "z")

	data := []byte{}
	err := error(nil)

	// Encode.
	{
		p := &MsgPayload{
			Msg:      msg,
			MetaData: md,
		}

		data, err = JSONMsgPayloadEncoder{}.EncodePayload(p)
		assert.NoError(err)
	}

	// Decode.
	{
		m := timestamp.Timestamp{}
		p := &MsgPayload{
			Msg: &m,
		}
		err = JSONMsgPayloadDecoder{}.DecodePayload(data, p)
		assert.NoError(err)

		assert.Equal(msg.Seconds, m.Seconds)
		assert.Equal(msg.Nanos, m.Nanos)
		assert.Equal(md, p.MetaData)
	}

	// Panic if Msg not set
	{
		p := &MsgPayload{}
		assert.Panics(func() {
			JSONMsgPayloadDecoder{}.DecodePayload(data, p)
		})
	}
}
