// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: video.proto

package videov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VideoService_Detect_FullMethodName = "/video.VideoService/Detect"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	Detect(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoResponse, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) Detect(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoResponse)
	err := c.cc.Invoke(ctx, VideoService_Detect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility.
type VideoServiceServer interface {
	Detect(context.Context, *VideoRequest) (*VideoResponse, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVideoServiceServer struct{}

func (UnimplementedVideoServiceServer) Detect(context.Context, *VideoRequest) (*VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}
func (UnimplementedVideoServiceServer) testEmbeddedByValue()                      {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	// If the following call pancis, it indicates UnimplementedVideoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_Detect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Detect(ctx, req.(*VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _VideoService_Detect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
