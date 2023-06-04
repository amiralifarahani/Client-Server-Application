// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: auth_server.proto

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
	AuthenticationService_ReqPq_FullMethodName        = "/authentication_service/req_pq"
	AuthenticationService_Req_DHParams_FullMethodName = "/authentication_service/req_DH_params"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	ReqPq(ctx context.Context, in *ReqPqRequest, opts ...grpc.CallOption) (*ReqPqReply, error)
	Req_DHParams(ctx context.Context, in *Req_DHParamsRequest, opts ...grpc.CallOption) (*Req_DHParamsReply, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) ReqPq(ctx context.Context, in *ReqPqRequest, opts ...grpc.CallOption) (*ReqPqReply, error) {
	out := new(ReqPqReply)
	err := c.cc.Invoke(ctx, AuthenticationService_ReqPq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) Req_DHParams(ctx context.Context, in *Req_DHParamsRequest, opts ...grpc.CallOption) (*Req_DHParamsReply, error) {
	out := new(Req_DHParamsReply)
	err := c.cc.Invoke(ctx, AuthenticationService_Req_DHParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	ReqPq(context.Context, *ReqPqRequest) (*ReqPqReply, error)
	Req_DHParams(context.Context, *Req_DHParamsRequest) (*Req_DHParamsReply, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) ReqPq(context.Context, *ReqPqRequest) (*ReqPqReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqPq not implemented")
}
func (UnimplementedAuthenticationServiceServer) Req_DHParams(context.Context, *Req_DHParamsRequest) (*Req_DHParamsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Req_DHParams not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ReqPq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ReqPq(ctx, req.(*ReqPqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_Req_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req_DHParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Req_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_Req_DHParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Req_DHParams(ctx, req.(*Req_DHParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication_service",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _AuthenticationService_ReqPq_Handler,
		},
		{
			MethodName: "req_DH_params",
			Handler:    _AuthenticationService_Req_DHParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_server.proto",
}