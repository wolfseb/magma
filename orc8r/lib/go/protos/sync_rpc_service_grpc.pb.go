// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: orc8r/protos/sync_rpc_service.proto

package protos

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

// SyncRPCServiceClient is the client API for SyncRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncRPCServiceClient interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error)
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error)
}

type syncRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncRPCServiceClient(cc grpc.ClientConnInterface) SyncRPCServiceClient {
	return &syncRPCServiceClient{cc}
}

func (c *syncRPCServiceClient) EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyncRPCService_ServiceDesc.Streams[0], "/magma.orc8r.SyncRPCService/EstablishSyncRPCStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceEstablishSyncRPCStreamClient{stream}
	return x, nil
}

type SyncRPCService_EstablishSyncRPCStreamClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceEstablishSyncRPCStreamClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncRPCServiceClient) SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyncRPCService_ServiceDesc.Streams[1], "/magma.orc8r.SyncRPCService/SyncRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceSyncRPCClient{stream}
	return x, nil
}

type SyncRPCService_SyncRPCClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceSyncRPCClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceSyncRPCClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncRPCServiceServer is the server API for SyncRPCService service.
// All implementations must embed UnimplementedSyncRPCServiceServer
// for forward compatibility
type SyncRPCServiceServer interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(SyncRPCService_EstablishSyncRPCStreamServer) error
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(SyncRPCService_SyncRPCServer) error
	mustEmbedUnimplementedSyncRPCServiceServer()
}

// UnimplementedSyncRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncRPCServiceServer struct {
}

func (UnimplementedSyncRPCServiceServer) EstablishSyncRPCStream(SyncRPCService_EstablishSyncRPCStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EstablishSyncRPCStream not implemented")
}
func (UnimplementedSyncRPCServiceServer) SyncRPC(SyncRPCService_SyncRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncRPC not implemented")
}
func (UnimplementedSyncRPCServiceServer) mustEmbedUnimplementedSyncRPCServiceServer() {}

// UnsafeSyncRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncRPCServiceServer will
// result in compilation errors.
type UnsafeSyncRPCServiceServer interface {
	mustEmbedUnimplementedSyncRPCServiceServer()
}

func RegisterSyncRPCServiceServer(s grpc.ServiceRegistrar, srv SyncRPCServiceServer) {
	s.RegisterService(&SyncRPCService_ServiceDesc, srv)
}

func _SyncRPCService_EstablishSyncRPCStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).EstablishSyncRPCStream(&syncRPCServiceEstablishSyncRPCStreamServer{stream})
}

type SyncRPCService_EstablishSyncRPCStreamServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceEstablishSyncRPCStreamServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SyncRPCService_SyncRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).SyncRPC(&syncRPCServiceSyncRPCServer{stream})
}

type SyncRPCService_SyncRPCServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceSyncRPCServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceSyncRPCServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncRPCService_ServiceDesc is the grpc.ServiceDesc for SyncRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.SyncRPCService",
	HandlerType: (*SyncRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishSyncRPCStream",
			Handler:       _SyncRPCService_EstablishSyncRPCStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SyncRPC",
			Handler:       _SyncRPCService_SyncRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orc8r/protos/sync_rpc_service.proto",
}
