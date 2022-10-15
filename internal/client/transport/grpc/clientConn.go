package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// ClientConnection client connections gRPC
type ClientConnection struct {
	Conn     *grpc.ClientConn
	Auth     pb.AuthServiceClient
	Version  pb.VersionServiceClient
	Lp       pb.LoginPassServiceClient
	Text     pb.TextServiceClient
	Binary   pb.BinaryServiceClient
	BankCard pb.BankCardServiceClient
}

func NewClientConnection(ctx context.Context, cfg *client.Config, log *zerolog.Logger, interceptors grpc.DialOption) (*ClientConnection, error) {
	conn, err := grpc.DialContext(
		ctx,
		cfg.GRPCClient.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		interceptors,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    cfg.GRPCClient.Time,
			Timeout: cfg.GRPCClient.Timeout,
		}))
	if err != nil {
		log.Err(err).Msg("grpc connect error")
		return nil, err
	}
	return &ClientConnection{
		Conn:     conn,
		Auth:     pb.NewAuthServiceClient(conn),
		Version:  pb.NewVersionServiceClient(conn),
		Lp:       pb.NewLoginPassServiceClient(conn),
		Text:     pb.NewTextServiceClient(conn),
		Binary:   pb.NewBinaryServiceClient(conn),
		BankCard: pb.NewBankCardServiceClient(conn),
	}, nil
}
