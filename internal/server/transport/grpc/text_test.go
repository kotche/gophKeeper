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

func TestHandler_CreateText(t *testing.T) {
	userID := 10
	text := "sdfstertetertert"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Text{
		UserID:   userID,
		Text:     text,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, repo, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.TextRequest{
		UserId:   int64(userID),
		Text:     text,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.CreateText(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_UpdateText(t *testing.T) {
	id := 20
	userID := 10
	text := "sdfstertetertert"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Text{
		ID:       id,
		UserID:   userID,
		Text:     text,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, repo, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.TextUpdateRequest{
		Id:       int64(id),
		UserId:   int64(userID),
		Text:     text,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.UpdateText(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_DeleteText(t *testing.T) {
	id := 20
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Text{
		UserID: userID,
		ID:     id,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, repo, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.TextDeleteRequest{
		UserId: int64(userID),
		Id:     int64(id),
	}

	ctx := context.Background()
	_, actualError := handler.DeleteText(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_GetAllText(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := []domain.Text{
		{ID: 1, UserID: userID, Text: "sdfsterteterterte4434", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Text: "sdfstertetertertsdfsd555f"},
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockITextRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	dataService := service.NewDataService(nil, repo, nil, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.TextGetAllRequest{UserId: int64(userID)}

	ctx := context.Background()
	resp, actualError := handler.GetAllText(ctx, &dataReq)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.Text, 0, len(resp.Texts))
	for _, v := range resp.Texts {
		dataResp = append(dataResp, domain.Text{
			ID:       int(v.Id),
			UserID:   userID,
			Text:     v.Text,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}
