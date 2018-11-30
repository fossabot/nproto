// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbenc.proto

package enc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

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
	// Timeout sets an optinal timeout for this RPC.
	Timeout *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Passthru is an optional dict carrying context values.
	Passthru             map[string]string `protobuf:"bytes,3,rep,name=passthru,proto3" json:"passthru,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PBRequest) Reset()         { *m = PBRequest{} }
func (m *PBRequest) String() string { return proto.CompactTextString(m) }
func (*PBRequest) ProtoMessage()    {}
func (*PBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbenc_b3ea0e8d16b158ce, []int{0}
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

func (m *PBRequest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *PBRequest) GetPassthru() map[string]string {
	if m != nil {
		return m.Passthru
	}
	return nil
}

// PBReply is reply of a RPC call encoded by protobuf.
type PBReply struct {
	// Types that are valid to be assigned to Reply:
	//	*PBReply_Result
	//	*PBReply_Error
	Reply                isPBReply_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PBReply) Reset()         { *m = PBReply{} }
func (m *PBReply) String() string { return proto.CompactTextString(m) }
func (*PBReply) ProtoMessage()    {}
func (*PBReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbenc_b3ea0e8d16b158ce, []int{1}
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

type isPBReply_Reply interface {
	isPBReply_Reply()
}

type PBReply_Result struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}
type PBReply_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*PBReply_Result) isPBReply_Reply() {}
func (*PBReply_Error) isPBReply_Reply()  {}

func (m *PBReply) GetReply() isPBReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *PBReply) GetResult() []byte {
	if x, ok := m.GetReply().(*PBReply_Result); ok {
		return x.Result
	}
	return nil
}

func (m *PBReply) GetError() string {
	if x, ok := m.GetReply().(*PBReply_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PBReply) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PBReply_OneofMarshaler, _PBReply_OneofUnmarshaler, _PBReply_OneofSizer, []interface{}{
		(*PBReply_Result)(nil),
		(*PBReply_Error)(nil),
	}
}

func _PBReply_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PBReply)
	// reply
	switch x := m.Reply.(type) {
	case *PBReply_Result:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Result)
	case *PBReply_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Error)
	case nil:
	default:
		return fmt.Errorf("PBReply.Reply has unexpected type %T", x)
	}
	return nil
}

func _PBReply_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PBReply)
	switch tag {
	case 1: // reply.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Reply = &PBReply_Result{x}
		return true, err
	case 2: // reply.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Reply = &PBReply_Error{x}
		return true, err
	default:
		return false, nil
	}
}

func _PBReply_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PBReply)
	// reply
	switch x := m.Reply.(type) {
	case *PBReply_Result:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Result)))
		n += len(x.Result)
	case *PBReply_Error:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Error)))
		n += len(x.Error)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*PBRequest)(nil), "enc.PBRequest")
	proto.RegisterMapType((map[string]string)(nil), "enc.PBRequest.PassthruEntry")
	proto.RegisterType((*PBReply)(nil), "enc.PBReply")
}

func init() { proto.RegisterFile("pbenc.proto", fileDescriptor_pbenc_b3ea0e8d16b158ce) }

var fileDescriptor_pbenc_b3ea0e8d16b158ce = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x97, 0x95, 0xae, 0xf6, 0x56, 0x41, 0x82, 0x48, 0x1d, 0x22, 0x65, 0x4f, 0x7d, 0xca,
	0x60, 0x7b, 0x19, 0xfa, 0x36, 0x26, 0xec, 0x71, 0xe4, 0x1f, 0xa4, 0xf3, 0x3a, 0x87, 0x5d, 0x13,
	0x6f, 0x13, 0xa1, 0x3f, 0xd2, 0xff, 0x24, 0x69, 0xba, 0xc1, 0xde, 0x72, 0x72, 0xcf, 0x3d, 0xf7,
	0x3b, 0x90, 0x99, 0x0a, 0x9b, 0xbd, 0x30, 0xa4, 0xad, 0xe6, 0x11, 0x36, 0xfb, 0xe9, 0xcb, 0x41,
	0xeb, 0x43, 0x8d, 0xf3, 0xfe, 0xab, 0x72, 0x9f, 0xf3, 0x0f, 0x47, 0xca, 0x1e, 0x75, 0x13, 0x4c,
	0xb3, 0x3f, 0x06, 0xe9, 0x6e, 0x2d, 0xf1, 0xc7, 0x61, 0x6b, 0xf9, 0x03, 0xc4, 0x46, 0x91, 0x3a,
	0xe5, 0xac, 0x60, 0xe5, 0xad, 0x0c, 0x82, 0x2f, 0x21, 0xb1, 0xc7, 0x13, 0x6a, 0x67, 0xf3, 0x71,
	0xc1, 0xca, 0x6c, 0xf1, 0x24, 0x42, 0xaa, 0x38, 0xa7, 0x8a, 0xcd, 0x90, 0x2a, 0xcf, 0x4e, 0xbe,
	0x82, 0x1b, 0xa3, 0xda, 0xd6, 0x7e, 0x91, 0xcb, 0xa3, 0x22, 0x2a, 0xb3, 0xc5, 0xb3, 0xf0, 0x6c,
	0x97, 0x63, 0x62, 0x37, 0x8c, 0xdf, 0x1b, 0x4b, 0x9d, 0xbc, 0xb8, 0xa7, 0x6f, 0x70, 0x77, 0x35,
	0xe2, 0xf7, 0x10, 0x7d, 0x63, 0xd7, 0x33, 0xa5, 0xd2, 0x3f, 0x3d, 0xe7, 0xaf, 0xaa, 0x1d, 0xf6,
	0x3c, 0xa9, 0x0c, 0xe2, 0x75, 0xbc, 0x62, 0xb3, 0x0d, 0x24, 0xfe, 0x82, 0xa9, 0x3b, 0x9e, 0xc3,
	0x84, 0xb0, 0x75, 0xb5, 0x0d, 0x6d, 0xb6, 0x23, 0x39, 0x68, 0xfe, 0x08, 0x31, 0x12, 0x69, 0x0a,
	0xeb, 0xdb, 0x91, 0x0c, 0x72, 0x9d, 0x40, 0x4c, 0x7e, 0xb5, 0x9a, 0xf4, 0xc5, 0x96, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0xf3, 0xf4, 0x48, 0x50, 0x01, 0x00, 0x00,
}