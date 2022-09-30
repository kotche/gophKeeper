package grpc

import (
	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/rs/zerolog"
)

type Handler struct {
	Service *service.Service
	pb.UnimplementedAuthServiceServer
	Conf *server.Config
	Log  *zerolog.Logger
}

func NewHandler(service *service.Service, conf *server.Config) *Handler {
	handler := &Handler{
		Service: service,
		Conf:    conf,
	}
	return handler
}
