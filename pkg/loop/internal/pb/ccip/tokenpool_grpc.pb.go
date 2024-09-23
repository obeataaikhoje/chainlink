// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: tokenpool.proto

package ccippb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TokenPoolBatcherReader_GetInboundTokenPoolRateLimits_FullMethodName = "/loop.internal.pb.ccip.TokenPoolBatcherReader/GetInboundTokenPoolRateLimits"
	TokenPoolBatcherReader_Close_FullMethodName                         = "/loop.internal.pb.ccip.TokenPoolBatcherReader/Close"
)

// TokenPoolBatcherReaderClient is the client API for TokenPoolBatcherReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenPoolBatcherReaderClient interface {
	GetInboundTokenPoolRateLimits(ctx context.Context, in *GetInboundTokenPoolRateLimitsRequest, opts ...grpc.CallOption) (*GetInboundTokenPoolRateLimitsResponse, error)
	Close(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tokenPoolBatcherReaderClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenPoolBatcherReaderClient(cc grpc.ClientConnInterface) TokenPoolBatcherReaderClient {
	return &tokenPoolBatcherReaderClient{cc}
}

func (c *tokenPoolBatcherReaderClient) GetInboundTokenPoolRateLimits(ctx context.Context, in *GetInboundTokenPoolRateLimitsRequest, opts ...grpc.CallOption) (*GetInboundTokenPoolRateLimitsResponse, error) {
	out := new(GetInboundTokenPoolRateLimitsResponse)
	err := c.cc.Invoke(ctx, TokenPoolBatcherReader_GetInboundTokenPoolRateLimits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenPoolBatcherReaderClient) Close(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TokenPoolBatcherReader_Close_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenPoolBatcherReaderServer is the server API for TokenPoolBatcherReader service.
// All implementations must embed UnimplementedTokenPoolBatcherReaderServer
// for forward compatibility
type TokenPoolBatcherReaderServer interface {
	GetInboundTokenPoolRateLimits(context.Context, *GetInboundTokenPoolRateLimitsRequest) (*GetInboundTokenPoolRateLimitsResponse, error)
	Close(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedTokenPoolBatcherReaderServer()
}

// UnimplementedTokenPoolBatcherReaderServer must be embedded to have forward compatible implementations.
type UnimplementedTokenPoolBatcherReaderServer struct {
}

func (UnimplementedTokenPoolBatcherReaderServer) GetInboundTokenPoolRateLimits(context.Context, *GetInboundTokenPoolRateLimitsRequest) (*GetInboundTokenPoolRateLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboundTokenPoolRateLimits not implemented")
}
func (UnimplementedTokenPoolBatcherReaderServer) Close(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedTokenPoolBatcherReaderServer) mustEmbedUnimplementedTokenPoolBatcherReaderServer() {
}

// UnsafeTokenPoolBatcherReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenPoolBatcherReaderServer will
// result in compilation errors.
type UnsafeTokenPoolBatcherReaderServer interface {
	mustEmbedUnimplementedTokenPoolBatcherReaderServer()
}

func RegisterTokenPoolBatcherReaderServer(s grpc.ServiceRegistrar, srv TokenPoolBatcherReaderServer) {
	s.RegisterService(&TokenPoolBatcherReader_ServiceDesc, srv)
}

func _TokenPoolBatcherReader_GetInboundTokenPoolRateLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInboundTokenPoolRateLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenPoolBatcherReaderServer).GetInboundTokenPoolRateLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenPoolBatcherReader_GetInboundTokenPoolRateLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenPoolBatcherReaderServer).GetInboundTokenPoolRateLimits(ctx, req.(*GetInboundTokenPoolRateLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenPoolBatcherReader_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenPoolBatcherReaderServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenPoolBatcherReader_Close_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenPoolBatcherReaderServer).Close(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenPoolBatcherReader_ServiceDesc is the grpc.ServiceDesc for TokenPoolBatcherReader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenPoolBatcherReader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.ccip.TokenPoolBatcherReader",
	HandlerType: (*TokenPoolBatcherReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInboundTokenPoolRateLimits",
			Handler:    _TokenPoolBatcherReader_GetInboundTokenPoolRateLimits_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _TokenPoolBatcherReader_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenpool.proto",
}
