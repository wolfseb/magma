// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: dp/protos/enodebd_dp.proto

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

// DPServiceClient is the client API for DPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DPServiceClient interface {
	GetCBSDState(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDRegister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDDeregister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDRelinquish(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
}

type dPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDPServiceClient(cc grpc.ClientConnInterface) DPServiceClient {
	return &dPServiceClient{cc}
}

func (c *dPServiceClient) GetCBSDState(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/GetCBSDState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDRegister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDDeregister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDDeregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDRelinquish(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDRelinquish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DPServiceServer is the server API for DPService service.
// All implementations must embed UnimplementedDPServiceServer
// for forward compatibility
type DPServiceServer interface {
	GetCBSDState(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDRegister(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDDeregister(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDRelinquish(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	mustEmbedUnimplementedDPServiceServer()
}

// UnimplementedDPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDPServiceServer struct {
}

func (UnimplementedDPServiceServer) GetCBSDState(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCBSDState not implemented")
}
func (UnimplementedDPServiceServer) CBSDRegister(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDRegister not implemented")
}
func (UnimplementedDPServiceServer) CBSDDeregister(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDDeregister not implemented")
}
func (UnimplementedDPServiceServer) CBSDRelinquish(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDRelinquish not implemented")
}
func (UnimplementedDPServiceServer) mustEmbedUnimplementedDPServiceServer() {}

// UnsafeDPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DPServiceServer will
// result in compilation errors.
type UnsafeDPServiceServer interface {
	mustEmbedUnimplementedDPServiceServer()
}

func RegisterDPServiceServer(s grpc.ServiceRegistrar, srv DPServiceServer) {
	s.RegisterService(&DPService_ServiceDesc, srv)
}

func _DPService_GetCBSDState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).GetCBSDState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/GetCBSDState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).GetCBSDState(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDRegister(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDDeregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDDeregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDDeregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDDeregister(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDRelinquish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDRelinquish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDRelinquish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDRelinquish(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DPService_ServiceDesc is the grpc.ServiceDesc for DPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DPService",
	HandlerType: (*DPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCBSDState",
			Handler:    _DPService_GetCBSDState_Handler,
		},
		{
			MethodName: "CBSDRegister",
			Handler:    _DPService_CBSDRegister_Handler,
		},
		{
			MethodName: "CBSDDeregister",
			Handler:    _DPService_CBSDDeregister_Handler,
		},
		{
			MethodName: "CBSDRelinquish",
			Handler:    _DPService_CBSDRelinquish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dp/protos/enodebd_dp.proto",
}
