// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: feg/protos/s8_proxy.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// S8ProxyClient is the client API for S8Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S8ProxyClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequestPgw, opts ...grpc.CallOption) (*CreateSessionResponsePgw, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequestPgw, opts ...grpc.CallOption) (*DeleteSessionResponsePgw, error)
	SendEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	CreateBearerResponse(ctx context.Context, in *CreateBearerResponsePgw, opts ...grpc.CallOption) (*protos.Void, error)
	DeleteBearerResponse(ctx context.Context, in *DeleteBearerResponsePgw, opts ...grpc.CallOption) (*protos.Void, error)
}

type s8ProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewS8ProxyClient(cc grpc.ClientConnInterface) S8ProxyClient {
	return &s8ProxyClient{cc}
}

func (c *s8ProxyClient) CreateSession(ctx context.Context, in *CreateSessionRequestPgw, opts ...grpc.CallOption) (*CreateSessionResponsePgw, error) {
	out := new(CreateSessionResponsePgw)
	err := c.cc.Invoke(ctx, "/magma.feg.S8Proxy/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s8ProxyClient) DeleteSession(ctx context.Context, in *DeleteSessionRequestPgw, opts ...grpc.CallOption) (*DeleteSessionResponsePgw, error) {
	out := new(DeleteSessionResponsePgw)
	err := c.cc.Invoke(ctx, "/magma.feg.S8Proxy/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s8ProxyClient) SendEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/magma.feg.S8Proxy/SendEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s8ProxyClient) CreateBearerResponse(ctx context.Context, in *CreateBearerResponsePgw, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.S8Proxy/CreateBearerResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s8ProxyClient) DeleteBearerResponse(ctx context.Context, in *DeleteBearerResponsePgw, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.S8Proxy/DeleteBearerResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S8ProxyServer is the server API for S8Proxy service.
// All implementations must embed UnimplementedS8ProxyServer
// for forward compatibility
type S8ProxyServer interface {
	CreateSession(context.Context, *CreateSessionRequestPgw) (*CreateSessionResponsePgw, error)
	DeleteSession(context.Context, *DeleteSessionRequestPgw) (*DeleteSessionResponsePgw, error)
	SendEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	CreateBearerResponse(context.Context, *CreateBearerResponsePgw) (*protos.Void, error)
	DeleteBearerResponse(context.Context, *DeleteBearerResponsePgw) (*protos.Void, error)
	mustEmbedUnimplementedS8ProxyServer()
}

// UnimplementedS8ProxyServer must be embedded to have forward compatible implementations.
type UnimplementedS8ProxyServer struct {
}

func (UnimplementedS8ProxyServer) CreateSession(context.Context, *CreateSessionRequestPgw) (*CreateSessionResponsePgw, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedS8ProxyServer) DeleteSession(context.Context, *DeleteSessionRequestPgw) (*DeleteSessionResponsePgw, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedS8ProxyServer) SendEcho(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEcho not implemented")
}
func (UnimplementedS8ProxyServer) CreateBearerResponse(context.Context, *CreateBearerResponsePgw) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBearerResponse not implemented")
}
func (UnimplementedS8ProxyServer) DeleteBearerResponse(context.Context, *DeleteBearerResponsePgw) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBearerResponse not implemented")
}
func (UnimplementedS8ProxyServer) mustEmbedUnimplementedS8ProxyServer() {}

// UnsafeS8ProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S8ProxyServer will
// result in compilation errors.
type UnsafeS8ProxyServer interface {
	mustEmbedUnimplementedS8ProxyServer()
}

func RegisterS8ProxyServer(s grpc.ServiceRegistrar, srv S8ProxyServer) {
	s.RegisterService(&S8Proxy_ServiceDesc, srv)
}

func _S8Proxy_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequestPgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8Proxy/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyServer).CreateSession(ctx, req.(*CreateSessionRequestPgw))
	}
	return interceptor(ctx, in, info, handler)
}

func _S8Proxy_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequestPgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8Proxy/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyServer).DeleteSession(ctx, req.(*DeleteSessionRequestPgw))
	}
	return interceptor(ctx, in, info, handler)
}

func _S8Proxy_SendEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyServer).SendEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8Proxy/SendEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyServer).SendEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S8Proxy_CreateBearerResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBearerResponsePgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyServer).CreateBearerResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8Proxy/CreateBearerResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyServer).CreateBearerResponse(ctx, req.(*CreateBearerResponsePgw))
	}
	return interceptor(ctx, in, info, handler)
}

func _S8Proxy_DeleteBearerResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBearerResponsePgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyServer).DeleteBearerResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8Proxy/DeleteBearerResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyServer).DeleteBearerResponse(ctx, req.(*DeleteBearerResponsePgw))
	}
	return interceptor(ctx, in, info, handler)
}

// S8Proxy_ServiceDesc is the grpc.ServiceDesc for S8Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S8Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.S8Proxy",
	HandlerType: (*S8ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _S8Proxy_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _S8Proxy_DeleteSession_Handler,
		},
		{
			MethodName: "SendEcho",
			Handler:    _S8Proxy_SendEcho_Handler,
		},
		{
			MethodName: "CreateBearerResponse",
			Handler:    _S8Proxy_CreateBearerResponse_Handler,
		},
		{
			MethodName: "DeleteBearerResponse",
			Handler:    _S8Proxy_DeleteBearerResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/s8_proxy.proto",
}

// S8ProxyResponderClient is the client API for S8ProxyResponder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S8ProxyResponderClient interface {
	CreateBearer(ctx context.Context, in *CreateBearerRequestPgw, opts ...grpc.CallOption) (*protos.Void, error)
	DeleteBearerRequest(ctx context.Context, in *DeleteBearerRequestPgw, opts ...grpc.CallOption) (*protos.Void, error)
}

type s8ProxyResponderClient struct {
	cc grpc.ClientConnInterface
}

func NewS8ProxyResponderClient(cc grpc.ClientConnInterface) S8ProxyResponderClient {
	return &s8ProxyResponderClient{cc}
}

func (c *s8ProxyResponderClient) CreateBearer(ctx context.Context, in *CreateBearerRequestPgw, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.S8ProxyResponder/CreateBearer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s8ProxyResponderClient) DeleteBearerRequest(ctx context.Context, in *DeleteBearerRequestPgw, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.S8ProxyResponder/DeleteBearerRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S8ProxyResponderServer is the server API for S8ProxyResponder service.
// All implementations must embed UnimplementedS8ProxyResponderServer
// for forward compatibility
type S8ProxyResponderServer interface {
	CreateBearer(context.Context, *CreateBearerRequestPgw) (*protos.Void, error)
	DeleteBearerRequest(context.Context, *DeleteBearerRequestPgw) (*protos.Void, error)
	mustEmbedUnimplementedS8ProxyResponderServer()
}

// UnimplementedS8ProxyResponderServer must be embedded to have forward compatible implementations.
type UnimplementedS8ProxyResponderServer struct {
}

func (UnimplementedS8ProxyResponderServer) CreateBearer(context.Context, *CreateBearerRequestPgw) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBearer not implemented")
}
func (UnimplementedS8ProxyResponderServer) DeleteBearerRequest(context.Context, *DeleteBearerRequestPgw) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBearerRequest not implemented")
}
func (UnimplementedS8ProxyResponderServer) mustEmbedUnimplementedS8ProxyResponderServer() {}

// UnsafeS8ProxyResponderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S8ProxyResponderServer will
// result in compilation errors.
type UnsafeS8ProxyResponderServer interface {
	mustEmbedUnimplementedS8ProxyResponderServer()
}

func RegisterS8ProxyResponderServer(s grpc.ServiceRegistrar, srv S8ProxyResponderServer) {
	s.RegisterService(&S8ProxyResponder_ServiceDesc, srv)
}

func _S8ProxyResponder_CreateBearer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBearerRequestPgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyResponderServer).CreateBearer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8ProxyResponder/CreateBearer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyResponderServer).CreateBearer(ctx, req.(*CreateBearerRequestPgw))
	}
	return interceptor(ctx, in, info, handler)
}

func _S8ProxyResponder_DeleteBearerRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBearerRequestPgw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S8ProxyResponderServer).DeleteBearerRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S8ProxyResponder/DeleteBearerRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S8ProxyResponderServer).DeleteBearerRequest(ctx, req.(*DeleteBearerRequestPgw))
	}
	return interceptor(ctx, in, info, handler)
}

// S8ProxyResponder_ServiceDesc is the grpc.ServiceDesc for S8ProxyResponder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S8ProxyResponder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.S8ProxyResponder",
	HandlerType: (*S8ProxyResponderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBearer",
			Handler:    _S8ProxyResponder_CreateBearer_Handler,
		},
		{
			MethodName: "DeleteBearerRequest",
			Handler:    _S8ProxyResponder_DeleteBearerRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/s8_proxy.proto",
}
