package storage

import (
	"context"
	"database/sql"

	"github.com/kotche/gophKeeper/internal/server/domain"
)

// IVersionRepo data version repository api
type IVersionRepo interface {
	InsertVersion(ctx context.Context, userID int, tx *sql.Tx) error
	UpdateVersion(ctx context.Context, userID int, tx *sql.Tx) error
	GetVersion(ctx context.Context, userID int) (uint, error)
}

// IAuthRepo authentication repository api
type IAuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserID(ctx context.Context, user *domain.User) (int, error)
}

// ILoginPassRepo repository login password data api
type ILoginPassRepo interface {
	Create(ctx context.Context, lp *domain.LoginPass) error
	Update(ctx context.Context, lp *domain.LoginPass) error
	Delete(ctx context.Context, lp *domain.LoginPass) error
	GetAll(ctx context.Context, userID int) ([]domain.LoginPass, error)
}

// ITextRepo text data repository api
type ITextRepo interface {
	Create(ctx context.Context, text *domain.Text) error
	Update(ctx context.Context, text *domain.Text) error
	Delete(ctx context.Context, text *domain.Text) error
	GetAll(ctx context.Context, userID int) ([]domain.Text, error)
}

// IBinaryRepo binary data repository api
type IBinaryRepo interface {
	Create(ctx context.Context, text *domain.Binary) error
	Update(ctx context.Context, text *domain.Binary) error
	Delete(ctx context.Context, text *domain.Binary) error
	GetAll(ctx context.Context, userID int) ([]domain.Binary, error)
}

// IBankCardRepo bank card data repository api
type IBankCardRepo interface {
	Create(ctx context.Context, text *domain.BankCard) error
	Update(ctx context.Context, text *domain.BankCard) error
	Delete(ctx context.Context, text *domain.BankCard) error
	GetAll(ctx context.Context, userID int) ([]domain.BankCard, error)
}

// Repository manager repositories
type Repository struct {
	Version   IVersionRepo
	Auth      IAuthRepo
	LoginPass ILoginPassRepo
	Text      ITextRepo
	Binary    IBinaryRepo
	BankCard  IBankCardRepo
}

func NewRepository(ver IVersionRepo, auth IAuthRepo, loginPass ILoginPassRepo, text ITextRepo,
	binary IBinaryRepo, bankCard IBankCardRepo) *Repository {

	return &Repository{
		Version:   ver,
		Auth:      auth,
		LoginPass: loginPass,
		Text:      text,
		Binary:    binary,
		BankCard:  bankCard,
	}
}
