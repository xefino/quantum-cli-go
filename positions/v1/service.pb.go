// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: protos/edge/positions/v1/service.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// PositionCalculationRequest contains the data necessary to calculate the values that should represent a shared
// state between the positions that may be modified. No modifications made by the endpoint will be reflected in the data.
type PositionCalculationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *data.StrategyContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"` // The strategy context that contains information associated with the currently-running strategy
}

func (x *PositionCalculationRequest) Reset() {
	*x = PositionCalculationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_positions_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCalculationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCalculationRequest) ProtoMessage() {}

func (x *PositionCalculationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_positions_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionCalculationRequest.ProtoReflect.Descriptor instead.
func (*PositionCalculationRequest) Descriptor() ([]byte, []int) {
	return file_protos_edge_positions_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PositionCalculationRequest) GetContext() *data.StrategyContext {
	if x != nil {
		return x.Context
	}
	return nil
}

// PositionCalculationResponse contains the results of the calculation, and will be used to inform the modification
// of every position.
type PositionCalculationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results *_struct.Struct `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"` // The results of the calculation
}

func (x *PositionCalculationResponse) Reset() {
	*x = PositionCalculationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_positions_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCalculationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCalculationResponse) ProtoMessage() {}

func (x *PositionCalculationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_positions_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionCalculationResponse.ProtoReflect.Descriptor instead.
func (*PositionCalculationResponse) Descriptor() ([]byte, []int) {
	return file_protos_edge_positions_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *PositionCalculationResponse) GetResults() *_struct.Struct {
	if x != nil {
		return x.Results
	}
	return nil
}

// PositionModificationRequest contains the information that may be used to modify a position contained in the request.
// No modifications made by the endpoint will be reflected in the data.
type PositionModificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context          *data.StrategyContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`                                           // The strategy context that contains information associated with the currently-running strategy
	CalculatedFields *_struct.Struct       `protobuf:"bytes,2,opt,name=calculated_fields,json=calculatedFields,proto3" json:"calculated_fields,omitempty"` // The results of the calculation step, to be submitted along with this position
	Position         *data.Position        `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`                                         // The ID of the position that may be exited
}

func (x *PositionModificationRequest) Reset() {
	*x = PositionModificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_positions_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionModificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionModificationRequest) ProtoMessage() {}

func (x *PositionModificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_positions_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionModificationRequest.ProtoReflect.Descriptor instead.
func (*PositionModificationRequest) Descriptor() ([]byte, []int) {
	return file_protos_edge_positions_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *PositionModificationRequest) GetContext() *data.StrategyContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *PositionModificationRequest) GetCalculatedFields() *_struct.Struct {
	if x != nil {
		return x.CalculatedFields
	}
	return nil
}

func (x *PositionModificationRequest) GetPosition() *data.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

// PositionModificationResponse describes the information that will be returned from the service. This will contain
// the trade request that will be submitted to the trade system on behalf of the strategy.
type PositionModificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldModify bool               `protobuf:"varint,1,opt,name=should_modify,json=shouldModify,proto3" json:"should_modify,omitempty"` // Whether or not the position should be modified or cancelled
	Order        *data.TradeRequest `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`                                    // The trade request that represents the position that should be modified or cancelled.
}

func (x *PositionModificationResponse) Reset() {
	*x = PositionModificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_positions_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionModificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionModificationResponse) ProtoMessage() {}

func (x *PositionModificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_positions_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionModificationResponse.ProtoReflect.Descriptor instead.
func (*PositionModificationResponse) Descriptor() ([]byte, []int) {
	return file_protos_edge_positions_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *PositionModificationResponse) GetShouldModify() bool {
	if x != nil {
		return x.ShouldModify
	}
	return false
}

func (x *PositionModificationResponse) GetOrder() *data.TradeRequest {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_protos_edge_positions_v1_service_proto protoreflect.FileDescriptor

var file_protos_edge_positions_v1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1a, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x50, 0x0a, 0x1b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x1b, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x1c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x9b,
	0x03, 0x0a, 0x19, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x42, 0x0a,
	0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x78, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x34,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x0e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e,
	0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edge_positions_v1_service_proto_rawDescOnce sync.Once
	file_protos_edge_positions_v1_service_proto_rawDescData = file_protos_edge_positions_v1_service_proto_rawDesc
)

func file_protos_edge_positions_v1_service_proto_rawDescGZIP() []byte {
	file_protos_edge_positions_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_edge_positions_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_positions_v1_service_proto_rawDescData)
	})
	return file_protos_edge_positions_v1_service_proto_rawDescData
}

var file_protos_edge_positions_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_edge_positions_v1_service_proto_goTypes = []interface{}{
	(*PositionCalculationRequest)(nil),   // 0: protos.edge.positions.v1.PositionCalculationRequest
	(*PositionCalculationResponse)(nil),  // 1: protos.edge.positions.v1.PositionCalculationResponse
	(*PositionModificationRequest)(nil),  // 2: protos.edge.positions.v1.PositionModificationRequest
	(*PositionModificationResponse)(nil), // 3: protos.edge.positions.v1.PositionModificationResponse
	(*data.StrategyContext)(nil),         // 4: protos.edge.data.StrategyContext
	(*_struct.Struct)(nil),               // 5: google.protobuf.Struct
	(*data.Position)(nil),                // 6: protos.edge.data.Position
	(*data.TradeRequest)(nil),            // 7: protos.edge.data.TradeRequest
	(*empty.Empty)(nil),                  // 8: google.protobuf.Empty
	(*data.NameResult)(nil),              // 9: protos.edge.data.NameResult
	(*data.EnabledResult)(nil),           // 10: protos.edge.data.EnabledResult
}
var file_protos_edge_positions_v1_service_proto_depIdxs = []int32{
	4,  // 0: protos.edge.positions.v1.PositionCalculationRequest.context:type_name -> protos.edge.data.StrategyContext
	5,  // 1: protos.edge.positions.v1.PositionCalculationResponse.results:type_name -> google.protobuf.Struct
	4,  // 2: protos.edge.positions.v1.PositionModificationRequest.context:type_name -> protos.edge.data.StrategyContext
	5,  // 3: protos.edge.positions.v1.PositionModificationRequest.calculated_fields:type_name -> google.protobuf.Struct
	6,  // 4: protos.edge.positions.v1.PositionModificationRequest.position:type_name -> protos.edge.data.Position
	7,  // 5: protos.edge.positions.v1.PositionModificationResponse.order:type_name -> protos.edge.data.TradeRequest
	8,  // 6: protos.edge.positions.v1.PositionManagementService.GetName:input_type -> google.protobuf.Empty
	8,  // 7: protos.edge.positions.v1.PositionManagementService.Enabled:input_type -> google.protobuf.Empty
	0,  // 8: protos.edge.positions.v1.PositionManagementService.Calculate:input_type -> protos.edge.positions.v1.PositionCalculationRequest
	2,  // 9: protos.edge.positions.v1.PositionManagementService.ModifyPosition:input_type -> protos.edge.positions.v1.PositionModificationRequest
	9,  // 10: protos.edge.positions.v1.PositionManagementService.GetName:output_type -> protos.edge.data.NameResult
	10, // 11: protos.edge.positions.v1.PositionManagementService.Enabled:output_type -> protos.edge.data.EnabledResult
	1,  // 12: protos.edge.positions.v1.PositionManagementService.Calculate:output_type -> protos.edge.positions.v1.PositionCalculationResponse
	3,  // 13: protos.edge.positions.v1.PositionManagementService.ModifyPosition:output_type -> protos.edge.positions.v1.PositionModificationResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_protos_edge_positions_v1_service_proto_init() }
func file_protos_edge_positions_v1_service_proto_init() {
	if File_protos_edge_positions_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_positions_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionCalculationRequest); i {
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
		file_protos_edge_positions_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionCalculationResponse); i {
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
		file_protos_edge_positions_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionModificationRequest); i {
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
		file_protos_edge_positions_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionModificationResponse); i {
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
			RawDescriptor: file_protos_edge_positions_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_edge_positions_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_edge_positions_v1_service_proto_depIdxs,
		MessageInfos:      file_protos_edge_positions_v1_service_proto_msgTypes,
	}.Build()
	File_protos_edge_positions_v1_service_proto = out.File
	file_protos_edge_positions_v1_service_proto_rawDesc = nil
	file_protos_edge_positions_v1_service_proto_goTypes = nil
	file_protos_edge_positions_v1_service_proto_depIdxs = nil
}