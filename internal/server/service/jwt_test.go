package service

import (
	"testing"

	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestJWTManager_Generate(t *testing.T) {
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

	_, err := jwt.Generate(user)
	assert.Equal(t, nil, err)
}

func TestJWTManager_Verify(t *testing.T) {
	userID := 10
	log := logger.Init("")
	cfg := &server.Config{}

	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)

	user := &domain.User{
		ID: userID,
	}
	token, _ := jwt.Generate(user)
	resp, err := jwt.Verify(token)

	assert.Equal(t, nil, err)
	assert.Equal(t, userID, resp.ID)
}

func TestJWTManager_VerifyFail(t *testing.T) {
	log := logger.Init("")
	cfg := &server.Config{}
	jwt := NewJWTManager(cfg.SecretKeyToken, cfg.TokenDuration, log)
	_, err := jwt.Verify("token_test")
	assert.Error(t, err)
}
