package service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

const (
	secretKey = "be55d1079e6c6167118ac91318fe"
)

type AuthRepoContract interface {
	CreateUser(ctx context.Context, user *domain.User) (int32, error)
	GetUserID(ctx context.Context, user *domain.User) (int32, error)
}

type AuthService struct {
	repo AuthRepoContract
	log  *zerolog.Logger
}

func NewAuthService(repo AuthRepoContract, log *zerolog.Logger) *AuthService {
	return &AuthService{
		repo: repo,
		log:  log,
	}
}

func (auth *AuthService) CreateUser(ctx context.Context, user *domain.User) error {
	user.Password = auth.generatePasswordHash(user.Password)
	userID, err := auth.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

func (auth *AuthService) AuthenticationUser(ctx context.Context, user *domain.User) error {
	user.Password = auth.generatePasswordHash(user.Password)
	userID, err := auth.repo.GetUserID(ctx, user)
	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

func (auth *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(secretKey)))
}
