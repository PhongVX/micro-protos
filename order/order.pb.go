// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: order/order.proto

package order

import (
	transaction "github.com/PhongVX/micro-protos/transaction"
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

type OrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID  int32   `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity   int32   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price      float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	TotalPrice float64 `protobuf:"fixed64,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	OrderID    string  `protobuf:"bytes,5,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderDetail) Reset() {
	*x = OrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetail) ProtoMessage() {}

func (x *OrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetail.ProtoReflect.Descriptor instead.
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderDetail) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *OrderDetail) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderDetail) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderDetail) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type InsertOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string                       `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	PhoneNumber   string                       `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	TotalPrice    float64                      `protobuf:"fixed64,3,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Name          string                       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                       `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	BeginTxRes    *transaction.BeginTxResponse `protobuf:"bytes,6,opt,name=beginTxRes,proto3" json:"beginTxRes,omitempty"`
}

func (x *InsertOrderRequest) Reset() {
	*x = InsertOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOrderRequest) ProtoMessage() {}

func (x *InsertOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOrderRequest.ProtoReflect.Descriptor instead.
func (*InsertOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *InsertOrderRequest) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *InsertOrderRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *InsertOrderRequest) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *InsertOrderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertOrderRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InsertOrderRequest) GetBeginTxRes() *transaction.BeginTxResponse {
	if x != nil {
		return x.BeginTxRes
	}
	return nil
}

type InsertOrderDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string                       `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	TxRandomID    string                       `protobuf:"bytes,2,opt,name=txRandomID,proto3" json:"txRandomID,omitempty"`
	OrderDetails  []*OrderDetail               `protobuf:"bytes,3,rep,name=orderDetails,proto3" json:"orderDetails,omitempty"`
	BeginTxRes    *transaction.BeginTxResponse `protobuf:"bytes,6,opt,name=beginTxRes,proto3" json:"beginTxRes,omitempty"`
}

func (x *InsertOrderDetailRequest) Reset() {
	*x = InsertOrderDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOrderDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOrderDetailRequest) ProtoMessage() {}

func (x *InsertOrderDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOrderDetailRequest.ProtoReflect.Descriptor instead.
func (*InsertOrderDetailRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *InsertOrderDetailRequest) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *InsertOrderDetailRequest) GetTxRandomID() string {
	if x != nil {
		return x.TxRandomID
	}
	return ""
}

func (x *InsertOrderDetailRequest) GetOrderDetails() []*OrderDetail {
	if x != nil {
		return x.OrderDetails
	}
	return nil
}

func (x *InsertOrderDetailRequest) GetBeginTxRes() *transaction.BeginTxResponse {
	if x != nil {
		return x.BeginTxRes
	}
	return nil
}

type InsertOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowAffected int32  `protobuf:"varint,1,opt,name=rowAffected,proto3" json:"rowAffected,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertOrderResponse) Reset() {
	*x = InsertOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOrderResponse) ProtoMessage() {}

func (x *InsertOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOrderResponse.ProtoReflect.Descriptor instead.
func (*InsertOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *InsertOrderResponse) GetRowAffected() int32 {
	if x != nil {
		return x.RowAffected
	}
	return 0
}

func (x *InsertOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InsertOrderDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowAffected int32  `protobuf:"varint,1,opt,name=rowAffected,proto3" json:"rowAffected,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertOrderDetailResponse) Reset() {
	*x = InsertOrderDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOrderDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOrderDetailResponse) ProtoMessage() {}

func (x *InsertOrderDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOrderDetailResponse.ProtoReflect.Descriptor instead.
func (*InsertOrderDetailResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *InsertOrderDetailResponse) GetRowAffected() int32 {
	if x != nil {
		return x.RowAffected
	}
	return 0
}

func (x *InsertOrderDetailResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_order_order_proto protoreflect.FileDescriptor

var file_order_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x22, 0xe8, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3c, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x22, 0xd6,
	0x01, 0x0a, 0x18, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x78, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x78, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49,
	0x44, 0x12, 0x36, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4d, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xa9, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0b, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x68, 0x6f, 0x6e, 0x67, 0x56,
	0x58, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_order_order_proto_rawDescOnce sync.Once
	file_order_order_proto_rawDescData = file_order_order_proto_rawDesc
)

func file_order_order_proto_rawDescGZIP() []byte {
	file_order_order_proto_rawDescOnce.Do(func() {
		file_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_order_proto_rawDescData)
	})
	return file_order_order_proto_rawDescData
}

var file_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_order_proto_goTypes = []interface{}{
	(*OrderDetail)(nil),                 // 0: order.OrderDetail
	(*InsertOrderRequest)(nil),          // 1: order.InsertOrderRequest
	(*InsertOrderDetailRequest)(nil),    // 2: order.InsertOrderDetailRequest
	(*InsertOrderResponse)(nil),         // 3: order.InsertOrderResponse
	(*InsertOrderDetailResponse)(nil),   // 4: order.InsertOrderDetailResponse
	(*transaction.BeginTxResponse)(nil), // 5: transaction.BeginTxResponse
}
var file_order_order_proto_depIdxs = []int32{
	5, // 0: order.InsertOrderRequest.beginTxRes:type_name -> transaction.BeginTxResponse
	0, // 1: order.InsertOrderDetailRequest.orderDetails:type_name -> order.OrderDetail
	5, // 2: order.InsertOrderDetailRequest.beginTxRes:type_name -> transaction.BeginTxResponse
	1, // 3: order.Order.InsertOrder:input_type -> order.InsertOrderRequest
	2, // 4: order.Order.InsertOrderDetail:input_type -> order.InsertOrderDetailRequest
	3, // 5: order.Order.InsertOrder:output_type -> order.InsertOrderResponse
	4, // 6: order.Order.InsertOrderDetail:output_type -> order.InsertOrderDetailResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_order_proto_init() }
func file_order_order_proto_init() {
	if File_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetail); i {
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
		file_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOrderRequest); i {
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
		file_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOrderDetailRequest); i {
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
		file_order_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOrderResponse); i {
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
		file_order_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOrderDetailResponse); i {
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
			RawDescriptor: file_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_order_proto_goTypes,
		DependencyIndexes: file_order_order_proto_depIdxs,
		MessageInfos:      file_order_order_proto_msgTypes,
	}.Build()
	File_order_order_proto = out.File
	file_order_order_proto_rawDesc = nil
	file_order_order_proto_goTypes = nil
	file_order_order_proto_depIdxs = nil
}
