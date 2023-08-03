package grpc

import (
	"context"

	server "github.com/MurashovVen/outsider-proto/whether/golang"
	"google.golang.org/protobuf/types/known/emptypb"

	"outsider-whether/internal/service"
)

type Server struct {
	*server.UnimplementedWhetherServer

	service *service.WhetherService
}

var (
	_ server.WhetherServer = (*Server)(nil)
)

func New(service *service.WhetherService) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) ActionProcess(ctx context.Context, request *server.ActionProcessRequest) (*emptypb.Empty, error) {
	err := s.service.ActionProcess(ctx, request.FromChatId, request.Action)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
