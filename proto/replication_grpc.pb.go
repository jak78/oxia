// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/replication.proto

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

// OxiaCoordinationClient is the client API for OxiaCoordination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OxiaCoordinationClient interface {
	PushShardAssignments(ctx context.Context, opts ...grpc.CallOption) (OxiaCoordination_PushShardAssignmentsClient, error)
	Fence(ctx context.Context, in *FenceRequest, opts ...grpc.CallOption) (*FenceResponse, error)
	BecomeLeader(ctx context.Context, in *BecomeLeaderRequest, opts ...grpc.CallOption) (*BecomeLeaderResponse, error)
	AddFollower(ctx context.Context, in *AddFollowerRequest, opts ...grpc.CallOption) (*AddFollowerResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type oxiaCoordinationClient struct {
	cc grpc.ClientConnInterface
}

func NewOxiaCoordinationClient(cc grpc.ClientConnInterface) OxiaCoordinationClient {
	return &oxiaCoordinationClient{cc}
}

func (c *oxiaCoordinationClient) PushShardAssignments(ctx context.Context, opts ...grpc.CallOption) (OxiaCoordination_PushShardAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaCoordination_ServiceDesc.Streams[0], "/replication.OxiaCoordination/PushShardAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaCoordinationPushShardAssignmentsClient{stream}
	return x, nil
}

type OxiaCoordination_PushShardAssignmentsClient interface {
	Send(*ShardAssignments) error
	CloseAndRecv() (*CoordinationShardAssignmentsResponse, error)
	grpc.ClientStream
}

type oxiaCoordinationPushShardAssignmentsClient struct {
	grpc.ClientStream
}

func (x *oxiaCoordinationPushShardAssignmentsClient) Send(m *ShardAssignments) error {
	return x.ClientStream.SendMsg(m)
}

func (x *oxiaCoordinationPushShardAssignmentsClient) CloseAndRecv() (*CoordinationShardAssignmentsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CoordinationShardAssignmentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oxiaCoordinationClient) Fence(ctx context.Context, in *FenceRequest, opts ...grpc.CallOption) (*FenceResponse, error) {
	out := new(FenceResponse)
	err := c.cc.Invoke(ctx, "/replication.OxiaCoordination/Fence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaCoordinationClient) BecomeLeader(ctx context.Context, in *BecomeLeaderRequest, opts ...grpc.CallOption) (*BecomeLeaderResponse, error) {
	out := new(BecomeLeaderResponse)
	err := c.cc.Invoke(ctx, "/replication.OxiaCoordination/BecomeLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaCoordinationClient) AddFollower(ctx context.Context, in *AddFollowerRequest, opts ...grpc.CallOption) (*AddFollowerResponse, error) {
	out := new(AddFollowerResponse)
	err := c.cc.Invoke(ctx, "/replication.OxiaCoordination/AddFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaCoordinationClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/replication.OxiaCoordination/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OxiaCoordinationServer is the server API for OxiaCoordination service.
// All implementations must embed UnimplementedOxiaCoordinationServer
// for forward compatibility
type OxiaCoordinationServer interface {
	PushShardAssignments(OxiaCoordination_PushShardAssignmentsServer) error
	Fence(context.Context, *FenceRequest) (*FenceResponse, error)
	BecomeLeader(context.Context, *BecomeLeaderRequest) (*BecomeLeaderResponse, error)
	AddFollower(context.Context, *AddFollowerRequest) (*AddFollowerResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	mustEmbedUnimplementedOxiaCoordinationServer()
}

// UnimplementedOxiaCoordinationServer must be embedded to have forward compatible implementations.
type UnimplementedOxiaCoordinationServer struct {
}

func (UnimplementedOxiaCoordinationServer) PushShardAssignments(OxiaCoordination_PushShardAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method PushShardAssignments not implemented")
}
func (UnimplementedOxiaCoordinationServer) Fence(context.Context, *FenceRequest) (*FenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fence not implemented")
}
func (UnimplementedOxiaCoordinationServer) BecomeLeader(context.Context, *BecomeLeaderRequest) (*BecomeLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BecomeLeader not implemented")
}
func (UnimplementedOxiaCoordinationServer) AddFollower(context.Context, *AddFollowerRequest) (*AddFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollower not implemented")
}
func (UnimplementedOxiaCoordinationServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedOxiaCoordinationServer) mustEmbedUnimplementedOxiaCoordinationServer() {}

// UnsafeOxiaCoordinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OxiaCoordinationServer will
// result in compilation errors.
type UnsafeOxiaCoordinationServer interface {
	mustEmbedUnimplementedOxiaCoordinationServer()
}

func RegisterOxiaCoordinationServer(s grpc.ServiceRegistrar, srv OxiaCoordinationServer) {
	s.RegisterService(&OxiaCoordination_ServiceDesc, srv)
}

func _OxiaCoordination_PushShardAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OxiaCoordinationServer).PushShardAssignments(&oxiaCoordinationPushShardAssignmentsServer{stream})
}

type OxiaCoordination_PushShardAssignmentsServer interface {
	SendAndClose(*CoordinationShardAssignmentsResponse) error
	Recv() (*ShardAssignments, error)
	grpc.ServerStream
}

type oxiaCoordinationPushShardAssignmentsServer struct {
	grpc.ServerStream
}

func (x *oxiaCoordinationPushShardAssignmentsServer) SendAndClose(m *CoordinationShardAssignmentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *oxiaCoordinationPushShardAssignmentsServer) Recv() (*ShardAssignments, error) {
	m := new(ShardAssignments)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OxiaCoordination_Fence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaCoordinationServer).Fence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replication.OxiaCoordination/Fence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaCoordinationServer).Fence(ctx, req.(*FenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaCoordination_BecomeLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BecomeLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaCoordinationServer).BecomeLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replication.OxiaCoordination/BecomeLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaCoordinationServer).BecomeLeader(ctx, req.(*BecomeLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaCoordination_AddFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaCoordinationServer).AddFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replication.OxiaCoordination/AddFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaCoordinationServer).AddFollower(ctx, req.(*AddFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaCoordination_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaCoordinationServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replication.OxiaCoordination/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaCoordinationServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OxiaCoordination_ServiceDesc is the grpc.ServiceDesc for OxiaCoordination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OxiaCoordination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.OxiaCoordination",
	HandlerType: (*OxiaCoordinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fence",
			Handler:    _OxiaCoordination_Fence_Handler,
		},
		{
			MethodName: "BecomeLeader",
			Handler:    _OxiaCoordination_BecomeLeader_Handler,
		},
		{
			MethodName: "AddFollower",
			Handler:    _OxiaCoordination_AddFollower_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _OxiaCoordination_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushShardAssignments",
			Handler:       _OxiaCoordination_PushShardAssignments_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/replication.proto",
}

// OxiaLogReplicationClient is the client API for OxiaLogReplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OxiaLogReplicationClient interface {
	Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateResponse, error)
	Replicate(ctx context.Context, opts ...grpc.CallOption) (OxiaLogReplication_ReplicateClient, error)
	SendSnapshot(ctx context.Context, opts ...grpc.CallOption) (OxiaLogReplication_SendSnapshotClient, error)
}

type oxiaLogReplicationClient struct {
	cc grpc.ClientConnInterface
}

func NewOxiaLogReplicationClient(cc grpc.ClientConnInterface) OxiaLogReplicationClient {
	return &oxiaLogReplicationClient{cc}
}

func (c *oxiaLogReplicationClient) Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateResponse, error) {
	out := new(TruncateResponse)
	err := c.cc.Invoke(ctx, "/replication.OxiaLogReplication/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaLogReplicationClient) Replicate(ctx context.Context, opts ...grpc.CallOption) (OxiaLogReplication_ReplicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaLogReplication_ServiceDesc.Streams[0], "/replication.OxiaLogReplication/Replicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaLogReplicationReplicateClient{stream}
	return x, nil
}

type OxiaLogReplication_ReplicateClient interface {
	Send(*Append) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type oxiaLogReplicationReplicateClient struct {
	grpc.ClientStream
}

func (x *oxiaLogReplicationReplicateClient) Send(m *Append) error {
	return x.ClientStream.SendMsg(m)
}

func (x *oxiaLogReplicationReplicateClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oxiaLogReplicationClient) SendSnapshot(ctx context.Context, opts ...grpc.CallOption) (OxiaLogReplication_SendSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaLogReplication_ServiceDesc.Streams[1], "/replication.OxiaLogReplication/SendSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaLogReplicationSendSnapshotClient{stream}
	return x, nil
}

type OxiaLogReplication_SendSnapshotClient interface {
	Send(*SnapshotChunk) error
	CloseAndRecv() (*SnapshotResponse, error)
	grpc.ClientStream
}

type oxiaLogReplicationSendSnapshotClient struct {
	grpc.ClientStream
}

func (x *oxiaLogReplicationSendSnapshotClient) Send(m *SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *oxiaLogReplicationSendSnapshotClient) CloseAndRecv() (*SnapshotResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OxiaLogReplicationServer is the server API for OxiaLogReplication service.
// All implementations must embed UnimplementedOxiaLogReplicationServer
// for forward compatibility
type OxiaLogReplicationServer interface {
	Truncate(context.Context, *TruncateRequest) (*TruncateResponse, error)
	Replicate(OxiaLogReplication_ReplicateServer) error
	SendSnapshot(OxiaLogReplication_SendSnapshotServer) error
	mustEmbedUnimplementedOxiaLogReplicationServer()
}

// UnimplementedOxiaLogReplicationServer must be embedded to have forward compatible implementations.
type UnimplementedOxiaLogReplicationServer struct {
}

func (UnimplementedOxiaLogReplicationServer) Truncate(context.Context, *TruncateRequest) (*TruncateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (UnimplementedOxiaLogReplicationServer) Replicate(OxiaLogReplication_ReplicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Replicate not implemented")
}
func (UnimplementedOxiaLogReplicationServer) SendSnapshot(OxiaLogReplication_SendSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSnapshot not implemented")
}
func (UnimplementedOxiaLogReplicationServer) mustEmbedUnimplementedOxiaLogReplicationServer() {}

// UnsafeOxiaLogReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OxiaLogReplicationServer will
// result in compilation errors.
type UnsafeOxiaLogReplicationServer interface {
	mustEmbedUnimplementedOxiaLogReplicationServer()
}

func RegisterOxiaLogReplicationServer(s grpc.ServiceRegistrar, srv OxiaLogReplicationServer) {
	s.RegisterService(&OxiaLogReplication_ServiceDesc, srv)
}

func _OxiaLogReplication_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaLogReplicationServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replication.OxiaLogReplication/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaLogReplicationServer).Truncate(ctx, req.(*TruncateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaLogReplication_Replicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OxiaLogReplicationServer).Replicate(&oxiaLogReplicationReplicateServer{stream})
}

type OxiaLogReplication_ReplicateServer interface {
	Send(*Ack) error
	Recv() (*Append, error)
	grpc.ServerStream
}

type oxiaLogReplicationReplicateServer struct {
	grpc.ServerStream
}

func (x *oxiaLogReplicationReplicateServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *oxiaLogReplicationReplicateServer) Recv() (*Append, error) {
	m := new(Append)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OxiaLogReplication_SendSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OxiaLogReplicationServer).SendSnapshot(&oxiaLogReplicationSendSnapshotServer{stream})
}

type OxiaLogReplication_SendSnapshotServer interface {
	SendAndClose(*SnapshotResponse) error
	Recv() (*SnapshotChunk, error)
	grpc.ServerStream
}

type oxiaLogReplicationSendSnapshotServer struct {
	grpc.ServerStream
}

func (x *oxiaLogReplicationSendSnapshotServer) SendAndClose(m *SnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *oxiaLogReplicationSendSnapshotServer) Recv() (*SnapshotChunk, error) {
	m := new(SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OxiaLogReplication_ServiceDesc is the grpc.ServiceDesc for OxiaLogReplication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OxiaLogReplication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.OxiaLogReplication",
	HandlerType: (*OxiaLogReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Truncate",
			Handler:    _OxiaLogReplication_Truncate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Replicate",
			Handler:       _OxiaLogReplication_Replicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendSnapshot",
			Handler:       _OxiaLogReplication_SendSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/replication.proto",
}
