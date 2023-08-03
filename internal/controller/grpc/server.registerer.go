package grpc

import (
	server "github.com/MurashovVen/outsider-proto/whether/golang"
	"github.com/MurashovVen/outsider-sdk/grpc"
)

var (
	_ grpc.ServerRegisterer = (*Server)(nil)
)

func (s *Server) Register(serv *grpc.Server) {
	server.RegisterWhetherServer(serv, s)
}
