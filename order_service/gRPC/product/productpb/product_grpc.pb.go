// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: product.proto

package productpb

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// unary - get single product details
	GetProductDetails(ctx context.Context, in *GetProductDetailsRequest, opts ...grpc.CallOption) (*GetProductDetailsResponse, error)
	// unary - get list of product details
	ListProductDetails(ctx context.Context, in *ListProductDetailsRequest, opts ...grpc.CallOption) (*ListProductDetailsResponse, error)
	// unary - update the product quantity
	UpdateProductQuantity(ctx context.Context, in *UpdateProductQuantityRequest, opts ...grpc.CallOption) (*UpdateProductQuantityResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductDetails(ctx context.Context, in *GetProductDetailsRequest, opts ...grpc.CallOption) (*GetProductDetailsResponse, error) {
	out := new(GetProductDetailsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProductDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProductDetails(ctx context.Context, in *ListProductDetailsRequest, opts ...grpc.CallOption) (*ListProductDetailsResponse, error) {
	out := new(ListProductDetailsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/ListProductDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductQuantity(ctx context.Context, in *UpdateProductQuantityRequest, opts ...grpc.CallOption) (*UpdateProductQuantityResponse, error) {
	out := new(UpdateProductQuantityResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/UpdateProductQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// unary - get single product details
	GetProductDetails(context.Context, *GetProductDetailsRequest) (*GetProductDetailsResponse, error)
	// unary - get list of product details
	ListProductDetails(context.Context, *ListProductDetailsRequest) (*ListProductDetailsResponse, error)
	// unary - update the product quantity
	UpdateProductQuantity(context.Context, *UpdateProductQuantityRequest) (*UpdateProductQuantityResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProductDetails(context.Context, *GetProductDetailsRequest) (*GetProductDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetails not implemented")
}
func (UnimplementedProductServiceServer) ListProductDetails(context.Context, *ListProductDetailsRequest) (*ListProductDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductDetails not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductQuantity(context.Context, *UpdateProductQuantityRequest) (*UpdateProductQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductQuantity not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProductDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProductDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductDetails(ctx, req.(*GetProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProductDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProductDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/ListProductDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProductDetails(ctx, req.(*ListProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProductQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/UpdateProductQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProductQuantity(ctx, req.(*UpdateProductQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductDetails",
			Handler:    _ProductService_GetProductDetails_Handler,
		},
		{
			MethodName: "ListProductDetails",
			Handler:    _ProductService_ListProductDetails_Handler,
		},
		{
			MethodName: "UpdateProductQuantity",
			Handler:    _ProductService_UpdateProductQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}