package service

import (
	"context"
	"errors"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

type IBinaryRepo interface {
	Create(ctx context.Context, lp *domain.Binary) error
	Update(ctx context.Context, lp *domain.Binary) error
	Delete(ctx context.Context, lp *domain.Binary) error
	GetAll(ctx context.Context, userID int) ([]domain.Binary, error)
}

type BinaryService struct {
	repo IBinaryRepo
	log  *zerolog.Logger
}

func NewBinaryService(repo IBinaryRepo, log *zerolog.Logger) *BinaryService {
	return &BinaryService{
		repo: repo,
		log:  log,
	}
}

func (b *BinaryService) Create(ctx context.Context, bin *domain.Binary) error {
	err := b.repo.Create(ctx, bin)
	if err != nil {
		return errors.New("create binary data error")
	}
	return nil
}
func (b *BinaryService) Update(ctx context.Context, bin *domain.Binary) error {
	err := b.repo.Update(ctx, bin)
	if err != nil {
		return errors.New("update binary data error")
	}
	return nil
}
func (b *BinaryService) Delete(ctx context.Context, bin *domain.Binary) error {
	err := b.repo.Delete(ctx, bin)
	if err != nil {
		return errors.New("delete binary data error")
	}
	return nil
}

func (b *BinaryService) GetAll(ctx context.Context, userID int) ([]domain.Binary, error) {
	return b.repo.GetAll(ctx, userID)
}
