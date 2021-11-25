// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortalClient is the client API for Portal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*SetInfoResponse, error)
	GetUptime(ctx context.Context, in *GetUptimeRequest, opts ...grpc.CallOption) (*GetUptimeResponse, error)
	GetRequests(ctx context.Context, in *GetRequestsRequest, opts ...grpc.CallOption) (*GetRequestsResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	GetMode(ctx context.Context, in *GetModeRequest, opts ...grpc.CallOption) (*GetModeResponse, error)
	SetMode(ctx context.Context, in *SetModeRequest, opts ...grpc.CallOption) (*SetModeResponse, error)
}

type portalClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalClient(cc grpc.ClientConnInterface) PortalClient {
	return &portalClient{cc}
}

func (c *portalClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*SetInfoResponse, error) {
	out := new(SetInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetUptime(ctx context.Context, in *GetUptimeRequest, opts ...grpc.CallOption) (*GetUptimeResponse, error) {
	out := new(GetUptimeResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/GetUptime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetRequests(ctx context.Context, in *GetRequestsRequest, opts ...grpc.CallOption) (*GetRequestsResponse, error) {
	out := new(GetRequestsResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/GetRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetMode(ctx context.Context, in *GetModeRequest, opts ...grpc.CallOption) (*GetModeResponse, error) {
	out := new(GetModeResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/GetMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) SetMode(ctx context.Context, in *SetModeRequest, opts ...grpc.CallOption) (*SetModeResponse, error) {
	out := new(SetModeResponse)
	err := c.cc.Invoke(ctx, "/pb.Portal/SetMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalServer is the server API for Portal service.
// All implementations must embed UnimplementedPortalServer
// for forward compatibility
type PortalServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error)
	GetUptime(context.Context, *GetUptimeRequest) (*GetUptimeResponse, error)
	GetRequests(context.Context, *GetRequestsRequest) (*GetRequestsResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	GetMode(context.Context, *GetModeRequest) (*GetModeResponse, error)
	SetMode(context.Context, *SetModeRequest) (*SetModeResponse, error)
	mustEmbedUnimplementedPortalServer()
}

// UnimplementedPortalServer must be embedded to have forward compatible implementations.
type UnimplementedPortalServer struct {
}

func (UnimplementedPortalServer) GetVersion(context.Context, *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedPortalServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedPortalServer) SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (UnimplementedPortalServer) GetUptime(context.Context, *GetUptimeRequest) (*GetUptimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUptime not implemented")
}
func (UnimplementedPortalServer) GetRequests(context.Context, *GetRequestsRequest) (*GetRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequests not implemented")
}
func (UnimplementedPortalServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedPortalServer) GetMode(context.Context, *GetModeRequest) (*GetModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMode not implemented")
}
func (UnimplementedPortalServer) SetMode(context.Context, *SetModeRequest) (*SetModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMode not implemented")
}
func (UnimplementedPortalServer) mustEmbedUnimplementedPortalServer() {}

// UnsafePortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalServer will
// result in compilation errors.
type UnsafePortalServer interface {
	mustEmbedUnimplementedPortalServer()
}

func RegisterPortalServer(s grpc.ServiceRegistrar, srv PortalServer) {
	s.RegisterService(&Portal_ServiceDesc, srv)
}

func _Portal_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).SetInfo(ctx, req.(*SetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetUptime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUptimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetUptime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/GetUptime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetUptime(ctx, req.(*GetUptimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/GetRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetRequests(ctx, req.(*GetRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/GetMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetMode(ctx, req.(*GetModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_SetMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).SetMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Portal/SetMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).SetMode(ctx, req.(*SetModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Portal_ServiceDesc is the grpc.ServiceDesc for Portal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Portal",
	HandlerType: (*PortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Portal_GetVersion_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Portal_GetInfo_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _Portal_SetInfo_Handler,
		},
		{
			MethodName: "GetUptime",
			Handler:    _Portal_GetUptime_Handler,
		},
		{
			MethodName: "GetRequests",
			Handler:    _Portal_GetRequests_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Portal_Reset_Handler,
		},
		{
			MethodName: "GetMode",
			Handler:    _Portal_GetMode_Handler,
		},
		{
			MethodName: "SetMode",
			Handler:    _Portal_SetMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
