// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_node_price_item.proto

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
	NodePriceItemService_CreateNodePriceItem_FullMethodName            = "/pb.NodePriceItemService/createNodePriceItem"
	NodePriceItemService_UpdateNodePriceItem_FullMethodName            = "/pb.NodePriceItemService/updateNodePriceItem"
	NodePriceItemService_DeleteNodePriceItem_FullMethodName            = "/pb.NodePriceItemService/deleteNodePriceItem"
	NodePriceItemService_FindAllEnabledNodePriceItems_FullMethodName   = "/pb.NodePriceItemService/findAllEnabledNodePriceItems"
	NodePriceItemService_FindAllAvailableNodePriceItems_FullMethodName = "/pb.NodePriceItemService/findAllAvailableNodePriceItems"
	NodePriceItemService_FindEnabledNodePriceItem_FullMethodName       = "/pb.NodePriceItemService/findEnabledNodePriceItem"
)

// NodePriceItemServiceClient is the client API for NodePriceItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodePriceItemServiceClient interface {
	// 创建区域价格
	CreateNodePriceItem(ctx context.Context, in *CreateNodePriceItemRequest, opts ...grpc.CallOption) (*CreateNodePriceItemResponse, error)
	// 修改区域价格
	UpdateNodePriceItem(ctx context.Context, in *UpdateNodePriceItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除区域价格
	DeleteNodePriceItem(ctx context.Context, in *DeleteNodePriceItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找所有区域价格
	FindAllEnabledNodePriceItems(ctx context.Context, in *FindAllEnabledNodePriceItemsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodePriceItemsResponse, error)
	// 查找所有启用的区域价格
	FindAllAvailableNodePriceItems(ctx context.Context, in *FindAllAvailableNodePriceItemsRequest, opts ...grpc.CallOption) (*FindAllAvailableNodePriceItemsResponse, error)
	// 查找单个区域信息
	FindEnabledNodePriceItem(ctx context.Context, in *FindEnabledNodePriceItemRequest, opts ...grpc.CallOption) (*FindEnabledNodePriceItemResponse, error)
}

type nodePriceItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodePriceItemServiceClient(cc grpc.ClientConnInterface) NodePriceItemServiceClient {
	return &nodePriceItemServiceClient{cc}
}

func (c *nodePriceItemServiceClient) CreateNodePriceItem(ctx context.Context, in *CreateNodePriceItemRequest, opts ...grpc.CallOption) (*CreateNodePriceItemResponse, error) {
	out := new(CreateNodePriceItemResponse)
	err := c.cc.Invoke(ctx, NodePriceItemService_CreateNodePriceItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePriceItemServiceClient) UpdateNodePriceItem(ctx context.Context, in *UpdateNodePriceItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodePriceItemService_UpdateNodePriceItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePriceItemServiceClient) DeleteNodePriceItem(ctx context.Context, in *DeleteNodePriceItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodePriceItemService_DeleteNodePriceItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePriceItemServiceClient) FindAllEnabledNodePriceItems(ctx context.Context, in *FindAllEnabledNodePriceItemsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodePriceItemsResponse, error) {
	out := new(FindAllEnabledNodePriceItemsResponse)
	err := c.cc.Invoke(ctx, NodePriceItemService_FindAllEnabledNodePriceItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePriceItemServiceClient) FindAllAvailableNodePriceItems(ctx context.Context, in *FindAllAvailableNodePriceItemsRequest, opts ...grpc.CallOption) (*FindAllAvailableNodePriceItemsResponse, error) {
	out := new(FindAllAvailableNodePriceItemsResponse)
	err := c.cc.Invoke(ctx, NodePriceItemService_FindAllAvailableNodePriceItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePriceItemServiceClient) FindEnabledNodePriceItem(ctx context.Context, in *FindEnabledNodePriceItemRequest, opts ...grpc.CallOption) (*FindEnabledNodePriceItemResponse, error) {
	out := new(FindEnabledNodePriceItemResponse)
	err := c.cc.Invoke(ctx, NodePriceItemService_FindEnabledNodePriceItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodePriceItemServiceServer is the server API for NodePriceItemService service.
// All implementations should embed UnimplementedNodePriceItemServiceServer
// for forward compatibility
type NodePriceItemServiceServer interface {
	// 创建区域价格
	CreateNodePriceItem(context.Context, *CreateNodePriceItemRequest) (*CreateNodePriceItemResponse, error)
	// 修改区域价格
	UpdateNodePriceItem(context.Context, *UpdateNodePriceItemRequest) (*RPCSuccess, error)
	// 删除区域价格
	DeleteNodePriceItem(context.Context, *DeleteNodePriceItemRequest) (*RPCSuccess, error)
	// 查找所有区域价格
	FindAllEnabledNodePriceItems(context.Context, *FindAllEnabledNodePriceItemsRequest) (*FindAllEnabledNodePriceItemsResponse, error)
	// 查找所有启用的区域价格
	FindAllAvailableNodePriceItems(context.Context, *FindAllAvailableNodePriceItemsRequest) (*FindAllAvailableNodePriceItemsResponse, error)
	// 查找单个区域信息
	FindEnabledNodePriceItem(context.Context, *FindEnabledNodePriceItemRequest) (*FindEnabledNodePriceItemResponse, error)
}

// UnimplementedNodePriceItemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNodePriceItemServiceServer struct {
}

func (UnimplementedNodePriceItemServiceServer) CreateNodePriceItem(context.Context, *CreateNodePriceItemRequest) (*CreateNodePriceItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodePriceItem not implemented")
}
func (UnimplementedNodePriceItemServiceServer) UpdateNodePriceItem(context.Context, *UpdateNodePriceItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodePriceItem not implemented")
}
func (UnimplementedNodePriceItemServiceServer) DeleteNodePriceItem(context.Context, *DeleteNodePriceItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodePriceItem not implemented")
}
func (UnimplementedNodePriceItemServiceServer) FindAllEnabledNodePriceItems(context.Context, *FindAllEnabledNodePriceItemsRequest) (*FindAllEnabledNodePriceItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledNodePriceItems not implemented")
}
func (UnimplementedNodePriceItemServiceServer) FindAllAvailableNodePriceItems(context.Context, *FindAllAvailableNodePriceItemsRequest) (*FindAllAvailableNodePriceItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAvailableNodePriceItems not implemented")
}
func (UnimplementedNodePriceItemServiceServer) FindEnabledNodePriceItem(context.Context, *FindEnabledNodePriceItemRequest) (*FindEnabledNodePriceItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledNodePriceItem not implemented")
}

// UnsafeNodePriceItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodePriceItemServiceServer will
// result in compilation errors.
type UnsafeNodePriceItemServiceServer interface {
	mustEmbedUnimplementedNodePriceItemServiceServer()
}

func RegisterNodePriceItemServiceServer(s grpc.ServiceRegistrar, srv NodePriceItemServiceServer) {
	s.RegisterService(&NodePriceItemService_ServiceDesc, srv)
}

func _NodePriceItemService_CreateNodePriceItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodePriceItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).CreateNodePriceItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_CreateNodePriceItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).CreateNodePriceItem(ctx, req.(*CreateNodePriceItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePriceItemService_UpdateNodePriceItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodePriceItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).UpdateNodePriceItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_UpdateNodePriceItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).UpdateNodePriceItem(ctx, req.(*UpdateNodePriceItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePriceItemService_DeleteNodePriceItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodePriceItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).DeleteNodePriceItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_DeleteNodePriceItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).DeleteNodePriceItem(ctx, req.(*DeleteNodePriceItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePriceItemService_FindAllEnabledNodePriceItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledNodePriceItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).FindAllEnabledNodePriceItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_FindAllEnabledNodePriceItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).FindAllEnabledNodePriceItems(ctx, req.(*FindAllEnabledNodePriceItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePriceItemService_FindAllAvailableNodePriceItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAvailableNodePriceItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).FindAllAvailableNodePriceItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_FindAllAvailableNodePriceItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).FindAllAvailableNodePriceItems(ctx, req.(*FindAllAvailableNodePriceItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePriceItemService_FindEnabledNodePriceItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledNodePriceItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePriceItemServiceServer).FindEnabledNodePriceItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodePriceItemService_FindEnabledNodePriceItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePriceItemServiceServer).FindEnabledNodePriceItem(ctx, req.(*FindEnabledNodePriceItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodePriceItemService_ServiceDesc is the grpc.ServiceDesc for NodePriceItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodePriceItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodePriceItemService",
	HandlerType: (*NodePriceItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodePriceItem",
			Handler:    _NodePriceItemService_CreateNodePriceItem_Handler,
		},
		{
			MethodName: "updateNodePriceItem",
			Handler:    _NodePriceItemService_UpdateNodePriceItem_Handler,
		},
		{
			MethodName: "deleteNodePriceItem",
			Handler:    _NodePriceItemService_DeleteNodePriceItem_Handler,
		},
		{
			MethodName: "findAllEnabledNodePriceItems",
			Handler:    _NodePriceItemService_FindAllEnabledNodePriceItems_Handler,
		},
		{
			MethodName: "findAllAvailableNodePriceItems",
			Handler:    _NodePriceItemService_FindAllAvailableNodePriceItems_Handler,
		},
		{
			MethodName: "findEnabledNodePriceItem",
			Handler:    _NodePriceItemService_FindEnabledNodePriceItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_price_item.proto",
}
