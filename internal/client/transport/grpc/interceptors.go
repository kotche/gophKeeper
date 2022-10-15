package grpc

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type IService interface {
	GetToken() string
}

type Interceptors struct {
	Service IService
}

func NewInterceptors(s IService) *Interceptors {
	return &Interceptors{Service: s}
}

func (i *Interceptors) GetInterceptors() grpc.DialOption {
	auth := grpc.WithChainUnaryInterceptor(i.authorize)
	return auth
}

func (i *Interceptors) authorize(ctx context.Context, method string, req interface{},
	reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	if strings.Contains(method, "/Login") || strings.Contains(method, "/Authentication") {
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	token := i.Service.GetToken()
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	
	return invoker(ctx, method, req, reply, cc, opts...)
}
