package storage

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

type IAuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) (int32, error)
	GetUserID(ctx context.Context, user *domain.User) (int32, error)
}

type Repository struct {
	Auth IAuthRepo
}

func NewRepository(auth IAuthRepo) *Repository {
	return &Repository{
		Auth: auth,
	}
}
