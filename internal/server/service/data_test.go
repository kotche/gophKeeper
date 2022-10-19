package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestDataService_Create(t *testing.T) {
	dataLp := &domain.LoginPass{
		UserID:   2,
		Login:    "login",
		Password: "password",
		MetaInfo: "meta",
	}

	dataText := &domain.Text{
		UserID:   2,
		Text:     "text",
		MetaInfo: "meta",
	}

	dataBinary := &domain.Binary{
		UserID:   2,
		Binary:   "binary",
		MetaInfo: "meta",
	}
	dataBankCard := &domain.BankCard{
		UserID:   2,
		Number:   "555555",
		MetaInfo: "meta",
	}

	control := gomock.NewController(t)
	repoLp := mock_server.NewMockILoginPassRepo(control)
	repoText := mock_server.NewMockITextRepo(control)
	repoBinary := mock_server.NewMockIBinaryRepo(control)
	repoBankCard := mock_server.NewMockIBankCardRepo(control)

	repoLp.EXPECT().Create(gomock.Any(), dataLp).Return(nil).Times(1)
	repoText.EXPECT().Create(gomock.Any(), dataText).Return(nil).Times(1)
	repoBinary.EXPECT().Create(gomock.Any(), dataBinary).Return(nil).Times(1)
	repoBankCard.EXPECT().Create(gomock.Any(), dataBankCard).Return(nil).Times(1)

	log := logger.Init("")
	dataService := NewDataService(repoLp, repoText, repoBinary, repoBankCard, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	errLp := srvc.Data.Create(ctx, dataLp)
	errText := srvc.Data.Create(ctx, dataText)
	errBinary := srvc.Data.Create(ctx, dataBinary)
	errBankCard := srvc.Data.Create(ctx, dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestDataService_CreateFail(t *testing.T) {
	data := "some_data"

	log := logger.Init("")
	dataService := NewDataService(nil, nil, nil, nil, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	err := srvc.Data.Create(ctx, data)
	assert.Error(t, err)
}

func TestDataService_Update(t *testing.T) {
	dataLp := &domain.LoginPass{
		ID:       1,
		UserID:   2,
		Login:    "login",
		Password: "password",
		MetaInfo: "meta",
	}

	dataText := &domain.Text{
		ID:       1,
		UserID:   2,
		Text:     "text",
		MetaInfo: "meta",
	}

	dataBinary := &domain.Binary{
		ID:       1,
		UserID:   2,
		Binary:   "binary",
		MetaInfo: "meta",
	}
	dataBankCard := &domain.BankCard{
		ID:       1,
		UserID:   2,
		Number:   "555555",
		MetaInfo: "meta",
	}

	control := gomock.NewController(t)
	repoLp := mock_server.NewMockILoginPassRepo(control)
	repoText := mock_server.NewMockITextRepo(control)
	repoBinary := mock_server.NewMockIBinaryRepo(control)
	repoBankCard := mock_server.NewMockIBankCardRepo(control)

	repoLp.EXPECT().Update(gomock.Any(), dataLp).Return(nil).Times(1)
	repoText.EXPECT().Update(gomock.Any(), dataText).Return(nil).Times(1)
	repoBinary.EXPECT().Update(gomock.Any(), dataBinary).Return(nil).Times(1)
	repoBankCard.EXPECT().Update(gomock.Any(), dataBankCard).Return(nil).Times(1)

	log := logger.Init("")
	dataService := NewDataService(repoLp, repoText, repoBinary, repoBankCard, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	errLp := srvc.Data.Update(ctx, dataLp)
	errText := srvc.Data.Update(ctx, dataText)
	errBinary := srvc.Data.Update(ctx, dataBinary)
	errBankCard := srvc.Data.Update(ctx, dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestDataService_UpdateFail(t *testing.T) {
	data := "some_data"

	log := logger.Init("")
	dataService := NewDataService(nil, nil, nil, nil, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	err := srvc.Data.Update(ctx, data)
	assert.Error(t, err)
}

func TestDataService_Delete(t *testing.T) {
	dataLp := &domain.LoginPass{
		UserID:   2,
		Login:    "login",
		Password: "password",
		MetaInfo: "meta",
	}

	dataText := &domain.Text{
		UserID:   2,
		Text:     "text",
		MetaInfo: "meta",
	}

	dataBinary := &domain.Binary{
		UserID:   2,
		Binary:   "binary",
		MetaInfo: "meta",
	}
	dataBankCard := &domain.BankCard{
		UserID:   2,
		Number:   "555555",
		MetaInfo: "meta",
	}

	control := gomock.NewController(t)
	repoLp := mock_server.NewMockILoginPassRepo(control)
	repoText := mock_server.NewMockITextRepo(control)
	repoBinary := mock_server.NewMockIBinaryRepo(control)
	repoBankCard := mock_server.NewMockIBankCardRepo(control)

	repoLp.EXPECT().Delete(gomock.Any(), dataLp).Return(nil).Times(1)
	repoText.EXPECT().Delete(gomock.Any(), dataText).Return(nil).Times(1)
	repoBinary.EXPECT().Delete(gomock.Any(), dataBinary).Return(nil).Times(1)
	repoBankCard.EXPECT().Delete(gomock.Any(), dataBankCard).Return(nil).Times(1)

	log := logger.Init("")
	dataService := NewDataService(repoLp, repoText, repoBinary, repoBankCard, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	errLp := srvc.Data.Delete(ctx, dataLp)
	errText := srvc.Data.Delete(ctx, dataText)
	errBinary := srvc.Data.Delete(ctx, dataBinary)
	errBankCard := srvc.Data.Delete(ctx, dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestDataService_GetAll(t *testing.T) {
	userID := 2

	dataLp := []domain.LoginPass{
		{UserID: userID, Login: "login1", Password: "password1", MetaInfo: "meta"},
		{UserID: userID, Login: "login2", Password: "password2", MetaInfo: "meta"},
	}
	dataText := []domain.Text{
		{UserID: userID, Text: "text1", MetaInfo: "meta"},
		{UserID: userID, Text: "text2", MetaInfo: "meta"},
	}
	dataBinary := []domain.Binary{
		{UserID: userID, Binary: "binary1", MetaInfo: "meta"},
		{UserID: userID, Binary: "binary2", MetaInfo: "meta"},
	}
	dataBankCard := []domain.BankCard{
		{UserID: userID, Number: "5555", MetaInfo: "meta"},
		{UserID: userID, Number: "6666", MetaInfo: "meta"},
	}

	control := gomock.NewController(t)
	repoLp := mock_server.NewMockILoginPassRepo(control)
	repoText := mock_server.NewMockITextRepo(control)
	repoBinary := mock_server.NewMockIBinaryRepo(control)
	repoBankCard := mock_server.NewMockIBankCardRepo(control)

	repoLp.EXPECT().GetAll(gomock.Any(), userID).Return(dataLp, nil).Times(1)
	repoText.EXPECT().GetAll(gomock.Any(), userID).Return(dataText, nil).Times(1)
	repoBinary.EXPECT().GetAll(gomock.Any(), userID).Return(dataBinary, nil).Times(1)
	repoBankCard.EXPECT().GetAll(gomock.Any(), userID).Return(dataBankCard, nil).Times(1)

	log := logger.Init("")
	dataService := NewDataService(repoLp, repoText, repoBinary, repoBankCard, log)
	srvc := NewService(nil, dataService, nil)

	ctx := context.Background()
	dataLpResp, errLp := srvc.Data.GetAll(ctx, userID, dataType.LP)
	dataTextResp, errText := srvc.Data.GetAll(ctx, userID, dataType.TEXT)
	dataBinaryResp, errBinary := srvc.Data.GetAll(ctx, userID, dataType.BINARY)
	dataBankCardResp, errBankCard := srvc.Data.GetAll(ctx, userID, dataType.BANKCARD)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)

	assert.EqualValues(t, dataLp, dataLpResp)
	assert.EqualValues(t, dataText, dataTextResp)
	assert.EqualValues(t, dataBinary, dataBinaryResp)
	assert.EqualValues(t, dataBankCard, dataBankCardResp)
}
