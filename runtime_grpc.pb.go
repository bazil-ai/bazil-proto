// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: runtime.proto

package bazil_proto

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

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	Info(ctx context.Context, in *ExpertUID, opts ...grpc.CallOption) (*ExpertInfo, error)
	Backward(ctx context.Context, opts ...grpc.CallOption) (Node_BackwardClient, error)
	Forward(ctx context.Context, opts ...grpc.CallOption) (Node_ForwardClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Info(ctx context.Context, in *ExpertUID, opts ...grpc.CallOption) (*ExpertInfo, error) {
	out := new(ExpertInfo)
	err := c.cc.Invoke(ctx, "/bazil.Node/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Backward(ctx context.Context, opts ...grpc.CallOption) (Node_BackwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], "/bazil.Node/Backward", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeBackwardClient{stream}
	return x, nil
}

type Node_BackwardClient interface {
	Send(*ExpertRequest) error
	Recv() (*Tensor, error)
	grpc.ClientStream
}

type nodeBackwardClient struct {
	grpc.ClientStream
}

func (x *nodeBackwardClient) Send(m *ExpertRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeBackwardClient) Recv() (*Tensor, error) {
	m := new(Tensor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) Forward(ctx context.Context, opts ...grpc.CallOption) (Node_ForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[1], "/bazil.Node/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeForwardClient{stream}
	return x, nil
}

type Node_ForwardClient interface {
	Send(*ExpertRequest) error
	Recv() (*Tensor, error)
	grpc.ClientStream
}

type nodeForwardClient struct {
	grpc.ClientStream
}

func (x *nodeForwardClient) Send(m *ExpertRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeForwardClient) Recv() (*Tensor, error) {
	m := new(Tensor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	Info(context.Context, *ExpertUID) (*ExpertInfo, error)
	Backward(Node_BackwardServer) error
	Forward(Node_ForwardServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Info(context.Context, *ExpertUID) (*ExpertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedNodeServer) Backward(Node_BackwardServer) error {
	return status.Errorf(codes.Unimplemented, "method Backward not implemented")
}
func (UnimplementedNodeServer) Forward(Node_ForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpertUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bazil.Node/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Info(ctx, req.(*ExpertUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Backward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Backward(&nodeBackwardServer{stream})
}

type Node_BackwardServer interface {
	Send(*Tensor) error
	Recv() (*ExpertRequest, error)
	grpc.ServerStream
}

type nodeBackwardServer struct {
	grpc.ServerStream
}

func (x *nodeBackwardServer) Send(m *Tensor) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeBackwardServer) Recv() (*ExpertRequest, error) {
	m := new(ExpertRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Forward(&nodeForwardServer{stream})
}

type Node_ForwardServer interface {
	Send(*Tensor) error
	Recv() (*ExpertRequest, error)
	grpc.ServerStream
}

type nodeForwardServer struct {
	grpc.ServerStream
}

func (x *nodeForwardServer) Send(m *Tensor) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeForwardServer) Recv() (*ExpertRequest, error) {
	m := new(ExpertRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bazil.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Node_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Backward",
			Handler:       _Node_Backward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Forward",
			Handler:       _Node_Forward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "runtime.proto",
}
