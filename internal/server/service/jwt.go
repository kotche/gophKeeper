package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

// JWTManager controls user authorization
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
	log           *zerolog.Logger
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey string, tokenDuration time.Duration, log *zerolog.Logger) *JWTManager {
	return &JWTManager{secretKey, tokenDuration, log}
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(user *domain.User) (string, error) {
	claims := domain.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		ID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*domain.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&domain.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				manager.log.Debug().Msg("JWTManager verify unexpected token")
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		manager.log.Err(err).Msg("JWTManager verify invalid token")
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*domain.UserClaims)
	if !ok {
		manager.log.Debug().Msg("JWTManager verify invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
