// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/usershareprovider/v0alpha/resources.proto

package usershareproviderv0alphapb

import (
	fmt "fmt"
	sharetypes "github.com/cs3org/go-cs3apis/cs3/sharetypes"
	v0alpha "github.com/cs3org/go-cs3apis/cs3/storageprovider/v0alpha"
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

// The state of the share.
type ShareState int32

const (
	// The share is no longer valid, for example, the share expired.
	ShareState_SHARE_STATE_INVALID ShareState = 0
	// New shares MUST be created in the SHARE_STATE_PENDING state.
	// This state means the share is pending to be accepted or rejected
	// by the recipient of the share.
	ShareState_SHARE_STATE_PENDING ShareState = 1
	// The recipient of the share has accepted the share.
	ShareState_SHARE_STATE_ACCEPTED ShareState = 2
	// The recipient of the share has rejected the share.
	// Do not means the share is removed, the recipient MAY
	// change the state to accepted or pending.
	ShareState_SHARE_STATE_REJECTED ShareState = 3
)

var ShareState_name = map[int32]string{
	0: "SHARE_STATE_INVALID",
	1: "SHARE_STATE_PENDING",
	2: "SHARE_STATE_ACCEPTED",
	3: "SHARE_STATE_REJECTED",
}

var ShareState_value = map[string]int32{
	"SHARE_STATE_INVALID":  0,
	"SHARE_STATE_PENDING":  1,
	"SHARE_STATE_ACCEPTED": 2,
	"SHARE_STATE_REJECTED": 3,
}

func (x ShareState) String() string {
	return proto.EnumName(ShareState_name, int32(x))
}

func (ShareState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{0}
}

// Represents the information of the share provider.
type ProviderInfo struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The share type that will become part of the
	// share id.
	// For example, if the share_type is "user", shares obtained
	// from this storage provider will have a share id like "user:1234"
	// or "ocm:1234" or "group:8838".
	ShareType sharetypes.ShareType `protobuf:"varint,2,opt,name=share_type,json=shareType,proto3,enum=cs3.sharetypes.ShareType" json:"share_type,omitempty"`
	// REQUIRED.
	// The address where the share provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the share provider. Meant to be read
	// by humans.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{0}
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

func (m *ProviderInfo) GetShareType() sharetypes.ShareType {
	if m != nil {
		return m.ShareType
	}
	return sharetypes.ShareType_SHARE_TYPE_INVALID
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

// Shares are relationships between a resource owner
// (usually the authenticated user) who grants permissions to a recipient (grantee)
// on a specified resource (resource_id). UserShares represents both user and groups.
type Share struct {
	// REQUIRED.
	// Opaque unique identifier of the share.
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// Unique identifier of the shared resource.
	ResourceId *v0alpha.ResourceId `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	// Permissions for the grantee to use
	// the resource.
	Permissions *SharePermissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// REQUIRED.
	// The receiver of the share, like a user, group ...
	Grantee *v0alpha.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies the owner of the share
	// (the resource owner at the time of creating the share).
	// In case the ownership of the underlying resource changes
	// the share owner field MAY change to reflect the change of ownsership.
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the share creation.
	// A creator can create shares on behalf of the owner (because of re-sharing,
	// because belonging to special groups, ...).
	// Creator and owner often result in being the same principal.
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the share.
	Ctime *types.Timestamp `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// REQUIRED.
	// Last modification time of the share.
	Mtime *types.Timestamp `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// OPTIONAL.
	// Display name for the shared resource (such as file, directory basename or any
	// user defined name).
	// The display name MAY be different than the actual resource basename.
	// This field is only useful for informational purposes.
	DisplayName          string   `protobuf:"bytes,9,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Share) Reset()         { *m = Share{} }
func (m *Share) String() string { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()    {}
func (*Share) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{1}
}

func (m *Share) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Share.Unmarshal(m, b)
}
func (m *Share) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Share.Marshal(b, m, deterministic)
}
func (m *Share) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Share.Merge(m, src)
}
func (m *Share) XXX_Size() int {
	return xxx_messageInfo_Share.Size(m)
}
func (m *Share) XXX_DiscardUnknown() {
	xxx_messageInfo_Share.DiscardUnknown(m)
}

var xxx_messageInfo_Share proto.InternalMessageInfo

func (m *Share) GetId() *ShareId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Share) GetResourceId() *v0alpha.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *Share) GetPermissions() *SharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Share) GetGrantee() *v0alpha.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *Share) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Share) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Share) GetCtime() *types.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Share) GetMtime() *types.Timestamp {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func (m *Share) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// The permissions for a share.
type SharePermissions struct {
	Permissions          *v0alpha.ResourcePermissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SharePermissions) Reset()         { *m = SharePermissions{} }
func (m *SharePermissions) String() string { return proto.CompactTextString(m) }
func (*SharePermissions) ProtoMessage()    {}
func (*SharePermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{2}
}

func (m *SharePermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharePermissions.Unmarshal(m, b)
}
func (m *SharePermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharePermissions.Marshal(b, m, deterministic)
}
func (m *SharePermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharePermissions.Merge(m, src)
}
func (m *SharePermissions) XXX_Size() int {
	return xxx_messageInfo_SharePermissions.Size(m)
}
func (m *SharePermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_SharePermissions.DiscardUnknown(m)
}

var xxx_messageInfo_SharePermissions proto.InternalMessageInfo

func (m *SharePermissions) GetPermissions() *v0alpha.ResourcePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// A received share is the share that a grantee will receive.
// It expands the original share by adding state to the share,
// a display name from the perspective of the grantee and a
// resource mount path in case the share will be mounted
// in a path in a storage provider.
type ReceivedShare struct {
	// REQUIRED.
	Share *Share `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	// REQUIRED.
	// The state of the share.
	State ShareState `protobuf:"varint,2,opt,name=state,proto3,enum=cs3.usershareproviderv0alpha.ShareState" json:"state,omitempty"`
	// OPTIONAL.
	// The storage provider path for the share to be mounted
	// in one of the available storage providers of the user.
	MountPath string `protobuf:"bytes,3,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// OPTIONAL.
	// Display name for the shared resource (such as file, directory basename or any
	// user defined name).
	// The display name MAY be different than the actual resource basename.
	// This field is only useful for informational purposes.
	DisplayName          string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceivedShare) Reset()         { *m = ReceivedShare{} }
func (m *ReceivedShare) String() string { return proto.CompactTextString(m) }
func (*ReceivedShare) ProtoMessage()    {}
func (*ReceivedShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{3}
}

func (m *ReceivedShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedShare.Unmarshal(m, b)
}
func (m *ReceivedShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedShare.Marshal(b, m, deterministic)
}
func (m *ReceivedShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedShare.Merge(m, src)
}
func (m *ReceivedShare) XXX_Size() int {
	return xxx_messageInfo_ReceivedShare.Size(m)
}
func (m *ReceivedShare) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedShare.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedShare proto.InternalMessageInfo

func (m *ReceivedShare) GetShare() *Share {
	if m != nil {
		return m.Share
	}
	return nil
}

func (m *ReceivedShare) GetState() ShareState {
	if m != nil {
		return m.State
	}
	return ShareState_SHARE_STATE_INVALID
}

func (m *ReceivedShare) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *ReceivedShare) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// Uniquely identifies a share in the share provider.
// A share MUST be uniquely identify by four (4) elements:
// 1) The share provider id
// 2) The owner of the share
// 3) The resource id
// 4) The grantee for the share
// This 4-tuple MUST be unique.
// For example, owner Alice shares the resource /home/docs with id
// home:1234 to an user named Bob. The 4-tuple will consist of
// 1) The share provider id = "user"
// 2) The owner of the share = "Alice"
// 3) The resource id = "home:1234"
// 4) The grantee for the share = Grantee("type" = "user", "" => "Bob")
type ShareKey struct {
	// REQUIRED.
	Type sharetypes.ShareType `protobuf:"varint,1,opt,name=type,proto3,enum=cs3.sharetypes.ShareType" json:"type,omitempty"`
	// REQUIRED.
	Owner *types.UserId `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	ResourceId *v0alpha.ResourceId `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	Grantee              *v0alpha.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ShareKey) Reset()         { *m = ShareKey{} }
func (m *ShareKey) String() string { return proto.CompactTextString(m) }
func (*ShareKey) ProtoMessage()    {}
func (*ShareKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{4}
}

func (m *ShareKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareKey.Unmarshal(m, b)
}
func (m *ShareKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareKey.Marshal(b, m, deterministic)
}
func (m *ShareKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareKey.Merge(m, src)
}
func (m *ShareKey) XXX_Size() int {
	return xxx_messageInfo_ShareKey.Size(m)
}
func (m *ShareKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareKey.DiscardUnknown(m)
}

var xxx_messageInfo_ShareKey proto.InternalMessageInfo

func (m *ShareKey) GetType() sharetypes.ShareType {
	if m != nil {
		return m.Type
	}
	return sharetypes.ShareType_SHARE_TYPE_INVALID
}

func (m *ShareKey) GetOwner() *types.UserId {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ShareKey) GetResourceId() *v0alpha.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *ShareKey) GetGrantee() *v0alpha.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

// A share id identifies uniquely a // share in the share provider namespace.
// A ShareId MUST be unique inside the share provider.
type ShareId struct {
	// REQUIRED.
	// The share type.
	Type sharetypes.ShareType `protobuf:"varint,1,opt,name=type,proto3,enum=cs3.sharetypes.ShareType" json:"type,omitempty"`
	// REQUIRED.
	// The internal id used by service implementor to
	// uniquely identity the share in the internal
	// implementation of the service.
	OpaqueId             string   `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareId) Reset()         { *m = ShareId{} }
func (m *ShareId) String() string { return proto.CompactTextString(m) }
func (*ShareId) ProtoMessage()    {}
func (*ShareId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{5}
}

func (m *ShareId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareId.Unmarshal(m, b)
}
func (m *ShareId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareId.Marshal(b, m, deterministic)
}
func (m *ShareId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareId.Merge(m, src)
}
func (m *ShareId) XXX_Size() int {
	return xxx_messageInfo_ShareId.Size(m)
}
func (m *ShareId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareId.DiscardUnknown(m)
}

var xxx_messageInfo_ShareId proto.InternalMessageInfo

func (m *ShareId) GetType() sharetypes.ShareType {
	if m != nil {
		return m.Type
	}
	return sharetypes.ShareType_SHARE_TYPE_INVALID
}

func (m *ShareId) GetOpaqueId() string {
	if m != nil {
		return m.OpaqueId
	}
	return ""
}

// The mechanism to identify a share
// in the share provider namespace.
type ShareReference struct {
	// REQUIRED.
	// One of the specifications MUST be specified.
	//
	// Types that are valid to be assigned to Spec:
	//	*ShareReference_Id
	//	*ShareReference_Key
	Spec                 isShareReference_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ShareReference) Reset()         { *m = ShareReference{} }
func (m *ShareReference) String() string { return proto.CompactTextString(m) }
func (*ShareReference) ProtoMessage()    {}
func (*ShareReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{6}
}

func (m *ShareReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareReference.Unmarshal(m, b)
}
func (m *ShareReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareReference.Marshal(b, m, deterministic)
}
func (m *ShareReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareReference.Merge(m, src)
}
func (m *ShareReference) XXX_Size() int {
	return xxx_messageInfo_ShareReference.Size(m)
}
func (m *ShareReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareReference.DiscardUnknown(m)
}

var xxx_messageInfo_ShareReference proto.InternalMessageInfo

type isShareReference_Spec interface {
	isShareReference_Spec()
}

type ShareReference_Id struct {
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type ShareReference_Key struct {
	Key *ShareKey `protobuf:"bytes,2,opt,name=key,proto3,oneof"`
}

func (*ShareReference_Id) isShareReference_Spec() {}

func (*ShareReference_Key) isShareReference_Spec() {}

func (m *ShareReference) GetSpec() isShareReference_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ShareReference) GetId() *ShareId {
	if x, ok := m.GetSpec().(*ShareReference_Id); ok {
		return x.Id
	}
	return nil
}

func (m *ShareReference) GetKey() *ShareKey {
	if x, ok := m.GetSpec().(*ShareReference_Key); ok {
		return x.Key
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ShareReference) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ShareReference_Id)(nil),
		(*ShareReference_Key)(nil),
	}
}

// A share grant specifies the share permissions
// for a grantee.
type ShareGrant struct {
	// REQUIRED.
	// The grantee of the grant.
	Grantee *v0alpha.Grantee `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// The share permissions for the grant.
	Permissions          *SharePermissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ShareGrant) Reset()         { *m = ShareGrant{} }
func (m *ShareGrant) String() string { return proto.CompactTextString(m) }
func (*ShareGrant) ProtoMessage()    {}
func (*ShareGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6733f43a2d3ad9, []int{7}
}

func (m *ShareGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareGrant.Unmarshal(m, b)
}
func (m *ShareGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareGrant.Marshal(b, m, deterministic)
}
func (m *ShareGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareGrant.Merge(m, src)
}
func (m *ShareGrant) XXX_Size() int {
	return xxx_messageInfo_ShareGrant.Size(m)
}
func (m *ShareGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareGrant.DiscardUnknown(m)
}

var xxx_messageInfo_ShareGrant proto.InternalMessageInfo

func (m *ShareGrant) GetGrantee() *v0alpha.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *ShareGrant) GetPermissions() *SharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterEnum("cs3.usershareproviderv0alpha.ShareState", ShareState_name, ShareState_value)
	proto.RegisterType((*ProviderInfo)(nil), "cs3.usershareproviderv0alpha.ProviderInfo")
	proto.RegisterType((*Share)(nil), "cs3.usershareproviderv0alpha.Share")
	proto.RegisterType((*SharePermissions)(nil), "cs3.usershareproviderv0alpha.SharePermissions")
	proto.RegisterType((*ReceivedShare)(nil), "cs3.usershareproviderv0alpha.ReceivedShare")
	proto.RegisterType((*ShareKey)(nil), "cs3.usershareproviderv0alpha.ShareKey")
	proto.RegisterType((*ShareId)(nil), "cs3.usershareproviderv0alpha.ShareId")
	proto.RegisterType((*ShareReference)(nil), "cs3.usershareproviderv0alpha.ShareReference")
	proto.RegisterType((*ShareGrant)(nil), "cs3.usershareproviderv0alpha.ShareGrant")
}

func init() {
	proto.RegisterFile("cs3/usershareprovider/v0alpha/resources.proto", fileDescriptor_ae6733f43a2d3ad9)
}

var fileDescriptor_ae6733f43a2d3ad9 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdb, 0x6e, 0xeb, 0x44,
	0x14, 0xad, 0x9d, 0x5b, 0xb3, 0x73, 0xa8, 0xc2, 0xd0, 0x23, 0x4c, 0x28, 0x22, 0xa4, 0x02, 0x42,
	0x51, 0x93, 0x2a, 0x11, 0xe2, 0x22, 0x81, 0x94, 0x8b, 0xd5, 0x9a, 0xa2, 0xd4, 0x4c, 0xd2, 0x0a,
	0x21, 0xa4, 0xc8, 0xb5, 0x77, 0x1b, 0x8b, 0xfa, 0xc2, 0x8c, 0x53, 0x94, 0x8f, 0xe0, 0x13, 0x78,
	0xe1, 0x81, 0x07, 0x5e, 0xf8, 0x08, 0xde, 0x10, 0x5f, 0xc2, 0x57, 0x20, 0xcf, 0xd8, 0x1c, 0x37,
	0x51, 0xd3, 0xb4, 0x7d, 0x1b, 0xcf, 0x5e, 0x6b, 0x7b, 0x79, 0xed, 0x8b, 0xe1, 0xd0, 0xe6, 0xdd,
	0xf6, 0x9c, 0x23, 0xe3, 0x33, 0x8b, 0x61, 0xc8, 0x82, 0x5b, 0xd7, 0x41, 0xd6, 0xbe, 0x3d, 0xb2,
	0x6e, 0xc2, 0x99, 0xd5, 0x66, 0xc8, 0x83, 0x39, 0xb3, 0x91, 0xb7, 0x42, 0x16, 0x44, 0x01, 0xd9,
	0xb3, 0x79, 0xb7, 0xb5, 0x02, 0x4f, 0xd0, 0xb5, 0x77, 0xe3, 0x64, 0x22, 0x12, 0x2d, 0x42, 0xe4,
	0x99, 0xa3, 0xa4, 0xd7, 0x3e, 0x16, 0x80, 0x28, 0x60, 0xd6, 0xf5, 0x83, 0xef, 0xaa, 0xbd, 0x8c,
	0xc1, 0x32, 0x51, 0x26, 0x47, 0xe3, 0x4f, 0x05, 0x5e, 0x98, 0x09, 0xd7, 0xf0, 0xaf, 0x02, 0xf2,
	0x11, 0x14, 0x83, 0xd0, 0xfa, 0x69, 0x8e, 0x9a, 0x52, 0x57, 0x9a, 0x95, 0xce, 0xeb, 0xad, 0x58,
	0xa4, 0xa4, 0x9c, 0x89, 0x00, 0x4d, 0x00, 0xe4, 0x33, 0x00, 0xa1, 0x69, 0x1a, 0x47, 0x35, 0xb5,
	0xae, 0x34, 0x77, 0x3a, 0x6f, 0x09, 0x78, 0x46, 0xea, 0x38, 0x3e, 0x4e, 0x16, 0x21, 0xd2, 0x32,
	0x4f, 0x8f, 0x44, 0x83, 0x92, 0xe5, 0x38, 0x0c, 0x39, 0xd7, 0x72, 0x75, 0xa5, 0x59, 0xa6, 0xe9,
	0x23, 0xa9, 0x43, 0xc5, 0x41, 0x6e, 0x33, 0x37, 0x8c, 0xdc, 0xc0, 0xd7, 0xf2, 0x22, 0x9a, 0xbd,
	0x6a, 0xfc, 0x95, 0x83, 0x82, 0x48, 0x4a, 0x3e, 0x01, 0xd5, 0x75, 0x12, 0x99, 0xef, 0xb7, 0xd6,
	0x79, 0x29, 0x55, 0x18, 0x0e, 0x55, 0x5d, 0x87, 0x1c, 0x43, 0x25, 0x35, 0x67, 0xea, 0x3a, 0x42,
	0x77, 0xa5, 0xf3, 0x81, 0xd4, 0x7d, 0xd7, 0xcc, 0x94, 0x4d, 0x13, 0xb8, 0xe1, 0x50, 0x60, 0xff,
	0x9f, 0x89, 0x09, 0x95, 0x10, 0x99, 0xe7, 0x72, 0xee, 0x06, 0xbe, 0xfc, 0x92, 0x4a, 0xa7, 0xb5,
	0x81, 0x10, 0xf3, 0x15, 0x8b, 0x66, 0x53, 0x90, 0x2f, 0xa1, 0x74, 0xcd, 0x2c, 0x3f, 0x42, 0x14,
	0x5f, 0x5e, 0xe9, 0xec, 0xaf, 0x93, 0x75, 0x2c, 0xa1, 0x34, 0xe5, 0x90, 0x5d, 0x28, 0x04, 0x3f,
	0xfb, 0xc8, 0xb4, 0x82, 0xb0, 0x4d, 0x3e, 0xc4, 0x66, 0xdb, 0x0c, 0xad, 0x28, 0x60, 0x5a, 0x51,
	0x9a, 0x9d, 0x3c, 0x92, 0x03, 0x28, 0xd8, 0x91, 0xeb, 0xa1, 0x56, 0x12, 0x2f, 0xdb, 0xcd, 0x94,
	0x7a, 0xe2, 0x7a, 0xc8, 0x23, 0xcb, 0x0b, 0xa9, 0x84, 0xc4, 0x58, 0x4f, 0x60, 0xb7, 0xd7, 0x61,
	0x05, 0x84, 0xbc, 0x07, 0x2f, 0x1c, 0x97, 0x87, 0x37, 0xd6, 0x62, 0xea, 0x5b, 0x1e, 0x6a, 0xe5,
	0xa4, 0x8a, 0xf2, 0x6e, 0x64, 0x79, 0xd8, 0x40, 0xa8, 0x2e, 0x5b, 0x41, 0xbe, 0xbd, 0xeb, 0xa7,
	0x2c, 0x6c, 0x7b, 0x93, 0xc2, 0xdc, 0x67, 0x68, 0xe3, 0x1f, 0x05, 0x5e, 0xa3, 0x68, 0xa3, 0x7b,
	0x8b, 0x8e, 0x6c, 0x9a, 0xcf, 0xa1, 0x20, 0x0a, 0x93, 0xa4, 0xdf, 0xdf, 0xa0, 0x5c, 0x54, 0x32,
	0xc8, 0x57, 0x50, 0xe0, 0x91, 0x15, 0xa5, 0xad, 0xde, 0xdc, 0x80, 0x3a, 0x8e, 0xf1, 0x54, 0xd2,
	0xc8, 0x3b, 0x00, 0x5e, 0x30, 0xf7, 0xa3, 0x69, 0x68, 0x45, 0xb3, 0xa4, 0xf1, 0xcb, 0xe2, 0xc6,
	0xb4, 0xa2, 0xd9, 0x8a, 0x6b, 0xf9, 0x55, 0xd7, 0xfe, 0x55, 0x60, 0x5b, 0xe4, 0x3d, 0xc5, 0x05,
	0x39, 0x84, 0xbc, 0x18, 0x3c, 0xe5, 0xa1, 0xc1, 0x13, 0x30, 0xf2, 0x61, 0xda, 0x1c, 0xea, 0xca,
	0x5c, 0x9f, 0x73, 0x64, 0x86, 0x93, 0xf6, 0xcb, 0xd2, 0x7c, 0xe4, 0x9e, 0x3c, 0x1f, 0xcf, 0xeb,
	0xe6, 0xc6, 0x39, 0x94, 0x92, 0xb1, 0x7d, 0xec, 0xa7, 0xbe, 0x0d, 0x65, 0xb9, 0xa2, 0xd2, 0xf9,
	0x2e, 0xd3, 0x6d, 0x79, 0x61, 0x38, 0x8d, 0x5f, 0x14, 0xd8, 0x91, 0x65, 0xc5, 0x2b, 0x64, 0xe8,
	0xdb, 0x48, 0x3e, 0x7d, 0xf4, 0x22, 0x39, 0xd9, 0x12, 0xab, 0xe4, 0x0b, 0xc8, 0xfd, 0x88, 0x8b,
	0x3b, 0x2b, 0x64, 0x3d, 0xf3, 0x14, 0x17, 0x27, 0x5b, 0x34, 0x26, 0xf5, 0x8b, 0x90, 0xe7, 0x21,
	0xda, 0x8d, 0x5f, 0x15, 0x00, 0x11, 0x13, 0x06, 0x64, 0x4d, 0x53, 0x9e, 0xb0, 0x02, 0x96, 0x76,
	0x92, 0xfa, 0xec, 0x9d, 0x74, 0xc0, 0x12, 0x79, 0xa2, 0x95, 0xc9, 0x9b, 0xf0, 0xc6, 0xf8, 0xa4,
	0x47, 0xf5, 0xe9, 0x78, 0xd2, 0x9b, 0xe8, 0x53, 0x63, 0x74, 0xd1, 0xfb, 0xc6, 0x18, 0x56, 0xb7,
	0x96, 0x03, 0xa6, 0x3e, 0x1a, 0x1a, 0xa3, 0xe3, 0xaa, 0x42, 0x34, 0xd8, 0xcd, 0x06, 0x7a, 0x83,
	0x81, 0x6e, 0x4e, 0xf4, 0x61, 0x55, 0x5d, 0x8e, 0x50, 0xfd, 0x6b, 0x7d, 0x10, 0x47, 0x72, 0xfd,
	0xdf, 0x15, 0xa8, 0xdb, 0x81, 0xb7, 0x56, 0x76, 0x7f, 0x27, 0x6d, 0x3b, 0x6e, 0xc6, 0xbf, 0x32,
	0x53, 0xf9, 0xbe, 0x76, 0x1f, 0x36, 0xbc, 0xfc, 0x4d, 0x7d, 0x39, 0xe8, 0x9f, 0x7d, 0x77, 0x3e,
	0xd6, 0xa9, 0x78, 0xa7, 0x49, 0xcf, 0x2e, 0x8c, 0xa1, 0x4e, 0xff, 0x50, 0xf7, 0x06, 0xe3, 0xae,
	0x98, 0x00, 0x69, 0x43, 0x42, 0xbc, 0x38, 0xea, 0xc5, 0xc4, 0xbf, 0x45, 0xf8, 0x87, 0xfb, 0xc2,
	0x97, 0x45, 0xf1, 0x17, 0xed, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x21, 0xf8, 0xdf, 0xf9,
	0x07, 0x00, 0x00,
}
