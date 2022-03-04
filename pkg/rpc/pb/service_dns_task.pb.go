// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_dns_task.proto

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

// 检查是否有正在执行的任务
type ExistsDNSTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExistsDNSTasksRequest) Reset() {
	*x = ExistsDNSTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_dns_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsDNSTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsDNSTasksRequest) ProtoMessage() {}

func (x *ExistsDNSTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_dns_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsDNSTasksRequest.ProtoReflect.Descriptor instead.
func (*ExistsDNSTasksRequest) Descriptor() ([]byte, []int) {
	return file_service_dns_task_proto_rawDescGZIP(), []int{0}
}

type ExistsDNSTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExistTasks bool `protobuf:"varint,1,opt,name=existTasks,proto3" json:"existTasks,omitempty"`
	ExistError bool `protobuf:"varint,2,opt,name=existError,proto3" json:"existError,omitempty"`
}

func (x *ExistsDNSTasksResponse) Reset() {
	*x = ExistsDNSTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_dns_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsDNSTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsDNSTasksResponse) ProtoMessage() {}

func (x *ExistsDNSTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_dns_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsDNSTasksResponse.ProtoReflect.Descriptor instead.
func (*ExistsDNSTasksResponse) Descriptor() ([]byte, []int) {
	return file_service_dns_task_proto_rawDescGZIP(), []int{1}
}

func (x *ExistsDNSTasksResponse) GetExistTasks() bool {
	if x != nil {
		return x.ExistTasks
	}
	return false
}

func (x *ExistsDNSTasksResponse) GetExistError() bool {
	if x != nil {
		return x.ExistError
	}
	return false
}

// 查找需要通知的任务
type FindAllDoingDNSTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
}

func (x *FindAllDoingDNSTasksRequest) Reset() {
	*x = FindAllDoingDNSTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_dns_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllDoingDNSTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllDoingDNSTasksRequest) ProtoMessage() {}

func (x *FindAllDoingDNSTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_dns_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllDoingDNSTasksRequest.ProtoReflect.Descriptor instead.
func (*FindAllDoingDNSTasksRequest) Descriptor() ([]byte, []int) {
	return file_service_dns_task_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllDoingDNSTasksRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

type FindAllDoingDNSTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsTasks []*DNSTask `protobuf:"bytes,1,rep,name=dnsTasks,proto3" json:"dnsTasks,omitempty"`
}

func (x *FindAllDoingDNSTasksResponse) Reset() {
	*x = FindAllDoingDNSTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_dns_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllDoingDNSTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllDoingDNSTasksResponse) ProtoMessage() {}

func (x *FindAllDoingDNSTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_dns_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllDoingDNSTasksResponse.ProtoReflect.Descriptor instead.
func (*FindAllDoingDNSTasksResponse) Descriptor() ([]byte, []int) {
	return file_service_dns_task_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllDoingDNSTasksResponse) GetDnsTasks() []*DNSTask {
	if x != nil {
		return x.DnsTasks
	}
	return nil
}

// 删除任务
type DeleteDNSTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsTaskId int64 `protobuf:"varint,1,opt,name=dnsTaskId,proto3" json:"dnsTaskId,omitempty"`
}

func (x *DeleteDNSTaskRequest) Reset() {
	*x = DeleteDNSTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_dns_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDNSTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDNSTaskRequest) ProtoMessage() {}

func (x *DeleteDNSTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_dns_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDNSTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteDNSTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_dns_task_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDNSTaskRequest) GetDnsTaskId() int64 {
	if x != nil {
		return x.DnsTaskId
	}
	return 0
}

var File_service_dns_task_proto protoreflect.FileDescriptor

var file_service_dns_task_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x4e,
	0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a,
	0x16, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x1c,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x44, 0x4e, 0x53, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08,
	0x64, 0x6e, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x64, 0x6e, 0x73,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x34, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x6e, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x64, 0x6e, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x32, 0xef, 0x01, 0x0a, 0x0e,
	0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x4e, 0x53, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x6f, 0x69, 0x6e,
	0x67, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x6f, 0x69,
	0x6e, 0x67, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x4e, 0x53, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x4e, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_dns_task_proto_rawDescOnce sync.Once
	file_service_dns_task_proto_rawDescData = file_service_dns_task_proto_rawDesc
)

func file_service_dns_task_proto_rawDescGZIP() []byte {
	file_service_dns_task_proto_rawDescOnce.Do(func() {
		file_service_dns_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_dns_task_proto_rawDescData)
	})
	return file_service_dns_task_proto_rawDescData
}

var file_service_dns_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_dns_task_proto_goTypes = []interface{}{
	(*ExistsDNSTasksRequest)(nil),        // 0: pb.ExistsDNSTasksRequest
	(*ExistsDNSTasksResponse)(nil),       // 1: pb.ExistsDNSTasksResponse
	(*FindAllDoingDNSTasksRequest)(nil),  // 2: pb.FindAllDoingDNSTasksRequest
	(*FindAllDoingDNSTasksResponse)(nil), // 3: pb.FindAllDoingDNSTasksResponse
	(*DeleteDNSTaskRequest)(nil),         // 4: pb.DeleteDNSTaskRequest
	(*DNSTask)(nil),                      // 5: pb.DNSTask
	(*RPCSuccess)(nil),                   // 6: pb.RPCSuccess
}
var file_service_dns_task_proto_depIdxs = []int32{
	5, // 0: pb.FindAllDoingDNSTasksResponse.dnsTasks:type_name -> pb.DNSTask
	0, // 1: pb.DNSTaskService.existsDNSTasks:input_type -> pb.ExistsDNSTasksRequest
	2, // 2: pb.DNSTaskService.findAllDoingDNSTasks:input_type -> pb.FindAllDoingDNSTasksRequest
	4, // 3: pb.DNSTaskService.deleteDNSTask:input_type -> pb.DeleteDNSTaskRequest
	1, // 4: pb.DNSTaskService.existsDNSTasks:output_type -> pb.ExistsDNSTasksResponse
	3, // 5: pb.DNSTaskService.findAllDoingDNSTasks:output_type -> pb.FindAllDoingDNSTasksResponse
	6, // 6: pb.DNSTaskService.deleteDNSTask:output_type -> pb.RPCSuccess
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_dns_task_proto_init() }
func file_service_dns_task_proto_init() {
	if File_service_dns_task_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_dns_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_dns_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsDNSTasksRequest); i {
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
		file_service_dns_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsDNSTasksResponse); i {
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
		file_service_dns_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllDoingDNSTasksRequest); i {
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
		file_service_dns_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllDoingDNSTasksResponse); i {
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
		file_service_dns_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDNSTaskRequest); i {
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
			RawDescriptor: file_service_dns_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_dns_task_proto_goTypes,
		DependencyIndexes: file_service_dns_task_proto_depIdxs,
		MessageInfos:      file_service_dns_task_proto_msgTypes,
	}.Build()
	File_service_dns_task_proto = out.File
	file_service_dns_task_proto_rawDesc = nil
	file_service_dns_task_proto_goTypes = nil
	file_service_dns_task_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DNSTaskServiceClient is the client API for DNSTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DNSTaskServiceClient interface {
	// 检查是否有正在执行的任务
	ExistsDNSTasks(ctx context.Context, in *ExistsDNSTasksRequest, opts ...grpc.CallOption) (*ExistsDNSTasksResponse, error)
	// 查找正在执行的所有任务
	FindAllDoingDNSTasks(ctx context.Context, in *FindAllDoingDNSTasksRequest, opts ...grpc.CallOption) (*FindAllDoingDNSTasksResponse, error)
	// 删除任务
	DeleteDNSTask(ctx context.Context, in *DeleteDNSTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type dNSTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSTaskServiceClient(cc grpc.ClientConnInterface) DNSTaskServiceClient {
	return &dNSTaskServiceClient{cc}
}

func (c *dNSTaskServiceClient) ExistsDNSTasks(ctx context.Context, in *ExistsDNSTasksRequest, opts ...grpc.CallOption) (*ExistsDNSTasksResponse, error) {
	out := new(ExistsDNSTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSTaskService/existsDNSTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSTaskServiceClient) FindAllDoingDNSTasks(ctx context.Context, in *FindAllDoingDNSTasksRequest, opts ...grpc.CallOption) (*FindAllDoingDNSTasksResponse, error) {
	out := new(FindAllDoingDNSTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSTaskService/findAllDoingDNSTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSTaskServiceClient) DeleteDNSTask(ctx context.Context, in *DeleteDNSTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DNSTaskService/deleteDNSTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSTaskServiceServer is the server API for DNSTaskService service.
type DNSTaskServiceServer interface {
	// 检查是否有正在执行的任务
	ExistsDNSTasks(context.Context, *ExistsDNSTasksRequest) (*ExistsDNSTasksResponse, error)
	// 查找正在执行的所有任务
	FindAllDoingDNSTasks(context.Context, *FindAllDoingDNSTasksRequest) (*FindAllDoingDNSTasksResponse, error)
	// 删除任务
	DeleteDNSTask(context.Context, *DeleteDNSTaskRequest) (*RPCSuccess, error)
}

// UnimplementedDNSTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDNSTaskServiceServer struct {
}

func (*UnimplementedDNSTaskServiceServer) ExistsDNSTasks(context.Context, *ExistsDNSTasksRequest) (*ExistsDNSTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsDNSTasks not implemented")
}
func (*UnimplementedDNSTaskServiceServer) FindAllDoingDNSTasks(context.Context, *FindAllDoingDNSTasksRequest) (*FindAllDoingDNSTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDoingDNSTasks not implemented")
}
func (*UnimplementedDNSTaskServiceServer) DeleteDNSTask(context.Context, *DeleteDNSTaskRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDNSTask not implemented")
}

func RegisterDNSTaskServiceServer(s *grpc.Server, srv DNSTaskServiceServer) {
	s.RegisterService(&_DNSTaskService_serviceDesc, srv)
}

func _DNSTaskService_ExistsDNSTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsDNSTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).ExistsDNSTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSTaskService/ExistsDNSTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).ExistsDNSTasks(ctx, req.(*ExistsDNSTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSTaskService_FindAllDoingDNSTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDoingDNSTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).FindAllDoingDNSTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSTaskService/FindAllDoingDNSTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).FindAllDoingDNSTasks(ctx, req.(*FindAllDoingDNSTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSTaskService_DeleteDNSTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDNSTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).DeleteDNSTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSTaskService/DeleteDNSTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).DeleteDNSTask(ctx, req.(*DeleteDNSTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSTaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNSTaskService",
	HandlerType: (*DNSTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "existsDNSTasks",
			Handler:    _DNSTaskService_ExistsDNSTasks_Handler,
		},
		{
			MethodName: "findAllDoingDNSTasks",
			Handler:    _DNSTaskService_FindAllDoingDNSTasks_Handler,
		},
		{
			MethodName: "deleteDNSTask",
			Handler:    _DNSTaskService_DeleteDNSTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dns_task.proto",
}
