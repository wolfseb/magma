// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.10.0
// source: orc8r/cloud/go/services/state/protos/indexer.proto

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

// IndexerClient is the client API for Indexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexerClient interface {
	// Index a set of states by forwarding to locally-registered indexers.
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	// DeIndex removes a set of states by forwarding to locally-registered indexers.
	DeIndex(ctx context.Context, in *DeIndexRequest, opts ...grpc.CallOption) (*DeIndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error)
}

type indexerClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerClient(cc grpc.ClientConnInterface) IndexerClient {
	return &indexerClient{cc}
}

func (c *indexerClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) DeIndex(ctx context.Context, in *DeIndexRequest, opts ...grpc.CallOption) (*DeIndexResponse, error) {
	out := new(DeIndexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/DeIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error) {
	out := new(PrepareReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/PrepareReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error) {
	out := new(CompleteReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/CompleteReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerServer is the server API for Indexer service.
// All implementations must embed UnimplementedIndexerServer
// for forward compatibility
type IndexerServer interface {
	// Index a set of states by forwarding to locally-registered indexers.
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	// DeIndex removes a set of states by forwarding to locally-registered indexers.
	DeIndex(context.Context, *DeIndexRequest) (*DeIndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(context.Context, *PrepareReindexRequest) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(context.Context, *CompleteReindexRequest) (*CompleteReindexResponse, error)
	mustEmbedUnimplementedIndexerServer()
}

// UnimplementedIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedIndexerServer struct {
}

func (UnimplementedIndexerServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedIndexerServer) DeIndex(context.Context, *DeIndexRequest) (*DeIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeIndex not implemented")
}
func (UnimplementedIndexerServer) PrepareReindex(context.Context, *PrepareReindexRequest) (*PrepareReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareReindex not implemented")
}
func (UnimplementedIndexerServer) CompleteReindex(context.Context, *CompleteReindexRequest) (*CompleteReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteReindex not implemented")
}
func (UnimplementedIndexerServer) mustEmbedUnimplementedIndexerServer() {}

// UnsafeIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexerServer will
// result in compilation errors.
type UnsafeIndexerServer interface {
	mustEmbedUnimplementedIndexerServer()
}

func RegisterIndexerServer(s grpc.ServiceRegistrar, srv IndexerServer) {
	s.RegisterService(&Indexer_ServiceDesc, srv)
}

func _Indexer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_DeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).DeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/DeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).DeIndex(ctx, req.(*DeIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_PrepareReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).PrepareReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/PrepareReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).PrepareReindex(ctx, req.(*PrepareReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_CompleteReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).CompleteReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/CompleteReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).CompleteReindex(ctx, req.(*CompleteReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Indexer_ServiceDesc is the grpc.ServiceDesc for Indexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Indexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.state.Indexer",
	HandlerType: (*IndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Indexer_Index_Handler,
		},
		{
			MethodName: "DeIndex",
			Handler:    _Indexer_DeIndex_Handler,
		},
		{
			MethodName: "PrepareReindex",
			Handler:    _Indexer_PrepareReindex_Handler,
		},
		{
			MethodName: "CompleteReindex",
			Handler:    _Indexer_CompleteReindex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/state/protos/indexer.proto",
}
