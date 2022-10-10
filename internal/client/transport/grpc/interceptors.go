package grpc

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Sender) authorize(ctx context.Context, method string, req interface{},
	reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	//s.Log.Debug().Msgf("authorize interceptors method: %s", method)

	if strings.Contains(method, "/Login") || strings.Contains(method, "/Authentication") {
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	token, err := s.Service.GetToken()
	if err != nil {
		s.Log.Err(err).Msg("sender client authorize token error")
	}
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	//s.Log.Debug().Msgf("authorize interceptors method: %s token %s", method, token)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func (s *Sender) getInterceptors() grpc.DialOption {
	auth := grpc.WithChainUnaryInterceptor(s.authorize)
	return auth
}
