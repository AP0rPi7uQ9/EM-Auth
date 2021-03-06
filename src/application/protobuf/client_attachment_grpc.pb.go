// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AttachmentClient is the client API for Attachment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachmentClient interface {
	Create(ctx context.Context, in *AttachmentCreate, opts ...grpc.CallOption) (*protobuf.Response, error)
	GetOne(ctx context.Context, in *AttachmentGetOne, opts ...grpc.CallOption) (*protobuf.Response, error)
	DiskCleanUp(ctx context.Context, in *AttachmentDiskCleanUp, opts ...grpc.CallOption) (*protobuf.Response, error)
}

type attachmentClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachmentClient(cc grpc.ClientConnInterface) AttachmentClient {
	return &attachmentClient{cc}
}

func (c *attachmentClient) Create(ctx context.Context, in *AttachmentCreate, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Attachment/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentClient) GetOne(ctx context.Context, in *AttachmentGetOne, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Attachment/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentClient) DiskCleanUp(ctx context.Context, in *AttachmentDiskCleanUp, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Attachment/DiskCleanUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentServer is the server API for Attachment service.
// All implementations must embed UnimplementedAttachmentServer
// for forward compatibility
type AttachmentServer interface {
	Create(context.Context, *AttachmentCreate) (*protobuf.Response, error)
	GetOne(context.Context, *AttachmentGetOne) (*protobuf.Response, error)
	DiskCleanUp(context.Context, *AttachmentDiskCleanUp) (*protobuf.Response, error)
	mustEmbedUnimplementedAttachmentServer()
}

// UnimplementedAttachmentServer must be embedded to have forward compatible implementations.
type UnimplementedAttachmentServer struct {
}

func (UnimplementedAttachmentServer) Create(context.Context, *AttachmentCreate) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttachmentServer) GetOne(context.Context, *AttachmentGetOne) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedAttachmentServer) DiskCleanUp(context.Context, *AttachmentDiskCleanUp) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskCleanUp not implemented")
}
func (UnimplementedAttachmentServer) mustEmbedUnimplementedAttachmentServer() {}

// UnsafeAttachmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachmentServer will
// result in compilation errors.
type UnsafeAttachmentServer interface {
	mustEmbedUnimplementedAttachmentServer()
}

func RegisterAttachmentServer(s grpc.ServiceRegistrar, srv AttachmentServer) {
	s.RegisterService(&_Attachment_serviceDesc, srv)
}

func _Attachment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Attachment/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServer).Create(ctx, req.(*AttachmentCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachment_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentGetOne)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Attachment/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServer).GetOne(ctx, req.(*AttachmentGetOne))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachment_DiskCleanUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentDiskCleanUp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServer).DiskCleanUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Attachment/DiskCleanUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServer).DiskCleanUp(ctx, req.(*AttachmentDiskCleanUp))
	}
	return interceptor(ctx, in, info, handler)
}

var _Attachment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Attachment",
	HandlerType: (*AttachmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Attachment_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Attachment_GetOne_Handler,
		},
		{
			MethodName: "DiskCleanUp",
			Handler:    _Attachment_DiskCleanUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_attachment.proto",
}
