// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/rpc/status.proto

package rpcpb

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

// The `Status` message contains two pieces of data: error code and error message.
// The error code should be an enum value of [cs3.rpc.code].
// The error message should be a developer-facing English
// message that helps developers *understand* and *resolve* the error.
type Status struct {
	// The status code, which should be an enum value of [cs3.rpc.code][cs3.rpc.code].
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=cs3.rpc.Code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_d143c7522a91ca77, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_CODE_INVALID
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "cs3.rpc.Status")
}

func init() { proto.RegisterFile("cs3/rpc/status.proto", fileDescriptor_d143c7522a91ca77) }

var fileDescriptor_d143c7522a91ca77 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2e, 0x36, 0xd6,
	0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0x2e, 0x36, 0xd6, 0x2b, 0x2a, 0x48, 0x96, 0x12, 0x82, 0x49, 0x27, 0xe7, 0xa7,
	0xa4, 0x42, 0x24, 0x95, 0x5c, 0xb9, 0xd8, 0x82, 0xc1, 0x8a, 0x85, 0x14, 0xb9, 0x58, 0x40, 0xe2,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xbc, 0x7a, 0x50, 0x5d, 0x7a, 0xce, 0xf9, 0x29, 0xa9,
	0x41, 0x60, 0x29, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0xd7, 0xc9, 0x9c, 0x8b, 0x3b, 0x39, 0x3f, 0x17, 0xa6, 0xc7,
	0x89, 0x1b, 0x62, 0x66, 0x00, 0xc8, 0x8a, 0x00, 0xc6, 0x28, 0xd6, 0xa2, 0x82, 0xe4, 0x82, 0xa4,
	0x45, 0x4c, 0xec, 0xce, 0x4e, 0xfe, 0x11, 0x41, 0x01, 0xce, 0xa7, 0x98, 0xd8, 0x9d, 0x83, 0x8d,
	0x63, 0x82, 0x02, 0x9c, 0x93, 0xd8, 0xc0, 0xce, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51,
	0x2f, 0x17, 0xc5, 0xbb, 0x00, 0x00, 0x00,
}
