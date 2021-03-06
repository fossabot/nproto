// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bench.proto

package benchapi // import "github.com/huangjunwen/nproto/tests/bench/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoMsg struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMsg) Reset()         { *m = EchoMsg{} }
func (m *EchoMsg) String() string { return proto.CompactTextString(m) }
func (*EchoMsg) ProtoMessage()    {}
func (*EchoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_bench_ec478ae3bbc0e85f, []int{0}
}
func (m *EchoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMsg.Unmarshal(m, b)
}
func (m *EchoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMsg.Marshal(b, m, deterministic)
}
func (dst *EchoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMsg.Merge(dst, src)
}
func (m *EchoMsg) XXX_Size() int {
	return xxx_messageInfo_EchoMsg.Size(m)
}
func (m *EchoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMsg proto.InternalMessageInfo

func (m *EchoMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoMsg)(nil), "huangjunwen.nproto.tests.benchapi.EchoMsg")
}

func init() { proto.RegisterFile("bench.proto", fileDescriptor_bench_ec478ae3bbc0e85f) }

var fileDescriptor_bench_ec478ae3bbc0e85f = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4a, 0xcd, 0x4b,
	0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xcc, 0x28, 0x4d, 0xcc, 0x4b, 0xcf, 0x2a,
	0xcd, 0x2b, 0x4f, 0xcd, 0xd3, 0xcb, 0x03, 0x8b, 0xe9, 0x95, 0xa4, 0x16, 0x97, 0x14, 0xeb, 0x81,
	0x55, 0x25, 0x16, 0x64, 0x2a, 0x29, 0x73, 0xb1, 0xbb, 0x26, 0x67, 0xe4, 0xfb, 0x16, 0xa7, 0x0b,
	0x49, 0x70, 0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x04, 0xc1, 0xb8, 0x46, 0xe9, 0x5c, 0xac, 0x4e, 0x20, 0x0d, 0x42, 0x71, 0x5c, 0x2c, 0x20, 0xd5,
	0x42, 0x5a, 0x7a, 0x04, 0x4d, 0xd6, 0x83, 0x1a, 0x2b, 0x45, 0x82, 0x5a, 0x27, 0x8b, 0x28, 0xb3,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x24, 0x7d, 0xfa, 0x10, 0x7d,
	0xfa, 0x60, 0x7d, 0xfa, 0x60, 0x7d, 0xfa, 0x89, 0x05, 0x99, 0xd6, 0x30, 0x13, 0x92, 0xd8, 0xc0,
	0xf2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x6b, 0xd2, 0x93, 0x00, 0x01, 0x00, 0x00,
}
