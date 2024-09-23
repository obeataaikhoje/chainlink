// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: pricegetter.proto

package ccippb

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
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

// FilterConfiguredTokensRequest is a request for which tokens of a list of addresses are configured and which aren't.. It is a gRPC adapter for the input arguments of
// [github.com/smartcontractkit/chainlink-common/chainlink-common/pkg/types/ccip/PriceGetter.FilterConfiguredTokens]
type FilterConfiguredTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"` // []Address
}

func (x *FilterConfiguredTokensRequest) Reset() {
	*x = FilterConfiguredTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricegetter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfiguredTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfiguredTokensRequest) ProtoMessage() {}

func (x *FilterConfiguredTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricegetter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfiguredTokensRequest.ProtoReflect.Descriptor instead.
func (*FilterConfiguredTokensRequest) Descriptor() ([]byte, []int) {
	return file_pricegetter_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfiguredTokensRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

// FilterConfiguredTokensResponse is a response for which tokens of a list of addresses are configured and which aren't. It is a gRPC adapter for the return values of
// [github.com/smartcontractkit/chainlink-common/chainlink-common/pkg/types/ccip/PriceGetter.FilterConfiguredTokens]
type FilterConfiguredTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configured   []string `protobuf:"bytes,1,rep,name=configured,proto3" json:"configured,omitempty"`     // []Address
	Unconfigured []string `protobuf:"bytes,2,rep,name=unconfigured,proto3" json:"unconfigured,omitempty"` // []Address
}

func (x *FilterConfiguredTokensResponse) Reset() {
	*x = FilterConfiguredTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricegetter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfiguredTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfiguredTokensResponse) ProtoMessage() {}

func (x *FilterConfiguredTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricegetter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfiguredTokensResponse.ProtoReflect.Descriptor instead.
func (*FilterConfiguredTokensResponse) Descriptor() ([]byte, []int) {
	return file_pricegetter_proto_rawDescGZIP(), []int{1}
}

func (x *FilterConfiguredTokensResponse) GetConfigured() []string {
	if x != nil {
		return x.Configured
	}
	return nil
}

func (x *FilterConfiguredTokensResponse) GetUnconfigured() []string {
	if x != nil {
		return x.Unconfigured
	}
	return nil
}

// TokenPricesRequest is a request for the price of a token in USD. It is a gRPC adapter for the input arguments of
// [github.com/smartcontractkit/chainlink-common/chainlink-common/pkg/types/ccip/PriceGetter.TokenPricesUSD]]
type TokenPricesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"` // []Address
}

func (x *TokenPricesRequest) Reset() {
	*x = TokenPricesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricegetter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPricesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPricesRequest) ProtoMessage() {}

func (x *TokenPricesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricegetter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPricesRequest.ProtoReflect.Descriptor instead.
func (*TokenPricesRequest) Descriptor() ([]byte, []int) {
	return file_pricegetter_proto_rawDescGZIP(), []int{2}
}

func (x *TokenPricesRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

// TokenPricesResponse is a response for the price of a token in USD. It is a gRPC adapter for the return values of
// [github.com/smartcontractkit/chainlink-common/chainlink-common/pkg/types/ccip/CommitStoreReader]
type TokenPricesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices map[string]*pb.BigInt `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[Address]price
}

func (x *TokenPricesResponse) Reset() {
	*x = TokenPricesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricegetter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPricesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPricesResponse) ProtoMessage() {}

func (x *TokenPricesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricegetter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPricesResponse.ProtoReflect.Descriptor instead.
func (*TokenPricesResponse) Descriptor() ([]byte, []int) {
	return file_pricegetter_proto_rawDescGZIP(), []int{3}
}

func (x *TokenPricesResponse) GetPrices() map[string]*pb.BigInt {
	if x != nil {
		return x.Prices
	}
	return nil
}

var File_pricegetter_proto protoreflect.FileDescriptor

var file_pricegetter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x69, 0x63, 0x65, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x1a, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x1d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22,
	0x64, 0x0a, 0x1e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63,
	0x63, 0x69, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x47, 0x0a, 0x0b, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xbd, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x47, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x87, 0x01, 0x0a, 0x16, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x34, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69,
	0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x55, 0x53, 0x44,
	0x12, 0x29, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63,
	0x63, 0x69, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x63, 0x69, 0x70, 0x3b, 0x63,
	0x63, 0x69, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pricegetter_proto_rawDescOnce sync.Once
	file_pricegetter_proto_rawDescData = file_pricegetter_proto_rawDesc
)

func file_pricegetter_proto_rawDescGZIP() []byte {
	file_pricegetter_proto_rawDescOnce.Do(func() {
		file_pricegetter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pricegetter_proto_rawDescData)
	})
	return file_pricegetter_proto_rawDescData
}

var file_pricegetter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pricegetter_proto_goTypes = []any{
	(*FilterConfiguredTokensRequest)(nil),  // 0: loop.internal.pb.ccip.FilterConfiguredTokensRequest
	(*FilterConfiguredTokensResponse)(nil), // 1: loop.internal.pb.ccip.FilterConfiguredTokensResponse
	(*TokenPricesRequest)(nil),             // 2: loop.internal.pb.ccip.TokenPricesRequest
	(*TokenPricesResponse)(nil),            // 3: loop.internal.pb.ccip.TokenPricesResponse
	nil,                                    // 4: loop.internal.pb.ccip.TokenPricesResponse.PricesEntry
	(*pb.BigInt)(nil),                      // 5: loop.BigInt
	(*emptypb.Empty)(nil),                  // 6: google.protobuf.Empty
}
var file_pricegetter_proto_depIdxs = []int32{
	4, // 0: loop.internal.pb.ccip.TokenPricesResponse.prices:type_name -> loop.internal.pb.ccip.TokenPricesResponse.PricesEntry
	5, // 1: loop.internal.pb.ccip.TokenPricesResponse.PricesEntry.value:type_name -> loop.BigInt
	0, // 2: loop.internal.pb.ccip.PriceGetter.FilterConfiguredTokens:input_type -> loop.internal.pb.ccip.FilterConfiguredTokensRequest
	2, // 3: loop.internal.pb.ccip.PriceGetter.TokenPricesUSD:input_type -> loop.internal.pb.ccip.TokenPricesRequest
	6, // 4: loop.internal.pb.ccip.PriceGetter.Close:input_type -> google.protobuf.Empty
	1, // 5: loop.internal.pb.ccip.PriceGetter.FilterConfiguredTokens:output_type -> loop.internal.pb.ccip.FilterConfiguredTokensResponse
	3, // 6: loop.internal.pb.ccip.PriceGetter.TokenPricesUSD:output_type -> loop.internal.pb.ccip.TokenPricesResponse
	6, // 7: loop.internal.pb.ccip.PriceGetter.Close:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pricegetter_proto_init() }
func file_pricegetter_proto_init() {
	if File_pricegetter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pricegetter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FilterConfiguredTokensRequest); i {
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
		file_pricegetter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FilterConfiguredTokensResponse); i {
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
		file_pricegetter_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TokenPricesRequest); i {
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
		file_pricegetter_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TokenPricesResponse); i {
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
			RawDescriptor: file_pricegetter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pricegetter_proto_goTypes,
		DependencyIndexes: file_pricegetter_proto_depIdxs,
		MessageInfos:      file_pricegetter_proto_msgTypes,
	}.Build()
	File_pricegetter_proto = out.File
	file_pricegetter_proto_rawDesc = nil
	file_pricegetter_proto_goTypes = nil
	file_pricegetter_proto_depIdxs = nil
}
