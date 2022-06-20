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

// SctpdDownlinkClient is the client API for SctpdDownlink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SctpdDownlinkClient interface {
	// Init - initialize sctp connection according to InitReq
	// @param InitReq request specifying desired sctp configuration
	// @return InitRes response w/ init success status
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitRes, error)
	// SendDl - send a downlink packet to eNB
	// @param SendDlReq request specifying packet data and destination
	// @return SendDlRes response w/ send success status
	SendDl(ctx context.Context, in *SendDlReq, opts ...grpc.CallOption) (*SendDlRes, error)
}

type sctpdDownlinkClient struct {
	cc grpc.ClientConnInterface
}

func NewSctpdDownlinkClient(cc grpc.ClientConnInterface) SctpdDownlinkClient {
	return &sctpdDownlinkClient{cc}
}

func (c *sctpdDownlinkClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitRes, error) {
	out := new(InitRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdDownlink/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdDownlinkClient) SendDl(ctx context.Context, in *SendDlReq, opts ...grpc.CallOption) (*SendDlRes, error) {
	out := new(SendDlRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdDownlink/SendDl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SctpdDownlinkServer is the server API for SctpdDownlink service.
// All implementations should embed UnimplementedSctpdDownlinkServer
// for forward compatibility
type SctpdDownlinkServer interface {
	// Init - initialize sctp connection according to InitReq
	// @param InitReq request specifying desired sctp configuration
	// @return InitRes response w/ init success status
	Init(context.Context, *InitReq) (*InitRes, error)
	// SendDl - send a downlink packet to eNB
	// @param SendDlReq request specifying packet data and destination
	// @return SendDlRes response w/ send success status
	SendDl(context.Context, *SendDlReq) (*SendDlRes, error)
}

// UnimplementedSctpdDownlinkServer should be embedded to have forward compatible implementations.
type UnimplementedSctpdDownlinkServer struct {
}

func (UnimplementedSctpdDownlinkServer) Init(context.Context, *InitReq) (*InitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedSctpdDownlinkServer) SendDl(context.Context, *SendDlReq) (*SendDlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDl not implemented")
}

// UnsafeSctpdDownlinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SctpdDownlinkServer will
// result in compilation errors.
type UnsafeSctpdDownlinkServer interface {
	mustEmbedUnimplementedSctpdDownlinkServer()
}

func RegisterSctpdDownlinkServer(s grpc.ServiceRegistrar, srv SctpdDownlinkServer) {
	s.RegisterService(&SctpdDownlink_ServiceDesc, srv)
}

func _SctpdDownlink_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdDownlinkServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdDownlink/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdDownlinkServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdDownlink_SendDl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdDownlinkServer).SendDl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdDownlink/SendDl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdDownlinkServer).SendDl(ctx, req.(*SendDlReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SctpdDownlink_ServiceDesc is the grpc.ServiceDesc for SctpdDownlink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SctpdDownlink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.sctpd.SctpdDownlink",
	HandlerType: (*SctpdDownlinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _SctpdDownlink_Init_Handler,
		},
		{
			MethodName: "SendDl",
			Handler:    _SctpdDownlink_SendDl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sctpd.proto",
}

// SctpdUplinkClient is the client API for SctpdUplink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SctpdUplinkClient interface {
	// SendUl - send an uplink packet to MME
	// @param SendUlReq request specifying packet data and destination
	// @return SendUlRes void response object
	SendUl(ctx context.Context, in *SendUlReq, opts ...grpc.CallOption) (*SendUlRes, error)
	// NewAssoc - notify MME of new eNB association
	// @param NewAssocReq request specifying new association's information
	// @return NewAssocRes void response object
	NewAssoc(ctx context.Context, in *NewAssocReq, opts ...grpc.CallOption) (*NewAssocRes, error)
	// CloseAssoc - notify MME of closing/resetting eNB association
	// @param CloseAssocReq request specifying closing association and close type
	// @return CloseAssocRes void response object
	CloseAssoc(ctx context.Context, in *CloseAssocReq, opts ...grpc.CallOption) (*CloseAssocRes, error)
}

type sctpdUplinkClient struct {
	cc grpc.ClientConnInterface
}

func NewSctpdUplinkClient(cc grpc.ClientConnInterface) SctpdUplinkClient {
	return &sctpdUplinkClient{cc}
}

func (c *sctpdUplinkClient) SendUl(ctx context.Context, in *SendUlReq, opts ...grpc.CallOption) (*SendUlRes, error) {
	out := new(SendUlRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/SendUl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdUplinkClient) NewAssoc(ctx context.Context, in *NewAssocReq, opts ...grpc.CallOption) (*NewAssocRes, error) {
	out := new(NewAssocRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/NewAssoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdUplinkClient) CloseAssoc(ctx context.Context, in *CloseAssocReq, opts ...grpc.CallOption) (*CloseAssocRes, error) {
	out := new(CloseAssocRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/CloseAssoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SctpdUplinkServer is the server API for SctpdUplink service.
// All implementations should embed UnimplementedSctpdUplinkServer
// for forward compatibility
type SctpdUplinkServer interface {
	// SendUl - send an uplink packet to MME
	// @param SendUlReq request specifying packet data and destination
	// @return SendUlRes void response object
	SendUl(context.Context, *SendUlReq) (*SendUlRes, error)
	// NewAssoc - notify MME of new eNB association
	// @param NewAssocReq request specifying new association's information
	// @return NewAssocRes void response object
	NewAssoc(context.Context, *NewAssocReq) (*NewAssocRes, error)
	// CloseAssoc - notify MME of closing/resetting eNB association
	// @param CloseAssocReq request specifying closing association and close type
	// @return CloseAssocRes void response object
	CloseAssoc(context.Context, *CloseAssocReq) (*CloseAssocRes, error)
}

// UnimplementedSctpdUplinkServer should be embedded to have forward compatible implementations.
type UnimplementedSctpdUplinkServer struct {
}

func (UnimplementedSctpdUplinkServer) SendUl(context.Context, *SendUlReq) (*SendUlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUl not implemented")
}
func (UnimplementedSctpdUplinkServer) NewAssoc(context.Context, *NewAssocReq) (*NewAssocRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAssoc not implemented")
}
func (UnimplementedSctpdUplinkServer) CloseAssoc(context.Context, *CloseAssocReq) (*CloseAssocRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseAssoc not implemented")
}

// UnsafeSctpdUplinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SctpdUplinkServer will
// result in compilation errors.
type UnsafeSctpdUplinkServer interface {
	mustEmbedUnimplementedSctpdUplinkServer()
}

func RegisterSctpdUplinkServer(s grpc.ServiceRegistrar, srv SctpdUplinkServer) {
	s.RegisterService(&SctpdUplink_ServiceDesc, srv)
}

func _SctpdUplink_SendUl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).SendUl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/SendUl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).SendUl(ctx, req.(*SendUlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdUplink_NewAssoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAssocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).NewAssoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/NewAssoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).NewAssoc(ctx, req.(*NewAssocReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdUplink_CloseAssoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseAssocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).CloseAssoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/CloseAssoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).CloseAssoc(ctx, req.(*CloseAssocReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SctpdUplink_ServiceDesc is the grpc.ServiceDesc for SctpdUplink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SctpdUplink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.sctpd.SctpdUplink",
	HandlerType: (*SctpdUplinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUl",
			Handler:    _SctpdUplink_SendUl_Handler,
		},
		{
			MethodName: "NewAssoc",
			Handler:    _SctpdUplink_NewAssoc_Handler,
		},
		{
			MethodName: "CloseAssoc",
			Handler:    _SctpdUplink_CloseAssoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sctpd.proto",
}
