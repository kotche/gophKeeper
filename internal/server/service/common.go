package service

import (
	"context"

	"github.com/rs/zerolog"
)

// ICommonRepo common api repository
type ICommonRepo interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

// CommonService common service
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

// GetVersion gets the data version service
func (c *CommonService) GetVersion(ctx context.Context, userID int) (uint, error) {
	return c.repo.GetVersion(ctx, userID)
}
