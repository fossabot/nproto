// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbenc.proto

package enc // import "github.com/huangjunwen/nproto/nproto/nprpc/enc"

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

// PBRequest is request of a RPC call encoded by protobuf.
type PBRequest struct {
	// Param is protobuf encoded param.
	Param []byte `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	// MetaData dict. NOTE: map value can't be repeated type
	// See: https://stackoverflow.com/questions/38886789/protobuf3-how-to-describe-map-of-repeated-string
	MetaData []*MetaDataKV `protobuf:"bytes,2,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	// Timeout is timeout in nanoseconds. Use int64 instead of wkt's duration to avoid an extra pointer.
	Timeout              int64    `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBRequest) Reset()         { *m = PBRequest{} }
func (m *PBRequest) String() string { return proto.CompactTextString(m) }
func (*PBRequest) ProtoMessage()    {}
func (*PBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbenc_0e66c4b1445ef4a9, []int{0}
}
func (m *PBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBRequest.Unmarshal(m, b)
}
func (m *PBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBRequest.Marshal(b, m, deterministic)
}
func (dst *PBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBRequest.Merge(dst, src)
}
func (m *PBRequest) XXX_Size() int {
	return xxx_messageInfo_PBRequest.Size(m)
}
func (m *PBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PBRequest proto.InternalMessageInfo

func (m *PBRequest) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *PBRequest) GetMetaData() []*MetaDataKV {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *PBRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// MetaDataKV is a kv pair of meta data.
type MetaDataKV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaDataKV) Reset()         { *m = MetaDataKV{} }
func (m *MetaDataKV) String() string { return proto.CompactTextString(m) }
func (*MetaDataKV) ProtoMessage()    {}
func (*MetaDataKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbenc_0e66c4b1445ef4a9, []int{1}
}
func (m *MetaDataKV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaDataKV.Unmarshal(m, b)
}
func (m *MetaDataKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaDataKV.Marshal(b, m, deterministic)
}
func (dst *MetaDataKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaDataKV.Merge(dst, src)
}
func (m *MetaDataKV) XXX_Size() int {
	return xxx_messageInfo_MetaDataKV.Size(m)
}
func (m *MetaDataKV) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaDataKV.DiscardUnknown(m)
}

var xxx_messageInfo_MetaDataKV proto.InternalMessageInfo

func (m *MetaDataKV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MetaDataKV) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// PBReply is reply of a RPC call encoded by protobuf.
type PBReply struct {
	// Result is protobuf encoded result.
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// Error is the error result of this rpc.
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBReply) Reset()         { *m = PBReply{} }
func (m *PBReply) String() string { return proto.CompactTextString(m) }
func (*PBReply) ProtoMessage()    {}
func (*PBReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbenc_0e66c4b1445ef4a9, []int{2}
}
func (m *PBReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBReply.Unmarshal(m, b)
}
func (m *PBReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBReply.Marshal(b, m, deterministic)
}
func (dst *PBReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBReply.Merge(dst, src)
}
func (m *PBReply) XXX_Size() int {
	return xxx_messageInfo_PBReply.Size(m)
}
func (m *PBReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PBReply.DiscardUnknown(m)
}

var xxx_messageInfo_PBReply proto.InternalMessageInfo

func (m *PBReply) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *PBReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*PBRequest)(nil), "huangjunwen.nproto.nprpc.enc.PBRequest")
	proto.RegisterType((*MetaDataKV)(nil), "huangjunwen.nproto.nprpc.enc.MetaDataKV")
	proto.RegisterType((*PBReply)(nil), "huangjunwen.nproto.nprpc.enc.PBReply")
}

func init() { proto.RegisterFile("pbenc.proto", fileDescriptor_pbenc_0e66c4b1445ef4a9) }

var fileDescriptor_pbenc_0e66c4b1445ef4a9 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0xc9, 0x05, 0xef, 0xcc, 0x9c, 0x85, 0x2c, 0x22, 0x29, 0x2c, 0x8e, 0x54, 0xa9, 0x36,
	0xa2, 0xa0, 0xfd, 0xa1, 0x95, 0x08, 0xb2, 0x85, 0x85, 0x8d, 0x4c, 0xd6, 0xe1, 0xee, 0x34, 0xfb,
	0xc7, 0xcd, 0xac, 0x62, 0xeb, 0x27, 0x97, 0x6c, 0x22, 0x5e, 0x65, 0x35, 0xf3, 0x7b, 0xec, 0xbe,
	0x79, 0x0f, 0x96, 0xbe, 0x25, 0xab, 0xa5, 0x0f, 0x8e, 0x9d, 0x38, 0xdb, 0x46, 0xb4, 0x9b, 0xd7,
	0x68, 0x3f, 0xc9, 0x4a, 0x9b, 0xb4, 0x61, 0x78, 0x2d, 0xc9, 0xea, 0xea, 0x3b, 0x83, 0xe2, 0x61,
	0xad, 0xe8, 0x3d, 0x52, 0xcf, 0xe2, 0x04, 0x0e, 0x3c, 0x06, 0x34, 0x65, 0xb6, 0xca, 0xea, 0x23,
	0x35, 0x82, 0xb8, 0x85, 0xc2, 0x10, 0xe3, 0xf3, 0x0b, 0x32, 0x96, 0xb3, 0x55, 0x5e, 0x2f, 0x2f,
	0x6a, 0xf9, 0x9f, 0xab, 0xbc, 0x27, 0xc6, 0x1b, 0x64, 0xbc, 0x7b, 0x54, 0x87, 0x66, 0xda, 0x45,
	0x09, 0x0b, 0xde, 0x19, 0x72, 0x91, 0xcb, 0x7c, 0x95, 0xd5, 0xb9, 0xfa, 0xc5, 0xea, 0x0a, 0xe0,
	0xef, 0x87, 0x38, 0x86, 0xfc, 0x8d, 0xbe, 0x52, 0x84, 0x42, 0x0d, 0xab, 0x38, 0x85, 0xf9, 0x07,
	0x76, 0x91, 0xfa, 0x74, 0xbd, 0x50, 0x13, 0x55, 0xd7, 0xb0, 0x18, 0xb2, 0xfb, 0x2e, 0x3d, 0x09,
	0xd4, 0xc7, 0x8e, 0xa7, 0xe8, 0x13, 0x0d, 0x8d, 0x28, 0x04, 0x17, 0xca, 0x59, 0xb2, 0x1b, 0x61,
	0x7d, 0xfe, 0x24, 0x37, 0x3b, 0xde, 0xc6, 0x56, 0x6a, 0x67, 0x9a, 0xbd, 0x2a, 0xcd, 0x58, 0x65,
	0x6f, 0x78, 0xdd, 0x90, 0xd5, 0xed, 0x3c, 0x09, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01,
	0xfb, 0xf6, 0x0e, 0x5b, 0x01, 0x00, 0x00,
}
