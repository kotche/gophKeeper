package service

import (
	"context"

	"github.com/rs/zerolog"
)

type ICommonRepo interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

type CommonService struct {
	repo ICommonRepo
	log  *zerolog.Logger
}

func NewCommonService(repo ICommonRepo, log *zerolog.Logger) *CommonService {
	return &CommonService{
		repo: repo,
		log:  log,
	}
}

func (c *CommonService) GetVersion(ctx context.Context, userID int) (uint, error) {
	return c.repo.GetVersion(ctx, userID)
}
