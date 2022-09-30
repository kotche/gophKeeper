package service

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

type IAuthService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticationUser(ctx context.Context, user *domain.User) error
}

type Service struct {
	Auth IAuthService
}

func NewService(auth IAuthService) *Service {
	return &Service{
		Auth: auth,
	}
}
