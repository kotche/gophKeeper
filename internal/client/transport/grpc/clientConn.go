package grpc

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct {
	conn *grpc.ClientConn
}

func (c Connection) GetClientConn(address string, log *zerolog.Logger, interceptors grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), interceptors)
	if err != nil {
		log.Warn().Err(err).Msg("grpc connect error")
		return nil, err
	}
	return conn, nil
}
