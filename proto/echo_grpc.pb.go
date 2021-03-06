// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	StreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_StreamingEchoClient, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) StreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_StreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[0], "/EchoService/StreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceStreamingEchoClient{stream}
	return x, nil
}

type EchoService_StreamingEchoClient interface {
	Send(*StreamingEchoRequest) error
	Recv() (*StreamingEchoResponse, error)
	grpc.ClientStream
}

type echoServiceStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceStreamingEchoClient) Send(m *StreamingEchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceStreamingEchoClient) Recv() (*StreamingEchoResponse, error) {
	m := new(StreamingEchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	StreamingEcho(EchoService_StreamingEchoServer) error
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) StreamingEcho(EchoService_StreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingEcho not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_StreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).StreamingEcho(&echoServiceStreamingEchoServer{stream})
}

type EchoService_StreamingEchoServer interface {
	Send(*StreamingEchoResponse) error
	Recv() (*StreamingEchoRequest, error)
	grpc.ServerStream
}

type echoServiceStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceStreamingEchoServer) Send(m *StreamingEchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceStreamingEchoServer) Recv() (*StreamingEchoRequest, error) {
	m := new(StreamingEchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingEcho",
			Handler:       _EchoService_StreamingEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/echo.proto",
}
