// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/reset_password.proto

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

// ResetPasswordServiceClient is the client API for ResetPasswordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResetPasswordServiceClient interface {
	RequestPassReset(ctx context.Context, in *RequestResetPassReq, opts ...grpc.CallOption) (*RequestResetPassRes, error)
	ApplyPassReset(ctx context.Context, in *ApplyResetPassReq, opts ...grpc.CallOption) (*ApplyResetPassRes, error)
}

type resetPasswordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResetPasswordServiceClient(cc grpc.ClientConnInterface) ResetPasswordServiceClient {
	return &resetPasswordServiceClient{cc}
}

func (c *resetPasswordServiceClient) RequestPassReset(ctx context.Context, in *RequestResetPassReq, opts ...grpc.CallOption) (*RequestResetPassRes, error) {
	out := new(RequestResetPassRes)
	err := c.cc.Invoke(ctx, "/reset_password.ResetPasswordService/RequestPassReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resetPasswordServiceClient) ApplyPassReset(ctx context.Context, in *ApplyResetPassReq, opts ...grpc.CallOption) (*ApplyResetPassRes, error) {
	out := new(ApplyResetPassRes)
	err := c.cc.Invoke(ctx, "/reset_password.ResetPasswordService/ApplyPassReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResetPasswordServiceServer is the server API for ResetPasswordService service.
// All implementations must embed UnimplementedResetPasswordServiceServer
// for forward compatibility
type ResetPasswordServiceServer interface {
	RequestPassReset(context.Context, *RequestResetPassReq) (*RequestResetPassRes, error)
	ApplyPassReset(context.Context, *ApplyResetPassReq) (*ApplyResetPassRes, error)
	mustEmbedUnimplementedResetPasswordServiceServer()
}

// UnimplementedResetPasswordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResetPasswordServiceServer struct {
}

func (UnimplementedResetPasswordServiceServer) RequestPassReset(context.Context, *RequestResetPassReq) (*RequestResetPassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPassReset not implemented")
}
func (UnimplementedResetPasswordServiceServer) ApplyPassReset(context.Context, *ApplyResetPassReq) (*ApplyResetPassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyPassReset not implemented")
}
func (UnimplementedResetPasswordServiceServer) mustEmbedUnimplementedResetPasswordServiceServer() {}

// UnsafeResetPasswordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResetPasswordServiceServer will
// result in compilation errors.
type UnsafeResetPasswordServiceServer interface {
	mustEmbedUnimplementedResetPasswordServiceServer()
}

func RegisterResetPasswordServiceServer(s grpc.ServiceRegistrar, srv ResetPasswordServiceServer) {
	s.RegisterService(&ResetPasswordService_ServiceDesc, srv)
}

func _ResetPasswordService_RequestPassReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResetPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResetPasswordServiceServer).RequestPassReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reset_password.ResetPasswordService/RequestPassReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResetPasswordServiceServer).RequestPassReset(ctx, req.(*RequestResetPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResetPasswordService_ApplyPassReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyResetPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResetPasswordServiceServer).ApplyPassReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reset_password.ResetPasswordService/ApplyPassReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResetPasswordServiceServer).ApplyPassReset(ctx, req.(*ApplyResetPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ResetPasswordService_ServiceDesc is the grpc.ServiceDesc for ResetPasswordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResetPasswordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reset_password.ResetPasswordService",
	HandlerType: (*ResetPasswordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestPassReset",
			Handler:    _ResetPasswordService_RequestPassReset_Handler,
		},
		{
			MethodName: "ApplyPassReset",
			Handler:    _ResetPasswordService_ApplyPassReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reset_password.proto",
}
