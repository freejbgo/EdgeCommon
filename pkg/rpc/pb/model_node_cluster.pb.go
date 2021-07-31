// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: models/model_node_cluster.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type NodeCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	NodeGrantId          int64  `protobuf:"varint,4,opt,name=nodeGrantId,proto3" json:"nodeGrantId,omitempty"`
	InstallDir           string `protobuf:"bytes,5,opt,name=installDir,proto3" json:"installDir,omitempty"`
	UniqueId             string `protobuf:"bytes,6,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret               string `protobuf:"bytes,7,opt,name=secret,proto3" json:"secret,omitempty"`
	DnsName              string `protobuf:"bytes,8,opt,name=dnsName,proto3" json:"dnsName,omitempty"`
	DnsDomainId          int64  `protobuf:"varint,9,opt,name=dnsDomainId,proto3" json:"dnsDomainId,omitempty"`
	HttpCachePolicyId    int64  `protobuf:"varint,10,opt,name=httpCachePolicyId,proto3" json:"httpCachePolicyId,omitempty"`
	HttpFirewallPolicyId int64  `protobuf:"varint,11,opt,name=httpFirewallPolicyId,proto3" json:"httpFirewallPolicyId,omitempty"`
	IsOn                 bool   `protobuf:"varint,12,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *NodeCluster) Reset() {
	*x = NodeCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_node_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeCluster) ProtoMessage() {}

func (x *NodeCluster) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_node_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeCluster.ProtoReflect.Descriptor instead.
func (*NodeCluster) Descriptor() ([]byte, []int) {
	return file_models_model_node_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *NodeCluster) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeCluster) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NodeCluster) GetNodeGrantId() int64 {
	if x != nil {
		return x.NodeGrantId
	}
	return 0
}

func (x *NodeCluster) GetInstallDir() string {
	if x != nil {
		return x.InstallDir
	}
	return ""
}

func (x *NodeCluster) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *NodeCluster) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *NodeCluster) GetDnsName() string {
	if x != nil {
		return x.DnsName
	}
	return ""
}

func (x *NodeCluster) GetDnsDomainId() int64 {
	if x != nil {
		return x.DnsDomainId
	}
	return 0
}

func (x *NodeCluster) GetHttpCachePolicyId() int64 {
	if x != nil {
		return x.HttpCachePolicyId
	}
	return 0
}

func (x *NodeCluster) GetHttpFirewallPolicyId() int64 {
	if x != nil {
		return x.HttpFirewallPolicyId
	}
	return 0
}

func (x *NodeCluster) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

var File_models_model_node_cluster_proto protoreflect.FileDescriptor

var file_models_model_node_cluster_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xf7, 0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x6e,
	0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x68, 0x74, 0x74,
	0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x68, 0x74, 0x74, 0x70, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_node_cluster_proto_rawDescOnce sync.Once
	file_models_model_node_cluster_proto_rawDescData = file_models_model_node_cluster_proto_rawDesc
)

func file_models_model_node_cluster_proto_rawDescGZIP() []byte {
	file_models_model_node_cluster_proto_rawDescOnce.Do(func() {
		file_models_model_node_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_node_cluster_proto_rawDescData)
	})
	return file_models_model_node_cluster_proto_rawDescData
}

var file_models_model_node_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_node_cluster_proto_goTypes = []interface{}{
	(*NodeCluster)(nil), // 0: pb.NodeCluster
}
var file_models_model_node_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_node_cluster_proto_init() }
func file_models_model_node_cluster_proto_init() {
	if File_models_model_node_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_node_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeCluster); i {
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
			RawDescriptor: file_models_model_node_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_node_cluster_proto_goTypes,
		DependencyIndexes: file_models_model_node_cluster_proto_depIdxs,
		MessageInfos:      file_models_model_node_cluster_proto_msgTypes,
	}.Build()
	File_models_model_node_cluster_proto = out.File
	file_models_model_node_cluster_proto_rawDesc = nil
	file_models_model_node_cluster_proto_goTypes = nil
	file_models_model_node_cluster_proto_depIdxs = nil
}
