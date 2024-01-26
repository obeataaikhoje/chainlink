// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: error_log.proto

package pb

import (
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

// SaveErrorRequest has arguments for [github.com/smartcontractkit/chainlink-common/pkg/loop.ErrorLog.SaveErrorRequest].
type SaveErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SaveErrorRequest) Reset() {
	*x = SaveErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveErrorRequest) ProtoMessage() {}

func (x *SaveErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_error_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveErrorRequest.ProtoReflect.Descriptor instead.
func (*SaveErrorRequest) Descriptor() ([]byte, []int) {
	return file_error_log_proto_rawDescGZIP(), []int{0}
}

func (x *SaveErrorRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_error_log_proto protoreflect.FileDescriptor

var file_error_log_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x49, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x3d,
	0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_log_proto_rawDescOnce sync.Once
	file_error_log_proto_rawDescData = file_error_log_proto_rawDesc
)

func file_error_log_proto_rawDescGZIP() []byte {
	file_error_log_proto_rawDescOnce.Do(func() {
		file_error_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_log_proto_rawDescData)
	})
	return file_error_log_proto_rawDescData
}

var file_error_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_error_log_proto_goTypes = []interface{}{
	(*SaveErrorRequest)(nil), // 0: loop.SaveErrorRequest
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_error_log_proto_depIdxs = []int32{
	0, // 0: loop.ErrorLog.SaveError:input_type -> loop.SaveErrorRequest
	1, // 1: loop.ErrorLog.SaveError:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_log_proto_init() }
func file_error_log_proto_init() {
	if File_error_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveErrorRequest); i {
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
			RawDescriptor: file_error_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_error_log_proto_goTypes,
		DependencyIndexes: file_error_log_proto_depIdxs,
		MessageInfos:      file_error_log_proto_msgTypes,
	}.Build()
	File_error_log_proto = out.File
	file_error_log_proto_rawDesc = nil
	file_error_log_proto_goTypes = nil
	file_error_log_proto_depIdxs = nil
}
