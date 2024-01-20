// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_http_firewall_policy.proto

package pb

import (
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

// WAF策略
type HTTPFirewallPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                 // 策略ID
	Name               string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                              // 名称
	Mode               string `protobuf:"bytes,7,opt,name=mode,proto3" json:"mode,omitempty"`                              // 模式
	IsOn               bool   `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`                             // 是否启用
	Description        string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`                // 描述
	InboundJSON        []byte `protobuf:"bytes,5,opt,name=inboundJSON,proto3" json:"inboundJSON,omitempty"`                // 入站配置
	OutboundJSON       []byte `protobuf:"bytes,6,opt,name=outboundJSON,proto3" json:"outboundJSON,omitempty"`              // 出站配置
	ServerId           int64  `protobuf:"varint,8,opt,name=serverId,proto3" json:"serverId,omitempty"`                     // 所属网站ID（如果为0表示公共策略）
	UseLocalFirewall   bool   `protobuf:"varint,9,opt,name=useLocalFirewall,proto3" json:"useLocalFirewall,omitempty"`     // 是否使用本机防火墙
	SynFloodJSON       []byte `protobuf:"bytes,10,opt,name=synFloodJSON,proto3" json:"synFloodJSON,omitempty"`             // synflood配置
	BlockOptionsJSON   []byte `protobuf:"bytes,11,opt,name=blockOptionsJSON,proto3" json:"blockOptionsJSON,omitempty"`     // 阻止动作配置
	PageOptionsJSON    []byte `protobuf:"bytes,13,opt,name=pageOptionsJSON,proto3" json:"pageOptionsJSON,omitempty"`       // 显示网页动作配置
	CaptchaOptionsJSON []byte `protobuf:"bytes,12,opt,name=captchaOptionsJSON,proto3" json:"captchaOptionsJSON,omitempty"` // 人机识别配置
}

func (x *HTTPFirewallPolicy) Reset() {
	*x = HTTPFirewallPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_firewall_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFirewallPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFirewallPolicy) ProtoMessage() {}

func (x *HTTPFirewallPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_firewall_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFirewallPolicy.ProtoReflect.Descriptor instead.
func (*HTTPFirewallPolicy) Descriptor() ([]byte, []int) {
	return file_models_model_http_firewall_policy_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPFirewallPolicy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HTTPFirewallPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HTTPFirewallPolicy) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *HTTPFirewallPolicy) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *HTTPFirewallPolicy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HTTPFirewallPolicy) GetInboundJSON() []byte {
	if x != nil {
		return x.InboundJSON
	}
	return nil
}

func (x *HTTPFirewallPolicy) GetOutboundJSON() []byte {
	if x != nil {
		return x.OutboundJSON
	}
	return nil
}

func (x *HTTPFirewallPolicy) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *HTTPFirewallPolicy) GetUseLocalFirewall() bool {
	if x != nil {
		return x.UseLocalFirewall
	}
	return false
}

func (x *HTTPFirewallPolicy) GetSynFloodJSON() []byte {
	if x != nil {
		return x.SynFloodJSON
	}
	return nil
}

func (x *HTTPFirewallPolicy) GetBlockOptionsJSON() []byte {
	if x != nil {
		return x.BlockOptionsJSON
	}
	return nil
}

func (x *HTTPFirewallPolicy) GetPageOptionsJSON() []byte {
	if x != nil {
		return x.PageOptionsJSON
	}
	return nil
}

func (x *HTTPFirewallPolicy) GetCaptchaOptionsJSON() []byte {
	if x != nil {
		return x.CaptchaOptionsJSON
	}
	return nil
}

var File_models_model_http_firewall_policy_proto protoreflect.FileDescriptor

var file_models_model_http_firewall_policy_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xba, 0x03,
	0x0a, 0x12, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x75, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x46, 0x6c, 0x6f, 0x6f, 0x64, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x46, 0x6c, 0x6f, 0x6f, 0x64,
	0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_http_firewall_policy_proto_rawDescOnce sync.Once
	file_models_model_http_firewall_policy_proto_rawDescData = file_models_model_http_firewall_policy_proto_rawDesc
)

func file_models_model_http_firewall_policy_proto_rawDescGZIP() []byte {
	file_models_model_http_firewall_policy_proto_rawDescOnce.Do(func() {
		file_models_model_http_firewall_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_http_firewall_policy_proto_rawDescData)
	})
	return file_models_model_http_firewall_policy_proto_rawDescData
}

var file_models_model_http_firewall_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_http_firewall_policy_proto_goTypes = []interface{}{
	(*HTTPFirewallPolicy)(nil), // 0: pb.HTTPFirewallPolicy
}
var file_models_model_http_firewall_policy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_http_firewall_policy_proto_init() }
func file_models_model_http_firewall_policy_proto_init() {
	if File_models_model_http_firewall_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_http_firewall_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFirewallPolicy); i {
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
			RawDescriptor: file_models_model_http_firewall_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_http_firewall_policy_proto_goTypes,
		DependencyIndexes: file_models_model_http_firewall_policy_proto_depIdxs,
		MessageInfos:      file_models_model_http_firewall_policy_proto_msgTypes,
	}.Build()
	File_models_model_http_firewall_policy_proto = out.File
	file_models_model_http_firewall_policy_proto_rawDesc = nil
	file_models_model_http_firewall_policy_proto_goTypes = nil
	file_models_model_http_firewall_policy_proto_depIdxs = nil
}
