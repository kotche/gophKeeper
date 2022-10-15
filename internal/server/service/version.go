package service

import (
	"context"

	"github.com/rs/zerolog"
)

// IVersionRepo common api repository
type IVersionRepo interface {
	GetVersion(ctx context.Context, userID int) (uint, error)
}

// VersionService data version service
type VersionService struct {
	repo IVersionRepo
	log  *zerolog.Logger
}

func NewVersionService(repo IVersionRepo, log *zerolog.Logger) *VersionService {
	return &VersionService{
		repo: repo,
		log:  log,
	}
}

// GetVersion gets the data version service
func (v *VersionService) GetVersion(ctx context.Context, userID int) (uint, error) {
	return v.repo.GetVersion(ctx, userID)
}
