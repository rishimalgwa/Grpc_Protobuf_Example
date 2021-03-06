// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: heartbeat_pb/heartbeat.proto

package heartbeat

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

// HeartBeatServiceClient is the client API for HeartBeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeatServiceClient interface {
	UserHeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
	LiveUserHeartBeat(ctx context.Context, opts ...grpc.CallOption) (HeartBeatService_LiveUserHeartBeatClient, error)
	UserHeartBeatHistory(ctx context.Context, in *HeartBeatHistoryRequest, opts ...grpc.CallOption) (HeartBeatService_UserHeartBeatHistoryClient, error)
	NormalAbnormalHeartBeat(ctx context.Context, opts ...grpc.CallOption) (HeartBeatService_NormalAbnormalHeartBeatClient, error)
}

type heartBeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeatServiceClient(cc grpc.ClientConnInterface) HeartBeatServiceClient {
	return &heartBeatServiceClient{cc}
}

func (c *heartBeatServiceClient) UserHeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/heartbeat_pb.HeartBeatService/UserHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartBeatServiceClient) LiveUserHeartBeat(ctx context.Context, opts ...grpc.CallOption) (HeartBeatService_LiveUserHeartBeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &HeartBeatService_ServiceDesc.Streams[0], "/heartbeat_pb.HeartBeatService/LiveUserHeartBeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &heartBeatServiceLiveUserHeartBeatClient{stream}
	return x, nil
}

type HeartBeatService_LiveUserHeartBeatClient interface {
	Send(*LiveHeartBeatRequest) error
	CloseAndRecv() (*LiveHeartBeatResponse, error)
	grpc.ClientStream
}

type heartBeatServiceLiveUserHeartBeatClient struct {
	grpc.ClientStream
}

func (x *heartBeatServiceLiveUserHeartBeatClient) Send(m *LiveHeartBeatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *heartBeatServiceLiveUserHeartBeatClient) CloseAndRecv() (*LiveHeartBeatResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LiveHeartBeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heartBeatServiceClient) UserHeartBeatHistory(ctx context.Context, in *HeartBeatHistoryRequest, opts ...grpc.CallOption) (HeartBeatService_UserHeartBeatHistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &HeartBeatService_ServiceDesc.Streams[1], "/heartbeat_pb.HeartBeatService/UserHeartBeatHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &heartBeatServiceUserHeartBeatHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HeartBeatService_UserHeartBeatHistoryClient interface {
	Recv() (*HeartBeatHistoryResponse, error)
	grpc.ClientStream
}

type heartBeatServiceUserHeartBeatHistoryClient struct {
	grpc.ClientStream
}

func (x *heartBeatServiceUserHeartBeatHistoryClient) Recv() (*HeartBeatHistoryResponse, error) {
	m := new(HeartBeatHistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heartBeatServiceClient) NormalAbnormalHeartBeat(ctx context.Context, opts ...grpc.CallOption) (HeartBeatService_NormalAbnormalHeartBeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &HeartBeatService_ServiceDesc.Streams[2], "/heartbeat_pb.HeartBeatService/NormalAbnormalHeartBeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &heartBeatServiceNormalAbnormalHeartBeatClient{stream}
	return x, nil
}

type HeartBeatService_NormalAbnormalHeartBeatClient interface {
	Send(*NormalAbnormalHeartBeatRequest) error
	Recv() (*NormalAbnormalHeartBeatResponse, error)
	grpc.ClientStream
}

type heartBeatServiceNormalAbnormalHeartBeatClient struct {
	grpc.ClientStream
}

func (x *heartBeatServiceNormalAbnormalHeartBeatClient) Send(m *NormalAbnormalHeartBeatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *heartBeatServiceNormalAbnormalHeartBeatClient) Recv() (*NormalAbnormalHeartBeatResponse, error) {
	m := new(NormalAbnormalHeartBeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeartBeatServiceServer is the server API for HeartBeatService service.
// All implementations must embed UnimplementedHeartBeatServiceServer
// for forward compatibility
type HeartBeatServiceServer interface {
	UserHeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
	LiveUserHeartBeat(HeartBeatService_LiveUserHeartBeatServer) error
	UserHeartBeatHistory(*HeartBeatHistoryRequest, HeartBeatService_UserHeartBeatHistoryServer) error
	NormalAbnormalHeartBeat(HeartBeatService_NormalAbnormalHeartBeatServer) error
	mustEmbedUnimplementedHeartBeatServiceServer()
}

// UnimplementedHeartBeatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServiceServer struct {
}

func (UnimplementedHeartBeatServiceServer) UserHeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHeartBeat not implemented")
}
func (UnimplementedHeartBeatServiceServer) LiveUserHeartBeat(HeartBeatService_LiveUserHeartBeatServer) error {
	return status.Errorf(codes.Unimplemented, "method LiveUserHeartBeat not implemented")
}
func (UnimplementedHeartBeatServiceServer) UserHeartBeatHistory(*HeartBeatHistoryRequest, HeartBeatService_UserHeartBeatHistoryServer) error {
	return status.Errorf(codes.Unimplemented, "method UserHeartBeatHistory not implemented")
}
func (UnimplementedHeartBeatServiceServer) NormalAbnormalHeartBeat(HeartBeatService_NormalAbnormalHeartBeatServer) error {
	return status.Errorf(codes.Unimplemented, "method NormalAbnormalHeartBeat not implemented")
}
func (UnimplementedHeartBeatServiceServer) mustEmbedUnimplementedHeartBeatServiceServer() {}

// UnsafeHeartBeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeatServiceServer will
// result in compilation errors.
type UnsafeHeartBeatServiceServer interface {
	mustEmbedUnimplementedHeartBeatServiceServer()
}

func RegisterHeartBeatServiceServer(s grpc.ServiceRegistrar, srv HeartBeatServiceServer) {
	s.RegisterService(&HeartBeatService_ServiceDesc, srv)
}

func _HeartBeatService_UserHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServiceServer).UserHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat_pb.HeartBeatService/UserHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServiceServer).UserHeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeartBeatService_LiveUserHeartBeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HeartBeatServiceServer).LiveUserHeartBeat(&heartBeatServiceLiveUserHeartBeatServer{stream})
}

type HeartBeatService_LiveUserHeartBeatServer interface {
	SendAndClose(*LiveHeartBeatResponse) error
	Recv() (*LiveHeartBeatRequest, error)
	grpc.ServerStream
}

type heartBeatServiceLiveUserHeartBeatServer struct {
	grpc.ServerStream
}

func (x *heartBeatServiceLiveUserHeartBeatServer) SendAndClose(m *LiveHeartBeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *heartBeatServiceLiveUserHeartBeatServer) Recv() (*LiveHeartBeatRequest, error) {
	m := new(LiveHeartBeatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HeartBeatService_UserHeartBeatHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HeartBeatHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeartBeatServiceServer).UserHeartBeatHistory(m, &heartBeatServiceUserHeartBeatHistoryServer{stream})
}

type HeartBeatService_UserHeartBeatHistoryServer interface {
	Send(*HeartBeatHistoryResponse) error
	grpc.ServerStream
}

type heartBeatServiceUserHeartBeatHistoryServer struct {
	grpc.ServerStream
}

func (x *heartBeatServiceUserHeartBeatHistoryServer) Send(m *HeartBeatHistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HeartBeatService_NormalAbnormalHeartBeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HeartBeatServiceServer).NormalAbnormalHeartBeat(&heartBeatServiceNormalAbnormalHeartBeatServer{stream})
}

type HeartBeatService_NormalAbnormalHeartBeatServer interface {
	Send(*NormalAbnormalHeartBeatResponse) error
	Recv() (*NormalAbnormalHeartBeatRequest, error)
	grpc.ServerStream
}

type heartBeatServiceNormalAbnormalHeartBeatServer struct {
	grpc.ServerStream
}

func (x *heartBeatServiceNormalAbnormalHeartBeatServer) Send(m *NormalAbnormalHeartBeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *heartBeatServiceNormalAbnormalHeartBeatServer) Recv() (*NormalAbnormalHeartBeatRequest, error) {
	m := new(NormalAbnormalHeartBeatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeartBeatService_ServiceDesc is the grpc.ServiceDesc for HeartBeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heartbeat_pb.HeartBeatService",
	HandlerType: (*HeartBeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserHeartBeat",
			Handler:    _HeartBeatService_UserHeartBeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LiveUserHeartBeat",
			Handler:       _HeartBeatService_LiveUserHeartBeat_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UserHeartBeatHistory",
			Handler:       _HeartBeatService_UserHeartBeatHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NormalAbnormalHeartBeat",
			Handler:       _HeartBeatService_NormalAbnormalHeartBeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "heartbeat_pb/heartbeat.proto",
}
