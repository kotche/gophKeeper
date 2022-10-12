package service

import (
	"context"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

type ICommon interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

type IAuthService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticationUser(ctx context.Context, user *domain.User) error
	GenerateToken(user *domain.User) (string, error)
	Verify(accessToken string) (*domain.UserClaims, error)
}

type ILoginPassService interface {
	Create(ctx context.Context, lp *domain.LoginPass) error
	Update(ctx context.Context, lp *domain.LoginPass) error
	Delete(ctx context.Context, lp *domain.LoginPass) error
	GetAll(ctx context.Context, userID int) ([]domain.LoginPass, error)
}

type ITextService interface {
	Create(ctx context.Context, lp *domain.Text) error
	Update(ctx context.Context, lp *domain.Text) error
	Delete(ctx context.Context, lp *domain.Text) error
	GetAll(ctx context.Context, userID int) ([]domain.Text, error)
}

type IBinaryService interface {
	Create(ctx context.Context, lp *domain.Binary) error
	Update(ctx context.Context, lp *domain.Binary) error
	Delete(ctx context.Context, lp *domain.Binary) error
	GetAll(ctx context.Context, userID int) ([]domain.Binary, error)
}

type IBankCardService interface {
	Create(ctx context.Context, lp *domain.BankCard) error
	Update(ctx context.Context, lp *domain.BankCard) error
	Delete(ctx context.Context, lp *domain.BankCard) error
	GetAll(ctx context.Context, userID int) ([]domain.BankCard, error)
}

type Service struct {
	Common    ICommon
	Auth      IAuthService
	LoginPass ILoginPassService
	Text      ITextService
	Binary    IBinaryService
	BankCard  IBankCardService
}

func NewService(com ICommon, auth IAuthService, lp ILoginPassService, text ITextService,
	binary IBinaryService, bankCard IBankCardService) *Service {
	return &Service{
		Common:    com,
		Auth:      auth,
		LoginPass: lp,
		Text:      text,
		Binary:    binary,
		BankCard:  bankCard,
	}
}
