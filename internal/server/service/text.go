package service

import (
	"context"
	"errors"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

type ITextRepo interface {
	Create(ctx context.Context, lp *domain.Text) error
	Update(ctx context.Context, lp *domain.Text) error
	Delete(ctx context.Context, lp *domain.Text) error
	GetAll(ctx context.Context, userID int) ([]domain.Text, error)
}

type TextService struct {
	repo ITextRepo
	log  *zerolog.Logger
}

func NewTextService(repo ITextRepo, log *zerolog.Logger) *TextService {
	return &TextService{
		repo: repo,
		log:  log,
	}
}

func (t *TextService) Create(ctx context.Context, text *domain.Text) error {
	err := t.repo.Create(ctx, text)
	if err != nil {
		return errors.New("create text data error")
	}
	return nil
}
func (t *TextService) Update(ctx context.Context, text *domain.Text) error {
	err := t.repo.Update(ctx, text)
	if err != nil {
		return errors.New("update text data error")
	}
	return nil
}
func (t *TextService) Delete(ctx context.Context, text *domain.Text) error {
	err := t.repo.Delete(ctx, text)
	if err != nil {
		return errors.New("delete text data error")
	}
	return nil
}

func (t *TextService) GetAll(ctx context.Context, userID int) ([]domain.Text, error) {
	return t.repo.GetAll(ctx, userID)
}
