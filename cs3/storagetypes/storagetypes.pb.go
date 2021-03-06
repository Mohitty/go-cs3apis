// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/storagetypes/storagetypes.proto

// This package contains common types
// used for sharing.

package storagetypespb

import (
	fmt "fmt"
	types "github.com/cs3org/go-cs3apis/cs3/types"
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

type ProviderInfo struct {
	// OPTIONAL.
	// Opaque information (containing storage-specific information).
	// For example, additional metadata attached to the resource.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage provider id that will become part of the
	// resource id.
	// For example, if the provider_id is "home", resources obtained
	// from this storage provider will have a resource id like "home:1234".
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// REQUIRED.
	// The public path prefix this storage provider handles.
	// In Unix literature, the mount path.
	// For example, if the provider_path is "/home", resources obtained
	// from this storage provier will have a resource path lik "/home/docs".
	ProviderPath string `protobuf:"bytes,3,opt,name=provider_path,json=providerPath,proto3" json:"provider_path,omitempty"`
	// REQUIRED.
	// The address where the storage provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the storage provider. Meant to be read
	// by humans.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// List of available methods.
	Features             *ProviderInfo_Features `protobuf:"bytes,6,opt,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e90c6b3f0c92b987, []int{0}
}

func (m *ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo.Unmarshal(m, b)
}
func (m *ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo.Merge(m, src)
}
func (m *ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo.Size(m)
}
func (m *ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo proto.InternalMessageInfo

func (m *ProviderInfo) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ProviderInfo) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *ProviderInfo) GetProviderPath() string {
	if m != nil {
		return m.ProviderPath
	}
	return ""
}

func (m *ProviderInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProviderInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProviderInfo) GetFeatures() *ProviderInfo_Features {
	if m != nil {
		return m.Features
	}
	return nil
}

// REQUIRED.
// Represents the list of features available
// on this storage provider. If the feature is not supported,
// the related service methods MUST return CODE_UNIMPLEMENTED.
type ProviderInfo_Features struct {
	Recycle              bool     `protobuf:"varint,1,opt,name=recycle,proto3" json:"recycle,omitempty"`
	FileVersions         bool     `protobuf:"varint,2,opt,name=file_versions,json=fileVersions,proto3" json:"file_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo_Features) Reset()         { *m = ProviderInfo_Features{} }
func (m *ProviderInfo_Features) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo_Features) ProtoMessage()    {}
func (*ProviderInfo_Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_e90c6b3f0c92b987, []int{0, 0}
}

func (m *ProviderInfo_Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo_Features.Unmarshal(m, b)
}
func (m *ProviderInfo_Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo_Features.Marshal(b, m, deterministic)
}
func (m *ProviderInfo_Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo_Features.Merge(m, src)
}
func (m *ProviderInfo_Features) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo_Features.Size(m)
}
func (m *ProviderInfo_Features) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo_Features.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo_Features proto.InternalMessageInfo

func (m *ProviderInfo_Features) GetRecycle() bool {
	if m != nil {
		return m.Recycle
	}
	return false
}

func (m *ProviderInfo_Features) GetFileVersions() bool {
	if m != nil {
		return m.FileVersions
	}
	return false
}

func init() {
	proto.RegisterType((*ProviderInfo)(nil), "cs3.storagetypes.ProviderInfo")
	proto.RegisterType((*ProviderInfo_Features)(nil), "cs3.storagetypes.ProviderInfo.Features")
}

func init() {
	proto.RegisterFile("cs3/storagetypes/storagetypes.proto", fileDescriptor_e90c6b3f0c92b987)
}

var fileDescriptor_e90c6b3f0c92b987 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4e, 0xf2, 0x30,
	0x18, 0x86, 0xb3, 0xfd, 0xbf, 0x38, 0x0b, 0x12, 0x68, 0x34, 0x69, 0x38, 0x91, 0xc8, 0x81, 0x78,
	0x52, 0x13, 0x77, 0x05, 0x42, 0x30, 0x72, 0xc4, 0xb2, 0x11, 0xa3, 0xc6, 0x84, 0x8c, 0xb6, 0x48,
	0x13, 0xa4, 0xb5, 0x2d, 0x24, 0xdc, 0x8e, 0x87, 0x5e, 0x8a, 0xde, 0x94, 0x69, 0xb7, 0x91, 0xaa,
	0x87, 0xdf, 0xf3, 0x3d, 0xef, 0xf2, 0xbd, 0x2b, 0xe8, 0x11, 0x1d, 0x5f, 0x69, 0x23, 0x54, 0xfe,
	0xc2, 0xcc, 0x4e, 0x32, 0xfd, 0x63, 0xc0, 0x52, 0x09, 0x23, 0x60, 0x8b, 0xe8, 0x18, 0xfb, 0xbc,
	0x73, 0x6a, 0x63, 0x85, 0xef, 0x89, 0xe7, 0x5f, 0x21, 0x68, 0x24, 0x4a, 0x6c, 0x39, 0x65, 0x6a,
	0xbc, 0x5e, 0x08, 0x78, 0x09, 0x6a, 0x42, 0xe6, 0x6f, 0x1b, 0x86, 0x82, 0x6e, 0xd0, 0xaf, 0x5f,
	0xb7, 0xb1, 0xfd, 0x54, 0x11, 0x99, 0xb8, 0x45, 0x5a, 0x0a, 0xf0, 0x0c, 0xd4, 0x65, 0x19, 0x9d,
	0x71, 0x8a, 0xc2, 0x6e, 0xd0, 0x3f, 0x4a, 0x41, 0x85, 0xc6, 0x14, 0xf6, 0xc0, 0xf1, 0x5e, 0x90,
	0xb9, 0x59, 0xa2, 0x7f, 0x4e, 0x69, 0x54, 0x30, 0xc9, 0xcd, 0x12, 0x22, 0x70, 0x98, 0x53, 0xaa,
	0x98, 0xd6, 0xe8, 0xbf, 0x5b, 0x57, 0x23, 0xec, 0x82, 0x3a, 0x65, 0x9a, 0x28, 0x2e, 0x0d, 0x17,
	0x6b, 0x74, 0xe0, 0xb6, 0x3e, 0x82, 0x43, 0x10, 0x2d, 0x58, 0x6e, 0x36, 0x8a, 0x69, 0x54, 0x73,
	0xe7, 0x5e, 0xe0, 0xdf, 0xcd, 0xb1, 0x5f, 0x0f, 0xdf, 0x96, 0x7a, 0xba, 0x0f, 0x76, 0xc6, 0x20,
	0xaa, 0xa8, 0x3d, 0x46, 0x31, 0xb2, 0x23, 0xab, 0xa2, 0x7e, 0x94, 0x56, 0xa3, 0xed, 0xb2, 0xe0,
	0x2b, 0x36, 0xdb, 0x32, 0xa5, 0xb9, 0x58, 0x6b, 0x57, 0x37, 0x4a, 0x1b, 0x16, 0xde, 0x97, 0x6c,
	0xa0, 0xc0, 0x09, 0x11, 0xaf, 0x7f, 0x4e, 0x18, 0xb4, 0x33, 0x6f, 0x4a, 0xec, 0x8f, 0x4f, 0x82,
	0xa7, 0xa6, 0xaf, 0xc8, 0xf9, 0x7b, 0xd8, 0x1c, 0x0e, 0x26, 0x0f, 0xd9, 0xdd, 0x4d, 0x3a, 0x9a,
	0x3e, 0x26, 0xa3, 0xec, 0x23, 0x6c, 0x0d, 0xb3, 0x18, 0x97, 0xd1, 0xa9, 0xb5, 0x3e, 0x1d, 0x7a,
	0xf6, 0xd1, 0xbc, 0xe6, 0x1e, 0x32, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x41, 0x8b, 0x26, 0x70,
	0x18, 0x02, 0x00, 0x00,
}
