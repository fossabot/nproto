package enc

import (
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestPBEncodeDecode(t *testing.T) {

	assert := assert.New(t)
	msg := ptypes.TimestampNow()
	passthru := map[string]string{"a": "z"}

	data := []byte{}
	err := error(nil)

	// Encode.
	{
		p := &MsgPayload{
			Msg:      msg,
			Passthru: passthru,
		}

		data, err = PBPublisherEncoder{}.EncodePayload(p)
		assert.NoError(err)
	}

	// Decode.
	{
		m := timestamp.Timestamp{}
		p := &MsgPayload{
			Msg: &m,
		}
		err = PBSubscriberEncoder{}.DecodePayload(data, p)
		assert.NoError(err)

		assert.Equal(msg.Seconds, m.Seconds)
		assert.Equal(msg.Nanos, m.Nanos)
		assert.Equal(passthru, p.Passthru)
	}

	// Panic if Msg not set
	{
		p := &MsgPayload{}
		assert.Panics(func() {
			PBSubscriberEncoder{}.DecodePayload(data, p)
		})
	}
}
