package enc

import (
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestJSONEncodeDecode(t *testing.T) {

	assert := assert.New(t)
	msg := ptypes.TimestampNow()
	passthru := map[string]string{"a": "z"}

	data := []byte{}
	err := error(nil)

	// Encode.
	{
		payload := &MsgPayload{
			Msg:      msg,
			Passthru: passthru,
		}

		data, err = JSONPublisherEncoder{}.EncodePayload(payload)
		assert.NoError(err)
	}

	// Decode.
	{
		m := timestamp.Timestamp{}
		payload := &MsgPayload{
			Msg: &m,
		}
		err = JSONSubscriberEncoder{}.DecodePayload(data, payload)
		assert.NoError(err)

		assert.Equal(m.Seconds, msg.Seconds)
		assert.Equal(m.Nanos, msg.Nanos)
		assert.Equal(passthru, payload.Passthru)
	}

	// Panic if Msg not set
	{
		payload := &MsgPayload{}
		assert.Panics(func() {
			JSONSubscriberEncoder{}.DecodePayload(data, payload)
		})
	}
}
