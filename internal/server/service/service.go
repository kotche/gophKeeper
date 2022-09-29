package service

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

type AuthServiceContract interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticationUser(ctx context.Context, user *domain.User) error
}

type Service struct {
	Auth AuthServiceContract
}

func NewService(auth AuthServiceContract) *Service {
	return &Service{
		Auth: auth,
	}
}
