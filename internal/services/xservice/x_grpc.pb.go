// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: service/x.proto

package xservice

import (
	context "context"
	productservice "github.com/jigadhirasu/demo/internal/services/productservice"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// XServiceClient is the client API for XService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XServiceClient interface {
	Create(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.Product, error)
	List(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.ProductList, error)
	Get(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.Product, error)
}

type xServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXServiceClient(cc grpc.ClientConnInterface) XServiceClient {
	return &xServiceClient{cc}
}

func (c *xServiceClient) Create(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.Product, error) {
	out := new(productservice.Product)
	err := c.cc.Invoke(ctx, "/xservice.XService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xServiceClient) List(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.ProductList, error) {
	out := new(productservice.ProductList)
	err := c.cc.Invoke(ctx, "/xservice.XService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xServiceClient) Get(ctx context.Context, in *productservice.Product, opts ...grpc.CallOption) (*productservice.Product, error) {
	out := new(productservice.Product)
	err := c.cc.Invoke(ctx, "/xservice.XService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XServiceServer is the server API for XService service.
// All implementations must embed UnimplementedXServiceServer
// for forward compatibility
type XServiceServer interface {
	Create(context.Context, *productservice.Product) (*productservice.Product, error)
	List(context.Context, *productservice.Product) (*productservice.ProductList, error)
	Get(context.Context, *productservice.Product) (*productservice.Product, error)
	mustEmbedUnimplementedXServiceServer()
}

// UnimplementedXServiceServer must be embedded to have forward compatible implementations.
type UnimplementedXServiceServer struct {
}

func (UnimplementedXServiceServer) Create(context.Context, *productservice.Product) (*productservice.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedXServiceServer) List(context.Context, *productservice.Product) (*productservice.ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedXServiceServer) Get(context.Context, *productservice.Product) (*productservice.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedXServiceServer) mustEmbedUnimplementedXServiceServer() {}

// UnsafeXServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XServiceServer will
// result in compilation errors.
type UnsafeXServiceServer interface {
	mustEmbedUnimplementedXServiceServer()
}

func RegisterXServiceServer(s grpc.ServiceRegistrar, srv XServiceServer) {
	s.RegisterService(&XService_ServiceDesc, srv)
}

func _XService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productservice.Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xservice.XService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XServiceServer).Create(ctx, req.(*productservice.Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _XService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productservice.Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xservice.XService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XServiceServer).List(ctx, req.(*productservice.Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _XService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productservice.Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xservice.XService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XServiceServer).Get(ctx, req.(*productservice.Product))
	}
	return interceptor(ctx, in, info, handler)
}

// XService_ServiceDesc is the grpc.ServiceDesc for XService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xservice.XService",
	HandlerType: (*XServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _XService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _XService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _XService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/x.proto",
}
