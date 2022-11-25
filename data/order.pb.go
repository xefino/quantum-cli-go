// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: protos/edge/data/order.proto

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

// Type describes the order type, which is required by some trade request operations
type Order_Type int32

const (
	Order_Buy           Order_Type = 0 // Market Buy order
	Order_Sell          Order_Type = 1 // Market Sell order
	Order_BuyLimit      Order_Type = 2 // Buy Limit pending order
	Order_SellLimit     Order_Type = 3 // Sell Limit pending order
	Order_BuyStop       Order_Type = 4 // Buy Stop pending order
	Order_SellStop      Order_Type = 5 // Sell Stop pending order
	Order_BuyStopLimit  Order_Type = 6 // Upon reaching the order price, a pending buy-limit order is placed at the StopLimit price
	Order_SellStopLimit Order_Type = 7 // Upon reaching the order price, a pending sell-limit order is placed at the StopLimit price
	Order_ClosedBy      Order_Type = 8 // Order to close a position by an opposite one
)

// Enum value maps for Order_Type.
var (
	Order_Type_name = map[int32]string{
		0: "Buy",
		1: "Sell",
		2: "BuyLimit",
		3: "SellLimit",
		4: "BuyStop",
		5: "SellStop",
		6: "BuyStopLimit",
		7: "SellStopLimit",
		8: "ClosedBy",
	}
	Order_Type_value = map[string]int32{
		"Buy":           0,
		"Sell":          1,
		"BuyLimit":      2,
		"SellLimit":     3,
		"BuyStop":       4,
		"SellStop":      5,
		"BuyStopLimit":  6,
		"SellStopLimit": 7,
		"ClosedBy":      8,
	}
)

func (x Order_Type) Enum() *Order_Type {
	p := new(Order_Type)
	*p = x
	return p
}

func (x Order_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edge_data_order_proto_enumTypes[0].Descriptor()
}

func (Order_Type) Type() protoreflect.EnumType {
	return &file_protos_edge_data_order_proto_enumTypes[0]
}

func (x Order_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Type.Descriptor instead.
func (Order_Type) EnumDescriptor() ([]byte, []int) {
	return file_protos_edge_data_order_proto_rawDescGZIP(), []int{0, 0}
}

// Status describes the current state of the order it is associated with
type Order_Status int32

const (
	Order_Started       Order_Status = 0 // Order checked, but not yet accepted by broker
	Order_Placed        Order_Status = 1 // Order accepted
	Order_Cancelled     Order_Status = 2 // Order canceled by client
	Order_Partial       Order_Status = 3 // Order partially executed
	Order_Filled        Order_Status = 4 // Order fully executed
	Order_Rejected      Order_Status = 5 // Order rejected
	Order_Expired       Order_Status = 6 // Order expired
	Order_RequestAdd    Order_Status = 7 // Order is being registered (placing to the trading system)
	Order_RequestModify Order_Status = 8 // Order is being modified (changing its parameters)
	Order_RequestCancel Order_Status = 9 // Order is being deleted (deleting from the trading system)
)

// Enum value maps for Order_Status.
var (
	Order_Status_name = map[int32]string{
		0: "Started",
		1: "Placed",
		2: "Cancelled",
		3: "Partial",
		4: "Filled",
		5: "Rejected",
		6: "Expired",
		7: "RequestAdd",
		8: "RequestModify",
		9: "RequestCancel",
	}
	Order_Status_value = map[string]int32{
		"Started":       0,
		"Placed":        1,
		"Cancelled":     2,
		"Partial":       3,
		"Filled":        4,
		"Rejected":      5,
		"Expired":       6,
		"RequestAdd":    7,
		"RequestModify": 8,
		"RequestCancel": 9,
	}
)

func (x Order_Status) Enum() *Order_Status {
	p := new(Order_Status)
	*p = x
	return p
}

func (x Order_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edge_data_order_proto_enumTypes[1].Descriptor()
}

func (Order_Status) Type() protoreflect.EnumType {
	return &file_protos_edge_data_order_proto_enumTypes[1]
}

func (x Order_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Status.Descriptor instead.
func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return file_protos_edge_data_order_proto_rawDescGZIP(), []int{0, 1}
}

// Reason describes the originator of the order
type Order_Reason int32

const (
	Order_Client     Order_Reason = 0 // The order was placed from a desktop terminal
	Order_Mobile     Order_Reason = 1 // The order was placed from a mobile application
	Order_Web        Order_Reason = 2 // The order was placed from a web platform
	Order_Strategy   Order_Reason = 3 // The order was placed from a strategy, i.e. by an Expert Advisor or a script
	Order_StopLoss   Order_Reason = 4 // The order was placed as a result of Stop Loss activation
	Order_TakeProfit Order_Reason = 5 // The order was placed as a result of Take Profit activation
	Order_StopOut    Order_Reason = 6 // The order was placed as a result of the Stop Out event
)

// Enum value maps for Order_Reason.
var (
	Order_Reason_name = map[int32]string{
		0: "Client",
		1: "Mobile",
		2: "Web",
		3: "Strategy",
		4: "StopLoss",
		5: "TakeProfit",
		6: "StopOut",
	}
	Order_Reason_value = map[string]int32{
		"Client":     0,
		"Mobile":     1,
		"Web":        2,
		"Strategy":   3,
		"StopLoss":   4,
		"TakeProfit": 5,
		"StopOut":    6,
	}
)

func (x Order_Reason) Enum() *Order_Reason {
	p := new(Order_Reason)
	*p = x
	return p
}

func (x Order_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edge_data_order_proto_enumTypes[2].Descriptor()
}

func (Order_Reason) Type() protoreflect.EnumType {
	return &file_protos_edge_data_order_proto_enumTypes[2]
}

func (x Order_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Reason.Descriptor instead.
func (Order_Reason) EnumDescriptor() ([]byte, []int) {
	return file_protos_edge_data_order_proto_rawDescGZIP(), []int{0, 2}
}

// Order contains the data after a trade request is sent to the trading system. Some of this data is generated from the
// trade request and some is populated by the trading system. A successfully placed order always results in a deal
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                                       // Order ticket. Unique number assigned to each order
	Symbol         string              `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`                                                                                // Symbol of the order
	ExternalId     string              `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`                                                      // Order identifier in an external trading system (on the Exchange)
	SetupTime      *gopb.UnixTimestamp `protobuf:"bytes,4,opt,name=setup_time,json=setupTime,proto3" json:"setup_time,omitempty"`                                                         // Order setup time
	Type           Order_Type          `protobuf:"varint,5,opt,name=type,proto3,enum=protos.edge.data.Order_Type" json:"type,omitempty"`                                                  // Order type
	Status         Order_Status        `protobuf:"varint,6,opt,name=status,proto3,enum=protos.edge.data.Order_Status" json:"status,omitempty"`                                            // Order state
	InitialVolume  *gopb.Decimal       `protobuf:"bytes,7,opt,name=initial_volume,json=initialVolume,proto3" json:"initial_volume,omitempty"`                                             // Order initial volume
	CurrentVolume  *gopb.Decimal       `protobuf:"bytes,8,opt,name=current_volume,json=currentVolume,proto3" json:"current_volume,omitempty"`                                             // Order current volume
	PriceOpen      *gopb.Decimal       `protobuf:"bytes,9,opt,name=price_open,json=priceOpen,proto3" json:"price_open,omitempty"`                                                         // Price specified in the order
	StopLoss       *gopb.Decimal       `protobuf:"bytes,10,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`                                                           // Stop Loss value
	TakeProfit     *gopb.Decimal       `protobuf:"bytes,11,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`                                                     // Take Profit value
	StopLimit      *gopb.Decimal       `protobuf:"bytes,12,opt,name=stop_limit,json=stopLimit,proto3" json:"stop_limit,omitempty"`                                                        // The Limit order price for the StopLimit order
	ExpirationTime *gopb.UnixTimestamp `protobuf:"bytes,13,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`                                         // Order expiration time
	DoneTime       *gopb.UnixTimestamp `protobuf:"bytes,14,opt,name=done_time,json=doneTime,proto3" json:"done_time,omitempty"`                                                           // Order execution or cancellation time
	FillType       FillPolicy          `protobuf:"varint,15,opt,name=fill_type,json=fillType,proto3,enum=protos.edge.data.FillPolicy" json:"fill_type,omitempty"`                         // Order filling type
	ExpirationType ExpirationPolicy    `protobuf:"varint,16,opt,name=expiration_type,json=expirationType,proto3,enum=protos.edge.data.ExpirationPolicy" json:"expiration_type,omitempty"` // Order lifetime
	Originator     Order_Reason        `protobuf:"varint,17,opt,name=originator,proto3,enum=protos.edge.data.Order_Reason" json:"originator,omitempty"`                                   // The reason or source for placing an order
	Comment        string              `protobuf:"bytes,18,opt,name=comment,proto3" json:"comment,omitempty"`                                                                             // Order comment
	PositionId     uint64              `protobuf:"varint,19,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`                                                    // Position identifier that is set to an order as soon as it is executed. Each executed order results in a deal that opens or modifies an already existing position. The identifier of exactly this position
	// is set to the executed order at this moment.
	OppositePositionId uint64 `protobuf:"varint,20,opt,name=opposite_position_id,json=oppositePositionId,proto3" json:"opposite_position_id,omitempty"` // Identifier of an opposite position used for closing by order OrderType.ClosedBy
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edge_data_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edge_data_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_protos_edge_data_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Order) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Order) GetSetupTime() *gopb.UnixTimestamp {
	if x != nil {
		return x.SetupTime
	}
	return nil
}

func (x *Order) GetType() Order_Type {
	if x != nil {
		return x.Type
	}
	return Order_Buy
}

func (x *Order) GetStatus() Order_Status {
	if x != nil {
		return x.Status
	}
	return Order_Started
}

func (x *Order) GetInitialVolume() *gopb.Decimal {
	if x != nil {
		return x.InitialVolume
	}
	return nil
}

func (x *Order) GetCurrentVolume() *gopb.Decimal {
	if x != nil {
		return x.CurrentVolume
	}
	return nil
}

func (x *Order) GetPriceOpen() *gopb.Decimal {
	if x != nil {
		return x.PriceOpen
	}
	return nil
}

func (x *Order) GetStopLoss() *gopb.Decimal {
	if x != nil {
		return x.StopLoss
	}
	return nil
}

func (x *Order) GetTakeProfit() *gopb.Decimal {
	if x != nil {
		return x.TakeProfit
	}
	return nil
}

func (x *Order) GetStopLimit() *gopb.Decimal {
	if x != nil {
		return x.StopLimit
	}
	return nil
}

func (x *Order) GetExpirationTime() *gopb.UnixTimestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *Order) GetDoneTime() *gopb.UnixTimestamp {
	if x != nil {
		return x.DoneTime
	}
	return nil
}

func (x *Order) GetFillType() FillPolicy {
	if x != nil {
		return x.FillType
	}
	return FillPolicy_FillOrKill
}

func (x *Order) GetExpirationType() ExpirationPolicy {
	if x != nil {
		return x.ExpirationType
	}
	return ExpirationPolicy_UntilCancelled
}

func (x *Order) GetOriginator() Order_Reason {
	if x != nil {
		return x.Originator
	}
	return Order_Client
}

func (x *Order) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Order) GetPositionId() uint64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *Order) GetOppositePositionId() uint64 {
	if x != nil {
		return x.OppositePositionId
	}
	return 0
}

var File_protos_edge_data_order_proto protoreflect.FileDescriptor

var file_protos_edge_data_order_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x0b, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x65, 0x74,
	0x75, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e,
	0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x65, 0x74,
	0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x6f,
	0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x61,
	0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x09, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x09, 0x64, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x64, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x6f, 0x70, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6f,
	0x70, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x75,
	0x79, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x42, 0x75, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x65, 0x6c, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x75,
	0x79, 0x53, 0x74, 0x6f, 0x70, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x6c, 0x53,
	0x74, 0x6f, 0x70, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x75, 0x79, 0x53, 0x74, 0x6f, 0x70,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x6c, 0x53,
	0x74, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x42, 0x79, 0x10, 0x08, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x06, 0x12,
	0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x10, 0x07, 0x12,
	0x11, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x10, 0x09, 0x22, 0x62, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x65, 0x62, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x74, 0x6f, 0x70, 0x4f, 0x75, 0x74, 0x10, 0x06, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edge_data_order_proto_rawDescOnce sync.Once
	file_protos_edge_data_order_proto_rawDescData = file_protos_edge_data_order_proto_rawDesc
)

func file_protos_edge_data_order_proto_rawDescGZIP() []byte {
	file_protos_edge_data_order_proto_rawDescOnce.Do(func() {
		file_protos_edge_data_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edge_data_order_proto_rawDescData)
	})
	return file_protos_edge_data_order_proto_rawDescData
}

var file_protos_edge_data_order_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protos_edge_data_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_edge_data_order_proto_goTypes = []interface{}{
	(Order_Type)(0),            // 0: protos.edge.data.Order.Type
	(Order_Status)(0),          // 1: protos.edge.data.Order.Status
	(Order_Reason)(0),          // 2: protos.edge.data.Order.Reason
	(*Order)(nil),              // 3: protos.edge.data.Order
	(*gopb.UnixTimestamp)(nil), // 4: protos.common.UnixTimestamp
	(*gopb.Decimal)(nil),       // 5: protos.common.Decimal
	(FillPolicy)(0),            // 6: protos.edge.data.FillPolicy
	(ExpirationPolicy)(0),      // 7: protos.edge.data.ExpirationPolicy
}
var file_protos_edge_data_order_proto_depIdxs = []int32{
	4,  // 0: protos.edge.data.Order.setup_time:type_name -> protos.common.UnixTimestamp
	0,  // 1: protos.edge.data.Order.type:type_name -> protos.edge.data.Order.Type
	1,  // 2: protos.edge.data.Order.status:type_name -> protos.edge.data.Order.Status
	5,  // 3: protos.edge.data.Order.initial_volume:type_name -> protos.common.Decimal
	5,  // 4: protos.edge.data.Order.current_volume:type_name -> protos.common.Decimal
	5,  // 5: protos.edge.data.Order.price_open:type_name -> protos.common.Decimal
	5,  // 6: protos.edge.data.Order.stop_loss:type_name -> protos.common.Decimal
	5,  // 7: protos.edge.data.Order.take_profit:type_name -> protos.common.Decimal
	5,  // 8: protos.edge.data.Order.stop_limit:type_name -> protos.common.Decimal
	4,  // 9: protos.edge.data.Order.expiration_time:type_name -> protos.common.UnixTimestamp
	4,  // 10: protos.edge.data.Order.done_time:type_name -> protos.common.UnixTimestamp
	6,  // 11: protos.edge.data.Order.fill_type:type_name -> protos.edge.data.FillPolicy
	7,  // 12: protos.edge.data.Order.expiration_type:type_name -> protos.edge.data.ExpirationPolicy
	2,  // 13: protos.edge.data.Order.originator:type_name -> protos.edge.data.Order.Reason
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_protos_edge_data_order_proto_init() }
func file_protos_edge_data_order_proto_init() {
	if File_protos_edge_data_order_proto != nil {
		return
	}
	file_protos_edge_data_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_edge_data_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_protos_edge_data_order_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_edge_data_order_proto_goTypes,
		DependencyIndexes: file_protos_edge_data_order_proto_depIdxs,
		EnumInfos:         file_protos_edge_data_order_proto_enumTypes,
		MessageInfos:      file_protos_edge_data_order_proto_msgTypes,
	}.Build()
	File_protos_edge_data_order_proto = out.File
	file_protos_edge_data_order_proto_rawDesc = nil
	file_protos_edge_data_order_proto_goTypes = nil
	file_protos_edge_data_order_proto_depIdxs = nil
}
