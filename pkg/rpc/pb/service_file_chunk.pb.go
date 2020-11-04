// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_file_chunk.proto

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

// 创建文件片段
type CreateFileChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId int64  `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateFileChunkRequest) Reset() {
	*x = CreateFileChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileChunkRequest) ProtoMessage() {}

func (x *CreateFileChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileChunkRequest.ProtoReflect.Descriptor instead.
func (*CreateFileChunkRequest) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFileChunkRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *CreateFileChunkRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateFileChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunkId int64 `protobuf:"varint,1,opt,name=fileChunkId,proto3" json:"fileChunkId,omitempty"`
}

func (x *CreateFileChunkResponse) Reset() {
	*x = CreateFileChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileChunkResponse) ProtoMessage() {}

func (x *CreateFileChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileChunkResponse.ProtoReflect.Descriptor instead.
func (*CreateFileChunkResponse) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFileChunkResponse) GetFileChunkId() int64 {
	if x != nil {
		return x.FileChunkId
	}
	return 0
}

// 获取的一个文件的所有片段IDs
type FindAllFileChunkIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId int64 `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *FindAllFileChunkIdsRequest) Reset() {
	*x = FindAllFileChunkIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllFileChunkIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllFileChunkIdsRequest) ProtoMessage() {}

func (x *FindAllFileChunkIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllFileChunkIdsRequest.ProtoReflect.Descriptor instead.
func (*FindAllFileChunkIdsRequest) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllFileChunkIdsRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FindAllFileChunkIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunkIds []int64 `protobuf:"varint,1,rep,packed,name=fileChunkIds,proto3" json:"fileChunkIds,omitempty"`
}

func (x *FindAllFileChunkIdsResponse) Reset() {
	*x = FindAllFileChunkIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllFileChunkIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllFileChunkIdsResponse) ProtoMessage() {}

func (x *FindAllFileChunkIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllFileChunkIdsResponse.ProtoReflect.Descriptor instead.
func (*FindAllFileChunkIdsResponse) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllFileChunkIdsResponse) GetFileChunkIds() []int64 {
	if x != nil {
		return x.FileChunkIds
	}
	return nil
}

// 下载文件片段
type DownloadFileChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunkId int64 `protobuf:"varint,1,opt,name=fileChunkId,proto3" json:"fileChunkId,omitempty"`
}

func (x *DownloadFileChunkRequest) Reset() {
	*x = DownloadFileChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileChunkRequest) ProtoMessage() {}

func (x *DownloadFileChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileChunkRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileChunkRequest) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadFileChunkRequest) GetFileChunkId() int64 {
	if x != nil {
		return x.FileChunkId
	}
	return 0
}

type DownloadFileChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk *FileChunk `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
}

func (x *DownloadFileChunkResponse) Reset() {
	*x = DownloadFileChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_chunk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileChunkResponse) ProtoMessage() {}

func (x *DownloadFileChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_chunk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileChunkResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileChunkResponse) Descriptor() ([]byte, []int) {
	return file_service_file_chunk_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadFileChunkResponse) GetFileChunk() *FileChunk {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

var File_service_file_chunk_proto protoreflect.FileDescriptor

var file_service_file_chunk_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x16,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x1a, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x64, 0x73, 0x22, 0x3c, 0x0a, 0x18, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64,
	0x22, 0x48, 0x0a, 0x19, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x88, 0x02, 0x0a, 0x10, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x66,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_file_chunk_proto_rawDescOnce sync.Once
	file_service_file_chunk_proto_rawDescData = file_service_file_chunk_proto_rawDesc
)

func file_service_file_chunk_proto_rawDescGZIP() []byte {
	file_service_file_chunk_proto_rawDescOnce.Do(func() {
		file_service_file_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_file_chunk_proto_rawDescData)
	})
	return file_service_file_chunk_proto_rawDescData
}

var file_service_file_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_file_chunk_proto_goTypes = []interface{}{
	(*CreateFileChunkRequest)(nil),      // 0: pb.CreateFileChunkRequest
	(*CreateFileChunkResponse)(nil),     // 1: pb.CreateFileChunkResponse
	(*FindAllFileChunkIdsRequest)(nil),  // 2: pb.FindAllFileChunkIdsRequest
	(*FindAllFileChunkIdsResponse)(nil), // 3: pb.FindAllFileChunkIdsResponse
	(*DownloadFileChunkRequest)(nil),    // 4: pb.DownloadFileChunkRequest
	(*DownloadFileChunkResponse)(nil),   // 5: pb.DownloadFileChunkResponse
	(*FileChunk)(nil),                   // 6: pb.FileChunk
}
var file_service_file_chunk_proto_depIdxs = []int32{
	6, // 0: pb.DownloadFileChunkResponse.fileChunk:type_name -> pb.FileChunk
	0, // 1: pb.FileChunkService.createFileChunk:input_type -> pb.CreateFileChunkRequest
	2, // 2: pb.FileChunkService.findAllFileChunkIds:input_type -> pb.FindAllFileChunkIdsRequest
	4, // 3: pb.FileChunkService.downloadFileChunk:input_type -> pb.DownloadFileChunkRequest
	1, // 4: pb.FileChunkService.createFileChunk:output_type -> pb.CreateFileChunkResponse
	3, // 5: pb.FileChunkService.findAllFileChunkIds:output_type -> pb.FindAllFileChunkIdsResponse
	5, // 6: pb.FileChunkService.downloadFileChunk:output_type -> pb.DownloadFileChunkResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_file_chunk_proto_init() }
func file_service_file_chunk_proto_init() {
	if File_service_file_chunk_proto != nil {
		return
	}
	file_model_file_chunk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_file_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileChunkRequest); i {
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
		file_service_file_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileChunkResponse); i {
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
		file_service_file_chunk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllFileChunkIdsRequest); i {
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
		file_service_file_chunk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllFileChunkIdsResponse); i {
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
		file_service_file_chunk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileChunkRequest); i {
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
		file_service_file_chunk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileChunkResponse); i {
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
			RawDescriptor: file_service_file_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_file_chunk_proto_goTypes,
		DependencyIndexes: file_service_file_chunk_proto_depIdxs,
		MessageInfos:      file_service_file_chunk_proto_msgTypes,
	}.Build()
	File_service_file_chunk_proto = out.File
	file_service_file_chunk_proto_rawDesc = nil
	file_service_file_chunk_proto_goTypes = nil
	file_service_file_chunk_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileChunkServiceClient is the client API for FileChunkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileChunkServiceClient interface {
	// 创建文件片段
	CreateFileChunk(ctx context.Context, in *CreateFileChunkRequest, opts ...grpc.CallOption) (*CreateFileChunkResponse, error)
	// 获取的一个文件的所有片段IDs
	FindAllFileChunkIds(ctx context.Context, in *FindAllFileChunkIdsRequest, opts ...grpc.CallOption) (*FindAllFileChunkIdsResponse, error)
	// 下载文件片段
	DownloadFileChunk(ctx context.Context, in *DownloadFileChunkRequest, opts ...grpc.CallOption) (*DownloadFileChunkResponse, error)
}

type fileChunkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileChunkServiceClient(cc grpc.ClientConnInterface) FileChunkServiceClient {
	return &fileChunkServiceClient{cc}
}

func (c *fileChunkServiceClient) CreateFileChunk(ctx context.Context, in *CreateFileChunkRequest, opts ...grpc.CallOption) (*CreateFileChunkResponse, error) {
	out := new(CreateFileChunkResponse)
	err := c.cc.Invoke(ctx, "/pb.FileChunkService/createFileChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileChunkServiceClient) FindAllFileChunkIds(ctx context.Context, in *FindAllFileChunkIdsRequest, opts ...grpc.CallOption) (*FindAllFileChunkIdsResponse, error) {
	out := new(FindAllFileChunkIdsResponse)
	err := c.cc.Invoke(ctx, "/pb.FileChunkService/findAllFileChunkIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileChunkServiceClient) DownloadFileChunk(ctx context.Context, in *DownloadFileChunkRequest, opts ...grpc.CallOption) (*DownloadFileChunkResponse, error) {
	out := new(DownloadFileChunkResponse)
	err := c.cc.Invoke(ctx, "/pb.FileChunkService/downloadFileChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileChunkServiceServer is the server API for FileChunkService service.
type FileChunkServiceServer interface {
	// 创建文件片段
	CreateFileChunk(context.Context, *CreateFileChunkRequest) (*CreateFileChunkResponse, error)
	// 获取的一个文件的所有片段IDs
	FindAllFileChunkIds(context.Context, *FindAllFileChunkIdsRequest) (*FindAllFileChunkIdsResponse, error)
	// 下载文件片段
	DownloadFileChunk(context.Context, *DownloadFileChunkRequest) (*DownloadFileChunkResponse, error)
}

// UnimplementedFileChunkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileChunkServiceServer struct {
}

func (*UnimplementedFileChunkServiceServer) CreateFileChunk(context.Context, *CreateFileChunkRequest) (*CreateFileChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileChunk not implemented")
}
func (*UnimplementedFileChunkServiceServer) FindAllFileChunkIds(context.Context, *FindAllFileChunkIdsRequest) (*FindAllFileChunkIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllFileChunkIds not implemented")
}
func (*UnimplementedFileChunkServiceServer) DownloadFileChunk(context.Context, *DownloadFileChunkRequest) (*DownloadFileChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFileChunk not implemented")
}

func RegisterFileChunkServiceServer(s *grpc.Server, srv FileChunkServiceServer) {
	s.RegisterService(&_FileChunkService_serviceDesc, srv)
}

func _FileChunkService_CreateFileChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileChunkServiceServer).CreateFileChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileChunkService/CreateFileChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileChunkServiceServer).CreateFileChunk(ctx, req.(*CreateFileChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileChunkService_FindAllFileChunkIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllFileChunkIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileChunkServiceServer).FindAllFileChunkIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileChunkService/FindAllFileChunkIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileChunkServiceServer).FindAllFileChunkIds(ctx, req.(*FindAllFileChunkIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileChunkService_DownloadFileChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileChunkServiceServer).DownloadFileChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileChunkService/DownloadFileChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileChunkServiceServer).DownloadFileChunk(ctx, req.(*DownloadFileChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileChunkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FileChunkService",
	HandlerType: (*FileChunkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createFileChunk",
			Handler:    _FileChunkService_CreateFileChunk_Handler,
		},
		{
			MethodName: "findAllFileChunkIds",
			Handler:    _FileChunkService_FindAllFileChunkIds_Handler,
		},
		{
			MethodName: "downloadFileChunk",
			Handler:    _FileChunkService_DownloadFileChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_file_chunk.proto",
}
