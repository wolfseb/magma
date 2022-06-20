// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: orc8r/cloud/go/services/accessd/protos/access.proto

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

// AccessControlManagerClient is the client API for AccessControlManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessControlManagerClient interface {
	// Overwrites Permissions for operator Identity to manage others
	// Request includes ACL to set for the Operator
	// If the Operator doesn't exist - creates a new operator with the given ACL
	SetOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Adds Permissions for one Identity to manage others
	// Request includes ACL to add (append to the existing ACL) for the Operator
	UpdateOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Removes all operator's permissions (the entire operator's ACL)
	DeleteOperator(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error)
	// Returns the managing Identity's permissions list
	GetOperatorACL(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*AccessControl_List, error)
	// Returns the managing Identity's permissions list
	GetOperatorsACLs(ctx context.Context, in *protos.Identity_List, opts ...grpc.CallOption) (*AccessControl_Lists, error)
	// Returns the managing Identity's permissions for a given entity
	// NOTE: Takes into account wildcards for the entity's type in the ACL
	GetPermissions(ctx context.Context, in *AccessControl_PermissionsRequest, opts ...grpc.CallOption) (*AccessControl_Entity, error)
	// CheckPermissions verifies Operator permissions for a list of given
	// Identities. AccessControl.ListRequest.entities is a list of
	// Identities and their corresponding permissions requested by the operator
	// CheckPermissions will return success only if all requested permissions
	// are satisfied (AND logic)
	// Intended to be used for multi-Identity requests such as Network Identity
	// AND REST API Identity, etc.
	CheckPermissions(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Lists all globally registered operators on the cloud
	ListOperators(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Identity_List, error)
	// Cleanup a given entity from all Operators' ACLs
	DeleteEntity(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error)
}

type accessControlManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessControlManagerClient(cc grpc.ClientConnInterface) AccessControlManagerClient {
	return &accessControlManagerClient{cc}
}

func (c *accessControlManagerClient) SetOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/SetOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) UpdateOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/UpdateOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) DeleteOperator(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/DeleteOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetOperatorACL(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*AccessControl_List, error) {
	out := new(AccessControl_List)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetOperatorACL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetOperatorsACLs(ctx context.Context, in *protos.Identity_List, opts ...grpc.CallOption) (*AccessControl_Lists, error) {
	out := new(AccessControl_Lists)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetOperatorsACLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetPermissions(ctx context.Context, in *AccessControl_PermissionsRequest, opts ...grpc.CallOption) (*AccessControl_Entity, error) {
	out := new(AccessControl_Entity)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) CheckPermissions(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/CheckPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) ListOperators(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Identity_List, error) {
	out := new(protos.Identity_List)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/ListOperators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) DeleteEntity(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessControlManagerServer is the server API for AccessControlManager service.
// All implementations must embed UnimplementedAccessControlManagerServer
// for forward compatibility
type AccessControlManagerServer interface {
	// Overwrites Permissions for operator Identity to manage others
	// Request includes ACL to set for the Operator
	// If the Operator doesn't exist - creates a new operator with the given ACL
	SetOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Adds Permissions for one Identity to manage others
	// Request includes ACL to add (append to the existing ACL) for the Operator
	UpdateOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Removes all operator's permissions (the entire operator's ACL)
	DeleteOperator(context.Context, *protos.Identity) (*protos.Void, error)
	// Returns the managing Identity's permissions list
	GetOperatorACL(context.Context, *protos.Identity) (*AccessControl_List, error)
	// Returns the managing Identity's permissions list
	GetOperatorsACLs(context.Context, *protos.Identity_List) (*AccessControl_Lists, error)
	// Returns the managing Identity's permissions for a given entity
	// NOTE: Takes into account wildcards for the entity's type in the ACL
	GetPermissions(context.Context, *AccessControl_PermissionsRequest) (*AccessControl_Entity, error)
	// CheckPermissions verifies Operator permissions for a list of given
	// Identities. AccessControl.ListRequest.entities is a list of
	// Identities and their corresponding permissions requested by the operator
	// CheckPermissions will return success only if all requested permissions
	// are satisfied (AND logic)
	// Intended to be used for multi-Identity requests such as Network Identity
	// AND REST API Identity, etc.
	CheckPermissions(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Lists all globally registered operators on the cloud
	ListOperators(context.Context, *protos.Void) (*protos.Identity_List, error)
	// Cleanup a given entity from all Operators' ACLs
	DeleteEntity(context.Context, *protos.Identity) (*protos.Void, error)
	mustEmbedUnimplementedAccessControlManagerServer()
}

// UnimplementedAccessControlManagerServer must be embedded to have forward compatible implementations.
type UnimplementedAccessControlManagerServer struct {
}

func (UnimplementedAccessControlManagerServer) SetOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOperator not implemented")
}
func (UnimplementedAccessControlManagerServer) UpdateOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOperator not implemented")
}
func (UnimplementedAccessControlManagerServer) DeleteOperator(context.Context, *protos.Identity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperator not implemented")
}
func (UnimplementedAccessControlManagerServer) GetOperatorACL(context.Context, *protos.Identity) (*AccessControl_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorACL not implemented")
}
func (UnimplementedAccessControlManagerServer) GetOperatorsACLs(context.Context, *protos.Identity_List) (*AccessControl_Lists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorsACLs not implemented")
}
func (UnimplementedAccessControlManagerServer) GetPermissions(context.Context, *AccessControl_PermissionsRequest) (*AccessControl_Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedAccessControlManagerServer) CheckPermissions(context.Context, *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermissions not implemented")
}
func (UnimplementedAccessControlManagerServer) ListOperators(context.Context, *protos.Void) (*protos.Identity_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperators not implemented")
}
func (UnimplementedAccessControlManagerServer) DeleteEntity(context.Context, *protos.Identity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}
func (UnimplementedAccessControlManagerServer) mustEmbedUnimplementedAccessControlManagerServer() {}

// UnsafeAccessControlManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessControlManagerServer will
// result in compilation errors.
type UnsafeAccessControlManagerServer interface {
	mustEmbedUnimplementedAccessControlManagerServer()
}

func RegisterAccessControlManagerServer(s grpc.ServiceRegistrar, srv AccessControlManagerServer) {
	s.RegisterService(&AccessControlManager_ServiceDesc, srv)
}

func _AccessControlManager_SetOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).SetOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/SetOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).SetOperator(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_UpdateOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).UpdateOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/UpdateOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).UpdateOperator(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_DeleteOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).DeleteOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/DeleteOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).DeleteOperator(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetOperatorACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetOperatorACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetOperatorACL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetOperatorACL(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetOperatorsACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity_List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetOperatorsACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetOperatorsACLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetOperatorsACLs(ctx, req.(*protos.Identity_List))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_PermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetPermissions(ctx, req.(*AccessControl_PermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_CheckPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).CheckPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/CheckPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).CheckPermissions(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_ListOperators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).ListOperators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/ListOperators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).ListOperators(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).DeleteEntity(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessControlManager_ServiceDesc is the grpc.ServiceDesc for AccessControlManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessControlManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.accessd.AccessControlManager",
	HandlerType: (*AccessControlManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOperator",
			Handler:    _AccessControlManager_SetOperator_Handler,
		},
		{
			MethodName: "UpdateOperator",
			Handler:    _AccessControlManager_UpdateOperator_Handler,
		},
		{
			MethodName: "DeleteOperator",
			Handler:    _AccessControlManager_DeleteOperator_Handler,
		},
		{
			MethodName: "GetOperatorACL",
			Handler:    _AccessControlManager_GetOperatorACL_Handler,
		},
		{
			MethodName: "GetOperatorsACLs",
			Handler:    _AccessControlManager_GetOperatorsACLs_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _AccessControlManager_GetPermissions_Handler,
		},
		{
			MethodName: "CheckPermissions",
			Handler:    _AccessControlManager_CheckPermissions_Handler,
		},
		{
			MethodName: "ListOperators",
			Handler:    _AccessControlManager_ListOperators_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _AccessControlManager_DeleteEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/accessd/protos/access.proto",
}
