package grpc

import (
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/config"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/rs/zerolog"
)

type Handler struct {
	Service *service.Service
	pb.UnimplementedKeeperServer
	Conf *config.Config
	Log  *zerolog.Logger
}

func NewHandler(service *service.Service, conf *config.Config) *Handler {
	handler := &Handler{
		Service: service,
		Conf:    conf,
	}
	return handler
}
