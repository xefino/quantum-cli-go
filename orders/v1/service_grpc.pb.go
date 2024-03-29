// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: protos/edge/orders/v1/service.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	data "github.com/xefino/quantum-cli-go/data"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrderManagementService_GetName_FullMethodName            = "/protos.edge.orders.v1.OrderManagementService/GetName"
	OrderManagementService_ModifyPendingOrder_FullMethodName = "/protos.edge.orders.v1.OrderManagementService/ModifyPendingOrder"
)

// OrderManagementServiceClient is the client API for OrderManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementServiceClient interface {
	// GetName retrieves an identifier associated with the service, allowing the strategy runner to
	// uniquely identify it when multiple order services are being referenced. It will return a
	// name result showing the address, ID and human-readable name of the order service that processed
	// the original request.
	GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*data.NameResult, error)
	// ModifyOrder takes an existing, pending order and may modify or cancel it based on updated market
	// conditions. This endpoint will be responsible for updating the price, stop-loss, take-profit, limits,
	// lot size, expiration or any other fields on the order before it is filled in full. This is done by
	// taking the order from the ModifyPendingOrderRequest object, making the necessary modifications, and
	// setting the associated TradeRequest on the ModifyPendingOrderResponse object.
	ModifyPendingOrder(ctx context.Context, in *ModifyPendingOrderRequest, opts ...grpc.CallOption) (*ModifyPendingOrderResponse, error)
}

type orderManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementServiceClient(cc grpc.ClientConnInterface) OrderManagementServiceClient {
	return &orderManagementServiceClient{cc}
}

func (c *orderManagementServiceClient) GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*data.NameResult, error) {
	out := new(data.NameResult)
	err := c.cc.Invoke(ctx, OrderManagementService_GetName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementServiceClient) ModifyPendingOrder(ctx context.Context, in *ModifyPendingOrderRequest, opts ...grpc.CallOption) (*ModifyPendingOrderResponse, error) {
	out := new(ModifyPendingOrderResponse)
	err := c.cc.Invoke(ctx, OrderManagementService_ModifyPendingOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServiceServer is the server API for OrderManagementService service.
// All implementations must embed UnimplementedOrderManagementServiceServer
// for forward compatibility
type OrderManagementServiceServer interface {
	// GetName retrieves an identifier associated with the service, allowing the strategy runner to
	// uniquely identify it when multiple order services are being referenced. It will return a
	// name result showing the address, ID and human-readable name of the order service that processed
	// the original request.
	GetName(context.Context, *empty.Empty) (*data.NameResult, error)
	// ModifyOrder takes an existing, pending order and may modify or cancel it based on updated market
	// conditions. This endpoint will be responsible for updating the price, stop-loss, take-profit, limits,
	// lot size, expiration or any other fields on the order before it is filled in full. This is done by
	// taking the order from the ModifyPendingOrderRequest object, making the necessary modifications, and
	// setting the associated TradeRequest on the ModifyPendingOrderResponse object.
	ModifyPendingOrder(context.Context, *ModifyPendingOrderRequest) (*ModifyPendingOrderResponse, error)
	mustEmbedUnimplementedOrderManagementServiceServer()
}

// UnimplementedOrderManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServiceServer struct {
}

func (UnimplementedOrderManagementServiceServer) GetName(context.Context, *empty.Empty) (*data.NameResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedOrderManagementServiceServer) ModifyPendingOrder(context.Context, *ModifyPendingOrderRequest) (*ModifyPendingOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPendingOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) mustEmbedUnimplementedOrderManagementServiceServer() {
}

// UnsafeOrderManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServiceServer will
// result in compilation errors.
type UnsafeOrderManagementServiceServer interface {
	mustEmbedUnimplementedOrderManagementServiceServer()
}

func RegisterOrderManagementServiceServer(s grpc.ServiceRegistrar, srv OrderManagementServiceServer) {
	s.RegisterService(&OrderManagementService_ServiceDesc, srv)
}

func _OrderManagementService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_GetName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).GetName(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagementService_ModifyPendingOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPendingOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).ModifyPendingOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_ModifyPendingOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).ModifyPendingOrder(ctx, req.(*ModifyPendingOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManagementService_ServiceDesc is the grpc.ServiceDesc for OrderManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.edge.orders.v1.OrderManagementService",
	HandlerType: (*OrderManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _OrderManagementService_GetName_Handler,
		},
		{
			MethodName: "ModifyPendingOrder",
			Handler:    _OrderManagementService_ModifyPendingOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/edge/orders/v1/service.proto",
}
