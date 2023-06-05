// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: req_pq.proto

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
	ReqPqAuthenticationService_ReqPq_FullMethodName = "/req_pq_authentication_service/req_pq"
)

// ReqPqAuthenticationServiceClient is the client API for ReqPqAuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReqPqAuthenticationServiceClient interface {
	ReqPq(ctx context.Context, in *ReqPq_Request, opts ...grpc.CallOption) (*ReqPq_Response, error)
}

type reqPqAuthenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReqPqAuthenticationServiceClient(cc grpc.ClientConnInterface) ReqPqAuthenticationServiceClient {
	return &reqPqAuthenticationServiceClient{cc}
}

func (c *reqPqAuthenticationServiceClient) ReqPq(ctx context.Context, in *ReqPq_Request, opts ...grpc.CallOption) (*ReqPq_Response, error) {
	out := new(ReqPq_Response)
	err := c.cc.Invoke(ctx, ReqPqAuthenticationService_ReqPq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReqPqAuthenticationServiceServer is the server API for ReqPqAuthenticationService service.
// All implementations must embed UnimplementedReqPqAuthenticationServiceServer
// for forward compatibility
type ReqPqAuthenticationServiceServer interface {
	ReqPq(context.Context, *ReqPq_Request) (*ReqPq_Response, error)
	mustEmbedUnimplementedReqPqAuthenticationServiceServer()
}

// UnimplementedReqPqAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReqPqAuthenticationServiceServer struct {
}

func (UnimplementedReqPqAuthenticationServiceServer) ReqPq(context.Context, *ReqPq_Request) (*ReqPq_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqPq not implemented")
}
func (UnimplementedReqPqAuthenticationServiceServer) mustEmbedUnimplementedReqPqAuthenticationServiceServer() {
}

// UnsafeReqPqAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReqPqAuthenticationServiceServer will
// result in compilation errors.
type UnsafeReqPqAuthenticationServiceServer interface {
	mustEmbedUnimplementedReqPqAuthenticationServiceServer()
}

func RegisterReqPqAuthenticationServiceServer(s grpc.ServiceRegistrar, srv ReqPqAuthenticationServiceServer) {
	s.RegisterService(&ReqPqAuthenticationService_ServiceDesc, srv)
}

func _ReqPqAuthenticationService_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPq_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReqPqAuthenticationServiceServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReqPqAuthenticationService_ReqPq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReqPqAuthenticationServiceServer).ReqPq(ctx, req.(*ReqPq_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ReqPqAuthenticationService_ServiceDesc is the grpc.ServiceDesc for ReqPqAuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReqPqAuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "req_pq_authentication_service",
	HandlerType: (*ReqPqAuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _ReqPqAuthenticationService_ReqPq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "req_pq.proto",
}