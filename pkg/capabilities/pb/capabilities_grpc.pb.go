// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: capabilities/pb/capabilities.proto

package pb

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
	BaseCapability_Info_FullMethodName = "/loop.BaseCapability/Info"
)

// BaseCapabilityClient is the client API for BaseCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseCapabilityClient interface {
	Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CapabilityInfoReply, error)
}

type baseCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseCapabilityClient(cc grpc.ClientConnInterface) BaseCapabilityClient {
	return &baseCapabilityClient{cc}
}

func (c *baseCapabilityClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CapabilityInfoReply, error) {
	out := new(CapabilityInfoReply)
	err := c.cc.Invoke(ctx, BaseCapability_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseCapabilityServer is the server API for BaseCapability service.
// All implementations must embed UnimplementedBaseCapabilityServer
// for forward compatibility
type BaseCapabilityServer interface {
	Info(context.Context, *emptypb.Empty) (*CapabilityInfoReply, error)
	mustEmbedUnimplementedBaseCapabilityServer()
}

// UnimplementedBaseCapabilityServer must be embedded to have forward compatible implementations.
type UnimplementedBaseCapabilityServer struct {
}

func (UnimplementedBaseCapabilityServer) Info(context.Context, *emptypb.Empty) (*CapabilityInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedBaseCapabilityServer) mustEmbedUnimplementedBaseCapabilityServer() {}

// UnsafeBaseCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseCapabilityServer will
// result in compilation errors.
type UnsafeBaseCapabilityServer interface {
	mustEmbedUnimplementedBaseCapabilityServer()
}

func RegisterBaseCapabilityServer(s grpc.ServiceRegistrar, srv BaseCapabilityServer) {
	s.RegisterService(&BaseCapability_ServiceDesc, srv)
}

func _BaseCapability_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseCapabilityServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseCapability_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseCapabilityServer).Info(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseCapability_ServiceDesc is the grpc.ServiceDesc for BaseCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.BaseCapability",
	HandlerType: (*BaseCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _BaseCapability_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "capabilities/pb/capabilities.proto",
}

const (
	TriggerExecutable_RegisterTrigger_FullMethodName   = "/loop.TriggerExecutable/RegisterTrigger"
	TriggerExecutable_UnregisterTrigger_FullMethodName = "/loop.TriggerExecutable/UnregisterTrigger"
)

// TriggerExecutableClient is the client API for TriggerExecutable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerExecutableClient interface {
	RegisterTrigger(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (TriggerExecutable_RegisterTriggerClient, error)
	UnregisterTrigger(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type triggerExecutableClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerExecutableClient(cc grpc.ClientConnInterface) TriggerExecutableClient {
	return &triggerExecutableClient{cc}
}

func (c *triggerExecutableClient) RegisterTrigger(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (TriggerExecutable_RegisterTriggerClient, error) {
	stream, err := c.cc.NewStream(ctx, &TriggerExecutable_ServiceDesc.Streams[0], TriggerExecutable_RegisterTrigger_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &triggerExecutableRegisterTriggerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TriggerExecutable_RegisterTriggerClient interface {
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type triggerExecutableRegisterTriggerClient struct {
	grpc.ClientStream
}

func (x *triggerExecutableRegisterTriggerClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *triggerExecutableClient) UnregisterTrigger(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TriggerExecutable_UnregisterTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerExecutableServer is the server API for TriggerExecutable service.
// All implementations must embed UnimplementedTriggerExecutableServer
// for forward compatibility
type TriggerExecutableServer interface {
	RegisterTrigger(*CapabilityRequest, TriggerExecutable_RegisterTriggerServer) error
	UnregisterTrigger(context.Context, *CapabilityRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTriggerExecutableServer()
}

// UnimplementedTriggerExecutableServer must be embedded to have forward compatible implementations.
type UnimplementedTriggerExecutableServer struct {
}

func (UnimplementedTriggerExecutableServer) RegisterTrigger(*CapabilityRequest, TriggerExecutable_RegisterTriggerServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterTrigger not implemented")
}
func (UnimplementedTriggerExecutableServer) UnregisterTrigger(context.Context, *CapabilityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterTrigger not implemented")
}
func (UnimplementedTriggerExecutableServer) mustEmbedUnimplementedTriggerExecutableServer() {}

// UnsafeTriggerExecutableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TriggerExecutableServer will
// result in compilation errors.
type UnsafeTriggerExecutableServer interface {
	mustEmbedUnimplementedTriggerExecutableServer()
}

func RegisterTriggerExecutableServer(s grpc.ServiceRegistrar, srv TriggerExecutableServer) {
	s.RegisterService(&TriggerExecutable_ServiceDesc, srv)
}

func _TriggerExecutable_RegisterTrigger_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CapabilityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TriggerExecutableServer).RegisterTrigger(m, &triggerExecutableRegisterTriggerServer{stream})
}

type TriggerExecutable_RegisterTriggerServer interface {
	Send(*ResponseMessage) error
	grpc.ServerStream
}

type triggerExecutableRegisterTriggerServer struct {
	grpc.ServerStream
}

func (x *triggerExecutableRegisterTriggerServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _TriggerExecutable_UnregisterTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerExecutableServer).UnregisterTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TriggerExecutable_UnregisterTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerExecutableServer).UnregisterTrigger(ctx, req.(*CapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TriggerExecutable_ServiceDesc is the grpc.ServiceDesc for TriggerExecutable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TriggerExecutable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.TriggerExecutable",
	HandlerType: (*TriggerExecutableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnregisterTrigger",
			Handler:    _TriggerExecutable_UnregisterTrigger_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterTrigger",
			Handler:       _TriggerExecutable_RegisterTrigger_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "capabilities/pb/capabilities.proto",
}

const (
	CallbackExecutable_RegisterToWorkflow_FullMethodName     = "/loop.CallbackExecutable/RegisterToWorkflow"
	CallbackExecutable_UnregisterFromWorkflow_FullMethodName = "/loop.CallbackExecutable/UnregisterFromWorkflow"
	CallbackExecutable_Execute_FullMethodName                = "/loop.CallbackExecutable/Execute"
)

// CallbackExecutableClient is the client API for CallbackExecutable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallbackExecutableClient interface {
	RegisterToWorkflow(ctx context.Context, in *RegisterToWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnregisterFromWorkflow(ctx context.Context, in *UnregisterFromWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Execute(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (CallbackExecutable_ExecuteClient, error)
}

type callbackExecutableClient struct {
	cc grpc.ClientConnInterface
}

func NewCallbackExecutableClient(cc grpc.ClientConnInterface) CallbackExecutableClient {
	return &callbackExecutableClient{cc}
}

func (c *callbackExecutableClient) RegisterToWorkflow(ctx context.Context, in *RegisterToWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CallbackExecutable_RegisterToWorkflow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackExecutableClient) UnregisterFromWorkflow(ctx context.Context, in *UnregisterFromWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CallbackExecutable_UnregisterFromWorkflow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackExecutableClient) Execute(ctx context.Context, in *CapabilityRequest, opts ...grpc.CallOption) (CallbackExecutable_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CallbackExecutable_ServiceDesc.Streams[0], CallbackExecutable_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &callbackExecutableExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CallbackExecutable_ExecuteClient interface {
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type callbackExecutableExecuteClient struct {
	grpc.ClientStream
}

func (x *callbackExecutableExecuteClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CallbackExecutableServer is the server API for CallbackExecutable service.
// All implementations must embed UnimplementedCallbackExecutableServer
// for forward compatibility
type CallbackExecutableServer interface {
	RegisterToWorkflow(context.Context, *RegisterToWorkflowRequest) (*emptypb.Empty, error)
	UnregisterFromWorkflow(context.Context, *UnregisterFromWorkflowRequest) (*emptypb.Empty, error)
	Execute(*CapabilityRequest, CallbackExecutable_ExecuteServer) error
	mustEmbedUnimplementedCallbackExecutableServer()
}

// UnimplementedCallbackExecutableServer must be embedded to have forward compatible implementations.
type UnimplementedCallbackExecutableServer struct {
}

func (UnimplementedCallbackExecutableServer) RegisterToWorkflow(context.Context, *RegisterToWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterToWorkflow not implemented")
}
func (UnimplementedCallbackExecutableServer) UnregisterFromWorkflow(context.Context, *UnregisterFromWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterFromWorkflow not implemented")
}
func (UnimplementedCallbackExecutableServer) Execute(*CapabilityRequest, CallbackExecutable_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedCallbackExecutableServer) mustEmbedUnimplementedCallbackExecutableServer() {}

// UnsafeCallbackExecutableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallbackExecutableServer will
// result in compilation errors.
type UnsafeCallbackExecutableServer interface {
	mustEmbedUnimplementedCallbackExecutableServer()
}

func RegisterCallbackExecutableServer(s grpc.ServiceRegistrar, srv CallbackExecutableServer) {
	s.RegisterService(&CallbackExecutable_ServiceDesc, srv)
}

func _CallbackExecutable_RegisterToWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterToWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackExecutableServer).RegisterToWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CallbackExecutable_RegisterToWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackExecutableServer).RegisterToWorkflow(ctx, req.(*RegisterToWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackExecutable_UnregisterFromWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterFromWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackExecutableServer).UnregisterFromWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CallbackExecutable_UnregisterFromWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackExecutableServer).UnregisterFromWorkflow(ctx, req.(*UnregisterFromWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackExecutable_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CapabilityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CallbackExecutableServer).Execute(m, &callbackExecutableExecuteServer{stream})
}

type CallbackExecutable_ExecuteServer interface {
	Send(*ResponseMessage) error
	grpc.ServerStream
}

type callbackExecutableExecuteServer struct {
	grpc.ServerStream
}

func (x *callbackExecutableExecuteServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

// CallbackExecutable_ServiceDesc is the grpc.ServiceDesc for CallbackExecutable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallbackExecutable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.CallbackExecutable",
	HandlerType: (*CallbackExecutableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterToWorkflow",
			Handler:    _CallbackExecutable_RegisterToWorkflow_Handler,
		},
		{
			MethodName: "UnregisterFromWorkflow",
			Handler:    _CallbackExecutable_UnregisterFromWorkflow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _CallbackExecutable_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "capabilities/pb/capabilities.proto",
}
