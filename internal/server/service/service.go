package service

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
)

// ICommon common service api
type ICommon interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

// IAuthService authorization service api
type IAuthService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticationUser(ctx context.Context, user *domain.User) error
	GenerateToken(user *domain.User) (string, error)
	Verify(accessToken string) (*domain.UserClaims, error)
}

// IData data service api
type IData interface {
	Create(ctx context.Context, dt any) error
	Update(ctx context.Context, dt any) error
	Delete(ctx context.Context, dt any) error
	GetAll(ctx context.Context, userID int, dt dataType.DataType) (any, error)
}

// Service manager service
type Service struct {
	Common ICommon
	Auth   IAuthService
	Data   IData
}

func NewService(com ICommon, auth IAuthService, data IData) *Service {
	return &Service{
		Common: com,
		Auth:   auth,
		Data:   data,
	}
}
