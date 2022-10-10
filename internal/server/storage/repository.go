package storage

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
)

type ICommonRepo interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

type IAuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserID(ctx context.Context, user *domain.User) (int, error)
}

type ILoginPassRepo interface {
	Create(ctx context.Context, lp *dataType.LoginPass) error
	GetAll(ctx context.Context, userID int) ([]dataType.LoginPass, error)
}

type Repository struct {
	Common    ICommonRepo
	Auth      IAuthRepo
	LoginPass ILoginPassRepo
}

func NewRepository(com ICommonRepo, auth IAuthRepo, loginPass ILoginPassRepo) *Repository {
	return &Repository{
		Common:    com,
		Auth:      auth,
		LoginPass: loginPass,
	}
}
