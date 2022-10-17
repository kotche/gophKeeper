package grpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/errs"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHandler_Login(t *testing.T) {
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := service.NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), &user).Return(nil).Times(1)

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(authRepo, jwt, pe, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	userReq := pb.UserRequest{
		Username: username,
		Password: password,
	}

	ctx := context.Background()
	_, actualError := handler.Login(ctx, &userReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_LoginFailed(t *testing.T) {
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := service.NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().CreateUser(gomock.Any(), &user).Return(errs.ConflictLoginError{}).Times(1)

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(authRepo, jwt, pe, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	userReq := pb.UserRequest{
		Username: username,
		Password: password,
	}

	ctx := context.Background()
	_, actualError := handler.Login(ctx, &userReq)
	var code codes.Code
	st, ok := status.FromError(actualError)
	if ok {
		code = st.Code()
	}
	assert.Equal(t, codes.AlreadyExists, code)
}

func TestHandler_Authentication(t *testing.T) {
	userID := 10
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := service.NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), &user).Return(userID, nil).Times(1)

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(authRepo, jwt, pe, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	userReq := pb.UserRequest{
		Username: username,
		Password: password,
	}

	ctx := context.Background()
	resp, actualError := handler.Authentication(ctx, &userReq)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, int64(userID), resp.Id)
}

func TestHandler_AuthenticationFailed(t *testing.T) {
	userID := 10
	username := "username"
	password := "password"

	log := logger.Init("")
	cfg := &server.Config{}

	pe := service.NewPasswordEncryptor(cfg.SecretKeyPassword)
	user := domain.User{
		Username: username,
		Password: pe.GeneratePasswordHash(password),
	}

	control := gomock.NewController(t)
	authRepo := mock_server.NewMockIAuthRepo(control)
	authRepo.EXPECT().GetUserID(gomock.Any(), &user).Return(userID, errs.AuthenticationError{}).Times(1)

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(authRepo, jwt, pe, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	userReq := pb.UserRequest{
		Username: username,
		Password: password,
	}

	ctx := context.Background()
	_, actualError := handler.Authentication(ctx, &userReq)
	var code codes.Code
	st, ok := status.FromError(actualError)
	if ok {
		code = st.Code()
	}
	assert.Equal(t, codes.Unauthenticated, code)
}
