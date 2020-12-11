// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_user_bill.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 手工生成订单
type GenerateAllUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *GenerateAllUserBillsRequest) Reset() {
	*x = GenerateAllUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAllUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAllUserBillsRequest) ProtoMessage() {}

func (x *GenerateAllUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAllUserBillsRequest.ProtoReflect.Descriptor instead.
func (*GenerateAllUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateAllUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

// 计算所有账单数量
type CountAllUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaidFlag int32  `protobuf:"varint,1,opt,name=paidFlag,proto3" json:"paidFlag,omitempty"` // 0|1|-1
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Month    string `protobuf:"bytes,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *CountAllUserBillsRequest) Reset() {
	*x = CountAllUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllUserBillsRequest) ProtoMessage() {}

func (x *CountAllUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllUserBillsRequest.ProtoReflect.Descriptor instead.
func (*CountAllUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{1}
}

func (x *CountAllUserBillsRequest) GetPaidFlag() int32 {
	if x != nil {
		return x.PaidFlag
	}
	return 0
}

func (x *CountAllUserBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountAllUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

// 列出单页账单
type ListUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaidFlag int32  `protobuf:"varint,1,opt,name=paidFlag,proto3" json:"paidFlag,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Month    string `protobuf:"bytes,5,opt,name=month,proto3" json:"month,omitempty"`
	Offset   int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUserBillsRequest) Reset() {
	*x = ListUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBillsRequest) ProtoMessage() {}

func (x *ListUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBillsRequest.ProtoReflect.Descriptor instead.
func (*ListUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserBillsRequest) GetPaidFlag() int32 {
	if x != nil {
		return x.PaidFlag
	}
	return 0
}

func (x *ListUserBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *ListUserBillsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserBillsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUserBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBills []*UserBill `protobuf:"bytes,1,rep,name=userBills,proto3" json:"userBills,omitempty"`
}

func (x *ListUserBillsResponse) Reset() {
	*x = ListUserBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBillsResponse) ProtoMessage() {}

func (x *ListUserBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBillsResponse.ProtoReflect.Descriptor instead.
func (*ListUserBillsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserBillsResponse) GetUserBills() []*UserBill {
	if x != nil {
		return x.UserBills
	}
	return nil
}

var File_service_user_bill_proto protoreflect.FileDescriptor

var file_service_user_bill_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x72,
	0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x69,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x1b, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x64, 0x0a,
	0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x69,
	0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x69,
	0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x32, 0xe9, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x14, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0d, 0x6c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_user_bill_proto_rawDescOnce sync.Once
	file_service_user_bill_proto_rawDescData = file_service_user_bill_proto_rawDesc
)

func file_service_user_bill_proto_rawDescGZIP() []byte {
	file_service_user_bill_proto_rawDescOnce.Do(func() {
		file_service_user_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_bill_proto_rawDescData)
	})
	return file_service_user_bill_proto_rawDescData
}

var file_service_user_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_user_bill_proto_goTypes = []interface{}{
	(*GenerateAllUserBillsRequest)(nil), // 0: pb.GenerateAllUserBillsRequest
	(*CountAllUserBillsRequest)(nil),    // 1: pb.CountAllUserBillsRequest
	(*ListUserBillsRequest)(nil),        // 2: pb.ListUserBillsRequest
	(*ListUserBillsResponse)(nil),       // 3: pb.ListUserBillsResponse
	(*UserBill)(nil),                    // 4: pb.UserBill
	(*RPCSuccess)(nil),                  // 5: pb.RPCSuccess
	(*RPCCountResponse)(nil),            // 6: pb.RPCCountResponse
}
var file_service_user_bill_proto_depIdxs = []int32{
	4, // 0: pb.ListUserBillsResponse.userBills:type_name -> pb.UserBill
	0, // 1: pb.UserBillService.generateAllUserBills:input_type -> pb.GenerateAllUserBillsRequest
	1, // 2: pb.UserBillService.countAllUserBills:input_type -> pb.CountAllUserBillsRequest
	2, // 3: pb.UserBillService.listUserBills:input_type -> pb.ListUserBillsRequest
	5, // 4: pb.UserBillService.generateAllUserBills:output_type -> pb.RPCSuccess
	6, // 5: pb.UserBillService.countAllUserBills:output_type -> pb.RPCCountResponse
	3, // 6: pb.UserBillService.listUserBills:output_type -> pb.ListUserBillsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_bill_proto_init() }
func file_service_user_bill_proto_init() {
	if File_service_user_bill_proto != nil {
		return
	}
	file_rpc_messages_proto_init()
	file_model_user_bill_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAllUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserBillsResponse); i {
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
			RawDescriptor: file_service_user_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_bill_proto_goTypes,
		DependencyIndexes: file_service_user_bill_proto_depIdxs,
		MessageInfos:      file_service_user_bill_proto_msgTypes,
	}.Build()
	File_service_user_bill_proto = out.File
	file_service_user_bill_proto_rawDesc = nil
	file_service_user_bill_proto_goTypes = nil
	file_service_user_bill_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserBillServiceClient is the client API for UserBillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserBillServiceClient interface {
	// 手工生成订单
	GenerateAllUserBills(ctx context.Context, in *GenerateAllUserBillsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算所有账单数量
	CountAllUserBills(ctx context.Context, in *CountAllUserBillsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页账单
	ListUserBills(ctx context.Context, in *ListUserBillsRequest, opts ...grpc.CallOption) (*ListUserBillsResponse, error)
}

type userBillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBillServiceClient(cc grpc.ClientConnInterface) UserBillServiceClient {
	return &userBillServiceClient{cc}
}

func (c *userBillServiceClient) GenerateAllUserBills(ctx context.Context, in *GenerateAllUserBillsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserBillService/generateAllUserBills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) CountAllUserBills(ctx context.Context, in *CountAllUserBillsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserBillService/countAllUserBills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) ListUserBills(ctx context.Context, in *ListUserBillsRequest, opts ...grpc.CallOption) (*ListUserBillsResponse, error) {
	out := new(ListUserBillsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserBillService/listUserBills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBillServiceServer is the server API for UserBillService service.
type UserBillServiceServer interface {
	// 手工生成订单
	GenerateAllUserBills(context.Context, *GenerateAllUserBillsRequest) (*RPCSuccess, error)
	// 计算所有账单数量
	CountAllUserBills(context.Context, *CountAllUserBillsRequest) (*RPCCountResponse, error)
	// 列出单页账单
	ListUserBills(context.Context, *ListUserBillsRequest) (*ListUserBillsResponse, error)
}

// UnimplementedUserBillServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserBillServiceServer struct {
}

func (*UnimplementedUserBillServiceServer) GenerateAllUserBills(context.Context, *GenerateAllUserBillsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAllUserBills not implemented")
}
func (*UnimplementedUserBillServiceServer) CountAllUserBills(context.Context, *CountAllUserBillsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllUserBills not implemented")
}
func (*UnimplementedUserBillServiceServer) ListUserBills(context.Context, *ListUserBillsRequest) (*ListUserBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserBills not implemented")
}

func RegisterUserBillServiceServer(s *grpc.Server, srv UserBillServiceServer) {
	s.RegisterService(&_UserBillService_serviceDesc, srv)
}

func _UserBillService_GenerateAllUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAllUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).GenerateAllUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserBillService/GenerateAllUserBills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).GenerateAllUserBills(ctx, req.(*GenerateAllUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_CountAllUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).CountAllUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserBillService/CountAllUserBills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).CountAllUserBills(ctx, req.(*CountAllUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_ListUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).ListUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserBillService/ListUserBills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).ListUserBills(ctx, req.(*ListUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserBillService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserBillService",
	HandlerType: (*UserBillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateAllUserBills",
			Handler:    _UserBillService_GenerateAllUserBills_Handler,
		},
		{
			MethodName: "countAllUserBills",
			Handler:    _UserBillService_CountAllUserBills_Handler,
		},
		{
			MethodName: "listUserBills",
			Handler:    _UserBillService_ListUserBills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_bill.proto",
}
