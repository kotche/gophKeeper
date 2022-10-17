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

func TestHandler_CreateBinary(t *testing.T) {
	userID := 10
	binary := "sdfstertetertert"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Binary{
		UserID:   userID,
		Binary:   binary,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, repo, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BinaryRequest{
		UserId:   int64(userID),
		Binary:   binary,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.CreateBinary(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_UpdateBinary(t *testing.T) {
	id := 20
	userID := 10
	binary := "sdfstertetertert"
	meta := "meta info"

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Binary{
		ID:       id,
		UserID:   userID,
		Binary:   binary,
		MetaInfo: meta,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, repo, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BinaryUpdateRequest{
		Id:       int64(id),
		UserId:   int64(userID),
		Binary:   binary,
		MetaInfo: meta,
	}

	ctx := context.Background()
	_, actualError := handler.UpdateBinary(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_DeleteBinary(t *testing.T) {
	id := 20
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := &domain.Binary{
		UserID: userID,
		ID:     id,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	dataService := service.NewDataService(nil, nil, repo, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BinaryDeleteRequest{
		UserId: int64(userID),
		Id:     int64(id),
	}

	ctx := context.Background()
	_, actualError := handler.DeleteBinary(ctx, &dataReq)
	assert.Equal(t, nil, actualError)
}

func TestHandler_GetAllBinary(t *testing.T) {
	userID := 10

	log := logger.Init("")
	cfg := &server.Config{}

	data := []domain.Binary{
		{ID: 1, UserID: userID, Binary: "sdfsterteterterte4434", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Binary: "sdfstertetertertsdfsd555f"},
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	dataService := service.NewDataService(nil, nil, repo, nil, log)
	srvc := service.NewService(nil, dataService, nil)
	handler := NewHandler(srvc, log, cfg)

	dataReq := pb.BinaryGetAllRequest{UserId: int64(userID)}

	ctx := context.Background()
	resp, actualError := handler.GetAllBinary(ctx, &dataReq)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.Binary, 0, len(resp.Binaries))
	for _, v := range resp.Binaries {
		dataResp = append(dataResp, domain.Binary{
			ID:       int(v.Id),
			UserID:   userID,
			Binary:   v.Binary,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}
