package storage

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
)

type IAuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserID(ctx context.Context, user *domain.User) (int, error)
}

type ILoginPassRepo interface {
	Create(ctx context.Context, lp *dataType.LoginPass) error
}

type Repository struct {
	Auth      IAuthRepo
	LoginPass ILoginPassRepo
}

func NewRepository(auth IAuthRepo, loginPass ILoginPassRepo) *Repository {
	return &Repository{
		Auth:      auth,
		LoginPass: loginPass,
	}
}
