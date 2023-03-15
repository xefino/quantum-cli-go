// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.20.3
// source: protos/edge/money/v1/service.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	data "github.com/xefino/quantum-cli-go/data"
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

// ModifyOrderRequest describes the information that may be used to assign an amount of capital to a
// position that will be entered or exited. No modifications made by the endpoint will be reflected
// in the data.
type ModifyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context       *data.StrategyContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`                                  // The strategy context that contains information associated with the currently-running strategy
	OriginalTrade *data.TradeRequest    `protobuf:"bytes,2,opt,name=original_trade,json=originalTrade,proto3" json:"original_trade,omitempty"` // The trade received by the service. Changes to this value will not be reflected in the sender
}

func (x *ModifyOrderRequest) Reset() {
	*x = ModifyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_money_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyOrderRequest) ProtoMessage() {}

func (x *ModifyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_money_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyOrderRequest.ProtoReflect.Descriptor instead.
func (*ModifyOrderRequest) Descriptor() ([]byte, []int) {
	return file_protos_edge_money_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ModifyOrderRequest) GetContext() *data.StrategyContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *ModifyOrderRequest) GetOriginalTrade() *data.TradeRequest {
	if x != nil {
		return x.OriginalTrade
	}
	return nil
}

// ModifyOrderResult describes the information that will be returned from the service. This will contain
// the updated trade request that will be submitted to the trade system on behalf of the strategy.
type ModifyOrderResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedTrade *data.TradeRequest `protobuf:"bytes,1,opt,name=updated_trade,json=updatedTrade,proto3" json:"updated_trade,omitempty"` // The trade request, updated by the service
}

func (x *ModifyOrderResult) Reset() {
	*x = ModifyOrderResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_money_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyOrderResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyOrderResult) ProtoMessage() {}

func (x *ModifyOrderResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_money_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyOrderResult.ProtoReflect.Descriptor instead.
func (*ModifyOrderResult) Descriptor() ([]byte, []int) {
	return file_protos_edge_money_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ModifyOrderResult) GetUpdatedTrade() *data.TradeRequest {
	if x != nil {
		return x.UpdatedTrade
	}
	return nil
}

var File_protos_edge_money_v1_service_proto protoreflect.FileDescriptor

var file_protos_edge_money_v1_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x45, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x22,
	0x58, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x32, 0xbb, 0x01, 0x0a, 0x16, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x60, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edge_money_v1_service_proto_rawDescOnce sync.Once
	file_protos_edge_money_v1_service_proto_rawDescData = file_protos_edge_money_v1_service_proto_rawDesc
)

func file_protos_edge_money_v1_service_proto_rawDescGZIP() []byte {
	file_protos_edge_money_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_edge_money_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_money_v1_service_proto_rawDescData)
	})
	return file_protos_edge_money_v1_service_proto_rawDescData
}

var file_protos_edge_money_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_edge_money_v1_service_proto_goTypes = []interface{}{
	(*ModifyOrderRequest)(nil),   // 0: protos.edge.money.v1.ModifyOrderRequest
	(*ModifyOrderResult)(nil),    // 1: protos.edge.money.v1.ModifyOrderResult
	(*data.StrategyContext)(nil), // 2: protos.edge.data.StrategyContext
	(*data.TradeRequest)(nil),    // 3: protos.edge.data.TradeRequest
	(*empty.Empty)(nil),          // 4: google.protobuf.Empty
	(*data.NameResult)(nil),      // 5: protos.edge.data.NameResult
}
var file_protos_edge_money_v1_service_proto_depIdxs = []int32{
	2, // 0: protos.edge.money.v1.ModifyOrderRequest.context:type_name -> protos.edge.data.StrategyContext
	3, // 1: protos.edge.money.v1.ModifyOrderRequest.original_trade:type_name -> protos.edge.data.TradeRequest
	3, // 2: protos.edge.money.v1.ModifyOrderResult.updated_trade:type_name -> protos.edge.data.TradeRequest
	4, // 3: protos.edge.money.v1.MoneyManagementService.GetName:input_type -> google.protobuf.Empty
	0, // 4: protos.edge.money.v1.MoneyManagementService.ModifyOrder:input_type -> protos.edge.money.v1.ModifyOrderRequest
	5, // 5: protos.edge.money.v1.MoneyManagementService.GetName:output_type -> protos.edge.data.NameResult
	1, // 6: protos.edge.money.v1.MoneyManagementService.ModifyOrder:output_type -> protos.edge.money.v1.ModifyOrderResult
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_edge_money_v1_service_proto_init() }
func file_protos_edge_money_v1_service_proto_init() {
	if File_protos_edge_money_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_money_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyOrderRequest); i {
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
		file_protos_edge_money_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyOrderResult); i {
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
			RawDescriptor: file_protos_edge_money_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_edge_money_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_edge_money_v1_service_proto_depIdxs,
		MessageInfos:      file_protos_edge_money_v1_service_proto_msgTypes,
	}.Build()
	File_protos_edge_money_v1_service_proto = out.File
	file_protos_edge_money_v1_service_proto_rawDesc = nil
	file_protos_edge_money_v1_service_proto_goTypes = nil
	file_protos_edge_money_v1_service_proto_depIdxs = nil
}
