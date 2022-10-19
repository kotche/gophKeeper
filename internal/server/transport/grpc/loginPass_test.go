package grpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateLoginPass(t *testing.T) {
	userID := 10
	login := "login"
	password := "password"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.LoginPass{
		UserID:   userID,
		Login:    login,
		Password: password,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(repo, nil, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.LoginPassRequest{
		UserId:   int64(userID),
		Username: login,
		Password: password,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.CreateLoginPass(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_UpdateLoginPass(t *testing.T) {
	id := 20
	userID := 10
	login := "login"
	password := "password"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.LoginPass{
		ID:       id,
		UserID:   userID,
		Login:    login,
		Password: password,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(repo, nil, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.LoginPassUpdateRequest{
		Id:       int64(id),
		UserId:   int64(userID),
		Username: login,
		Password: password,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.UpdateLoginPass(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_DeleteLoginPass(t *testing.T) {
	id := 20
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.LoginPass{
		UserID: userID,
		ID:     id,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(repo, nil, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.LoginPassDeleteRequest{
		UserId: int64(userID),
		Id:     int64(id),
	}

	ctx := context.Background()
	_, actualError := handler.DeleteLoginPass(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_GetAllLoginPass(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := []domain.LoginPass{
		{ID: 1, UserID: userID, Login: "login1", Password: "password1", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Login: "login1", Password: "password1"},
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	dataService := service.NewDataService(repo, nil, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.LoginPassGetAllRequest{UserId: int64(userID)}

	ctx := context.Background()
	resp, actualError := handler.GetAllLoginPass(ctx, &dataReq)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.LoginPass, 0, len(resp.LoginPassPairs))
	for _, v := range resp.LoginPassPairs {
		dataResp = append(dataResp, domain.LoginPass{
			ID:       int(v.Id),
			UserID:   userID,
			Login:    v.Login,
			Password: v.Password,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}
