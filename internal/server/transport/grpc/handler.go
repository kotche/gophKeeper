package grpc

import (
	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/rs/zerolog"
)

type Handler struct {
	Service *service.Service
	pb.UnimplementedVersionServiceServer
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedLoginPassServiceServer
	pb.UnimplementedTextServiceServer
	pb.UnimplementedBinaryServiceServer
	pb.UnimplementedBankCardServiceServer
	Conf *server.Config
	Log  *zerolog.Logger
}

func NewHandler(service *service.Service, log *zerolog.Logger, conf *server.Config) *Handler {
	handler := &Handler{
		Service: service,
		Log:     log,
		Conf:    conf,
	}
	return handler
}
