package service

import (
	"context"
	"fmt"
	"reflect"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"github.com/rs/zerolog"
)

// ILoginPassRepo login-password data repository api
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

// DataService data management service
type DataService struct {
	LpRepo     ILoginPassRepo
	TextRepo   ITextRepo
	BinaryRepo IBinaryRepo
	BankRepo   IBankCardRepo
	Log        *zerolog.Logger
}

func NewDataService(lpRepo ILoginPassRepo, textRepo ITextRepo,
	binaryRepo IBinaryRepo, bankRepo IBankCardRepo, log *zerolog.Logger) *DataService {
	return &DataService{
		LpRepo:     lpRepo,
		TextRepo:   textRepo,
		BinaryRepo: binaryRepo,
		BankRepo:   bankRepo,
		Log:        log,
	}
}

// Create passes the data to be created, depending on the data type
func (d *DataService) Create(ctx context.Context, dt any) error {
	switch data := dt.(type) {
	case *domain.LoginPass:
		return d.LpRepo.Create(ctx, data)
	case *domain.Text:
		return d.TextRepo.Create(ctx, data)
	case *domain.Binary:
		return d.BinaryRepo.Create(ctx, data)
	case *domain.BankCard:
		return d.BankRepo.Create(ctx, data)
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		d.Log.Err(err).Msg("dataService Create error")
		return err
	}
}

// Update passes the data to be updated, depending on the data type
func (d *DataService) Update(ctx context.Context, dt any) error {
	switch data := dt.(type) {
	case *domain.LoginPass:
		return d.LpRepo.Update(ctx, data)
	case *domain.Text:
		return d.TextRepo.Update(ctx, data)
	case *domain.Binary:
		return d.BinaryRepo.Update(ctx, data)
	case *domain.BankCard:
		return d.BankRepo.Update(ctx, data)
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		d.Log.Err(err).Msg("dataService Update error")
		return err
	}
}

// Delete passes the data to be deleted, depending on the data type
func (d *DataService) Delete(ctx context.Context, dt any) error {
	switch data := dt.(type) {
	case *domain.LoginPass:
		return d.LpRepo.Delete(ctx, data)
	case *domain.Text:
		return d.TextRepo.Delete(ctx, data)
	case *domain.Binary:
		return d.BinaryRepo.Delete(ctx, data)
	case *domain.BankCard:
		return d.BankRepo.Delete(ctx, data)
	default:
		err := fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		d.Log.Err(err).Msg("dataService Delete error")
		return err
	}
}

// GetAll gets data depending on the data type
func (d *DataService) GetAll(ctx context.Context, userID int, dt dataType.DataType) (any, error) {
	switch dt {
	case dataType.LP:
		return d.LpRepo.GetAll(ctx, userID)
	case dataType.TEXT:
		return d.TextRepo.GetAll(ctx, userID)
	case dataType.BINARY:
		return d.BinaryRepo.GetAll(ctx, userID)
	case dataType.BANKCARD:
		return d.BankRepo.GetAll(ctx, userID)
	default:
		err := fmt.Errorf("unsupported type '%v'", dt)
		d.Log.Err(err).Msg("dataService GetAll error")
		return nil, err
	}
}
