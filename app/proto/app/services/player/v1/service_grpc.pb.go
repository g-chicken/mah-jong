// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package player

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

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerServiceClient interface {
	// CreatePlayer creates a player.
	// If the name have already exist in DB, return conflict error.
	CreatePlayer(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// FetchPlayers search for players.
	// If it is no player in DB, return empty list.
	FetchPlayers(ctx context.Context, in *FetchPlayersRequest, opts ...grpc.CallOption) (*FetchPlayersResponse, error)
}

type playerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerServiceClient(cc grpc.ClientConnInterface) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) CreatePlayer(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/app.services.player.v1.PlayerService/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) FetchPlayers(ctx context.Context, in *FetchPlayersRequest, opts ...grpc.CallOption) (*FetchPlayersResponse, error) {
	out := new(FetchPlayersResponse)
	err := c.cc.Invoke(ctx, "/app.services.player.v1.PlayerService/FetchPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
// All implementations should embed UnimplementedPlayerServiceServer
// for forward compatibility
type PlayerServiceServer interface {
	// CreatePlayer creates a player.
	// If the name have already exist in DB, return conflict error.
	CreatePlayer(context.Context, *CreateRequest) (*CreateResponse, error)
	// FetchPlayers search for players.
	// If it is no player in DB, return empty list.
	FetchPlayers(context.Context, *FetchPlayersRequest) (*FetchPlayersResponse, error)
}

// UnimplementedPlayerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPlayerServiceServer struct {
}

func (UnimplementedPlayerServiceServer) CreatePlayer(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}
func (UnimplementedPlayerServiceServer) FetchPlayers(context.Context, *FetchPlayersRequest) (*FetchPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPlayers not implemented")
}

// UnsafePlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServiceServer will
// result in compilation errors.
type UnsafePlayerServiceServer interface {
	mustEmbedUnimplementedPlayerServiceServer()
}

func RegisterPlayerServiceServer(s grpc.ServiceRegistrar, srv PlayerServiceServer) {
	s.RegisterService(&PlayerService_ServiceDesc, srv)
}

func _PlayerService_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.services.player.v1.PlayerService/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).CreatePlayer(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_FetchPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).FetchPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.services.player.v1.PlayerService/FetchPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).FetchPlayers(ctx, req.(*FetchPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerService_ServiceDesc is the grpc.ServiceDesc for PlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.services.player.v1.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayer",
			Handler:    _PlayerService_CreatePlayer_Handler,
		},
		{
			MethodName: "FetchPlayers",
			Handler:    _PlayerService_FetchPlayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/services/player/v1/service.proto",
}
