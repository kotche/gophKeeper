package grpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetVersion(t *testing.T) {
	userID := 10
	var ver uint
	ver = 2567

	log := logger.Init("")
	cfg := &server.Config{}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIVersionRepo(control)
	repo.EXPECT().GetVersion(gomock.Any(), userID).Return(ver, nil).Times(1)

	verService := service.NewVersionService(repo, log)
	srvc := service.NewService(nil, nil, verService)
	handler := NewHandler(srvc, log, cfg)

	req := &pb.GetVersionRequest{UserId: int64(userID)}

	ctx := context.Background()
	resp, actualError := handler.GetVersion(ctx, req)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, uint64(ver), resp.Version)
}
