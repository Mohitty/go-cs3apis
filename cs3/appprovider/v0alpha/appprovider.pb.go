// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/appprovider/v0alpha/appprovider.proto

package appproviderv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc"
	v0alpha "github.com/cs3org/go-cs3apis/cs3/storageprovider/v0alpha"
	types "github.com/cs3org/go-cs3apis/cs3/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OpenRequest struct {
	// REQUIRED.
	// The id of the resource to be opened.
	ResourceId *v0alpha.ResourceId `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	// The access token the application provider will use when contacting
	// the storage provider storing the resource pointed by resource_id.
	// Service implementors MUST make sure that the access token only grants
	// access to the resource the open request is made.
	AccessToken          string   `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74a8f301591c1e7, []int{0}
}

func (m *OpenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenRequest.Unmarshal(m, b)
}
func (m *OpenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenRequest.Marshal(b, m, deterministic)
}
func (m *OpenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenRequest.Merge(m, src)
}
func (m *OpenRequest) XXX_Size() int {
	return xxx_messageInfo_OpenRequest.Size(m)
}
func (m *OpenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenRequest proto.InternalMessageInfo

func (m *OpenRequest) GetResourceId() *v0alpha.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *OpenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type OpenResponse struct {
	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url to render, usually inside an IFrame in the client.
	IframeUrl            string   `protobuf:"bytes,3,opt,name=iframe_url,json=iframeUrl,proto3" json:"iframe_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74a8f301591c1e7, []int{1}
}

func (m *OpenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenResponse.Unmarshal(m, b)
}
func (m *OpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenResponse.Marshal(b, m, deterministic)
}
func (m *OpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenResponse.Merge(m, src)
}
func (m *OpenResponse) XXX_Size() int {
	return xxx_messageInfo_OpenResponse.Size(m)
}
func (m *OpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenResponse proto.InternalMessageInfo

func (m *OpenResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenResponse) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenResponse) GetIframeUrl() string {
	if m != nil {
		return m.IframeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenRequest)(nil), "cs3.appproviderv0alpha.OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "cs3.appproviderv0alpha.OpenResponse")
}

func init() {
	proto.RegisterFile("cs3/appprovider/v0alpha/appprovider.proto", fileDescriptor_f74a8f301591c1e7)
}

var fileDescriptor_f74a8f301591c1e7 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x6a, 0xe2, 0x40,
	0x14, 0x86, 0x89, 0xbb, 0x04, 0x9c, 0x08, 0xbb, 0x3b, 0xb8, 0x22, 0x81, 0x82, 0xb5, 0xa5, 0x55,
	0x0a, 0xa3, 0x98, 0x27, 0x48, 0xbc, 0x28, 0xbd, 0xd2, 0xc6, 0xb6, 0x94, 0x52, 0x90, 0x71, 0x3c,
	0xb5, 0xa1, 0xea, 0x8c, 0x33, 0x89, 0xe0, 0xa5, 0xaf, 0xd2, 0xcb, 0x3e, 0x4a, 0x9f, 0xaa, 0x4c,
	0x66, 0x6c, 0x43, 0x6d, 0x7b, 0x13, 0xc2, 0x77, 0xbe, 0x9c, 0xff, 0x9c, 0x1c, 0xd4, 0x66, 0x2a,
	0xe8, 0x50, 0x21, 0x84, 0xe4, 0xeb, 0x64, 0x0a, 0xb2, 0xb3, 0xee, 0xd2, 0xb9, 0x78, 0xa4, 0x45,
	0x46, 0x84, 0xe4, 0x29, 0xc7, 0x35, 0xa6, 0x02, 0x52, 0xc0, 0xd6, 0xf4, 0xab, 0xba, 0x85, 0x14,
	0xac, 0xa3, 0x52, 0x9a, 0x66, 0xca, 0xd8, 0xfe, 0x99, 0xa6, 0x2a, 0xe5, 0x92, 0xce, 0x60, 0xaf,
	0xb9, 0x04, 0xc5, 0x33, 0xc9, 0x60, 0x27, 0xff, 0xd7, 0x72, 0xba, 0x11, 0xa0, 0xcc, 0xd3, 0xe0,
	0xe6, 0x06, 0x79, 0x03, 0x01, 0xcb, 0x18, 0x56, 0x19, 0xa8, 0x14, 0x9f, 0x23, 0x6f, 0xf7, 0xe1,
	0x38, 0x99, 0xd6, 0x9d, 0x86, 0xd3, 0xf2, 0x7a, 0x27, 0x44, 0x8f, 0xf5, 0x29, 0xc8, 0xe6, 0x90,
	0xd8, 0xea, 0x17, 0xd3, 0x18, 0xc9, 0xf7, 0x77, 0x7c, 0x88, 0x2a, 0x94, 0x31, 0x50, 0x6a, 0x9c,
	0xf2, 0x27, 0x58, 0xd6, 0x7f, 0x35, 0x9c, 0x56, 0x39, 0xf6, 0x0c, 0xbb, 0xd2, 0xa8, 0xb9, 0x75,
	0x50, 0xc5, 0x64, 0x2b, 0xc1, 0x97, 0x0a, 0xf0, 0x29, 0x72, 0xcd, 0x7e, 0x36, 0xf7, 0x4f, 0x9e,
	0x2b, 0x05, 0x23, 0xa3, 0x1c, 0xc7, 0xb6, 0x8c, 0xdb, 0xc8, 0xe5, 0x82, 0xae, 0x32, 0xa8, 0x97,
	0x72, 0xf1, 0x5f, 0x2e, 0x9a, 0xb5, 0x06, 0x79, 0x21, 0xb6, 0x02, 0x3e, 0x40, 0x28, 0x79, 0x90,
	0x74, 0x01, 0xe3, 0x4c, 0xce, 0xed, 0x14, 0x65, 0x43, 0xae, 0xe5, 0xbc, 0x37, 0x43, 0x38, 0x14,
	0x62, 0x68, 0x77, 0x1a, 0x81, 0x5c, 0x27, 0x0c, 0xf0, 0x25, 0xfa, 0xad, 0x07, 0xc3, 0x47, 0xe4,
	0xeb, 0x7b, 0x90, 0xc2, 0x2f, 0xf3, 0x8f, 0x7f, 0x96, 0xcc, 0x6e, 0xd1, 0xd6, 0x41, 0x3e, 0xe3,
	0x8b, 0x6f, 0xdc, 0xe8, 0x6f, 0xf8, 0xc1, 0x86, 0xfa, 0x30, 0x43, 0xe7, 0xae, 0xba, 0xef, 0x89,
	0xc9, 0x73, 0xc9, 0xed, 0x47, 0x83, 0xdb, 0x30, 0x7a, 0x29, 0xd5, 0xfa, 0xa3, 0x80, 0x14, 0xa6,
	0xbf, 0xe9, 0x86, 0xda, 0x79, 0xcd, 0x0b, 0xf7, 0xfb, 0x85, 0x89, 0x9b, 0x9f, 0x3c, 0x78, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x96, 0x74, 0x4c, 0x91, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppProviderServiceClient is the client API for AppProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppProviderServiceClient interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
}

type appProviderServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppProviderServiceClient(cc *grpc.ClientConn) AppProviderServiceClient {
	return &appProviderServiceClient{cc}
}

func (c *appProviderServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, "/cs3.appproviderv0alpha.AppProviderService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppProviderServiceServer is the server API for AppProviderService service.
type AppProviderServiceServer interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
}

// UnimplementedAppProviderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAppProviderServiceServer struct {
}

func (*UnimplementedAppProviderServiceServer) Open(ctx context.Context, req *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}

func RegisterAppProviderServiceServer(s *grpc.Server, srv AppProviderServiceServer) {
	s.RegisterService(&_AppProviderService_serviceDesc, srv)
}

func _AppProviderService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.appproviderv0alpha.AppProviderService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.appproviderv0alpha.AppProviderService",
	HandlerType: (*AppProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _AppProviderService_Open_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/appprovider/v0alpha/appprovider.proto",
}
