// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: media/v1/media.proto

package mediav1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MediaService_Image_FullMethodName = "/media.v1.MediaService/Image"
	MediaService_Csv_FullMethodName   = "/media.v1.MediaService/Csv"
)

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	Image(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	Csv(ctx context.Context, in *CsvRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) Image(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, MediaService_Image_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) Csv(ctx context.Context, in *CsvRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, MediaService_Csv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	Image(context.Context, *ImageRequest) (*httpbody.HttpBody, error)
	Csv(context.Context, *CsvRequest) (*httpbody.HttpBody, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) Image(context.Context, *ImageRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Image not implemented")
}
func (UnimplementedMediaServiceServer) Csv(context.Context, *CsvRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Csv not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_Image_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).Image(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_Image_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).Image(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_Csv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).Csv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_Csv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).Csv(ctx, req.(*CsvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "media.v1.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Image",
			Handler:    _MediaService_Image_Handler,
		},
		{
			MethodName: "Csv",
			Handler:    _MediaService_Csv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "media/v1/media.proto",
}
