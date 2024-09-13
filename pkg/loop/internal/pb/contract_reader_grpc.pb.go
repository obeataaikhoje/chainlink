// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: contract_reader.proto

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
	ContractReader_GetLatestValue_FullMethodName                = "/loop.ContractReader/GetLatestValue"
	ContractReader_GetLatestValueWithDefaultType_FullMethodName = "/loop.ContractReader/GetLatestValueWithDefaultType"
	ContractReader_BatchGetLatestValues_FullMethodName          = "/loop.ContractReader/BatchGetLatestValues"
	ContractReader_QueryKey_FullMethodName                      = "/loop.ContractReader/QueryKey"
	ContractReader_Bind_FullMethodName                          = "/loop.ContractReader/Bind"
	ContractReader_Unbind_FullMethodName                        = "/loop.ContractReader/Unbind"
)

// ContractReaderClient is the client API for ContractReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractReaderClient interface {
	GetLatestValue(ctx context.Context, in *GetLatestValueRequest, opts ...grpc.CallOption) (*GetLatestValueReply, error)
	GetLatestValueWithDefaultType(ctx context.Context, in *GetLatestValueWithDefaultTypeRequest, opts ...grpc.CallOption) (*GetLatestValueWithDefaultTypeReply, error)
	BatchGetLatestValues(ctx context.Context, in *BatchGetLatestValuesRequest, opts ...grpc.CallOption) (*BatchGetLatestValuesReply, error)
	QueryKey(ctx context.Context, in *QueryKeyRequest, opts ...grpc.CallOption) (*QueryKeyReply, error)
	Bind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unbind(ctx context.Context, in *UnbindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contractReaderClient struct {
	cc grpc.ClientConnInterface
}

func NewContractReaderClient(cc grpc.ClientConnInterface) ContractReaderClient {
	return &contractReaderClient{cc}
}

func (c *contractReaderClient) GetLatestValue(ctx context.Context, in *GetLatestValueRequest, opts ...grpc.CallOption) (*GetLatestValueReply, error) {
	out := new(GetLatestValueReply)
	err := c.cc.Invoke(ctx, ContractReader_GetLatestValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderClient) GetLatestValueWithDefaultType(ctx context.Context, in *GetLatestValueWithDefaultTypeRequest, opts ...grpc.CallOption) (*GetLatestValueWithDefaultTypeReply, error) {
	out := new(GetLatestValueWithDefaultTypeReply)
	err := c.cc.Invoke(ctx, ContractReader_GetLatestValueWithDefaultType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderClient) BatchGetLatestValues(ctx context.Context, in *BatchGetLatestValuesRequest, opts ...grpc.CallOption) (*BatchGetLatestValuesReply, error) {
	out := new(BatchGetLatestValuesReply)
	err := c.cc.Invoke(ctx, ContractReader_BatchGetLatestValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderClient) QueryKey(ctx context.Context, in *QueryKeyRequest, opts ...grpc.CallOption) (*QueryKeyReply, error) {
	out := new(QueryKeyReply)
	err := c.cc.Invoke(ctx, ContractReader_QueryKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderClient) Bind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContractReader_Bind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractReaderClient) Unbind(ctx context.Context, in *UnbindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContractReader_Unbind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractReaderServer is the server API for ContractReader service.
// All implementations must embed UnimplementedContractReaderServer
// for forward compatibility
type ContractReaderServer interface {
	GetLatestValue(context.Context, *GetLatestValueRequest) (*GetLatestValueReply, error)
	GetLatestValueWithDefaultType(context.Context, *GetLatestValueWithDefaultTypeRequest) (*GetLatestValueWithDefaultTypeReply, error)
	BatchGetLatestValues(context.Context, *BatchGetLatestValuesRequest) (*BatchGetLatestValuesReply, error)
	QueryKey(context.Context, *QueryKeyRequest) (*QueryKeyReply, error)
	Bind(context.Context, *BindRequest) (*emptypb.Empty, error)
	Unbind(context.Context, *UnbindRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedContractReaderServer()
}

// UnimplementedContractReaderServer must be embedded to have forward compatible implementations.
type UnimplementedContractReaderServer struct {
}

func (UnimplementedContractReaderServer) GetLatestValue(context.Context, *GetLatestValueRequest) (*GetLatestValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValue not implemented")
}
func (UnimplementedContractReaderServer) GetLatestValueWithDefaultType(context.Context, *GetLatestValueWithDefaultTypeRequest) (*GetLatestValueWithDefaultTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValueWithDefaultType not implemented")
}
func (UnimplementedContractReaderServer) BatchGetLatestValues(context.Context, *BatchGetLatestValuesRequest) (*BatchGetLatestValuesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetLatestValues not implemented")
}
func (UnimplementedContractReaderServer) QueryKey(context.Context, *QueryKeyRequest) (*QueryKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryKey not implemented")
}
func (UnimplementedContractReaderServer) Bind(context.Context, *BindRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (UnimplementedContractReaderServer) Unbind(context.Context, *UnbindRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unbind not implemented")
}
func (UnimplementedContractReaderServer) mustEmbedUnimplementedContractReaderServer() {}

// UnsafeContractReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractReaderServer will
// result in compilation errors.
type UnsafeContractReaderServer interface {
	mustEmbedUnimplementedContractReaderServer()
}

func RegisterContractReaderServer(s grpc.ServiceRegistrar, srv ContractReaderServer) {
	s.RegisterService(&ContractReader_ServiceDesc, srv)
}

func _ContractReader_GetLatestValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).GetLatestValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_GetLatestValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).GetLatestValue(ctx, req.(*GetLatestValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReader_GetLatestValueWithDefaultType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValueWithDefaultTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).GetLatestValueWithDefaultType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_GetLatestValueWithDefaultType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).GetLatestValueWithDefaultType(ctx, req.(*GetLatestValueWithDefaultTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReader_BatchGetLatestValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetLatestValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).BatchGetLatestValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_BatchGetLatestValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).BatchGetLatestValues(ctx, req.(*BatchGetLatestValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReader_QueryKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).QueryKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_QueryKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).QueryKey(ctx, req.(*QueryKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReader_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_Bind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).Bind(ctx, req.(*BindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractReader_Unbind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractReaderServer).Unbind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractReader_Unbind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractReaderServer).Unbind(ctx, req.(*UnbindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractReader_ServiceDesc is the grpc.ServiceDesc for ContractReader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractReader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.ContractReader",
	HandlerType: (*ContractReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestValue",
			Handler:    _ContractReader_GetLatestValue_Handler,
		},
		{
			MethodName: "GetLatestValueWithDefaultType",
			Handler:    _ContractReader_GetLatestValueWithDefaultType_Handler,
		},
		{
			MethodName: "BatchGetLatestValues",
			Handler:    _ContractReader_BatchGetLatestValues_Handler,
		},
		{
			MethodName: "QueryKey",
			Handler:    _ContractReader_QueryKey_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _ContractReader_Bind_Handler,
		},
		{
			MethodName: "Unbind",
			Handler:    _ContractReader_Unbind_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_reader.proto",
}
