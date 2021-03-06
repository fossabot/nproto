// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// MsgPayload is the payload of a message encoded by protobuf.
type MsgPayload struct {
	// MsgData is serialized message data.
	MsgData []byte `protobuf:"bytes,1,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	// MetaData dict.
	MetaData             []*MetaDataKV `protobuf:"bytes,2,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MsgPayload) Reset()         { *m = MsgPayload{} }
func (m *MsgPayload) String() string { return proto.CompactTextString(m) }
func (*MsgPayload) ProtoMessage()    {}
func (*MsgPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *MsgPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPayload.Unmarshal(m, b)
}
func (m *MsgPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPayload.Marshal(b, m, deterministic)
}
func (m *MsgPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayload.Merge(m, src)
}
func (m *MsgPayload) XXX_Size() int {
	return xxx_messageInfo_MsgPayload.Size(m)
}
func (m *MsgPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayload proto.InternalMessageInfo

func (m *MsgPayload) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *MsgPayload) GetMetaData() []*MetaDataKV {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgPayload)(nil), "nproto.pb.msg.MsgPayload")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0x03, 0xd3, 0x7a, 0x05, 0x49, 0x7a, 0xb9, 0xc5,
	0xe9, 0x52, 0x1c, 0xb9, 0x29, 0x10, 0x09, 0xa5, 0x38, 0x2e, 0x2e, 0xdf, 0xe2, 0xf4, 0x80, 0xc4,
	0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x21, 0x49, 0x2e, 0x8e, 0xdc, 0xe2, 0xf4, 0xf8, 0x94, 0xc4, 0x92,
	0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0xf6, 0xdc, 0xe2, 0x74, 0x97, 0xc4, 0x92, 0x44,
	0x21, 0x53, 0x2e, 0xce, 0xdc, 0xd4, 0x92, 0x44, 0x88, 0x1c, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x84, 0x1e, 0x92, 0xa9, 0x29, 0x7a, 0xbe, 0xa9, 0x25, 0x89, 0x20, 0xa5, 0xde, 0x61, 0x41, 0x1c,
	0xb9, 0x50, 0xb6, 0x93, 0x66, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x7e, 0x46, 0x69, 0x62, 0x5e, 0x7a, 0x56, 0x69, 0x5e, 0x79, 0x6a, 0x9e, 0x3e, 0x44, 0x2f,
	0x8c, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0x7d,
	0x51, 0x02, 0xb7, 0x00, 0x00, 0x00,
}
