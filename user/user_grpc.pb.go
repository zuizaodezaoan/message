// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: users.proto

package users

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
	User_UserCreate_FullMethodName        = "/users.User/UserCreate"
	User_UserUpdate_FullMethodName        = "/users.User/UserUpdate"
	User_UserDelete_FullMethodName        = "/users.User/UserDelete"
	User_UserGet_FullMethodName           = "/users.User/UserGet"
	User_UsersGet_FullMethodName          = "/users.User/UsersGet"
	User_UserGetByUsername_FullMethodName = "/users.User/UserGetByUsername"
	User_UsersMobile_FullMethodName       = "/users.User/UsersMobile"
	User_UserLogin_FullMethodName         = "/users.User/UserLogin"
	User_UserLoginMobile_FullMethodName   = "/users.User/UserLoginMobile"
	User_LoginMobile_FullMethodName       = "/users.User/LoginMobile"
	User_LoginMobiles_FullMethodName      = "/users.User/LoginMobiles"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserUpdate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	UserGet(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error)
	UsersGet(ctx context.Context, in *UsersGetRequest, opts ...grpc.CallOption) (*UsersGetResponse, error)
	UserGetByUsername(ctx context.Context, in *UserGetByUsernameRequest, opts ...grpc.CallOption) (*UserGetByUsernameResponse, error)
	UsersMobile(ctx context.Context, in *UserMobileRequest, opts ...grpc.CallOption) (*UserMobileResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserLoginMobile(ctx context.Context, in *UserLoginMobileRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	LoginMobile(ctx context.Context, in *LoginMobileRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	LoginMobiles(ctx context.Context, in *UserLoginsRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, User_UserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUpdate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, User_UserUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, User_UserDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserGet(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error) {
	out := new(UserGetResponse)
	err := c.cc.Invoke(ctx, User_UserGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UsersGet(ctx context.Context, in *UsersGetRequest, opts ...grpc.CallOption) (*UsersGetResponse, error) {
	out := new(UsersGetResponse)
	err := c.cc.Invoke(ctx, User_UsersGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserGetByUsername(ctx context.Context, in *UserGetByUsernameRequest, opts ...grpc.CallOption) (*UserGetByUsernameResponse, error) {
	out := new(UserGetByUsernameResponse)
	err := c.cc.Invoke(ctx, User_UserGetByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UsersMobile(ctx context.Context, in *UserMobileRequest, opts ...grpc.CallOption) (*UserMobileResponse, error) {
	out := new(UserMobileResponse)
	err := c.cc.Invoke(ctx, User_UsersMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLoginMobile(ctx context.Context, in *UserLoginMobileRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserLoginMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginMobile(ctx context.Context, in *LoginMobileRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_LoginMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginMobiles(ctx context.Context, in *UserLoginsRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_LoginMobiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserUpdate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	UserGet(context.Context, *UserGetRequest) (*UserGetResponse, error)
	UsersGet(context.Context, *UsersGetRequest) (*UsersGetResponse, error)
	UserGetByUsername(context.Context, *UserGetByUsernameRequest) (*UserGetByUsernameResponse, error)
	UsersMobile(context.Context, *UserMobileRequest) (*UserMobileResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserLoginMobile(context.Context, *UserLoginMobileRequest) (*UserLoginResponse, error)
	LoginMobile(context.Context, *LoginMobileRequest) (*UserLoginResponse, error)
	LoginMobiles(context.Context, *UserLoginsRequest) (*UserLoginResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedUserServer) UserUpdate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedUserServer) UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedUserServer) UserGet(context.Context, *UserGetRequest) (*UserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (UnimplementedUserServer) UsersGet(context.Context, *UsersGetRequest) (*UsersGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersGet not implemented")
}
func (UnimplementedUserServer) UserGetByUsername(context.Context, *UserGetByUsernameRequest) (*UserGetByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetByUsername not implemented")
}
func (UnimplementedUserServer) UsersMobile(context.Context, *UserMobileRequest) (*UserMobileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersMobile not implemented")
}
func (UnimplementedUserServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) UserLoginMobile(context.Context, *UserLoginMobileRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginMobile not implemented")
}
func (UnimplementedUserServer) LoginMobile(context.Context, *LoginMobileRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginMobile not implemented")
}
func (UnimplementedUserServer) LoginMobiles(context.Context, *UserLoginsRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginMobiles not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUpdate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserGet(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UsersGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UsersGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UsersGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UsersGet(ctx, req.(*UsersGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserGetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserGetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserGetByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserGetByUsername(ctx, req.(*UserGetByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UsersMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UsersMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UsersMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UsersMobile(ctx, req.(*UserMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLoginMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLoginMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLoginMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLoginMobile(ctx, req.(*UserLoginMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LoginMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginMobile(ctx, req.(*LoginMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginMobiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginMobiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LoginMobiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginMobiles(ctx, req.(*UserLoginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCreate",
			Handler:    _User_UserCreate_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _User_UserUpdate_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _User_UserDelete_Handler,
		},
		{
			MethodName: "UserGet",
			Handler:    _User_UserGet_Handler,
		},
		{
			MethodName: "UsersGet",
			Handler:    _User_UsersGet_Handler,
		},
		{
			MethodName: "UserGetByUsername",
			Handler:    _User_UserGetByUsername_Handler,
		},
		{
			MethodName: "UsersMobile",
			Handler:    _User_UsersMobile_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserLoginMobile",
			Handler:    _User_UserLoginMobile_Handler,
		},
		{
			MethodName: "LoginMobile",
			Handler:    _User_LoginMobile_Handler,
		},
		{
			MethodName: "LoginMobiles",
			Handler:    _User_LoginMobiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
