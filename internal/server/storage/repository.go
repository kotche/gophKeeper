package storage

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

type AuthRepoContract interface {
	CreateUser(ctx context.Context, user *domain.User) (int32, error)
	GetUserID(ctx context.Context, user *domain.User) (int32, error)
}

type Repository struct {
	Auth AuthRepoContract
}

func NewRepository(auth AuthRepoContract) *Repository {
	return &Repository{
		Auth: auth,
	}
}
