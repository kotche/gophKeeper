package service

import (
	"context"
	"errors"

	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"github.com/rs/zerolog"
)

type ILoginPassRepo interface {
	Create(ctx context.Context, lp *dataType.LoginPass) error
	Update(ctx context.Context, lp *dataType.LoginPass) error
	Delete(ctx context.Context, lp *dataType.LoginPass) error
	GetAll(ctx context.Context, userID int) ([]dataType.LoginPass, error)
}

type LoginPassService struct {
	repo ILoginPassRepo
	log  *zerolog.Logger
}

func NewLoginPassService(repo ILoginPassRepo, log *zerolog.Logger) *LoginPassService {
	return &LoginPassService{
		repo: repo,
		log:  log,
	}
}

func (l *LoginPassService) Create(ctx context.Context, lp *dataType.LoginPass) error {
	err := l.repo.Create(ctx, lp)
	if err != nil {
		return errors.New("create lp error")
	}
	return nil
}
func (l *LoginPassService) Update(ctx context.Context, lp *dataType.LoginPass) error {
	err := l.repo.Update(ctx, lp)
	if err != nil {
		return errors.New("update lp error")
	}
	return nil
}
func (l *LoginPassService) Delete(ctx context.Context, lp *dataType.LoginPass) error {
	err := l.repo.Delete(ctx, lp)
	if err != nil {
		return errors.New("delete lp error")
	}
	return nil
}

func (l *LoginPassService) GetAll(ctx context.Context, userID int) ([]dataType.LoginPass, error) {
	return l.repo.GetAll(ctx, userID)
}
