// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: protos/edge/filtering/v1/service.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	gopb "github.com/xefino/protobuf-gen-go/gopb"
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

// FilterResult describes whether or not the operation associated with the strategy context should be
//  paused or whether it may continue to run.
type FilterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stop bool `protobuf:"varint,1,opt,name=stop,proto3" json:"stop,omitempty"` // Whether or not the functionality should be stopped (true) or
	// whether it may continue to run (false)
	Until *gopb.UnixDuration `protobuf:"bytes,2,opt,name=until,proto3" json:"until,omitempty"` // A duration expressing how long the filter should wait until it
}

func (x *FilterResult) Reset() {
	*x = FilterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_filtering_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterResult) ProtoMessage() {}

func (x *FilterResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_filtering_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterResult.ProtoReflect.Descriptor instead.
func (*FilterResult) Descriptor() ([]byte, []int) {
	return file_protos_edge_filtering_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *FilterResult) GetStop() bool {
	if x != nil {
		return x.Stop
	}
	return false
}

func (x *FilterResult) GetUntil() *gopb.UnixDuration {
	if x != nil {
		return x.Until
	}
	return nil
}

var File_protos_edge_filtering_v1_service_proto protoreflect.FileDescriptor

var file_protos_edge_filtering_v1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x31, 0x0a, 0x05,
	0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x78,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x32,
	0xab, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69,
	0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edge_filtering_v1_service_proto_rawDescOnce sync.Once
	file_protos_edge_filtering_v1_service_proto_rawDescData = file_protos_edge_filtering_v1_service_proto_rawDesc
)

func file_protos_edge_filtering_v1_service_proto_rawDescGZIP() []byte {
	file_protos_edge_filtering_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_edge_filtering_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_filtering_v1_service_proto_rawDescData)
	})
	return file_protos_edge_filtering_v1_service_proto_rawDescData
}

var file_protos_edge_filtering_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_edge_filtering_v1_service_proto_goTypes = []interface{}{
	(*FilterResult)(nil),         // 0: protos.edge.filtering.v1.FilterResult
	(*gopb.UnixDuration)(nil),    // 1: protos.common.UnixDuration
	(*empty.Empty)(nil),          // 2: google.protobuf.Empty
	(*data.StrategyContext)(nil), // 3: protos.edge.data.StrategyContext
	(*data.NameResult)(nil),      // 4: protos.edge.data.NameResult
}
var file_protos_edge_filtering_v1_service_proto_depIdxs = []int32{
	1, // 0: protos.edge.filtering.v1.FilterResult.until:type_name -> protos.common.UnixDuration
	2, // 1: protos.edge.filtering.v1.FilterService.GetName:input_type -> google.protobuf.Empty
	3, // 2: protos.edge.filtering.v1.FilterService.ShouldFilter:input_type -> protos.edge.data.StrategyContext
	4, // 3: protos.edge.filtering.v1.FilterService.GetName:output_type -> protos.edge.data.NameResult
	0, // 4: protos.edge.filtering.v1.FilterService.ShouldFilter:output_type -> protos.edge.filtering.v1.FilterResult
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_edge_filtering_v1_service_proto_init() }
func file_protos_edge_filtering_v1_service_proto_init() {
	if File_protos_edge_filtering_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_filtering_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterResult); i {
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
			RawDescriptor: file_protos_edge_filtering_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_edge_filtering_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_edge_filtering_v1_service_proto_depIdxs,
		MessageInfos:      file_protos_edge_filtering_v1_service_proto_msgTypes,
	}.Build()
	File_protos_edge_filtering_v1_service_proto = out.File
	file_protos_edge_filtering_v1_service_proto_rawDesc = nil
	file_protos_edge_filtering_v1_service_proto_goTypes = nil
	file_protos_edge_filtering_v1_service_proto_depIdxs = nil
}
