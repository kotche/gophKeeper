package grpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestHandler_authorize(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	cfg.TokenDuration = 60 * time.Second

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(nil, jwt, nil, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	user := &domain.User{
		ID: userID,
	}

	ctx := context.Background()
	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewIncomingContext(ctx, md)

	err := handler.authorize(ctx, userID)
	assert.Equal(t, nil, err)
}

func TestHandler_authorizeFail(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	cfg.TokenDuration = 60 * time.Second

	jwt := service.NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	authService := service.NewAuthService(nil, jwt, nil, log)
	srvc := service.NewService(authService, nil, nil)
	handler := NewHandler(srvc, log, cfg)

	ctx := context.Background()

	user := &domain.User{
		ID: userID,
	}

	tests := []struct {
		name   string
		script func()
	}{
		{
			name:   "metadata_not_provided",
			script: func() {},
		},
		{
			name: "no_token",
			script: func() {
				ctx = context.Background()
				md := metadata.New(map[string]string{"token223": "token444"})
				ctx = metadata.NewIncomingContext(ctx, md)
			},
		},
		{
			name: "token_invalid",
			script: func() {
				ctx = context.Background()
				token, _ := authService.GenerateToken(user)
				token = fmt.Sprintf("%ssdfsdfsdfs", token)
				md := metadata.New(map[string]string{"token": token})
				ctx = metadata.NewIncomingContext(ctx, md)
			},
		},
		{
			name: "user_id_incorrect",
			script: func() {
				ctx = context.Background()
				user.ID = 20
				token, _ := authService.GenerateToken(user)
				md := metadata.New(map[string]string{"token": token})
				ctx = metadata.NewIncomingContext(ctx, md)

				user.ID = userID
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.script()

			err := handler.authorize(ctx, userID)
			var code codes.Code
			st, ok := status.FromError(err)
			if ok {
				code = st.Code()
			}
			assert.Equal(t, codes.Unauthenticated, code)
		})
	}
}

func TestHandler_getUserIDFromRequest(t *testing.T) {
	handler := NewHandler(nil, logger.Init(""), nil)

	tests := []struct {
		name       string
		data       any
		wantUserID int
	}{
		{
			name:       "loginPassRequest",
			data:       &pb.LoginPassRequest{UserId: 10},
			wantUserID: 10,
		},
		{
			name:       "textRequest",
			data:       &pb.TextRequest{UserId: 20},
			wantUserID: 20,
		},
		{
			name:       "binaryRequest",
			data:       &pb.BinaryRequest{UserId: 30},
			wantUserID: 30,
		},
		{
			name:       "bankCardRequest",
			data:       &pb.BankCardRequest{UserId: 40},
			wantUserID: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := handler.getUserIDFromRequest(tt.data)
			assert.Equal(t, nil, err)
			assert.Equal(t, tt.wantUserID, id)
		})
	}
}

func TestHandler_getUserIDFromRequestFail(t *testing.T) {
	handler := NewHandler(nil, logger.Init(""), nil)
	_, err := handler.getUserIDFromRequest("unsupported")
	var code codes.Code
	st, ok := status.FromError(err)
	if ok {
		code = st.Code()
	}
	assert.Equal(t, codes.Internal, code)
}
