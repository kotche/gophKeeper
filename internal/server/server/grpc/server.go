package grpc

import (
	"fmt"
	"net"

	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/pb"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	cfg        *server.Config
	handler    *grpcHandler.Handler
	grpcServer *grpc.Server
}

func NewServer(cfg *server.Config, handler *grpcHandler.Handler) *Server {
	return &Server{
		cfg:        cfg,
		handler:    handler,
		grpcServer: grpc.NewServer(),
	}
}

func (s *Server) Run() error {
	pb.RegisterAuthServiceServer(s.grpcServer, s.handler)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.cfg.TCP.Port))
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(listen)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}
