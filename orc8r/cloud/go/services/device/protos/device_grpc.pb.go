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

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	RegisterDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
	UpdateDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
	GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error)
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) RegisterDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/RegisterDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/UpdateDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error) {
	out := new(GetDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations should embed UnimplementedDeviceServer
// for forward compatibility
type DeviceServer interface {
	RegisterDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error)
	UpdateDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error)
	GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error)
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*protos.Void, error)
}

// UnimplementedDeviceServer should be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (UnimplementedDeviceServer) RegisterDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevices not implemented")
}
func (UnimplementedDeviceServer) UpdateDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevices not implemented")
}
func (UnimplementedDeviceServer) GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfo not implemented")
}
func (UnimplementedDeviceServer) DeleteDevices(context.Context, *DeleteDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_RegisterDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterOrUpdateDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).RegisterDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/RegisterDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).RegisterDevices(ctx, req.(*RegisterOrUpdateDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterOrUpdateDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/UpdateDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDevices(ctx, req.(*RegisterOrUpdateDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceInfo(ctx, req.(*GetDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.device.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevices",
			Handler:    _Device_RegisterDevices_Handler,
		},
		{
			MethodName: "UpdateDevices",
			Handler:    _Device_UpdateDevices_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _Device_GetDeviceInfo_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _Device_DeleteDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/device/protos/device.proto",
}
