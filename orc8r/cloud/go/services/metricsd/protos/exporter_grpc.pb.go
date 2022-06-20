// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: orc8r/cloud/go/services/metricsd/protos/exporter.proto

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

// MetricsExporterClient is the client API for MetricsExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsExporterClient interface {
	// Submit metrics to datasinks.
	Submit(ctx context.Context, in *SubmitMetricsRequest, opts ...grpc.CallOption) (*SubmitMetricsResponse, error)
}

type metricsExporterClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsExporterClient(cc grpc.ClientConnInterface) MetricsExporterClient {
	return &metricsExporterClient{cc}
}

func (c *metricsExporterClient) Submit(ctx context.Context, in *SubmitMetricsRequest, opts ...grpc.CallOption) (*SubmitMetricsResponse, error) {
	out := new(SubmitMetricsResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.metricsd.MetricsExporter/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsExporterServer is the server API for MetricsExporter service.
// All implementations must embed UnimplementedMetricsExporterServer
// for forward compatibility
type MetricsExporterServer interface {
	// Submit metrics to datasinks.
	Submit(context.Context, *SubmitMetricsRequest) (*SubmitMetricsResponse, error)
	mustEmbedUnimplementedMetricsExporterServer()
}

// UnimplementedMetricsExporterServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsExporterServer struct {
}

func (UnimplementedMetricsExporterServer) Submit(context.Context, *SubmitMetricsRequest) (*SubmitMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedMetricsExporterServer) mustEmbedUnimplementedMetricsExporterServer() {}

// UnsafeMetricsExporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsExporterServer will
// result in compilation errors.
type UnsafeMetricsExporterServer interface {
	mustEmbedUnimplementedMetricsExporterServer()
}

func RegisterMetricsExporterServer(s grpc.ServiceRegistrar, srv MetricsExporterServer) {
	s.RegisterService(&MetricsExporter_ServiceDesc, srv)
}

func _MetricsExporter_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsExporterServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.metricsd.MetricsExporter/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsExporterServer).Submit(ctx, req.(*SubmitMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsExporter_ServiceDesc is the grpc.ServiceDesc for MetricsExporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsExporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.metricsd.MetricsExporter",
	HandlerType: (*MetricsExporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _MetricsExporter_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/metricsd/protos/exporter.proto",
}
