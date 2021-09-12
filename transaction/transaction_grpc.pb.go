// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transaction

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

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	BeginTx(ctx context.Context, in *BeginTxRequest, opts ...grpc.CallOption) (*BeginTxResponse, error)
	Commit(ctx context.Context, in *CommonTxDoActionRequest, opts ...grpc.CallOption) (*CommonTxResponse, error)
	Rollback(ctx context.Context, in *CommonTxDoActionRequest, opts ...grpc.CallOption) (*CommonTxResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) BeginTx(ctx context.Context, in *BeginTxRequest, opts ...grpc.CallOption) (*BeginTxResponse, error) {
	out := new(BeginTxResponse)
	err := c.cc.Invoke(ctx, "/transaction.Transaction/BeginTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Commit(ctx context.Context, in *CommonTxDoActionRequest, opts ...grpc.CallOption) (*CommonTxResponse, error) {
	out := new(CommonTxResponse)
	err := c.cc.Invoke(ctx, "/transaction.Transaction/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Rollback(ctx context.Context, in *CommonTxDoActionRequest, opts ...grpc.CallOption) (*CommonTxResponse, error) {
	out := new(CommonTxResponse)
	err := c.cc.Invoke(ctx, "/transaction.Transaction/Rollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	BeginTx(context.Context, *BeginTxRequest) (*BeginTxResponse, error)
	Commit(context.Context, *CommonTxDoActionRequest) (*CommonTxResponse, error)
	Rollback(context.Context, *CommonTxDoActionRequest) (*CommonTxResponse, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) BeginTx(context.Context, *BeginTxRequest) (*BeginTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginTx not implemented")
}
func (UnimplementedTransactionServer) Commit(context.Context, *CommonTxDoActionRequest) (*CommonTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedTransactionServer) Rollback(context.Context, *CommonTxDoActionRequest) (*CommonTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_BeginTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).BeginTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.Transaction/BeginTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).BeginTx(ctx, req.(*BeginTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonTxDoActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.Transaction/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Commit(ctx, req.(*CommonTxDoActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonTxDoActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.Transaction/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Rollback(ctx, req.(*CommonTxDoActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeginTx",
			Handler:    _Transaction_BeginTx_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Transaction_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Transaction_Rollback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
