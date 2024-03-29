// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: protos/edge/entry/v1/service.proto

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

// EntryRequest describes the information that may be used to inform whether or not a position should
// be entered. No modifications made by the endpoint will be reflected in the data.
type EntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *data.StrategyContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`                                      // The strategy context that contains information associated with the currently-running strategy
	Allowed data.AllowedTrades    `protobuf:"varint,2,opt,name=allowed,proto3,enum=protos.edge.data.AllowedTrades" json:"allowed,omitempty"` // Which trades should be allowed by the filter
}

func (x *EntryRequest) Reset() {
	*x = EntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_entry_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryRequest) ProtoMessage() {}

func (x *EntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_entry_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryRequest.ProtoReflect.Descriptor instead.
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_edge_entry_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *EntryRequest) GetContext() *data.StrategyContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *EntryRequest) GetAllowed() data.AllowedTrades {
	if x != nil {
		return x.Allowed
	}
	return data.AllowedTrades(0)
}

// EntryResult contains the information that should be returned after an entry decision has been reached
// and will be used to inform the strategy runner about which action should be taken next.
type EntryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldEnter bool                 `protobuf:"varint,1,opt,name=should_enter,json=shouldEnter,proto3" json:"should_enter,omitempty"` // Whether or not the position should be entered. True, to enter. False to ignore.
	Order       []*data.TradeRequest `protobuf:"bytes,2,rep,name=order,proto3" json:"order,omitempty"`                                 // The order that represents the position that should be entered. This value will be empty if should_enter is false. Furthermore, the order may be modified by other modules before being submitted.
}

func (x *EntryResult) Reset() {
	*x = EntryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_entry_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryResult) ProtoMessage() {}

func (x *EntryResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_entry_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryResult.ProtoReflect.Descriptor instead.
func (*EntryResult) Descriptor() ([]byte, []int) {
	return file_protos_edge_entry_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *EntryResult) GetShouldEnter() bool {
	if x != nil {
		return x.ShouldEnter
	}
	return false
}

func (x *EntryResult) GetOrder() []*data.TradeRequest {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_protos_edge_entry_v1_service_proto protoreflect.FileDescriptor

var file_protos_edge_entry_v1_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x66, 0x0a,
	0x0b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xa5, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69,
	0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_edge_entry_v1_service_proto_rawDescOnce sync.Once
	file_protos_edge_entry_v1_service_proto_rawDescData = file_protos_edge_entry_v1_service_proto_rawDesc
)

func file_protos_edge_entry_v1_service_proto_rawDescGZIP() []byte {
	file_protos_edge_entry_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_edge_entry_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_entry_v1_service_proto_rawDescData)
	})
	return file_protos_edge_entry_v1_service_proto_rawDescData
}

var file_protos_edge_entry_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_edge_entry_v1_service_proto_goTypes = []interface{}{
	(*EntryRequest)(nil),         // 0: protos.edge.entry.v1.EntryRequest
	(*EntryResult)(nil),          // 1: protos.edge.entry.v1.EntryResult
	(*data.StrategyContext)(nil), // 2: protos.edge.data.StrategyContext
	(data.AllowedTrades)(0),      // 3: protos.edge.data.AllowedTrades
	(*data.TradeRequest)(nil),    // 4: protos.edge.data.TradeRequest
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
	(*data.NameResult)(nil),      // 6: protos.edge.data.NameResult
}
var file_protos_edge_entry_v1_service_proto_depIdxs = []int32{
	2, // 0: protos.edge.entry.v1.EntryRequest.context:type_name -> protos.edge.data.StrategyContext
	3, // 1: protos.edge.entry.v1.EntryRequest.allowed:type_name -> protos.edge.data.AllowedTrades
	4, // 2: protos.edge.entry.v1.EntryResult.order:type_name -> protos.edge.data.TradeRequest
	5, // 3: protos.edge.entry.v1.EntryService.GetName:input_type -> google.protobuf.Empty
	0, // 4: protos.edge.entry.v1.EntryService.ShouldEnter:input_type -> protos.edge.entry.v1.EntryRequest
	6, // 5: protos.edge.entry.v1.EntryService.GetName:output_type -> protos.edge.data.NameResult
	1, // 6: protos.edge.entry.v1.EntryService.ShouldEnter:output_type -> protos.edge.entry.v1.EntryResult
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_edge_entry_v1_service_proto_init() }
func file_protos_edge_entry_v1_service_proto_init() {
	if File_protos_edge_entry_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_entry_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryRequest); i {
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
		file_protos_edge_entry_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryResult); i {
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
			RawDescriptor: file_protos_edge_entry_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_edge_entry_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_edge_entry_v1_service_proto_depIdxs,
		MessageInfos:      file_protos_edge_entry_v1_service_proto_msgTypes,
	}.Build()
	File_protos_edge_entry_v1_service_proto = out.File
	file_protos_edge_entry_v1_service_proto_rawDesc = nil
	file_protos_edge_entry_v1_service_proto_goTypes = nil
	file_protos_edge_entry_v1_service_proto_depIdxs = nil
}
