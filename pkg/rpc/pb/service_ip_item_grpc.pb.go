// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ip_item.proto

package pb

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

const (
	IPItemService_CreateIPItem_FullMethodName            = "/pb.IPItemService/createIPItem"
	IPItemService_CreateIPItems_FullMethodName           = "/pb.IPItemService/createIPItems"
	IPItemService_UpdateIPItem_FullMethodName            = "/pb.IPItemService/updateIPItem"
	IPItemService_DeleteIPItem_FullMethodName            = "/pb.IPItemService/deleteIPItem"
	IPItemService_DeleteIPItems_FullMethodName           = "/pb.IPItemService/deleteIPItems"
	IPItemService_CountIPItemsWithListId_FullMethodName  = "/pb.IPItemService/countIPItemsWithListId"
	IPItemService_ListIPItemsWithListId_FullMethodName   = "/pb.IPItemService/listIPItemsWithListId"
	IPItemService_FindEnabledIPItem_FullMethodName       = "/pb.IPItemService/findEnabledIPItem"
	IPItemService_ListIPItemsAfterVersion_FullMethodName = "/pb.IPItemService/listIPItemsAfterVersion"
	IPItemService_CheckIPItemStatus_FullMethodName       = "/pb.IPItemService/checkIPItemStatus"
	IPItemService_ExistsEnabledIPItem_FullMethodName     = "/pb.IPItemService/existsEnabledIPItem"
	IPItemService_CountAllEnabledIPItems_FullMethodName  = "/pb.IPItemService/countAllEnabledIPItems"
	IPItemService_ListAllEnabledIPItems_FullMethodName   = "/pb.IPItemService/listAllEnabledIPItems"
	IPItemService_UpdateIPItemsRead_FullMethodName       = "/pb.IPItemService/updateIPItemsRead"
)

// IPItemServiceClient is the client API for IPItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPItemServiceClient interface {
	// 创建IP
	CreateIPItem(ctx context.Context, in *CreateIPItemRequest, opts ...grpc.CallOption) (*CreateIPItemResponse, error)
	// 创建一组IP
	CreateIPItems(ctx context.Context, in *CreateIPItemsRequest, opts ...grpc.CallOption) (*CreateIPItemsResponse, error)
	// 修改IP
	UpdateIPItem(ctx context.Context, in *UpdateIPItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除IP
	DeleteIPItem(ctx context.Context, in *DeleteIPItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 批量删除IP
	DeleteIPItems(ctx context.Context, in *DeleteIPItemsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算IP数量
	CountIPItemsWithListId(ctx context.Context, in *CountIPItemsWithListIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的IP
	ListIPItemsWithListId(ctx context.Context, in *ListIPItemsWithListIdRequest, opts ...grpc.CallOption) (*ListIPItemsWithListIdResponse, error)
	// 查找单个IP
	FindEnabledIPItem(ctx context.Context, in *FindEnabledIPItemRequest, opts ...grpc.CallOption) (*FindEnabledIPItemResponse, error)
	// 根据版本列出一组IP
	ListIPItemsAfterVersion(ctx context.Context, in *ListIPItemsAfterVersionRequest, opts ...grpc.CallOption) (*ListIPItemsAfterVersionResponse, error)
	// 检查IP状态
	CheckIPItemStatus(ctx context.Context, in *CheckIPItemStatusRequest, opts ...grpc.CallOption) (*CheckIPItemStatusResponse, error)
	// 检查IP是否存在
	ExistsEnabledIPItem(ctx context.Context, in *ExistsEnabledIPItemRequest, opts ...grpc.CallOption) (*ExistsEnabledIPItemResponse, error)
	// 计算所有IP数量
	CountAllEnabledIPItems(ctx context.Context, in *CountAllEnabledIPItemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出所有名单中的IP
	ListAllEnabledIPItems(ctx context.Context, in *ListAllEnabledIPItemsRequest, opts ...grpc.CallOption) (*ListAllEnabledIPItemsResponse, error)
	// 设置所有为已读
	UpdateIPItemsRead(ctx context.Context, in *UpdateIPItemsReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type iPItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPItemServiceClient(cc grpc.ClientConnInterface) IPItemServiceClient {
	return &iPItemServiceClient{cc}
}

func (c *iPItemServiceClient) CreateIPItem(ctx context.Context, in *CreateIPItemRequest, opts ...grpc.CallOption) (*CreateIPItemResponse, error) {
	out := new(CreateIPItemResponse)
	err := c.cc.Invoke(ctx, IPItemService_CreateIPItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) CreateIPItems(ctx context.Context, in *CreateIPItemsRequest, opts ...grpc.CallOption) (*CreateIPItemsResponse, error) {
	out := new(CreateIPItemsResponse)
	err := c.cc.Invoke(ctx, IPItemService_CreateIPItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) UpdateIPItem(ctx context.Context, in *UpdateIPItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPItemService_UpdateIPItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) DeleteIPItem(ctx context.Context, in *DeleteIPItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPItemService_DeleteIPItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) DeleteIPItems(ctx context.Context, in *DeleteIPItemsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPItemService_DeleteIPItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) CountIPItemsWithListId(ctx context.Context, in *CountIPItemsWithListIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, IPItemService_CountIPItemsWithListId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) ListIPItemsWithListId(ctx context.Context, in *ListIPItemsWithListIdRequest, opts ...grpc.CallOption) (*ListIPItemsWithListIdResponse, error) {
	out := new(ListIPItemsWithListIdResponse)
	err := c.cc.Invoke(ctx, IPItemService_ListIPItemsWithListId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) FindEnabledIPItem(ctx context.Context, in *FindEnabledIPItemRequest, opts ...grpc.CallOption) (*FindEnabledIPItemResponse, error) {
	out := new(FindEnabledIPItemResponse)
	err := c.cc.Invoke(ctx, IPItemService_FindEnabledIPItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) ListIPItemsAfterVersion(ctx context.Context, in *ListIPItemsAfterVersionRequest, opts ...grpc.CallOption) (*ListIPItemsAfterVersionResponse, error) {
	out := new(ListIPItemsAfterVersionResponse)
	err := c.cc.Invoke(ctx, IPItemService_ListIPItemsAfterVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) CheckIPItemStatus(ctx context.Context, in *CheckIPItemStatusRequest, opts ...grpc.CallOption) (*CheckIPItemStatusResponse, error) {
	out := new(CheckIPItemStatusResponse)
	err := c.cc.Invoke(ctx, IPItemService_CheckIPItemStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) ExistsEnabledIPItem(ctx context.Context, in *ExistsEnabledIPItemRequest, opts ...grpc.CallOption) (*ExistsEnabledIPItemResponse, error) {
	out := new(ExistsEnabledIPItemResponse)
	err := c.cc.Invoke(ctx, IPItemService_ExistsEnabledIPItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) CountAllEnabledIPItems(ctx context.Context, in *CountAllEnabledIPItemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, IPItemService_CountAllEnabledIPItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) ListAllEnabledIPItems(ctx context.Context, in *ListAllEnabledIPItemsRequest, opts ...grpc.CallOption) (*ListAllEnabledIPItemsResponse, error) {
	out := new(ListAllEnabledIPItemsResponse)
	err := c.cc.Invoke(ctx, IPItemService_ListAllEnabledIPItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPItemServiceClient) UpdateIPItemsRead(ctx context.Context, in *UpdateIPItemsReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPItemService_UpdateIPItemsRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPItemServiceServer is the server API for IPItemService service.
// All implementations should embed UnimplementedIPItemServiceServer
// for forward compatibility
type IPItemServiceServer interface {
	// 创建IP
	CreateIPItem(context.Context, *CreateIPItemRequest) (*CreateIPItemResponse, error)
	// 创建一组IP
	CreateIPItems(context.Context, *CreateIPItemsRequest) (*CreateIPItemsResponse, error)
	// 修改IP
	UpdateIPItem(context.Context, *UpdateIPItemRequest) (*RPCSuccess, error)
	// 删除IP
	DeleteIPItem(context.Context, *DeleteIPItemRequest) (*RPCSuccess, error)
	// 批量删除IP
	DeleteIPItems(context.Context, *DeleteIPItemsRequest) (*RPCSuccess, error)
	// 计算IP数量
	CountIPItemsWithListId(context.Context, *CountIPItemsWithListIdRequest) (*RPCCountResponse, error)
	// 列出单页的IP
	ListIPItemsWithListId(context.Context, *ListIPItemsWithListIdRequest) (*ListIPItemsWithListIdResponse, error)
	// 查找单个IP
	FindEnabledIPItem(context.Context, *FindEnabledIPItemRequest) (*FindEnabledIPItemResponse, error)
	// 根据版本列出一组IP
	ListIPItemsAfterVersion(context.Context, *ListIPItemsAfterVersionRequest) (*ListIPItemsAfterVersionResponse, error)
	// 检查IP状态
	CheckIPItemStatus(context.Context, *CheckIPItemStatusRequest) (*CheckIPItemStatusResponse, error)
	// 检查IP是否存在
	ExistsEnabledIPItem(context.Context, *ExistsEnabledIPItemRequest) (*ExistsEnabledIPItemResponse, error)
	// 计算所有IP数量
	CountAllEnabledIPItems(context.Context, *CountAllEnabledIPItemsRequest) (*RPCCountResponse, error)
	// 列出所有名单中的IP
	ListAllEnabledIPItems(context.Context, *ListAllEnabledIPItemsRequest) (*ListAllEnabledIPItemsResponse, error)
	// 设置所有为已读
	UpdateIPItemsRead(context.Context, *UpdateIPItemsReadRequest) (*RPCSuccess, error)
}

// UnimplementedIPItemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIPItemServiceServer struct {
}

func (UnimplementedIPItemServiceServer) CreateIPItem(context.Context, *CreateIPItemRequest) (*CreateIPItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPItem not implemented")
}
func (UnimplementedIPItemServiceServer) CreateIPItems(context.Context, *CreateIPItemsRequest) (*CreateIPItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPItems not implemented")
}
func (UnimplementedIPItemServiceServer) UpdateIPItem(context.Context, *UpdateIPItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPItem not implemented")
}
func (UnimplementedIPItemServiceServer) DeleteIPItem(context.Context, *DeleteIPItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPItem not implemented")
}
func (UnimplementedIPItemServiceServer) DeleteIPItems(context.Context, *DeleteIPItemsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPItems not implemented")
}
func (UnimplementedIPItemServiceServer) CountIPItemsWithListId(context.Context, *CountIPItemsWithListIdRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountIPItemsWithListId not implemented")
}
func (UnimplementedIPItemServiceServer) ListIPItemsWithListId(context.Context, *ListIPItemsWithListIdRequest) (*ListIPItemsWithListIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIPItemsWithListId not implemented")
}
func (UnimplementedIPItemServiceServer) FindEnabledIPItem(context.Context, *FindEnabledIPItemRequest) (*FindEnabledIPItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledIPItem not implemented")
}
func (UnimplementedIPItemServiceServer) ListIPItemsAfterVersion(context.Context, *ListIPItemsAfterVersionRequest) (*ListIPItemsAfterVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIPItemsAfterVersion not implemented")
}
func (UnimplementedIPItemServiceServer) CheckIPItemStatus(context.Context, *CheckIPItemStatusRequest) (*CheckIPItemStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIPItemStatus not implemented")
}
func (UnimplementedIPItemServiceServer) ExistsEnabledIPItem(context.Context, *ExistsEnabledIPItemRequest) (*ExistsEnabledIPItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsEnabledIPItem not implemented")
}
func (UnimplementedIPItemServiceServer) CountAllEnabledIPItems(context.Context, *CountAllEnabledIPItemsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledIPItems not implemented")
}
func (UnimplementedIPItemServiceServer) ListAllEnabledIPItems(context.Context, *ListAllEnabledIPItemsRequest) (*ListAllEnabledIPItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllEnabledIPItems not implemented")
}
func (UnimplementedIPItemServiceServer) UpdateIPItemsRead(context.Context, *UpdateIPItemsReadRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPItemsRead not implemented")
}

// UnsafeIPItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPItemServiceServer will
// result in compilation errors.
type UnsafeIPItemServiceServer interface {
	mustEmbedUnimplementedIPItemServiceServer()
}

func RegisterIPItemServiceServer(s grpc.ServiceRegistrar, srv IPItemServiceServer) {
	s.RegisterService(&IPItemService_ServiceDesc, srv)
}

func _IPItemService_CreateIPItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).CreateIPItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_CreateIPItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).CreateIPItem(ctx, req.(*CreateIPItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_CreateIPItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).CreateIPItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_CreateIPItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).CreateIPItems(ctx, req.(*CreateIPItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_UpdateIPItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).UpdateIPItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_UpdateIPItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).UpdateIPItem(ctx, req.(*UpdateIPItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_DeleteIPItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIPItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).DeleteIPItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_DeleteIPItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).DeleteIPItem(ctx, req.(*DeleteIPItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_DeleteIPItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIPItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).DeleteIPItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_DeleteIPItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).DeleteIPItems(ctx, req.(*DeleteIPItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_CountIPItemsWithListId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountIPItemsWithListIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).CountIPItemsWithListId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_CountIPItemsWithListId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).CountIPItemsWithListId(ctx, req.(*CountIPItemsWithListIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_ListIPItemsWithListId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIPItemsWithListIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).ListIPItemsWithListId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_ListIPItemsWithListId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).ListIPItemsWithListId(ctx, req.(*ListIPItemsWithListIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_FindEnabledIPItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledIPItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).FindEnabledIPItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_FindEnabledIPItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).FindEnabledIPItem(ctx, req.(*FindEnabledIPItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_ListIPItemsAfterVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIPItemsAfterVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).ListIPItemsAfterVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_ListIPItemsAfterVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).ListIPItemsAfterVersion(ctx, req.(*ListIPItemsAfterVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_CheckIPItemStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIPItemStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).CheckIPItemStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_CheckIPItemStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).CheckIPItemStatus(ctx, req.(*CheckIPItemStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_ExistsEnabledIPItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsEnabledIPItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).ExistsEnabledIPItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_ExistsEnabledIPItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).ExistsEnabledIPItem(ctx, req.(*ExistsEnabledIPItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_CountAllEnabledIPItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledIPItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).CountAllEnabledIPItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_CountAllEnabledIPItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).CountAllEnabledIPItems(ctx, req.(*CountAllEnabledIPItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_ListAllEnabledIPItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllEnabledIPItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).ListAllEnabledIPItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_ListAllEnabledIPItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).ListAllEnabledIPItems(ctx, req.(*ListAllEnabledIPItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPItemService_UpdateIPItemsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPItemsReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPItemServiceServer).UpdateIPItemsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPItemService_UpdateIPItemsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPItemServiceServer).UpdateIPItemsRead(ctx, req.(*UpdateIPItemsReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPItemService_ServiceDesc is the grpc.ServiceDesc for IPItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPItemService",
	HandlerType: (*IPItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createIPItem",
			Handler:    _IPItemService_CreateIPItem_Handler,
		},
		{
			MethodName: "createIPItems",
			Handler:    _IPItemService_CreateIPItems_Handler,
		},
		{
			MethodName: "updateIPItem",
			Handler:    _IPItemService_UpdateIPItem_Handler,
		},
		{
			MethodName: "deleteIPItem",
			Handler:    _IPItemService_DeleteIPItem_Handler,
		},
		{
			MethodName: "deleteIPItems",
			Handler:    _IPItemService_DeleteIPItems_Handler,
		},
		{
			MethodName: "countIPItemsWithListId",
			Handler:    _IPItemService_CountIPItemsWithListId_Handler,
		},
		{
			MethodName: "listIPItemsWithListId",
			Handler:    _IPItemService_ListIPItemsWithListId_Handler,
		},
		{
			MethodName: "findEnabledIPItem",
			Handler:    _IPItemService_FindEnabledIPItem_Handler,
		},
		{
			MethodName: "listIPItemsAfterVersion",
			Handler:    _IPItemService_ListIPItemsAfterVersion_Handler,
		},
		{
			MethodName: "checkIPItemStatus",
			Handler:    _IPItemService_CheckIPItemStatus_Handler,
		},
		{
			MethodName: "existsEnabledIPItem",
			Handler:    _IPItemService_ExistsEnabledIPItem_Handler,
		},
		{
			MethodName: "countAllEnabledIPItems",
			Handler:    _IPItemService_CountAllEnabledIPItems_Handler,
		},
		{
			MethodName: "listAllEnabledIPItems",
			Handler:    _IPItemService_ListAllEnabledIPItems_Handler,
		},
		{
			MethodName: "updateIPItemsRead",
			Handler:    _IPItemService_UpdateIPItemsRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ip_item.proto",
}
