// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: example.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Used for testing
type TestCustomMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolVal   bool    `protobuf:"varint,1,opt,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	IntVal    int64   `protobuf:"varint,2,opt,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	FloatVal  float32 `protobuf:"fixed32,3,opt,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	StringVal string  `protobuf:"bytes,4,opt,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	BytesVal  []byte  `protobuf:"bytes,5,opt,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
}

func (x *TestCustomMessage) Reset() {
	*x = TestCustomMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCustomMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCustomMessage) ProtoMessage() {}

func (x *TestCustomMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCustomMessage.ProtoReflect.Descriptor instead.
func (*TestCustomMessage) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *TestCustomMessage) GetBoolVal() bool {
	if x != nil {
		return x.BoolVal
	}
	return false
}

func (x *TestCustomMessage) GetIntVal() int64 {
	if x != nil {
		return x.IntVal
	}
	return 0
}

func (x *TestCustomMessage) GetFloatVal() float32 {
	if x != nil {
		return x.FloatVal
	}
	return 0
}

func (x *TestCustomMessage) GetStringVal() string {
	if x != nil {
		return x.StringVal
	}
	return ""
}

func (x *TestCustomMessage) GetBytesVal() []byte {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

type KeystoneCustomMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg                   string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	CapabilityID          string `protobuf:"bytes,2,opt,name=capabilityID,proto3" json:"capabilityID,omitempty"`
	CapabilityExecutionID string `protobuf:"bytes,3,opt,name=capabilityExecutionID,proto3" json:"capabilityExecutionID,omitempty"`
	TriggerID             string `protobuf:"bytes,4,opt,name=triggerID,proto3" json:"triggerID,omitempty"`
	WorkflowID            string `protobuf:"bytes,5,opt,name=workflowID,proto3" json:"workflowID,omitempty"`
	WorkflowExecutionID   string `protobuf:"bytes,6,opt,name=workflowExecutionID,proto3" json:"workflowExecutionID,omitempty"`
	WorkflowOwner         string `protobuf:"bytes,7,opt,name=workflowOwner,proto3" json:"workflowOwner,omitempty"`
	StepID                string `protobuf:"bytes,8,opt,name=stepID,proto3" json:"stepID,omitempty"`
	StepRef               string `protobuf:"bytes,9,opt,name=stepRef,proto3" json:"stepRef,omitempty"`
}

func (x *KeystoneCustomMessage) Reset() {
	*x = KeystoneCustomMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeystoneCustomMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeystoneCustomMessage) ProtoMessage() {}

func (x *KeystoneCustomMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeystoneCustomMessage.ProtoReflect.Descriptor instead.
func (*KeystoneCustomMessage) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *KeystoneCustomMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *KeystoneCustomMessage) GetCapabilityID() string {
	if x != nil {
		return x.CapabilityID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetCapabilityExecutionID() string {
	if x != nil {
		return x.CapabilityExecutionID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetTriggerID() string {
	if x != nil {
		return x.TriggerID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetWorkflowID() string {
	if x != nil {
		return x.WorkflowID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetWorkflowExecutionID() string {
	if x != nil {
		return x.WorkflowExecutionID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetWorkflowOwner() string {
	if x != nil {
		return x.WorkflowOwner
	}
	return ""
}

func (x *KeystoneCustomMessage) GetStepID() string {
	if x != nil {
		return x.StepID
	}
	return ""
}

func (x *KeystoneCustomMessage) GetStepRef() string {
	if x != nil {
		return x.StepRef
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x22, 0xcb, 0x02, 0x0a, 0x15, 0x4b, 0x65, 0x79, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x65, 0x70, 0x52, 0x65, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x65,
	0x70, 0x52, 0x65, 0x66, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_goTypes = []any{
	(*TestCustomMessage)(nil),     // 0: pb.TestCustomMessage
	(*KeystoneCustomMessage)(nil), // 1: pb.KeystoneCustomMessage
}
var file_example_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TestCustomMessage); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeystoneCustomMessage); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
