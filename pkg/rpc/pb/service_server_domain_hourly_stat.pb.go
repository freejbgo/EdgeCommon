// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_server_domain_hourly_stat.proto

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

// 读取域名排行
type ListTopServerDomainStatsWithServerIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64  `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	NodeId        int64  `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	ServerId      int64  `protobuf:"varint,3,opt,name=serverId,proto3" json:"serverId,omitempty"`
	HourFrom      string `protobuf:"bytes,4,opt,name=hourFrom,proto3" json:"hourFrom,omitempty"`
	HourTo        string `protobuf:"bytes,5,opt,name=hourTo,proto3" json:"hourTo,omitempty"`
	Size          int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListTopServerDomainStatsWithServerIdRequest) Reset() {
	*x = ListTopServerDomainStatsWithServerIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_domain_hourly_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopServerDomainStatsWithServerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopServerDomainStatsWithServerIdRequest) ProtoMessage() {}

func (x *ListTopServerDomainStatsWithServerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_domain_hourly_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopServerDomainStatsWithServerIdRequest.ProtoReflect.Descriptor instead.
func (*ListTopServerDomainStatsWithServerIdRequest) Descriptor() ([]byte, []int) {
	return file_service_server_domain_hourly_stat_proto_rawDescGZIP(), []int{0}
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetHourFrom() string {
	if x != nil {
		return x.HourFrom
	}
	return ""
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetHourTo() string {
	if x != nil {
		return x.HourTo
	}
	return ""
}

func (x *ListTopServerDomainStatsWithServerIdRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListTopServerDomainStatsWithServerIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainStats []*ServerDomainHourlyStat `protobuf:"bytes,1,rep,name=domainStats,proto3" json:"domainStats,omitempty"`
}

func (x *ListTopServerDomainStatsWithServerIdResponse) Reset() {
	*x = ListTopServerDomainStatsWithServerIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_domain_hourly_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopServerDomainStatsWithServerIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopServerDomainStatsWithServerIdResponse) ProtoMessage() {}

func (x *ListTopServerDomainStatsWithServerIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_domain_hourly_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopServerDomainStatsWithServerIdResponse.ProtoReflect.Descriptor instead.
func (*ListTopServerDomainStatsWithServerIdResponse) Descriptor() ([]byte, []int) {
	return file_service_server_domain_hourly_stat_proto_rawDescGZIP(), []int{1}
}

func (x *ListTopServerDomainStatsWithServerIdResponse) GetDomainStats() []*ServerDomainHourlyStat {
	if x != nil {
		return x.DomainStats
	}
	return nil
}

var File_service_server_domain_hourly_stat_proto protoreflect.FileDescriptor

var file_service_server_domain_hourly_stat_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x2c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x2b,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x75, 0x72, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x75, 0x72, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x75, 0x72, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x6f, 0x75, 0x72, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x6c, 0x0a,
	0x2c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0b,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x32, 0xab, 0x01, 0x0a, 0x1d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x24, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_domain_hourly_stat_proto_rawDescOnce sync.Once
	file_service_server_domain_hourly_stat_proto_rawDescData = file_service_server_domain_hourly_stat_proto_rawDesc
)

func file_service_server_domain_hourly_stat_proto_rawDescGZIP() []byte {
	file_service_server_domain_hourly_stat_proto_rawDescOnce.Do(func() {
		file_service_server_domain_hourly_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_domain_hourly_stat_proto_rawDescData)
	})
	return file_service_server_domain_hourly_stat_proto_rawDescData
}

var file_service_server_domain_hourly_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_server_domain_hourly_stat_proto_goTypes = []interface{}{
	(*ListTopServerDomainStatsWithServerIdRequest)(nil),  // 0: pb.ListTopServerDomainStatsWithServerIdRequest
	(*ListTopServerDomainStatsWithServerIdResponse)(nil), // 1: pb.ListTopServerDomainStatsWithServerIdResponse
	(*ServerDomainHourlyStat)(nil),                       // 2: pb.ServerDomainHourlyStat
}
var file_service_server_domain_hourly_stat_proto_depIdxs = []int32{
	2, // 0: pb.ListTopServerDomainStatsWithServerIdResponse.domainStats:type_name -> pb.ServerDomainHourlyStat
	0, // 1: pb.ServerDomainHourlyStatService.listTopServerDomainStatsWithServerId:input_type -> pb.ListTopServerDomainStatsWithServerIdRequest
	1, // 2: pb.ServerDomainHourlyStatService.listTopServerDomainStatsWithServerId:output_type -> pb.ListTopServerDomainStatsWithServerIdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_server_domain_hourly_stat_proto_init() }
func file_service_server_domain_hourly_stat_proto_init() {
	if File_service_server_domain_hourly_stat_proto != nil {
		return
	}
	file_models_model_server_domain_hourly_stat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_domain_hourly_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopServerDomainStatsWithServerIdRequest); i {
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
		file_service_server_domain_hourly_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopServerDomainStatsWithServerIdResponse); i {
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
			RawDescriptor: file_service_server_domain_hourly_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_domain_hourly_stat_proto_goTypes,
		DependencyIndexes: file_service_server_domain_hourly_stat_proto_depIdxs,
		MessageInfos:      file_service_server_domain_hourly_stat_proto_msgTypes,
	}.Build()
	File_service_server_domain_hourly_stat_proto = out.File
	file_service_server_domain_hourly_stat_proto_rawDesc = nil
	file_service_server_domain_hourly_stat_proto_goTypes = nil
	file_service_server_domain_hourly_stat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerDomainHourlyStatServiceClient is the client API for ServerDomainHourlyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerDomainHourlyStatServiceClient interface {
	// 读取服务的域名排行
	ListTopServerDomainStatsWithServerId(ctx context.Context, in *ListTopServerDomainStatsWithServerIdRequest, opts ...grpc.CallOption) (*ListTopServerDomainStatsWithServerIdResponse, error)
}

type serverDomainHourlyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerDomainHourlyStatServiceClient(cc grpc.ClientConnInterface) ServerDomainHourlyStatServiceClient {
	return &serverDomainHourlyStatServiceClient{cc}
}

func (c *serverDomainHourlyStatServiceClient) ListTopServerDomainStatsWithServerId(ctx context.Context, in *ListTopServerDomainStatsWithServerIdRequest, opts ...grpc.CallOption) (*ListTopServerDomainStatsWithServerIdResponse, error) {
	out := new(ListTopServerDomainStatsWithServerIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDomainHourlyStatService/listTopServerDomainStatsWithServerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerDomainHourlyStatServiceServer is the server API for ServerDomainHourlyStatService service.
type ServerDomainHourlyStatServiceServer interface {
	// 读取服务的域名排行
	ListTopServerDomainStatsWithServerId(context.Context, *ListTopServerDomainStatsWithServerIdRequest) (*ListTopServerDomainStatsWithServerIdResponse, error)
}

// UnimplementedServerDomainHourlyStatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerDomainHourlyStatServiceServer struct {
}

func (*UnimplementedServerDomainHourlyStatServiceServer) ListTopServerDomainStatsWithServerId(context.Context, *ListTopServerDomainStatsWithServerIdRequest) (*ListTopServerDomainStatsWithServerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopServerDomainStatsWithServerId not implemented")
}

func RegisterServerDomainHourlyStatServiceServer(s *grpc.Server, srv ServerDomainHourlyStatServiceServer) {
	s.RegisterService(&_ServerDomainHourlyStatService_serviceDesc, srv)
}

func _ServerDomainHourlyStatService_ListTopServerDomainStatsWithServerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopServerDomainStatsWithServerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDomainHourlyStatServiceServer).ListTopServerDomainStatsWithServerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDomainHourlyStatService/ListTopServerDomainStatsWithServerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDomainHourlyStatServiceServer).ListTopServerDomainStatsWithServerId(ctx, req.(*ListTopServerDomainStatsWithServerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerDomainHourlyStatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerDomainHourlyStatService",
	HandlerType: (*ServerDomainHourlyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listTopServerDomainStatsWithServerId",
			Handler:    _ServerDomainHourlyStatService_ListTopServerDomainStatsWithServerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_domain_hourly_stat.proto",
}
