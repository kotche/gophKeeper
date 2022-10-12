package service

import (
	"context"
	"errors"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

type IBankCardRepo interface {
	Create(ctx context.Context, lp *domain.BankCard) error
	Update(ctx context.Context, lp *domain.BankCard) error
	Delete(ctx context.Context, lp *domain.BankCard) error
	GetAll(ctx context.Context, userID int) ([]domain.BankCard, error)
}

type BankCardService struct {
	repo IBankCardRepo
	log  *zerolog.Logger
}

func NewBankCardService(repo IBankCardRepo, log *zerolog.Logger) *BankCardService {
	return &BankCardService{
		repo: repo,
		log:  log,
	}
}

func (b *BankCardService) Create(ctx context.Context, bank *domain.BankCard) error {
	err := b.repo.Create(ctx, bank)
	if err != nil {
		return errors.New("create bank card data error")
	}
	return nil
}
func (b *BankCardService) Update(ctx context.Context, bank *domain.BankCard) error {
	err := b.repo.Update(ctx, bank)
	if err != nil {
		return errors.New("update bank card data error")
	}
	return nil
}
func (b *BankCardService) Delete(ctx context.Context, bank *domain.BankCard) error {
	err := b.repo.Delete(ctx, bank)
	if err != nil {
		return errors.New("delete bank card data error")
	}
	return nil
}

func (b *BankCardService) GetAll(ctx context.Context, userID int) ([]domain.BankCard, error) {
	return b.repo.GetAll(ctx, userID)
}
