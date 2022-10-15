package grpc

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/rs/zerolog"
)

// Sender transport gRPC client
type Sender struct {
	Service    *service.Service
	ClientConn *ClientConnection
	Conf       *client.Config
	Log        *zerolog.Logger
}

func NewSender(service *service.Service, clientConn *ClientConnection, conf *client.Config, log *zerolog.Logger) *Sender {
	return &Sender{Service: service,
		ClientConn: clientConn,
		Conf:       conf,
		Log:        log,
	}
}
