// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_region_provider.proto

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

// 查找所有ISP
type FindAllEnabledRegionProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllEnabledRegionProvidersRequest) Reset() {
	*x = FindAllEnabledRegionProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvidersRequest) ProtoMessage() {}

func (x *FindAllEnabledRegionProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvidersRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvidersRequest) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{0}
}

type FindAllEnabledRegionProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProviders []*RegionProvider `protobuf:"bytes,1,rep,name=regionProviders,proto3" json:"regionProviders,omitempty"`
}

func (x *FindAllEnabledRegionProvidersResponse) Reset() {
	*x = FindAllEnabledRegionProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvidersResponse) ProtoMessage() {}

func (x *FindAllEnabledRegionProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvidersResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvidersResponse) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllEnabledRegionProvidersResponse) GetRegionProviders() []*RegionProvider {
	if x != nil {
		return x.RegionProviders
	}
	return nil
}

// 查找单个ISP信息
type FindEnabledRegionProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProviderId int64 `protobuf:"varint,1,opt,name=regionProviderId,proto3" json:"regionProviderId,omitempty"`
}

func (x *FindEnabledRegionProviderRequest) Reset() {
	*x = FindEnabledRegionProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProviderRequest) ProtoMessage() {}

func (x *FindEnabledRegionProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProviderRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProviderRequest) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{2}
}

func (x *FindEnabledRegionProviderRequest) GetRegionProviderId() int64 {
	if x != nil {
		return x.RegionProviderId
	}
	return 0
}

type FindEnabledRegionProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvider *RegionProvider `protobuf:"bytes,1,opt,name=regionProvider,proto3" json:"regionProvider,omitempty"`
}

func (x *FindEnabledRegionProviderResponse) Reset() {
	*x = FindEnabledRegionProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProviderResponse) ProtoMessage() {}

func (x *FindEnabledRegionProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProviderResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProviderResponse) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledRegionProviderResponse) GetRegionProvider() *RegionProvider {
	if x != nil {
		return x.RegionProvider
	}
	return nil
}

// 查找所有ISP
type FindAllRegionProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRegionProvidersRequest) Reset() {
	*x = FindAllRegionProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvidersRequest) ProtoMessage() {}

func (x *FindAllRegionProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvidersRequest.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvidersRequest) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{4}
}

type FindAllRegionProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProviders []*RegionProvider `protobuf:"bytes,1,rep,name=regionProviders,proto3" json:"regionProviders,omitempty"`
}

func (x *FindAllRegionProvidersResponse) Reset() {
	*x = FindAllRegionProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvidersResponse) ProtoMessage() {}

func (x *FindAllRegionProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvidersResponse.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvidersResponse) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllRegionProvidersResponse) GetRegionProviders() []*RegionProvider {
	if x != nil {
		return x.RegionProviders
	}
	return nil
}

// 查找单个ISP信息
type FindRegionProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProviderId int64 `protobuf:"varint,1,opt,name=regionProviderId,proto3" json:"regionProviderId,omitempty"`
}

func (x *FindRegionProviderRequest) Reset() {
	*x = FindRegionProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionProviderRequest) ProtoMessage() {}

func (x *FindRegionProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionProviderRequest.ProtoReflect.Descriptor instead.
func (*FindRegionProviderRequest) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{6}
}

func (x *FindRegionProviderRequest) GetRegionProviderId() int64 {
	if x != nil {
		return x.RegionProviderId
	}
	return 0
}

type FindRegionProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvider *RegionProvider `protobuf:"bytes,1,opt,name=regionProvider,proto3" json:"regionProvider,omitempty"`
}

func (x *FindRegionProviderResponse) Reset() {
	*x = FindRegionProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionProviderResponse) ProtoMessage() {}

func (x *FindRegionProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionProviderResponse.ProtoReflect.Descriptor instead.
func (*FindRegionProviderResponse) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{7}
}

func (x *FindRegionProviderResponse) GetRegionProvider() *RegionProvider {
	if x != nil {
		return x.RegionProvider
	}
	return nil
}

// 修改ISP定制信息
type UpdateRegionProviderCustomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProviderId int64    `protobuf:"varint,1,opt,name=regionProviderId,proto3" json:"regionProviderId,omitempty"`
	CustomName       string   `protobuf:"bytes,2,opt,name=customName,proto3" json:"customName,omitempty"`
	CustomCodes      []string `protobuf:"bytes,3,rep,name=customCodes,proto3" json:"customCodes,omitempty"`
}

func (x *UpdateRegionProviderCustomRequest) Reset() {
	*x = UpdateRegionProviderCustomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_provider_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegionProviderCustomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionProviderCustomRequest) ProtoMessage() {}

func (x *UpdateRegionProviderCustomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_provider_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionProviderCustomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionProviderCustomRequest) Descriptor() ([]byte, []int) {
	return file_service_region_provider_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRegionProviderCustomRequest) GetRegionProviderId() int64 {
	if x != nil {
		return x.RegionProviderId
	}
	return 0
}

func (x *UpdateRegionProviderCustomRequest) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

func (x *UpdateRegionProviderCustomRequest) GetCustomCodes() []string {
	if x != nil {
		return x.CustomCodes
	}
	return nil
}

var File_service_region_provider_proto protoreflect.FileDescriptor

var file_service_region_provider_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x25, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x4e, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x5f, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x1f, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x1a,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x91, 0x01, 0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x8c, 0x04, 0x0a, 0x15, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x1d, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12,
	0x6d, 0x0a, 0x19, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x5f,
	0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_region_provider_proto_rawDescOnce sync.Once
	file_service_region_provider_proto_rawDescData = file_service_region_provider_proto_rawDesc
)

func file_service_region_provider_proto_rawDescGZIP() []byte {
	file_service_region_provider_proto_rawDescOnce.Do(func() {
		file_service_region_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_region_provider_proto_rawDescData)
	})
	return file_service_region_provider_proto_rawDescData
}

var file_service_region_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_region_provider_proto_goTypes = []interface{}{
	(*FindAllEnabledRegionProvidersRequest)(nil),  // 0: pb.FindAllEnabledRegionProvidersRequest
	(*FindAllEnabledRegionProvidersResponse)(nil), // 1: pb.FindAllEnabledRegionProvidersResponse
	(*FindEnabledRegionProviderRequest)(nil),      // 2: pb.FindEnabledRegionProviderRequest
	(*FindEnabledRegionProviderResponse)(nil),     // 3: pb.FindEnabledRegionProviderResponse
	(*FindAllRegionProvidersRequest)(nil),         // 4: pb.FindAllRegionProvidersRequest
	(*FindAllRegionProvidersResponse)(nil),        // 5: pb.FindAllRegionProvidersResponse
	(*FindRegionProviderRequest)(nil),             // 6: pb.FindRegionProviderRequest
	(*FindRegionProviderResponse)(nil),            // 7: pb.FindRegionProviderResponse
	(*UpdateRegionProviderCustomRequest)(nil),     // 8: pb.UpdateRegionProviderCustomRequest
	(*RegionProvider)(nil),                        // 9: pb.RegionProvider
	(*RPCSuccess)(nil),                            // 10: pb.RPCSuccess
}
var file_service_region_provider_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllEnabledRegionProvidersResponse.regionProviders:type_name -> pb.RegionProvider
	9,  // 1: pb.FindEnabledRegionProviderResponse.regionProvider:type_name -> pb.RegionProvider
	9,  // 2: pb.FindAllRegionProvidersResponse.regionProviders:type_name -> pb.RegionProvider
	9,  // 3: pb.FindRegionProviderResponse.regionProvider:type_name -> pb.RegionProvider
	0,  // 4: pb.RegionProviderService.findAllEnabledRegionProviders:input_type -> pb.FindAllEnabledRegionProvidersRequest
	2,  // 5: pb.RegionProviderService.findEnabledRegionProvider:input_type -> pb.FindEnabledRegionProviderRequest
	4,  // 6: pb.RegionProviderService.findAllRegionProviders:input_type -> pb.FindAllRegionProvidersRequest
	6,  // 7: pb.RegionProviderService.findRegionProvider:input_type -> pb.FindRegionProviderRequest
	8,  // 8: pb.RegionProviderService.updateRegionProviderCustom:input_type -> pb.UpdateRegionProviderCustomRequest
	1,  // 9: pb.RegionProviderService.findAllEnabledRegionProviders:output_type -> pb.FindAllEnabledRegionProvidersResponse
	3,  // 10: pb.RegionProviderService.findEnabledRegionProvider:output_type -> pb.FindEnabledRegionProviderResponse
	5,  // 11: pb.RegionProviderService.findAllRegionProviders:output_type -> pb.FindAllRegionProvidersResponse
	7,  // 12: pb.RegionProviderService.findRegionProvider:output_type -> pb.FindRegionProviderResponse
	10, // 13: pb.RegionProviderService.updateRegionProviderCustom:output_type -> pb.RPCSuccess
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_service_region_provider_proto_init() }
func file_service_region_provider_proto_init() {
	if File_service_region_provider_proto != nil {
		return
	}
	file_models_model_region_provider_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_region_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvidersRequest); i {
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
		file_service_region_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvidersResponse); i {
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
		file_service_region_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProviderRequest); i {
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
		file_service_region_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProviderResponse); i {
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
		file_service_region_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvidersRequest); i {
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
		file_service_region_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvidersResponse); i {
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
		file_service_region_provider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionProviderRequest); i {
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
		file_service_region_provider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionProviderResponse); i {
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
		file_service_region_provider_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegionProviderCustomRequest); i {
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
			RawDescriptor: file_service_region_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_region_provider_proto_goTypes,
		DependencyIndexes: file_service_region_provider_proto_depIdxs,
		MessageInfos:      file_service_region_provider_proto_msgTypes,
	}.Build()
	File_service_region_provider_proto = out.File
	file_service_region_provider_proto_rawDesc = nil
	file_service_region_provider_proto_goTypes = nil
	file_service_region_provider_proto_depIdxs = nil
}
