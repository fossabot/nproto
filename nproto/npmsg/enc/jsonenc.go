package enc

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"

	"github.com/huangjunwen/nproto/nproto"
)

// JSONMsgPayloadEncoder is MsgPayloadEncoder using json encoding.
type JSONMsgPayloadEncoder struct{}

// JSONMsgPayloadDecoder is MsgPayloadDecoder using json encoding.
type JSONMsgPayloadDecoder struct{}

type JSONPayload struct {
	Msg      json.RawMessage `json:"msg"`
	MetaData nproto.MetaData `json:"metadata"`
}

var (
	_ MsgPayloadEncoder = JSONMsgPayloadEncoder{}
	_ MsgPayloadDecoder = JSONMsgPayloadDecoder{}
)

var (
	jsonUnmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	jsonMarshaler = jsonpb.Marshaler{
		EmitDefaults: true,
	}
)

// EncodePayload implements MsgPayloadEncoder interface.
func (e JSONMsgPayloadEncoder) EncodePayload(payload *MsgPayload) ([]byte, error) {
	p := &JSONPayload{}

	// Encode msg.
	buf := &bytes.Buffer{}
	if err := jsonMarshaler.Marshal(buf, payload.Msg); err != nil {
		return nil, err
	}
	p.Msg = json.RawMessage(buf.Bytes())

	// Meta data.
	p.MetaData = payload.MetaData

	// Encode payload.
	return json.Marshal(p)
}

// DecodePayload implements MsgPayloadDecoder interface.
func (e JSONMsgPayloadDecoder) DecodePayload(data []byte, payload *MsgPayload) error {
	// Decode payload.
	p := &JSONPayload{}
	if err := json.Unmarshal(data, p); err != nil {
		return err
	}

	// Decode msg.
	reader := bytes.NewReader(p.Msg)
	if err := jsonUnmarshaler.Unmarshal(reader, payload.Msg); err != nil {
		return err
	}

	// Meta data.
	payload.MetaData = p.MetaData
	return nil
}
