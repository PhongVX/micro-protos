// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: product/product.proto

package product

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

type UpdateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string                       `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	ID            int32                        `protobuf:"varint,2,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Quantity      int32                        `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	BeginTxRes    *transaction.BeginTxResponse `protobuf:"bytes,4,opt,name=beginTxRes,proto3" json:"beginTxRes,omitempty"`
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateProductRequest) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *UpdateProductRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateProductRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *UpdateProductRequest) GetBeginTxRes() *transaction.BeginTxResponse {
	if x != nil {
		return x.BeginTxRes
	}
	return nil
}

type UpdateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowAffected int32 `protobuf:"varint,1,opt,name=rowAffected,proto3" json:"rowAffected,omitempty"`
}

func (x *UpdateProductResponse) Reset() {
	*x = UpdateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductResponse) ProtoMessage() {}

func (x *UpdateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProductResponse) GetRowAffected() int32 {
	if x != nil {
		return x.RowAffected
	}
	return 0
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x1a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa6, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x32, 0x5b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x50,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50,
	0x68, 0x6f, 0x6e, 0x67, 0x56, 0x58, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData = file_product_product_proto_rawDesc
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_product_proto_rawDescData)
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_product_proto_goTypes = []interface{}{
	(*UpdateProductRequest)(nil),        // 0: product.UpdateProductRequest
	(*UpdateProductResponse)(nil),       // 1: product.UpdateProductResponse
	(*transaction.BeginTxResponse)(nil), // 2: transaction.BeginTxResponse
}
var file_product_product_proto_depIdxs = []int32{
	2, // 0: product.UpdateProductRequest.beginTxRes:type_name -> transaction.BeginTxResponse
	0, // 1: product.Product.UpdateProduct:input_type -> product.UpdateProductRequest
	1, // 2: product.Product.UpdateProduct:output_type -> product.UpdateProductResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductRequest); i {
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
		file_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductResponse); i {
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
			RawDescriptor: file_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_rawDesc = nil
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}