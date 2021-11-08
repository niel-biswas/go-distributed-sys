// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// EventStoreClient is the client API for EventStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventStoreClient interface {
	// Create a new event to the event repository
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	// Get all events for the given aggregate and event
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	//    Get stream of events for the given event
	GetEventsStream(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (EventStore_GetEventsStreamClient, error)
}

type eventStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStoreClient(cc grpc.ClientConnInterface) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/pb.EventStore/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/pb.EventStore/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) GetEventsStream(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (EventStore_GetEventsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventStore_ServiceDesc.Streams[0], "/pb.EventStore/GetEventsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStoreGetEventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStore_GetEventsStreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventStoreGetEventsStreamClient struct {
	grpc.ClientStream
}

func (x *eventStoreGetEventsStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventStoreServer is the server API for EventStore service.
// All implementations must embed UnimplementedEventStoreServer
// for forward compatibility
type EventStoreServer interface {
	// Create a new event to the event repository
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	// Get all events for the given aggregate and event
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	//    Get stream of events for the given event
	GetEventsStream(*GetEventsRequest, EventStore_GetEventsStreamServer) error
	mustEmbedUnimplementedEventStoreServer()
}

// UnimplementedEventStoreServer must be embedded to have forward compatible implementations.
type UnimplementedEventStoreServer struct {
}

func (UnimplementedEventStoreServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventStoreServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventStoreServer) GetEventsStream(*GetEventsRequest, EventStore_GetEventsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventsStream not implemented")
}
func (UnimplementedEventStoreServer) mustEmbedUnimplementedEventStoreServer() {}

// UnsafeEventStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventStoreServer will
// result in compilation errors.
type UnsafeEventStoreServer interface {
	mustEmbedUnimplementedEventStoreServer()
}

func RegisterEventStoreServer(s grpc.ServiceRegistrar, srv EventStoreServer) {
	s.RegisterService(&EventStore_ServiceDesc, srv)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EventStore/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_GetEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStoreServer).GetEventsStream(m, &eventStoreGetEventsStreamServer{stream})
}

type EventStore_GetEventsStreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventStoreGetEventsStreamServer struct {
	grpc.ServerStream
}

func (x *eventStoreGetEventsStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// EventStore_ServiceDesc is the grpc.ServiceDesc for EventStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _EventStore_GetEvents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEventsStream",
			Handler:       _EventStore_GetEventsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/eventstore.proto",
}
