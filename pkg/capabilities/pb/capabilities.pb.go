// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: capabilities/pb/capabilities.proto

package pb

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CapabilityType int32

const (
	CapabilityType_CAPABILITY_TYPE_UNKNOWN   CapabilityType = 0
	CapabilityType_CAPABILITY_TYPE_TRIGGER   CapabilityType = 1
	CapabilityType_CAPABILITY_TYPE_ACTION    CapabilityType = 2
	CapabilityType_CAPABILITY_TYPE_CONSENSUS CapabilityType = 3
	CapabilityType_CAPABILITY_TYPE_TARGET    CapabilityType = 4
)

// Enum value maps for CapabilityType.
var (
	CapabilityType_name = map[int32]string{
		0: "CAPABILITY_TYPE_UNKNOWN",
		1: "CAPABILITY_TYPE_TRIGGER",
		2: "CAPABILITY_TYPE_ACTION",
		3: "CAPABILITY_TYPE_CONSENSUS",
		4: "CAPABILITY_TYPE_TARGET",
	}
	CapabilityType_value = map[string]int32{
		"CAPABILITY_TYPE_UNKNOWN":   0,
		"CAPABILITY_TYPE_TRIGGER":   1,
		"CAPABILITY_TYPE_ACTION":    2,
		"CAPABILITY_TYPE_CONSENSUS": 3,
		"CAPABILITY_TYPE_TARGET":    4,
	}
)

func (x CapabilityType) Enum() *CapabilityType {
	p := new(CapabilityType)
	*p = x
	return p
}

func (x CapabilityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapabilityType) Descriptor() protoreflect.EnumDescriptor {
	return file_capabilities_pb_capabilities_proto_enumTypes[0].Descriptor()
}

func (CapabilityType) Type() protoreflect.EnumType {
	return &file_capabilities_pb_capabilities_proto_enumTypes[0]
}

func (x CapabilityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapabilityType.Descriptor instead.
func (CapabilityType) EnumDescriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{0}
}

type CapabilityInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CapabilityType CapabilityType `protobuf:"varint,2,opt,name=capability_type,json=capabilityType,proto3,enum=loop.CapabilityType" json:"capability_type,omitempty"`
	Description    string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Version        string         `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CapabilityInfoReply) Reset() {
	*x = CapabilityInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityInfoReply) ProtoMessage() {}

func (x *CapabilityInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityInfoReply.ProtoReflect.Descriptor instead.
func (*CapabilityInfoReply) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{0}
}

func (x *CapabilityInfoReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CapabilityInfoReply) GetCapabilityType() CapabilityType {
	if x != nil {
		return x.CapabilityType
	}
	return CapabilityType_CAPABILITY_TYPE_UNKNOWN
}

func (x *CapabilityInfoReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CapabilityInfoReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId          string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	WorkflowExecutionId string `protobuf:"bytes,2,opt,name=workflow_execution_id,json=workflowExecutionId,proto3" json:"workflow_execution_id,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMetadata) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *RequestMetadata) GetWorkflowExecutionId() string {
	if x != nil {
		return x.WorkflowExecutionId
	}
	return ""
}

type CapabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *RequestMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Config   *pb.Value        `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Inputs   *pb.Value        `protobuf:"bytes,3,opt,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *CapabilityRequest) Reset() {
	*x = CapabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityRequest) ProtoMessage() {}

func (x *CapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityRequest.ProtoReflect.Descriptor instead.
func (*CapabilityRequest) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{2}
}

func (x *CapabilityRequest) GetMetadata() *RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CapabilityRequest) GetConfig() *pb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *CapabilityRequest) GetInputs() *pb.Value {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type TriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Metadata *RequestMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Config   *pb.Value        `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	Inputs   *pb.Value        `protobuf:"bytes,5,opt,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *TriggerRequest) Reset() {
	*x = TriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerRequest) ProtoMessage() {}

func (x *TriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerRequest.ProtoReflect.Descriptor instead.
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TriggerRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TriggerRequest) GetMetadata() *RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TriggerRequest) GetConfig() *pb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *TriggerRequest) GetInputs() *pb.Value {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type TriggerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload   *pb.Value `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TriggerEvent) Reset() {
	*x = TriggerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerEvent) ProtoMessage() {}

func (x *TriggerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerEvent.ProtoReflect.Descriptor instead.
func (*TriggerEvent) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{4}
}

func (x *TriggerEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TriggerEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TriggerEvent) GetPayload() *pb.Value {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CapabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *pb.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CapabilityResponse) Reset() {
	*x = CapabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityResponse) ProtoMessage() {}

func (x *CapabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityResponse.ProtoReflect.Descriptor instead.
func (*CapabilityResponse) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{5}
}

func (x *CapabilityResponse) GetValue() *pb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CapabilityResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RegistrationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
}

func (x *RegistrationMetadata) Reset() {
	*x = RegistrationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationMetadata) ProtoMessage() {}

func (x *RegistrationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationMetadata.ProtoReflect.Descriptor instead.
func (*RegistrationMetadata) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{6}
}

func (x *RegistrationMetadata) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type RegisterToWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *RegistrationMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Config   *pb.Value             `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *RegisterToWorkflowRequest) Reset() {
	*x = RegisterToWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterToWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterToWorkflowRequest) ProtoMessage() {}

func (x *RegisterToWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterToWorkflowRequest.ProtoReflect.Descriptor instead.
func (*RegisterToWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterToWorkflowRequest) GetMetadata() *RegistrationMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RegisterToWorkflowRequest) GetConfig() *pb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

type UnregisterFromWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *RegistrationMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Config   *pb.Value             `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UnregisterFromWorkflowRequest) Reset() {
	*x = UnregisterFromWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_capabilities_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterFromWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterFromWorkflowRequest) ProtoMessage() {}

func (x *UnregisterFromWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_capabilities_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterFromWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UnregisterFromWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_capabilities_proto_rawDescGZIP(), []int{8}
}

func (x *UnregisterFromWorkflowRequest) GetMetadata() *RegistrationMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UnregisterFromWorkflowRequest) GetConfig() *pb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_capabilities_pb_capabilities_proto protoreflect.FileDescriptor

var file_capabilities_pb_capabilities_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x16, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x22, 0xb5, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x65, 0x0a, 0x0c, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x4f, 0x0a, 0x12, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x37, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x19, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x7e, 0x0a, 0x1d, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0xa1, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x50,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12,
	0x1d, 0x0a, 0x19, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x4e, 0x53, 0x55, 0x53, 0x10, 0x03, 0x12, 0x1a,
	0x0a, 0x16, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x04, 0x32, 0x4d, 0x0a, 0x0e, 0x42, 0x61,
	0x73, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x99, 0x01, 0x0a, 0x11, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x43, 0x0a, 0x11, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0x80, 0x02, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x12,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x1f, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x16, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x23, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x55,
	0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capabilities_pb_capabilities_proto_rawDescOnce sync.Once
	file_capabilities_pb_capabilities_proto_rawDescData = file_capabilities_pb_capabilities_proto_rawDesc
)

func file_capabilities_pb_capabilities_proto_rawDescGZIP() []byte {
	file_capabilities_pb_capabilities_proto_rawDescOnce.Do(func() {
		file_capabilities_pb_capabilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_capabilities_pb_capabilities_proto_rawDescData)
	})
	return file_capabilities_pb_capabilities_proto_rawDescData
}

var file_capabilities_pb_capabilities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_capabilities_pb_capabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_capabilities_pb_capabilities_proto_goTypes = []interface{}{
	(CapabilityType)(0),                   // 0: loop.CapabilityType
	(*CapabilityInfoReply)(nil),           // 1: loop.CapabilityInfoReply
	(*RequestMetadata)(nil),               // 2: loop.RequestMetadata
	(*CapabilityRequest)(nil),             // 3: loop.CapabilityRequest
	(*TriggerRequest)(nil),                // 4: loop.TriggerRequest
	(*TriggerEvent)(nil),                  // 5: loop.TriggerEvent
	(*CapabilityResponse)(nil),            // 6: loop.CapabilityResponse
	(*RegistrationMetadata)(nil),          // 7: loop.RegistrationMetadata
	(*RegisterToWorkflowRequest)(nil),     // 8: loop.RegisterToWorkflowRequest
	(*UnregisterFromWorkflowRequest)(nil), // 9: loop.UnregisterFromWorkflowRequest
	(*pb.Value)(nil),                      // 10: values.Value
	(*emptypb.Empty)(nil),                 // 11: google.protobuf.Empty
}
var file_capabilities_pb_capabilities_proto_depIdxs = []int32{
	0,  // 0: loop.CapabilityInfoReply.capability_type:type_name -> loop.CapabilityType
	2,  // 1: loop.CapabilityRequest.metadata:type_name -> loop.RequestMetadata
	10, // 2: loop.CapabilityRequest.config:type_name -> values.Value
	10, // 3: loop.CapabilityRequest.inputs:type_name -> values.Value
	2,  // 4: loop.TriggerRequest.metadata:type_name -> loop.RequestMetadata
	10, // 5: loop.TriggerRequest.config:type_name -> values.Value
	10, // 6: loop.TriggerRequest.inputs:type_name -> values.Value
	10, // 7: loop.TriggerEvent.payload:type_name -> values.Value
	10, // 8: loop.CapabilityResponse.value:type_name -> values.Value
	7,  // 9: loop.RegisterToWorkflowRequest.metadata:type_name -> loop.RegistrationMetadata
	10, // 10: loop.RegisterToWorkflowRequest.config:type_name -> values.Value
	7,  // 11: loop.UnregisterFromWorkflowRequest.metadata:type_name -> loop.RegistrationMetadata
	10, // 12: loop.UnregisterFromWorkflowRequest.config:type_name -> values.Value
	11, // 13: loop.BaseCapability.Info:input_type -> google.protobuf.Empty
	4,  // 14: loop.TriggerExecutable.RegisterTrigger:input_type -> loop.TriggerRequest
	4,  // 15: loop.TriggerExecutable.UnregisterTrigger:input_type -> loop.TriggerRequest
	8,  // 16: loop.CallbackExecutable.RegisterToWorkflow:input_type -> loop.RegisterToWorkflowRequest
	9,  // 17: loop.CallbackExecutable.UnregisterFromWorkflow:input_type -> loop.UnregisterFromWorkflowRequest
	3,  // 18: loop.CallbackExecutable.Execute:input_type -> loop.CapabilityRequest
	1,  // 19: loop.BaseCapability.Info:output_type -> loop.CapabilityInfoReply
	5,  // 20: loop.TriggerExecutable.RegisterTrigger:output_type -> loop.TriggerEvent
	11, // 21: loop.TriggerExecutable.UnregisterTrigger:output_type -> google.protobuf.Empty
	11, // 22: loop.CallbackExecutable.RegisterToWorkflow:output_type -> google.protobuf.Empty
	11, // 23: loop.CallbackExecutable.UnregisterFromWorkflow:output_type -> google.protobuf.Empty
	6,  // 24: loop.CallbackExecutable.Execute:output_type -> loop.CapabilityResponse
	19, // [19:25] is the sub-list for method output_type
	13, // [13:19] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_capabilities_pb_capabilities_proto_init() }
func file_capabilities_pb_capabilities_proto_init() {
	if File_capabilities_pb_capabilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capabilities_pb_capabilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterToWorkflowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_capabilities_pb_capabilities_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterFromWorkflowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_capabilities_pb_capabilities_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_capabilities_pb_capabilities_proto_goTypes,
		DependencyIndexes: file_capabilities_pb_capabilities_proto_depIdxs,
		EnumInfos:         file_capabilities_pb_capabilities_proto_enumTypes,
		MessageInfos:      file_capabilities_pb_capabilities_proto_msgTypes,
	}.Build()
	File_capabilities_pb_capabilities_proto = out.File
	file_capabilities_pb_capabilities_proto_rawDesc = nil
	file_capabilities_pb_capabilities_proto_goTypes = nil
	file_capabilities_pb_capabilities_proto_depIdxs = nil
}
