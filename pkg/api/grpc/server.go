package grpc

import (
	"context"
	"info/pkg/api"
	"info/pkg/api/grpc/pb"
	"info/pkg/domain/user"
	"net"

	"google.golang.org/grpc"
)

const port = ":50001"

type Server struct {
	pb.UnimplementedInfoServer
	ServiceProvider *api.ServiceProvider
}

func NewServer(serviceProvider *api.ServiceProvider) *Server {
	return &Server{
		ServiceProvider: serviceProvider,
	}
}

func (s Server) Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterInfoServer(srv, s)
	if err = srv.Serve(lis); err != nil {
		panic(err)
	}
}

func newActionResponse(uid string, aid string) *pb.ActionResponse {
	return &pb.ActionResponse{
		ActionId: aid,
		UserId:   uid,
	}
}

func (s Server) Action(ctx context.Context, empty *pb.Empty) (*pb.ActionResponse, error) {
	ctx = user.ContextWithValue(ctx)

	u, a := s.ServiceProvider.ActionFacade().Action(ctx)

	return newActionResponse(u, a), nil
}
