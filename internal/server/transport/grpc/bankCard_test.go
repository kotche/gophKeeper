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

func TestHandler_CreateBankCard(t *testing.T) {
	userID := 10
	number := "555555"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.BankCard{
		UserID:   userID,
		Number:   number,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, nil, repo, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BankCardRequest{
		UserId:   int64(userID),
		Number:   number,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.CreateBankCard(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_UpdateBankCard(t *testing.T) {
	id := 20
	userID := 10
	number := "555555"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.BankCard{
		ID:       id,
		UserID:   userID,
		Number:   number,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, nil, repo, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BankCardUpdateRequest{
		Id:       int64(id),
		UserId:   int64(userID),
		Number:   number,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.UpdateBankCard(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_DeleteBankCard(t *testing.T) {
	id := 20
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.BankCard{
		UserID: userID,
		ID:     id,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, nil, repo, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BankCardDeleteRequest{
		UserId: int64(userID),
		Id:     int64(id),
	}

	ctx := context.Background()
	_, actualError := handler.DeleteBankCard(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_GetAllBankCard(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := []domain.BankCard{
		{ID: 1, UserID: userID, Number: "5555", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Number: "6666"},
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBankCardRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	dataService := service.NewDataService(nil, nil, nil, repo, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BankCardGetAllRequest{UserId: int64(userID)}

	ctx := context.Background()
	resp, actualError := handler.GetAllBankCard(ctx, &dataReq)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.BankCard, 0, len(resp.BankCards))
	for _, v := range resp.BankCards {
		dataResp = append(dataResp, domain.BankCard{
			ID:       int(v.Id),
			UserID:   userID,
			Number:   v.Number,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}
