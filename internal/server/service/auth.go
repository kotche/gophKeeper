package service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

// IAuthRepo authorization repository api
type IAuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserID(ctx context.Context, user *domain.User) (int, error)
}

// AuthService authorization user service
type AuthService struct {
	repo        IAuthRepo
	log         *zerolog.Logger
	jwt         *JWTManager
	keyPassword string
}

func NewAuthService(repo IAuthRepo, log *zerolog.Logger, jwt *JWTManager, keyPassword string) *AuthService {
	return &AuthService{
		repo:        repo,
		log:         log,
		jwt:         jwt,
		keyPassword: keyPassword,
	}
}

// CreateUser creates a new user
func (auth *AuthService) CreateUser(ctx context.Context, user *domain.User) error {
	user.Password = auth.generatePasswordHash(user.Password)
	err := auth.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("user is not create: %w", err)
	}
	return nil
}

// AuthenticationUser user authentication
func (auth *AuthService) AuthenticationUser(ctx context.Context, user *domain.User) error {
	user.Password = auth.generatePasswordHash(user.Password)
	userID, err := auth.repo.GetUserID(ctx, user)
	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

// GenerateToken generates a token
func (auth *AuthService) GenerateToken(user *domain.User) (string, error) {
	return auth.jwt.Generate(user)
}

// Verify verifies the user by token
func (auth *AuthService) Verify(accessToken string) (*domain.UserClaims, error) {
	return auth.jwt.Verify(accessToken)
}

// generatePasswordHash generates an encrypted password
func (auth *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(auth.keyPassword)))
}
