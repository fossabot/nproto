// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trace.proto

package traceapi // import "github.com/huangjunwen/nproto/tests/trace/api"

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

type RecursiveRequest struct {
	Depth                int32    `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecursiveRequest) Reset()         { *m = RecursiveRequest{} }
func (m *RecursiveRequest) String() string { return proto.CompactTextString(m) }
func (*RecursiveRequest) ProtoMessage()    {}
func (*RecursiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_b08f2e409ba4f1c7, []int{0}
}
func (m *RecursiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecursiveRequest.Unmarshal(m, b)
}
func (m *RecursiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecursiveRequest.Marshal(b, m, deterministic)
}
func (dst *RecursiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecursiveRequest.Merge(dst, src)
}
func (m *RecursiveRequest) XXX_Size() int {
	return xxx_messageInfo_RecursiveRequest.Size(m)
}
func (m *RecursiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecursiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecursiveRequest proto.InternalMessageInfo

func (m *RecursiveRequest) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type RecursiveReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecursiveReply) Reset()         { *m = RecursiveReply{} }
func (m *RecursiveReply) String() string { return proto.CompactTextString(m) }
func (*RecursiveReply) ProtoMessage()    {}
func (*RecursiveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_b08f2e409ba4f1c7, []int{1}
}
func (m *RecursiveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecursiveReply.Unmarshal(m, b)
}
func (m *RecursiveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecursiveReply.Marshal(b, m, deterministic)
}
func (dst *RecursiveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecursiveReply.Merge(dst, src)
}
func (m *RecursiveReply) XXX_Size() int {
	return xxx_messageInfo_RecursiveReply.Size(m)
}
func (m *RecursiveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RecursiveReply.DiscardUnknown(m)
}

var xxx_messageInfo_RecursiveReply proto.InternalMessageInfo

func (m *RecursiveReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*RecursiveRequest)(nil), "huangjunwen.nproto.tests.traceapi.RecursiveRequest")
	proto.RegisterType((*RecursiveReply)(nil), "huangjunwen.nproto.tests.traceapi.RecursiveReply")
}

func init() { proto.RegisterFile("trace.proto", fileDescriptor_trace_b08f2e409ba4f1c7) }

var fileDescriptor_trace_b08f2e409ba4f1c7 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x29, 0x4a, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xcc, 0x28, 0x4d, 0xcc, 0x4b, 0xcf, 0x2a,
	0xcd, 0x2b, 0x4f, 0xcd, 0xd3, 0xcb, 0x03, 0x8b, 0xe9, 0x95, 0xa4, 0x16, 0x97, 0x14, 0xeb, 0x81,
	0x55, 0x25, 0x16, 0x64, 0x2a, 0x69, 0x70, 0x09, 0x04, 0xa5, 0x26, 0x97, 0x16, 0x15, 0x67, 0x96,
	0xa5, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0xa6, 0xa4, 0x16, 0x94,
	0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x4a, 0x1a, 0x5c, 0x7c, 0x48, 0x2a,
	0x0b, 0x72, 0x2a, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0, 0x0a, 0xa1,
	0x3c, 0xa3, 0x1a, 0x2e, 0xd6, 0x10, 0x90, 0xf9, 0x42, 0xc5, 0x5c, 0x9c, 0x70, 0x2d, 0x42, 0xc6,
	0x7a, 0x04, 0x5d, 0xa3, 0x87, 0xee, 0x14, 0x29, 0x43, 0xd2, 0x34, 0x15, 0xe4, 0x54, 0x3a, 0x59,
	0x44, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x69, 0xd7,
	0x87, 0x68, 0xd7, 0x07, 0x6b, 0xd7, 0x07, 0x6b, 0xd7, 0x4f, 0x2c, 0xc8, 0xb4, 0x86, 0x19, 0x94,
	0xc4, 0x06, 0x96, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xda, 0x8c, 0x95, 0x44, 0x01,
	0x00, 0x00,
}
