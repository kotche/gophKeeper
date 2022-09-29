package grpc

import (
	"github.com/kotche/gophKeeper/internal/client/config"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type IConnection interface {
	GetClientConn(address string, log *zerolog.Logger) (*grpc.ClientConn, error)
}

type Sender struct {
	Service    *service.Service
	ClientConn IConnection
	Conf       *config.Config
	Log        *zerolog.Logger
}

func NewSender(service *service.Service, clientConn IConnection, conf *config.Config, log *zerolog.Logger) *Sender {
	return &Sender{Service: service,
		ClientConn: clientConn,
		Conf:       conf,
		Log:        log,
	}
}
