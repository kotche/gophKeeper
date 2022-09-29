package grpc

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct{}

func (c Connection) GetClientConn(address string, log *zerolog.Logger) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warn().Err(err).Msg("grpc connect error")
		return nil, err
	}
	return conn, nil
}
