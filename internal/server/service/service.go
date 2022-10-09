package service

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
)

type IAuthService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticationUser(ctx context.Context, user *domain.User) error
	GenerateToken(user *domain.User) (string, error)
	Verify(accessToken string) (*domain.UserClaims, error)
}

type ILoginPassService interface {
	Create(ctx context.Context, lp *dataType.LoginPass) error
	GetAll(ctx context.Context, userID int) ([]dataType.LoginPass, error)
}

type Service struct {
	Auth      IAuthService
	LoginPass ILoginPassService
}

func NewService(auth IAuthService, lp ILoginPassService) *Service {
	return &Service{
		Auth:      auth,
		LoginPass: lp,
	}
}
