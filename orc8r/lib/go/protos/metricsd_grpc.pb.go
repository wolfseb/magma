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

// MetricsControllerClient is the client API for MetricsController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsControllerClient interface {
	// Report a collection of metrics from a service
	Collect(ctx context.Context, in *MetricsContainer, opts ...grpc.CallOption) (*Void, error)
}

type metricsControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsControllerClient(cc grpc.ClientConnInterface) MetricsControllerClient {
	return &metricsControllerClient{cc}
}

func (c *metricsControllerClient) Collect(ctx context.Context, in *MetricsContainer, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.MetricsController/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsControllerServer is the server API for MetricsController service.
// All implementations should embed UnimplementedMetricsControllerServer
// for forward compatibility
type MetricsControllerServer interface {
	// Report a collection of metrics from a service
	Collect(context.Context, *MetricsContainer) (*Void, error)
}

// UnimplementedMetricsControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMetricsControllerServer struct {
}

func (UnimplementedMetricsControllerServer) Collect(context.Context, *MetricsContainer) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

// UnsafeMetricsControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsControllerServer will
// result in compilation errors.
type UnsafeMetricsControllerServer interface {
	mustEmbedUnimplementedMetricsControllerServer()
}

func RegisterMetricsControllerServer(s grpc.ServiceRegistrar, srv MetricsControllerServer) {
	s.RegisterService(&MetricsController_ServiceDesc, srv)
}

func _MetricsController_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsContainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsControllerServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.MetricsController/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsControllerServer).Collect(ctx, req.(*MetricsContainer))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsController_ServiceDesc is the grpc.ServiceDesc for MetricsController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.MetricsController",
	HandlerType: (*MetricsControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _MetricsController_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/metricsd.proto",
}

// CloudMetricsControllerClient is the client API for CloudMetricsController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudMetricsControllerClient interface {
	// Push a collection of metrics to metricsd
	Push(ctx context.Context, in *PushedMetricsContainer, opts ...grpc.CallOption) (*Void, error)
}

type cloudMetricsControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudMetricsControllerClient(cc grpc.ClientConnInterface) CloudMetricsControllerClient {
	return &cloudMetricsControllerClient{cc}
}

func (c *cloudMetricsControllerClient) Push(ctx context.Context, in *PushedMetricsContainer, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.CloudMetricsController/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudMetricsControllerServer is the server API for CloudMetricsController service.
// All implementations should embed UnimplementedCloudMetricsControllerServer
// for forward compatibility
type CloudMetricsControllerServer interface {
	// Push a collection of metrics to metricsd
	Push(context.Context, *PushedMetricsContainer) (*Void, error)
}

// UnimplementedCloudMetricsControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCloudMetricsControllerServer struct {
}

func (UnimplementedCloudMetricsControllerServer) Push(context.Context, *PushedMetricsContainer) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

// UnsafeCloudMetricsControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudMetricsControllerServer will
// result in compilation errors.
type UnsafeCloudMetricsControllerServer interface {
	mustEmbedUnimplementedCloudMetricsControllerServer()
}

func RegisterCloudMetricsControllerServer(s grpc.ServiceRegistrar, srv CloudMetricsControllerServer) {
	s.RegisterService(&CloudMetricsController_ServiceDesc, srv)
}

func _CloudMetricsController_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushedMetricsContainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudMetricsControllerServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.CloudMetricsController/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudMetricsControllerServer).Push(ctx, req.(*PushedMetricsContainer))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudMetricsController_ServiceDesc is the grpc.ServiceDesc for CloudMetricsController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudMetricsController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.CloudMetricsController",
	HandlerType: (*CloudMetricsControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _CloudMetricsController_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/metricsd.proto",
}
