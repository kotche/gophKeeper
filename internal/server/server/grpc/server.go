package grpc

import (
	"net"

	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/pb"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"google.golang.org/grpc"
)

// Server server gRPC
type Server struct {
	cfg        *server.Config
	handler    *grpcHandler.Handler
	grpcServer *grpc.Server
}

// NewServer get gRPC server
func NewServer(cfg *server.Config, handler *grpcHandler.Handler) *Server {
	authInterceptor := grpc.UnaryInterceptor(handler.UnaryAuthorize)

	return &Server{
		cfg:        cfg,
		handler:    handler,
		grpcServer: grpc.NewServer(authInterceptor),
	}
}

// Run gRPC server with registration api
func (s *Server) Run() error {
	pb.RegisterVersionServiceServer(s.grpcServer, s.handler)
	pb.RegisterAuthServiceServer(s.grpcServer, s.handler)
	pb.RegisterLoginPassServiceServer(s.grpcServer, s.handler)
	pb.RegisterTextServiceServer(s.grpcServer, s.handler)
	pb.RegisterBinaryServiceServer(s.grpcServer, s.handler)
	pb.RegisterBankCardServiceServer(s.grpcServer, s.handler)

	listen, err := net.Listen("tcp", s.cfg.GRPCServer.Address)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(listen)
}

// Stop gRPC server
func (s *Server) Stop() {
	s.grpcServer.Stop()
}
