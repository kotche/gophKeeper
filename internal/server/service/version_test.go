package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestVersionService_GetVersion(t *testing.T) {
	userID := 10
	var ver uint
	ver = 2567

	log := logger.Init("")

	control := gomock.NewController(t)
	repo := mock_server.NewMockIVersionRepo(control)
	repo.EXPECT().GetVersion(gomock.Any(), userID).Return(ver, nil).Times(1)

	verService := NewVersionService(repo, log)
	srvc := NewService(nil, nil, verService)

	ctx := context.Background()
	verResp, err := srvc.Version.GetVersion(ctx, userID)
	assert.Equal(t, nil, err)
	assert.Equal(t, ver, verResp)
}
