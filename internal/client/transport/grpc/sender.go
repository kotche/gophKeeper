package grpc

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type IConnection interface {
	GetClientConn(address string, log *zerolog.Logger, interceptors grpc.DialOption) (*grpc.ClientConn, error)
}

type Sender struct {
	Service    *service.Service
	ClientConn IConnection
	Conf       *client.Config
	Log        *zerolog.Logger
}

func NewSender(service *service.Service, clientConn IConnection, conf *client.Config, log *zerolog.Logger) *Sender {
	return &Sender{Service: service,
		ClientConn: clientConn,
		Conf:       conf,
		Log:        log,
	}
}
