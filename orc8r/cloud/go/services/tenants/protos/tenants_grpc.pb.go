// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// TenantsServiceClient is the client API for TenantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantsServiceClient interface {
	GetAllTenants(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*TenantList, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error)
	SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error)
	DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*protos.Void, error)
	GetControlProxy(ctx context.Context, in *GetControlProxyRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error)
	GetControlProxyFromNetworkID(ctx context.Context, in *GetControlProxyFromNetworkIDRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error)
	CreateOrUpdateControlProxy(ctx context.Context, in *CreateOrUpdateControlProxyRequest, opts ...grpc.CallOption) (*CreateOrUpdateControlProxyResponse, error)
}

type tenantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantsServiceClient(cc grpc.ClientConnInterface) TenantsServiceClient {
	return &tenantsServiceClient{cc}
}

func (c *tenantsServiceClient) GetAllTenants(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetAllTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/SetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetControlProxy(ctx context.Context, in *GetControlProxyRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error) {
	out := new(GetControlProxyResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetControlProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetControlProxyFromNetworkID(ctx context.Context, in *GetControlProxyFromNetworkIDRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error) {
	out := new(GetControlProxyResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetControlProxyFromNetworkID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) CreateOrUpdateControlProxy(ctx context.Context, in *CreateOrUpdateControlProxyRequest, opts ...grpc.CallOption) (*CreateOrUpdateControlProxyResponse, error) {
	out := new(CreateOrUpdateControlProxyResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/CreateOrUpdateControlProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantsServiceServer is the server API for TenantsService service.
// All implementations should embed UnimplementedTenantsServiceServer
// for forward compatibility
type TenantsServiceServer interface {
	GetAllTenants(context.Context, *protos.Void) (*TenantList, error)
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	CreateTenant(context.Context, *IDAndTenant) (*protos.Void, error)
	SetTenant(context.Context, *IDAndTenant) (*protos.Void, error)
	DeleteTenant(context.Context, *GetTenantRequest) (*protos.Void, error)
	GetControlProxy(context.Context, *GetControlProxyRequest) (*GetControlProxyResponse, error)
	GetControlProxyFromNetworkID(context.Context, *GetControlProxyFromNetworkIDRequest) (*GetControlProxyResponse, error)
	CreateOrUpdateControlProxy(context.Context, *CreateOrUpdateControlProxyRequest) (*CreateOrUpdateControlProxyResponse, error)
}

// UnimplementedTenantsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantsServiceServer struct {
}

func (UnimplementedTenantsServiceServer) GetAllTenants(context.Context, *protos.Void) (*TenantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTenants not implemented")
}
func (UnimplementedTenantsServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantsServiceServer) CreateTenant(context.Context, *IDAndTenant) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantsServiceServer) SetTenant(context.Context, *IDAndTenant) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTenant not implemented")
}
func (UnimplementedTenantsServiceServer) DeleteTenant(context.Context, *GetTenantRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantsServiceServer) GetControlProxy(context.Context, *GetControlProxyRequest) (*GetControlProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControlProxy not implemented")
}
func (UnimplementedTenantsServiceServer) GetControlProxyFromNetworkID(context.Context, *GetControlProxyFromNetworkIDRequest) (*GetControlProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControlProxyFromNetworkID not implemented")
}
func (UnimplementedTenantsServiceServer) CreateOrUpdateControlProxy(context.Context, *CreateOrUpdateControlProxyRequest) (*CreateOrUpdateControlProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateControlProxy not implemented")
}

// UnsafeTenantsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantsServiceServer will
// result in compilation errors.
type UnsafeTenantsServiceServer interface {
	mustEmbedUnimplementedTenantsServiceServer()
}

func RegisterTenantsServiceServer(s grpc.ServiceRegistrar, srv TenantsServiceServer) {
	s.RegisterService(&TenantsService_ServiceDesc, srv)
}

func _TenantsService_GetAllTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetAllTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_SetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).SetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/SetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).SetTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetControlProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetControlProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetControlProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetControlProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetControlProxy(ctx, req.(*GetControlProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetControlProxyFromNetworkID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetControlProxyFromNetworkIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetControlProxyFromNetworkID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetControlProxyFromNetworkID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetControlProxyFromNetworkID(ctx, req.(*GetControlProxyFromNetworkIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_CreateOrUpdateControlProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateControlProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateOrUpdateControlProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/CreateOrUpdateControlProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateOrUpdateControlProxy(ctx, req.(*CreateOrUpdateControlProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantsService_ServiceDesc is the grpc.ServiceDesc for TenantsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.TenantsService",
	HandlerType: (*TenantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTenants",
			Handler:    _TenantsService_GetAllTenants_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantsService_GetTenant_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _TenantsService_CreateTenant_Handler,
		},
		{
			MethodName: "SetTenant",
			Handler:    _TenantsService_SetTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantsService_DeleteTenant_Handler,
		},
		{
			MethodName: "GetControlProxy",
			Handler:    _TenantsService_GetControlProxy_Handler,
		},
		{
			MethodName: "GetControlProxyFromNetworkID",
			Handler:    _TenantsService_GetControlProxyFromNetworkID_Handler,
		},
		{
			MethodName: "CreateOrUpdateControlProxy",
			Handler:    _TenantsService_CreateOrUpdateControlProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/tenants/protos/tenants.proto",
}
