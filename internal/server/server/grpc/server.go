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
	authInterceptor := grpc.UnaryInterceptor(handler.UnaryAuthorize)

	return &Server{
		cfg:        cfg,
		handler:    handler,
		grpcServer: grpc.NewServer(authInterceptor),
	}
}

func (s *Server) Run() error {
	pb.RegisterVersionServiceServer(s.grpcServer, s.handler)
	pb.RegisterAuthServiceServer(s.grpcServer, s.handler)
	pb.RegisterLoginPassServiceServer(s.grpcServer, s.handler)
	pb.RegisterTextServiceServer(s.grpcServer, s.handler)
	pb.RegisterBinaryServiceServer(s.grpcServer, s.handler)
	pb.RegisterBankCardServiceServer(s.grpcServer, s.handler)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.cfg.TCP.Port))
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(listen)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}
