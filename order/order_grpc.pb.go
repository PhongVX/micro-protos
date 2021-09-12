// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	InsertOrder(ctx context.Context, in *InsertOrderRequest, opts ...grpc.CallOption) (*InsertOrderResponse, error)
	InsertOrderDetail(ctx context.Context, in *InsertOrderDetailRequest, opts ...grpc.CallOption) (*InsertOrderDetailResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) InsertOrder(ctx context.Context, in *InsertOrderRequest, opts ...grpc.CallOption) (*InsertOrderResponse, error) {
	out := new(InsertOrderResponse)
	err := c.cc.Invoke(ctx, "/order.Order/InsertOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) InsertOrderDetail(ctx context.Context, in *InsertOrderDetailRequest, opts ...grpc.CallOption) (*InsertOrderDetailResponse, error) {
	out := new(InsertOrderDetailResponse)
	err := c.cc.Invoke(ctx, "/order.Order/InsertOrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	InsertOrder(context.Context, *InsertOrderRequest) (*InsertOrderResponse, error)
	InsertOrderDetail(context.Context, *InsertOrderDetailRequest) (*InsertOrderDetailResponse, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) InsertOrder(context.Context, *InsertOrderRequest) (*InsertOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertOrder not implemented")
}
func (UnimplementedOrderServer) InsertOrderDetail(context.Context, *InsertOrderDetailRequest) (*InsertOrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertOrderDetail not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_InsertOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).InsertOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/InsertOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).InsertOrder(ctx, req.(*InsertOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_InsertOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertOrderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).InsertOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/InsertOrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).InsertOrderDetail(ctx, req.(*InsertOrderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertOrder",
			Handler:    _Order_InsertOrder_Handler,
		},
		{
			MethodName: "InsertOrderDetail",
			Handler:    _Order_InsertOrderDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
