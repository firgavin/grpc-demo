// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package demopb

import (
	context "context"
	fmt "fmt"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "demo.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "demo.v1.HelloResponse")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x4e, 0x03, 0x31,
	0x10, 0x44, 0x31, 0x42, 0x01, 0x56, 0xa1, 0x59, 0x89, 0x88, 0x40, 0x83, 0xae, 0x21, 0x0d, 0x97,
	0x04, 0xca, 0x74, 0xa1, 0x80, 0x12, 0x25, 0x47, 0x43, 0xe7, 0x5c, 0x56, 0x91, 0x25, 0x9f, 0x6d,
	0xd6, 0xce, 0x49, 0xfe, 0x2a, 0x7e, 0x11, 0xd9, 0xb8, 0x40, 0x74, 0x90, 0x6e, 0x67, 0x34, 0x7a,
	0xaf, 0x58, 0x80, 0x2d, 0x75, 0xb6, 0x76, 0x6c, 0x83, 0xc5, 0xd3, 0x7c, 0xf7, 0xf3, 0xaa, 0x82,
	0xe1, 0x0b, 0x69, 0x6d, 0x57, 0xf4, 0xb1, 0x27, 0x1f, 0x10, 0xe1, 0xc4, 0xc8, 0x8e, 0xae, 0xc4,
	0xad, 0x98, 0x9c, 0xaf, 0xf2, 0x5d, 0xdd, 0xc1, 0x45, 0xd9, 0x78, 0x67, 0x8d, 0x27, 0x1c, 0xc1,
	0x80, 0xc9, 0xef, 0x75, 0x28, 0xb3, 0x92, 0x1e, 0x3e, 0x8f, 0x61, 0x98, 0xc0, 0x6b, 0xe2, 0x5e,
	0xb5, 0xe4, 0x71, 0x01, 0x67, 0x0d, 0xc7, 0x37, 0x23, 0x39, 0xe2, 0x65, 0x5d, 0x9c, 0xf5, 0x4f,
	0xe1, 0xf5, 0xe8, 0x77, 0xfd, 0xed, 0xa8, 0x8e, 0xf0, 0x19, 0xb0, 0xe1, 0x98, 0x58, 0xc4, 0xeb,
	0xc0, 0x24, 0x3b, 0x65, 0x76, 0x7f, 0xc6, 0xcc, 0x44, 0x01, 0x3d, 0x69, 0x45, 0x26, 0xfc, 0x1f,
	0x34, 0x11, 0xf8, 0x0a, 0xe3, 0x86, 0xe3, 0x52, 0x6d, 0x15, 0x53, 0x1b, 0x94, 0x35, 0x52, 0x1f,
	0xc2, 0x9b, 0x89, 0xe5, 0xcd, 0xfb, 0x78, 0xc7, 0xae, 0xbd, 0x4f, 0x9b, 0x69, 0x7e, 0xcd, 0xb4,
	0x9f, 0x2f, 0x52, 0x72, 0x9b, 0xcd, 0x20, 0x17, 0x8f, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x13, 0x43, 0x80, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServicesClient is the client API for DemoServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServicesClient interface {
	// unary RPC
	TryUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// server streaming RPC
	TryServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (DemoServices_TryServerStreamingClient, error)
	// client streaming RPC
	TryClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DemoServices_TryClientStreamingClient, error)
	// bidirectional streaming Rpc
	TryBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DemoServices_TryBidirectionalStreamingClient, error)
}

type demoServicesClient struct {
	cc *grpc.ClientConn
}

func NewDemoServicesClient(cc *grpc.ClientConn) DemoServicesClient {
	return &demoServicesClient{cc}
}

func (c *demoServicesClient) TryUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/demo.v1.demoServices/TryUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServicesClient) TryServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (DemoServices_TryServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DemoServices_serviceDesc.Streams[0], "/demo.v1.demoServices/TryServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServicesTryServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DemoServices_TryServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type demoServicesTryServerStreamingClient struct {
	grpc.ClientStream
}

func (x *demoServicesTryServerStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoServicesClient) TryClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DemoServices_TryClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DemoServices_serviceDesc.Streams[1], "/demo.v1.demoServices/TryClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServicesTryClientStreamingClient{stream}
	return x, nil
}

type DemoServices_TryClientStreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type demoServicesTryClientStreamingClient struct {
	grpc.ClientStream
}

func (x *demoServicesTryClientStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoServicesTryClientStreamingClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoServicesClient) TryBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DemoServices_TryBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DemoServices_serviceDesc.Streams[2], "/demo.v1.demoServices/TryBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServicesTryBidirectionalStreamingClient{stream}
	return x, nil
}

type DemoServices_TryBidirectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type demoServicesTryBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *demoServicesTryBidirectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoServicesTryBidirectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServicesServer is the server API for DemoServices service.
type DemoServicesServer interface {
	// unary RPC
	TryUnary(context.Context, *HelloRequest) (*HelloResponse, error)
	// server streaming RPC
	TryServerStreaming(*HelloRequest, DemoServices_TryServerStreamingServer) error
	// client streaming RPC
	TryClientStreaming(DemoServices_TryClientStreamingServer) error
	// bidirectional streaming Rpc
	TryBidirectionalStreaming(DemoServices_TryBidirectionalStreamingServer) error
}

// UnimplementedDemoServicesServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServicesServer struct {
}

func (*UnimplementedDemoServicesServer) TryUnary(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryUnary not implemented")
}
func (*UnimplementedDemoServicesServer) TryServerStreaming(req *HelloRequest, srv DemoServices_TryServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method TryServerStreaming not implemented")
}
func (*UnimplementedDemoServicesServer) TryClientStreaming(srv DemoServices_TryClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method TryClientStreaming not implemented")
}
func (*UnimplementedDemoServicesServer) TryBidirectionalStreaming(srv DemoServices_TryBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method TryBidirectionalStreaming not implemented")
}

func RegisterDemoServicesServer(s *grpc.Server, srv DemoServicesServer) {
	s.RegisterService(&_DemoServices_serviceDesc, srv)
}

func _DemoServices_TryUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServicesServer).TryUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.v1.demoServices/TryUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServicesServer).TryUnary(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoServices_TryServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServicesServer).TryServerStreaming(m, &demoServicesTryServerStreamingServer{stream})
}

type DemoServices_TryServerStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type demoServicesTryServerStreamingServer struct {
	grpc.ServerStream
}

func (x *demoServicesTryServerStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DemoServices_TryClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServicesServer).TryClientStreaming(&demoServicesTryClientStreamingServer{stream})
}

type DemoServices_TryClientStreamingServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type demoServicesTryClientStreamingServer struct {
	grpc.ServerStream
}

func (x *demoServicesTryClientStreamingServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoServicesTryClientStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DemoServices_TryBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServicesServer).TryBidirectionalStreaming(&demoServicesTryBidirectionalStreamingServer{stream})
}

type DemoServices_TryBidirectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type demoServicesTryBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *demoServicesTryBidirectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoServicesTryBidirectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DemoServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.v1.demoServices",
	HandlerType: (*DemoServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TryUnary",
			Handler:    _DemoServices_TryUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TryServerStreaming",
			Handler:       _DemoServices_TryServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TryClientStreaming",
			Handler:       _DemoServices_TryClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TryBidirectionalStreaming",
			Handler:       _DemoServices_TryBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo.proto",
}
