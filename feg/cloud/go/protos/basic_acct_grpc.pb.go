// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: feg/protos/basic_acct.proto

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

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingClient interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error) {
	out := new(AcctStopResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
// All implementations must embed UnimplementedAccountingServer
// for forward compatibility
type AccountingServer interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(context.Context, *AcctSession) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(context.Context, *AcctUpdateReq) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(context.Context, *AcctUpdateReq) (*AcctStopResp, error)
	mustEmbedUnimplementedAccountingServer()
}

// UnimplementedAccountingServer must be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (UnimplementedAccountingServer) Start(context.Context, *AcctSession) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedAccountingServer) Update(context.Context, *AcctUpdateReq) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccountingServer) Stop(context.Context, *AcctUpdateReq) (*AcctStopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedAccountingServer) mustEmbedUnimplementedAccountingServer() {}

// UnsafeAccountingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingServer will
// result in compilation errors.
type UnsafeAccountingServer interface {
	mustEmbedUnimplementedAccountingServer()
}

func RegisterAccountingServer(s grpc.ServiceRegistrar, srv AccountingServer) {
	s.RegisterService(&Accounting_ServiceDesc, srv)
}

func _Accounting_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Start(ctx, req.(*AcctSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Update(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Stop(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounting_ServiceDesc is the grpc.ServiceDesc for Accounting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Accounting_Start_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Accounting_Update_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Accounting_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/basic_acct.proto",
}
