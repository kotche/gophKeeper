package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/errs"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_CreateUser(t *testing.T) {
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := &domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), user).Return(nil).Times(1)
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(authRepo, jwt, pe, log)
	srvc := NewService(authService, nil, nil)

	userReq := &domain.User{
		Username: username,
		Password: password,
	}
	ctx := context.Background()
	err := srvc.Auth.CreateUser(ctx, userReq)
	assert.Equal(t, nil, err)
}

func TestAuthService_CreateUserFailed(t *testing.T) {
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := &domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), user).Return(errs.ConflictLoginError{Username: username}).Times(1)
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(authRepo, jwt, pe, log)
	srvc := NewService(authService, nil, nil)

	userReq := &domain.User{
		Username: username,
		Password: password,
	}
	ctx := context.Background()
	err := srvc.Auth.CreateUser(ctx, userReq)
	assert.Equal(t, errs.ConflictLoginError{Username: username}, errors.Unwrap(err))
}

func TestAuthService_AuthenticationUser(t *testing.T) {
	userID := 10
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := &domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), user).Return(userID, nil).Times(1)
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(authRepo, jwt, pe, log)
	srvc := NewService(authService, nil, nil)

	userReq := &domain.User{
		Username: username,
		Password: password,
	}
	ctx := context.Background()
	err := srvc.Auth.AuthenticationUser(ctx, userReq)
	assert.Equal(t, nil, err)
	assert.Equal(t, userID, userReq.ID)
}

func TestAuthService_AuthenticationUserFailed(t *testing.T) {
	userID := 10
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := &domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), user).Return(userID, errs.AuthenticationError{}).Times(1)
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(authRepo, jwt, pe, log)
	srvc := NewService(authService, nil, nil)

	userReq := &domain.User{
		Username: username,
		Password: password,
	}
	ctx := context.Background()
	err := srvc.Auth.AuthenticationUser(ctx, userReq)
	assert.Equal(t, errs.AuthenticationError{}, err)
}

func TestAuthService_GenerateToken(t *testing.T) {
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := &domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(nil, jwt, pe, log)
	srvc := NewService(authService, nil, nil)

	_, err := srvc.Auth.GenerateToken(user)
	assert.Equal(t, nil, err)
}

func TestAuthService_Verify(t *testing.T) {
	userID := 10
	log := logger.Init("")
	cfg := &server.Config{}

	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(nil, jwt, nil, log)
	srvc := NewService(authService, nil, nil)

	user := &domain.User{
		ID: userID,
	}
	token, _ := srvc.Auth.GenerateToken(user)
	resp, err := srvc.Auth.Verify(token)

	assert.Equal(t, nil, err)
	assert.Equal(t, userID, resp.ID)
}

func TestAuthService_VerifyFail(t *testing.T) {
	log := logger.Init("")
	cfg := &server.Config{}

	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := NewAuthService(nil, jwt, nil, log)
	srvc := NewService(authService, nil, nil)

	_, err := srvc.Auth.Verify("token_test")
	assert.Error(t, err)
}
