//go:generate protoc --go_out=paths=source_relative:. pbenc.proto

package enc

import (
	"github.com/golang/protobuf/proto"

	"github.com/huangjunwen/nproto/nproto"
)

// PBMsgPayloadEncoder is MsgPayloadEncoder using protobuf encoding.
type PBMsgPayloadEncoder struct{}

// PBMsgPayloadDecoder is MsgPayloadDecoder using protobuf encoding.
type PBMsgPayloadDecoder struct{}

var (
	_ MsgPayloadEncoder = PBMsgPayloadEncoder{}
	_ MsgPayloadDecoder = PBMsgPayloadDecoder{}
)

// EncodePayload implements MsgPayloadEncoder interface.
func (e PBMsgPayloadEncoder) EncodePayload(payload *MsgPayload) ([]byte, error) {
	var err error
	p := &PBPayload{}

	// Encode Msg.
	p.Msg, err = proto.Marshal(payload.Msg)
	if err != nil {
		return nil, err
	}

	// Meta data.
	if len(payload.MetaData) != 0 {
		for key, vals := range payload.MetaData {
			p.MetaData = append(p.MetaData, &MetaDataKV{
				Key:    key,
				Values: vals,
			})
		}
	}
	return proto.Marshal(p)
}

// DecodePayload implements MsgPayloadDecoder interface.
func (e PBMsgPayloadDecoder) DecodePayload(data []byte, payload *MsgPayload) error {
	// Decode payload.
	p := &PBPayload{}
	if err := proto.Unmarshal(data, p); err != nil {
		return err
	}

	// Decode msg.
	if err := proto.Unmarshal(p.Msg, payload.Msg); err != nil {
		return err
	}

	// Meta data.
	if len(p.MetaData) != 0 {
		payload.MetaData = nproto.MetaData{}
		for _, kv := range p.MetaData {
			payload.MetaData[kv.Key] = kv.Values
		}
	}
	return nil
}
