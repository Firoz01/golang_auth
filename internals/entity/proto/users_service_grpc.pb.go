// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: users_service.proto

package proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	VerifyPasswordAndGetClaimData(ctx context.Context, in *VerifyPasswordRequest, opts ...grpc.CallOption) (*VerifyPasswordResponse, error)
	GetClaimDataByUserID(ctx context.Context, in *ClaimDataByUserIDRequest, opts ...grpc.CallOption) (*UserClaimData, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) VerifyPasswordAndGetClaimData(ctx context.Context, in *VerifyPasswordRequest, opts ...grpc.CallOption) (*VerifyPasswordResponse, error) {
	out := new(VerifyPasswordResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/VerifyPasswordAndGetClaimData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetClaimDataByUserID(ctx context.Context, in *ClaimDataByUserIDRequest, opts ...grpc.CallOption) (*UserClaimData, error) {
	out := new(UserClaimData)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetClaimDataByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	VerifyPasswordAndGetClaimData(context.Context, *VerifyPasswordRequest) (*VerifyPasswordResponse, error)
	GetClaimDataByUserID(context.Context, *ClaimDataByUserIDRequest) (*UserClaimData, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) VerifyPasswordAndGetClaimData(context.Context, *VerifyPasswordRequest) (*VerifyPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPasswordAndGetClaimData not implemented")
}
func (UnimplementedUserServiceServer) GetClaimDataByUserID(context.Context, *ClaimDataByUserIDRequest) (*UserClaimData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimDataByUserID not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_VerifyPasswordAndGetClaimData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyPasswordAndGetClaimData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/VerifyPasswordAndGetClaimData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyPasswordAndGetClaimData(ctx, req.(*VerifyPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetClaimDataByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimDataByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetClaimDataByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetClaimDataByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetClaimDataByUserID(ctx, req.(*ClaimDataByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyPasswordAndGetClaimData",
			Handler:    _UserService_VerifyPasswordAndGetClaimData_Handler,
		},
		{
			MethodName: "GetClaimDataByUserID",
			Handler:    _UserService_GetClaimDataByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users_service.proto",
}
