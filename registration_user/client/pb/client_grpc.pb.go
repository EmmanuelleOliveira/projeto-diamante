// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/client.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	GetAllClients(ctx context.Context, in *EmptyField, opts ...grpc.CallOption) (*GetAllClientsResponse, error)
	GetClient(ctx context.Context, in *DocNumberRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	DeleteClient(ctx context.Context, in *DocNumberRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetAllClients(ctx context.Context, in *EmptyField, opts ...grpc.CallOption) (*GetAllClientsResponse, error) {
	out := new(GetAllClientsResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetAllClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClient(ctx context.Context, in *DocNumberRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClient(ctx context.Context, in *DocNumberRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateClient(context.Context, *ClientRequest) (*ClientResponse, error)
	GetAllClients(context.Context, *EmptyField) (*GetAllClientsResponse, error)
	GetClient(context.Context, *DocNumberRequest) (*ClientResponse, error)
	UpdateClient(context.Context, *ClientRequest) (*ErrorResponse, error)
	DeleteClient(context.Context, *DocNumberRequest) (*ErrorResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateClient(context.Context, *ClientRequest) (*ClientResponse, error) {
	return &ClientResponse{}, errors.New("entrei aqui")
}
func (UnimplementedClientServiceServer) GetAllClients(context.Context, *EmptyField) (*GetAllClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllClients not implemented")
}
func (UnimplementedClientServiceServer) GetClient(context.Context, *DocNumberRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedClientServiceServer) UpdateClient(context.Context, *ClientRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedClientServiceServer) DeleteClient(context.Context, *DocNumberRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetAllClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyField)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetAllClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetAllClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetAllClients(ctx, req.(*EmptyField))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClient(ctx, req.(*DocNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClient(ctx, req.(*DocNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientService_CreateClient_Handler,
		},
		{
			MethodName: "GetAllClients",
			Handler:    _ClientService_GetAllClients_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _ClientService_GetClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ClientService_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientService_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/client.proto",
}
