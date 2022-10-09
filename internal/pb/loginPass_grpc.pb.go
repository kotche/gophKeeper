// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/proto/loginPass.proto

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

// LoginPassServiceClient is the client API for LoginPassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginPassServiceClient interface {
	CreateLoginPass(ctx context.Context, in *LoginPassRequest, opts ...grpc.CallOption) (*LoginPassResponse, error)
	GetAllLoginPass(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type loginPassServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginPassServiceClient(cc grpc.ClientConnInterface) LoginPassServiceClient {
	return &loginPassServiceClient{cc}
}

func (c *loginPassServiceClient) CreateLoginPass(ctx context.Context, in *LoginPassRequest, opts ...grpc.CallOption) (*LoginPassResponse, error) {
	out := new(LoginPassResponse)
	err := c.cc.Invoke(ctx, "/keeper.LoginPassService/CreateLoginPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginPassServiceClient) GetAllLoginPass(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/keeper.LoginPassService/GetAllLoginPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginPassServiceServer is the server API for LoginPassService service.
// All implementations must embed UnimplementedLoginPassServiceServer
// for forward compatibility
type LoginPassServiceServer interface {
	CreateLoginPass(context.Context, *LoginPassRequest) (*LoginPassResponse, error)
	GetAllLoginPass(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedLoginPassServiceServer()
}

// UnimplementedLoginPassServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginPassServiceServer struct {
}

func (UnimplementedLoginPassServiceServer) CreateLoginPass(context.Context, *LoginPassRequest) (*LoginPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoginPass not implemented")
}
func (UnimplementedLoginPassServiceServer) GetAllLoginPass(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLoginPass not implemented")
}
func (UnimplementedLoginPassServiceServer) mustEmbedUnimplementedLoginPassServiceServer() {}

// UnsafeLoginPassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginPassServiceServer will
// result in compilation errors.
type UnsafeLoginPassServiceServer interface {
	mustEmbedUnimplementedLoginPassServiceServer()
}

func RegisterLoginPassServiceServer(s grpc.ServiceRegistrar, srv LoginPassServiceServer) {
	s.RegisterService(&LoginPassService_ServiceDesc, srv)
}

func _LoginPassService_CreateLoginPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginPassServiceServer).CreateLoginPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.LoginPassService/CreateLoginPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginPassServiceServer).CreateLoginPass(ctx, req.(*LoginPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginPassService_GetAllLoginPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginPassServiceServer).GetAllLoginPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.LoginPassService/GetAllLoginPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginPassServiceServer).GetAllLoginPass(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginPassService_ServiceDesc is the grpc.ServiceDesc for LoginPassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginPassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.LoginPassService",
	HandlerType: (*LoginPassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLoginPass",
			Handler:    _LoginPassService_CreateLoginPass_Handler,
		},
		{
			MethodName: "GetAllLoginPass",
			Handler:    _LoginPassService_GetAllLoginPass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/loginPass.proto",
}
