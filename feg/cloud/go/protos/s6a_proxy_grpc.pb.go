// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: feg/protos/s6a_proxy.proto

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

// S6AProxyClient is the client API for S6AProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S6AProxyClient interface {
	// Authentication-Information (Code 318)
	AuthenticationInformation(ctx context.Context, in *AuthenticationInformationRequest, opts ...grpc.CallOption) (*AuthenticationInformationAnswer, error)
	// Update-Location (Code 316)
	UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationAnswer, error)
	// Purge-UE (Code 321)
	PurgeUE(ctx context.Context, in *PurgeUERequest, opts ...grpc.CallOption) (*PurgeUEAnswer, error)
}

type s6AProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewS6AProxyClient(cc grpc.ClientConnInterface) S6AProxyClient {
	return &s6AProxyClient{cc}
}

func (c *s6AProxyClient) AuthenticationInformation(ctx context.Context, in *AuthenticationInformationRequest, opts ...grpc.CallOption) (*AuthenticationInformationAnswer, error) {
	out := new(AuthenticationInformationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.S6aProxy/AuthenticationInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AProxyClient) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationAnswer, error) {
	out := new(UpdateLocationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.S6aProxy/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AProxyClient) PurgeUE(ctx context.Context, in *PurgeUERequest, opts ...grpc.CallOption) (*PurgeUEAnswer, error) {
	out := new(PurgeUEAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.S6aProxy/PurgeUE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S6AProxyServer is the server API for S6AProxy service.
// All implementations must embed UnimplementedS6AProxyServer
// for forward compatibility
type S6AProxyServer interface {
	// Authentication-Information (Code 318)
	AuthenticationInformation(context.Context, *AuthenticationInformationRequest) (*AuthenticationInformationAnswer, error)
	// Update-Location (Code 316)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationAnswer, error)
	// Purge-UE (Code 321)
	PurgeUE(context.Context, *PurgeUERequest) (*PurgeUEAnswer, error)
	mustEmbedUnimplementedS6AProxyServer()
}

// UnimplementedS6AProxyServer must be embedded to have forward compatible implementations.
type UnimplementedS6AProxyServer struct {
}

func (UnimplementedS6AProxyServer) AuthenticationInformation(context.Context, *AuthenticationInformationRequest) (*AuthenticationInformationAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticationInformation not implemented")
}
func (UnimplementedS6AProxyServer) UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedS6AProxyServer) PurgeUE(context.Context, *PurgeUERequest) (*PurgeUEAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeUE not implemented")
}
func (UnimplementedS6AProxyServer) mustEmbedUnimplementedS6AProxyServer() {}

// UnsafeS6AProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S6AProxyServer will
// result in compilation errors.
type UnsafeS6AProxyServer interface {
	mustEmbedUnimplementedS6AProxyServer()
}

func RegisterS6AProxyServer(s grpc.ServiceRegistrar, srv S6AProxyServer) {
	s.RegisterService(&S6AProxy_ServiceDesc, srv)
}

func _S6AProxy_AuthenticationInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).AuthenticationInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S6aProxy/AuthenticationInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).AuthenticationInformation(ctx, req.(*AuthenticationInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AProxy_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S6aProxy/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).UpdateLocation(ctx, req.(*UpdateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AProxy_PurgeUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeUERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).PurgeUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S6aProxy/PurgeUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).PurgeUE(ctx, req.(*PurgeUERequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S6AProxy_ServiceDesc is the grpc.ServiceDesc for S6AProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S6AProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.S6aProxy",
	HandlerType: (*S6AProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticationInformation",
			Handler:    _S6AProxy_AuthenticationInformation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _S6AProxy_UpdateLocation_Handler,
		},
		{
			MethodName: "PurgeUE",
			Handler:    _S6AProxy_PurgeUE_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/s6a_proxy.proto",
}

// S6AGatewayServiceClient is the client API for S6AGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S6AGatewayServiceClient interface {
	// Cancel-Location (Code 317)
	CancelLocation(ctx context.Context, in *CancelLocationRequest, opts ...grpc.CallOption) (*CancelLocationAnswer, error)
	// Reset (Code 322)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetAnswer, error)
}

type s6AGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewS6AGatewayServiceClient(cc grpc.ClientConnInterface) S6AGatewayServiceClient {
	return &s6AGatewayServiceClient{cc}
}

func (c *s6AGatewayServiceClient) CancelLocation(ctx context.Context, in *CancelLocationRequest, opts ...grpc.CallOption) (*CancelLocationAnswer, error) {
	out := new(CancelLocationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.S6aGatewayService/CancelLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AGatewayServiceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetAnswer, error) {
	out := new(ResetAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.S6aGatewayService/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S6AGatewayServiceServer is the server API for S6AGatewayService service.
// All implementations must embed UnimplementedS6AGatewayServiceServer
// for forward compatibility
type S6AGatewayServiceServer interface {
	// Cancel-Location (Code 317)
	CancelLocation(context.Context, *CancelLocationRequest) (*CancelLocationAnswer, error)
	// Reset (Code 322)
	Reset(context.Context, *ResetRequest) (*ResetAnswer, error)
	mustEmbedUnimplementedS6AGatewayServiceServer()
}

// UnimplementedS6AGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedS6AGatewayServiceServer struct {
}

func (UnimplementedS6AGatewayServiceServer) CancelLocation(context.Context, *CancelLocationRequest) (*CancelLocationAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLocation not implemented")
}
func (UnimplementedS6AGatewayServiceServer) Reset(context.Context, *ResetRequest) (*ResetAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedS6AGatewayServiceServer) mustEmbedUnimplementedS6AGatewayServiceServer() {}

// UnsafeS6AGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S6AGatewayServiceServer will
// result in compilation errors.
type UnsafeS6AGatewayServiceServer interface {
	mustEmbedUnimplementedS6AGatewayServiceServer()
}

func RegisterS6AGatewayServiceServer(s grpc.ServiceRegistrar, srv S6AGatewayServiceServer) {
	s.RegisterService(&S6AGatewayService_ServiceDesc, srv)
}

func _S6AGatewayService_CancelLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AGatewayServiceServer).CancelLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S6aGatewayService/CancelLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AGatewayServiceServer).CancelLocation(ctx, req.(*CancelLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AGatewayService_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AGatewayServiceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.S6aGatewayService/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AGatewayServiceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S6AGatewayService_ServiceDesc is the grpc.ServiceDesc for S6AGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S6AGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.S6aGatewayService",
	HandlerType: (*S6AGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelLocation",
			Handler:    _S6AGatewayService_CancelLocation_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _S6AGatewayService_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/s6a_proxy.proto",
}
