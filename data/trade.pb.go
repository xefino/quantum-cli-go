// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: protos/edge/data/trade.proto

package data

import (
	gopb "github.com/xefino/protobuf-gen-go/gopb"
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

// Action describes the type of request being made.
type TradeRequest_Action int32

const (
	TradeRequest_Deal    TradeRequest_Action = 0 // Place a trade order for an immediate execution with the specified parameters (market order)
	TradeRequest_Pending TradeRequest_Action = 1 // Place a trade order for execution under specified conditions (pending order)
	TradeRequest_SLTP    TradeRequest_Action = 2 // Modify the stop-loss and/or take-profit values of an opened position
	TradeRequest_Modify  TradeRequest_Action = 3 // Modify the parameters of a previously placed order
	TradeRequest_Remove  TradeRequest_Action = 4 // Delete a pending order, placed previously
	TradeRequest_CloseBy TradeRequest_Action = 5 // Close a position by placing an order for an opposite one
)

// Enum value maps for TradeRequest_Action.
var (
	TradeRequest_Action_name = map[int32]string{
		0: "Deal",
		1: "Pending",
		2: "SLTP",
		3: "Modify",
		4: "Remove",
		5: "CloseBy",
	}
	TradeRequest_Action_value = map[string]int32{
		"Deal":    0,
		"Pending": 1,
		"SLTP":    2,
		"Modify":  3,
		"Remove":  4,
		"CloseBy": 5,
	}
)

func (x TradeRequest_Action) Enum() *TradeRequest_Action {
	p := new(TradeRequest_Action)
	*p = x
	return p
}

func (x TradeRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edge_data_trade_proto_enumTypes[0].Descriptor()
}

func (TradeRequest_Action) Type() protoreflect.EnumType {
	return &file_protos_edge_data_trade_proto_enumTypes[0]
}

func (x TradeRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradeRequest_Action.Descriptor instead.
func (TradeRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_protos_edge_data_trade_proto_rawDescGZIP(), []int{0, 0}
}

// Interaction between the client terminal and a trade server for executing the order placing operation is performed by using the TradeRequest type, which contain all the fields necessary to perform trade deals.
type TradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action             TradeRequest_Action `protobuf:"varint,1,opt,name=action,proto3,enum=protos.edge.data.TradeRequest_Action" json:"action,omitempty"`                                     // Trade operation type
	OrderId            uint64              `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                                                              // Order ID. It is used for modifying pending orders.
	Symbol             string              `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`                                                                                // Symbol of the order. It is not necessary for order modification and position close operations.
	Volume             *Decimal            `protobuf:"bytes,4,opt,name=volume,proto3" json:"volume,omitempty"`                                                                                // Requested order volume in lots. Note that the real volume of a deal will depend on the order execution type.
	Price              *Decimal            `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`                                                                                  // Price, reaching which the order must be executed. Market orders of symbols, whose execution type is "Market Execution", of OrderAction.Deal type, do not require specification of price.
	StopLimit          *Decimal            `protobuf:"bytes,6,opt,name=stop_limit,json=stopLimit,proto3" json:"stop_limit,omitempty"`                                                         // The price value, at which the Limit pending order will be placed, when price reaches the price value (this condition is obligatory). Until then the pending order is not placed.
	StopLoss           *Decimal            `protobuf:"bytes,7,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`                                                            // Stop Loss price in case of the unfavorable price movement
	TakeProfit         *Decimal            `protobuf:"bytes,8,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`                                                      // Take Profit price in the case of the favorable price movement
	Deviation          *Decimal            `protobuf:"bytes,9,opt,name=deviation,proto3" json:"deviation,omitempty"`                                                                          // The maximal price deviation
	Type               Order_Type          `protobuf:"varint,10,opt,name=type,proto3,enum=protos.edge.data.Order_Type" json:"type,omitempty"`                                                 // Order type. Can be one of the OrderType enumeration values.
	FillType           FillPolicy          `protobuf:"varint,11,opt,name=fill_type,json=fillType,proto3,enum=protos.edge.data.FillPolicy" json:"fill_type,omitempty"`                         // Order execution type. Can be one of the enumeration FillPolicy values.
	ExpirationType     ExpirationPolicy    `protobuf:"varint,12,opt,name=expiration_type,json=expirationType,proto3,enum=protos.edge.data.ExpirationPolicy" json:"expiration_type,omitempty"` // Order expiration type. Can be one of the enumeration ExpirationPolicy values.
	Expiration         *gopb.UnixDuration  `protobuf:"bytes,13,opt,name=expiration,proto3" json:"expiration,omitempty"`                                                                       // Order expiration time (for orders of ExpirationPolicy.TimeSpecified type)
	Comment            string              `protobuf:"bytes,14,opt,name=comment,proto3" json:"comment,omitempty"`                                                                             // Order comment
	PositionId         uint64              `protobuf:"varint,15,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`                                                    // ID of a position. Should be filled in when a position is modified or closed to identify the position. As a rule it is equal to the ID of the order, based on which the position was opened.
	OppositePositionId uint64              `protobuf:"varint,16,opt,name=opposite_position_id,json=oppositePositionId,proto3" json:"opposite_position_id,omitempty"`                          // ID of an opposite position. Used when a position is closed by an opposite one open for the same symbol in the opposite direction.
}

func (x *TradeRequest) Reset() {
	*x = TradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_data_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeRequest) ProtoMessage() {}

func (x *TradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_data_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeRequest.ProtoReflect.Descriptor instead.
func (*TradeRequest) Descriptor() ([]byte, []int) {
	return file_protos_edge_data_trade_proto_rawDescGZIP(), []int{0}
}

func (x *TradeRequest) GetAction() TradeRequest_Action {
	if x != nil {
		return x.Action
	}
	return TradeRequest_Deal
}

func (x *TradeRequest) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *TradeRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TradeRequest) GetVolume() *Decimal {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *TradeRequest) GetPrice() *Decimal {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *TradeRequest) GetStopLimit() *Decimal {
	if x != nil {
		return x.StopLimit
	}
	return nil
}

func (x *TradeRequest) GetStopLoss() *Decimal {
	if x != nil {
		return x.StopLoss
	}
	return nil
}

func (x *TradeRequest) GetTakeProfit() *Decimal {
	if x != nil {
		return x.TakeProfit
	}
	return nil
}

func (x *TradeRequest) GetDeviation() *Decimal {
	if x != nil {
		return x.Deviation
	}
	return nil
}

func (x *TradeRequest) GetType() Order_Type {
	if x != nil {
		return x.Type
	}
	return Order_Buy
}

func (x *TradeRequest) GetFillType() FillPolicy {
	if x != nil {
		return x.FillType
	}
	return FillPolicy_FillOrKill
}

func (x *TradeRequest) GetExpirationType() ExpirationPolicy {
	if x != nil {
		return x.ExpirationType
	}
	return ExpirationPolicy_UntilCancelled
}

func (x *TradeRequest) GetExpiration() *gopb.UnixDuration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *TradeRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *TradeRequest) GetPositionId() uint64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *TradeRequest) GetOppositePositionId() uint64 {
	if x != nil {
		return x.OppositePositionId
	}
	return 0
}

var File_protos_edge_data_trade_proto protoreflect.FileDescriptor

var file_protos_edge_data_trade_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x06, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x31, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2f, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x09, 0x73,
	0x74, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70,
	0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73,
	0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x37, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3b, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x70, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6f, 0x70, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c,
	0x54, 0x50, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x79, 0x10, 0x05, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edge_data_trade_proto_rawDescOnce sync.Once
	file_protos_edge_data_trade_proto_rawDescData = file_protos_edge_data_trade_proto_rawDesc
)

func file_protos_edge_data_trade_proto_rawDescGZIP() []byte {
	file_protos_edge_data_trade_proto_rawDescOnce.Do(func() {
		file_protos_edge_data_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_data_trade_proto_rawDescData)
	})
	return file_protos_edge_data_trade_proto_rawDescData
}

var file_protos_edge_data_trade_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_edge_data_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_edge_data_trade_proto_goTypes = []interface{}{
	(TradeRequest_Action)(0),  // 0: protos.edge.data.TradeRequest.Action
	(*TradeRequest)(nil),      // 1: protos.edge.data.TradeRequest
	(*Decimal)(nil),           // 2: protos.edge.data.Decimal
	(Order_Type)(0),           // 3: protos.edge.data.Order.Type
	(FillPolicy)(0),           // 4: protos.edge.data.FillPolicy
	(ExpirationPolicy)(0),     // 5: protos.edge.data.ExpirationPolicy
	(*gopb.UnixDuration)(nil), // 6: protos.common.UnixDuration
}
var file_protos_edge_data_trade_proto_depIdxs = []int32{
	0,  // 0: protos.edge.data.TradeRequest.action:type_name -> protos.edge.data.TradeRequest.Action
	2,  // 1: protos.edge.data.TradeRequest.volume:type_name -> protos.edge.data.Decimal
	2,  // 2: protos.edge.data.TradeRequest.price:type_name -> protos.edge.data.Decimal
	2,  // 3: protos.edge.data.TradeRequest.stop_limit:type_name -> protos.edge.data.Decimal
	2,  // 4: protos.edge.data.TradeRequest.stop_loss:type_name -> protos.edge.data.Decimal
	2,  // 5: protos.edge.data.TradeRequest.take_profit:type_name -> protos.edge.data.Decimal
	2,  // 6: protos.edge.data.TradeRequest.deviation:type_name -> protos.edge.data.Decimal
	3,  // 7: protos.edge.data.TradeRequest.type:type_name -> protos.edge.data.Order.Type
	4,  // 8: protos.edge.data.TradeRequest.fill_type:type_name -> protos.edge.data.FillPolicy
	5,  // 9: protos.edge.data.TradeRequest.expiration_type:type_name -> protos.edge.data.ExpirationPolicy
	6,  // 10: protos.edge.data.TradeRequest.expiration:type_name -> protos.common.UnixDuration
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protos_edge_data_trade_proto_init() }
func file_protos_edge_data_trade_proto_init() {
	if File_protos_edge_data_trade_proto != nil {
		return
	}
	file_protos_edge_data_common_proto_init()
	file_protos_edge_data_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_data_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeRequest); i {
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
			RawDescriptor: file_protos_edge_data_trade_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_edge_data_trade_proto_goTypes,
		DependencyIndexes: file_protos_edge_data_trade_proto_depIdxs,
		EnumInfos:         file_protos_edge_data_trade_proto_enumTypes,
		MessageInfos:      file_protos_edge_data_trade_proto_msgTypes,
	}.Build()
	File_protos_edge_data_trade_proto = out.File
	file_protos_edge_data_trade_proto_rawDesc = nil
	file_protos_edge_data_trade_proto_goTypes = nil
	file_protos_edge_data_trade_proto_depIdxs = nil
}
