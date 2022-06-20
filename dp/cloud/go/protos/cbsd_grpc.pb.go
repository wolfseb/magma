// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CbsdManagementClient is the client API for CbsdManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CbsdManagementClient interface {
	CreateCbsd(ctx context.Context, in *CreateCbsdRequest, opts ...grpc.CallOption) (*CreateCbsdResponse, error)
	UserUpdateCbsd(ctx context.Context, in *UpdateCbsdRequest, opts ...grpc.CallOption) (*UpdateCbsdResponse, error)
	EnodebdUpdateCbsd(ctx context.Context, in *EnodebdUpdateCbsdRequest, opts ...grpc.CallOption) (*UpdateCbsdResponse, error)
	DeleteCbsd(ctx context.Context, in *DeleteCbsdRequest, opts ...grpc.CallOption) (*DeleteCbsdResponse, error)
	FetchCbsd(ctx context.Context, in *FetchCbsdRequest, opts ...grpc.CallOption) (*FetchCbsdResponse, error)
	ListCbsds(ctx context.Context, in *ListCbsdRequest, opts ...grpc.CallOption) (*ListCbsdResponse, error)
	DeregisterCbsd(ctx context.Context, in *DeregisterCbsdRequest, opts ...grpc.CallOption) (*DeregisterCbsdResponse, error)
}

type cbsdManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCbsdManagementClient(cc grpc.ClientConnInterface) CbsdManagementClient {
	return &cbsdManagementClient{cc}
}

func (c *cbsdManagementClient) CreateCbsd(ctx context.Context, in *CreateCbsdRequest, opts ...grpc.CallOption) (*CreateCbsdResponse, error) {
	out := new(CreateCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/CreateCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) UserUpdateCbsd(ctx context.Context, in *UpdateCbsdRequest, opts ...grpc.CallOption) (*UpdateCbsdResponse, error) {
	out := new(UpdateCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/UserUpdateCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) EnodebdUpdateCbsd(ctx context.Context, in *EnodebdUpdateCbsdRequest, opts ...grpc.CallOption) (*UpdateCbsdResponse, error) {
	out := new(UpdateCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/EnodebdUpdateCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) DeleteCbsd(ctx context.Context, in *DeleteCbsdRequest, opts ...grpc.CallOption) (*DeleteCbsdResponse, error) {
	out := new(DeleteCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/DeleteCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) FetchCbsd(ctx context.Context, in *FetchCbsdRequest, opts ...grpc.CallOption) (*FetchCbsdResponse, error) {
	out := new(FetchCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/FetchCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) ListCbsds(ctx context.Context, in *ListCbsdRequest, opts ...grpc.CallOption) (*ListCbsdResponse, error) {
	out := new(ListCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/ListCbsds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbsdManagementClient) DeregisterCbsd(ctx context.Context, in *DeregisterCbsdRequest, opts ...grpc.CallOption) (*DeregisterCbsdResponse, error) {
	out := new(DeregisterCbsdResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.CbsdManagement/DeregisterCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CbsdManagementServer is the server API for CbsdManagement service.
// All implementations should embed UnimplementedCbsdManagementServer
// for forward compatibility
type CbsdManagementServer interface {
	CreateCbsd(context.Context, *CreateCbsdRequest) (*CreateCbsdResponse, error)
	UserUpdateCbsd(context.Context, *UpdateCbsdRequest) (*UpdateCbsdResponse, error)
	EnodebdUpdateCbsd(context.Context, *EnodebdUpdateCbsdRequest) (*UpdateCbsdResponse, error)
	DeleteCbsd(context.Context, *DeleteCbsdRequest) (*DeleteCbsdResponse, error)
	FetchCbsd(context.Context, *FetchCbsdRequest) (*FetchCbsdResponse, error)
	ListCbsds(context.Context, *ListCbsdRequest) (*ListCbsdResponse, error)
	DeregisterCbsd(context.Context, *DeregisterCbsdRequest) (*DeregisterCbsdResponse, error)
}

// UnimplementedCbsdManagementServer should be embedded to have forward compatible implementations.
type UnimplementedCbsdManagementServer struct {
}

func (UnimplementedCbsdManagementServer) CreateCbsd(context.Context, *CreateCbsdRequest) (*CreateCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCbsd not implemented")
}
func (UnimplementedCbsdManagementServer) UserUpdateCbsd(context.Context, *UpdateCbsdRequest) (*UpdateCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateCbsd not implemented")
}
func (UnimplementedCbsdManagementServer) EnodebdUpdateCbsd(context.Context, *EnodebdUpdateCbsdRequest) (*UpdateCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnodebdUpdateCbsd not implemented")
}
func (UnimplementedCbsdManagementServer) DeleteCbsd(context.Context, *DeleteCbsdRequest) (*DeleteCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCbsd not implemented")
}
func (UnimplementedCbsdManagementServer) FetchCbsd(context.Context, *FetchCbsdRequest) (*FetchCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCbsd not implemented")
}
func (UnimplementedCbsdManagementServer) ListCbsds(context.Context, *ListCbsdRequest) (*ListCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCbsds not implemented")
}
func (UnimplementedCbsdManagementServer) DeregisterCbsd(context.Context, *DeregisterCbsdRequest) (*DeregisterCbsdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterCbsd not implemented")
}

// UnsafeCbsdManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CbsdManagementServer will
// result in compilation errors.
type UnsafeCbsdManagementServer interface {
	mustEmbedUnimplementedCbsdManagementServer()
}

func RegisterCbsdManagementServer(s grpc.ServiceRegistrar, srv CbsdManagementServer) {
	s.RegisterService(&CbsdManagement_ServiceDesc, srv)
}

func _CbsdManagement_CreateCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).CreateCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/CreateCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).CreateCbsd(ctx, req.(*CreateCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_UserUpdateCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).UserUpdateCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/UserUpdateCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).UserUpdateCbsd(ctx, req.(*UpdateCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_EnodebdUpdateCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebdUpdateCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).EnodebdUpdateCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/EnodebdUpdateCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).EnodebdUpdateCbsd(ctx, req.(*EnodebdUpdateCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_DeleteCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).DeleteCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/DeleteCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).DeleteCbsd(ctx, req.(*DeleteCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_FetchCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).FetchCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/FetchCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).FetchCbsd(ctx, req.(*FetchCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_ListCbsds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).ListCbsds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/ListCbsds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).ListCbsds(ctx, req.(*ListCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbsdManagement_DeregisterCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbsdManagementServer).DeregisterCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.CbsdManagement/DeregisterCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbsdManagementServer).DeregisterCbsd(ctx, req.(*DeregisterCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CbsdManagement_ServiceDesc is the grpc.ServiceDesc for CbsdManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CbsdManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.dp.CbsdManagement",
	HandlerType: (*CbsdManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCbsd",
			Handler:    _CbsdManagement_CreateCbsd_Handler,
		},
		{
			MethodName: "UserUpdateCbsd",
			Handler:    _CbsdManagement_UserUpdateCbsd_Handler,
		},
		{
			MethodName: "EnodebdUpdateCbsd",
			Handler:    _CbsdManagement_EnodebdUpdateCbsd_Handler,
		},
		{
			MethodName: "DeleteCbsd",
			Handler:    _CbsdManagement_DeleteCbsd_Handler,
		},
		{
			MethodName: "FetchCbsd",
			Handler:    _CbsdManagement_FetchCbsd_Handler,
		},
		{
			MethodName: "ListCbsds",
			Handler:    _CbsdManagement_ListCbsds_Handler,
		},
		{
			MethodName: "DeregisterCbsd",
			Handler:    _CbsdManagement_DeregisterCbsd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dp/protos/cbsd.proto",
}
