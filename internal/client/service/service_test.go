package service

import (
	"testing"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/client/storage"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestService_Save(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := srvc.Save(dataLp)
	errText := srvc.Save(dataText)
	errBinary := srvc.Save(dataBinary)
	errBankCard := srvc.Save(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestService_SaveFail(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	data := "some data"
	err := srvc.Save(data)

	assert.Error(t, err)
}

func TestService_Update(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := srvc.Update(dataLp)
	errText := srvc.Update(dataText)
	errBinary := srvc.Update(dataBinary)
	errBankCard := srvc.Update(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestService_UpdateFail(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	data := "some data"
	err := srvc.Update(data)

	assert.Error(t, err)
}

func TestService_UpdateAll(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	dataLp := []*domain.LoginPass{
		{ID: 1, Login: "login1", Password: "password1", MetaInfo: "meta"},
		{ID: 2, Login: "login2", Password: "password2", MetaInfo: "meta"},
	}
	dataText := []*domain.Text{
		{ID: 1, Text: "text1", MetaInfo: "meta"},
		{ID: 2, Text: "text2", MetaInfo: "meta"},
	}
	dataBinary := []*domain.Binary{
		{ID: 1, Binary: "binary1", MetaInfo: "meta"},
		{ID: 2, Binary: "binary2", MetaInfo: "meta"},
	}
	dataBankCard := []*domain.BankCard{
		{ID: 1, Number: "5555", MetaInfo: "meta"},
		{ID: 2, Number: "6666", MetaInfo: "meta"},
	}

	errLp := srvc.UpdateAll(dataLp)
	errText := srvc.UpdateAll(dataText)
	errBinary := srvc.UpdateAll(dataBinary)
	errBankCard := srvc.UpdateAll(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)

	dataLpResp, _ := srvc.GetAll(dataType.LP)
	dataTextResp, _ := srvc.GetAll(dataType.TEXT)
	dataBinaryResp, _ := srvc.GetAll(dataType.BINARY)
	dataBankCardResp, _ := srvc.GetAll(dataType.BANKCARD)

	assert.EqualValues(t, dataLp, dataLpResp)
	assert.EqualValues(t, dataText, dataTextResp)
	assert.EqualValues(t, dataBinary, dataBinaryResp)
	assert.EqualValues(t, dataBankCard, dataBankCardResp)
}

func TestService_UpdateAllFail(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	data := "some data"
	err := srvc.UpdateAll(data)

	assert.Error(t, err)
}

func TestService_Delete(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := srvc.Delete(dataLp)
	errText := srvc.Delete(dataText)
	errBinary := srvc.Delete(dataBinary)
	errBankCard := srvc.Delete(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestService_DeleteFail(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	data := "some data"
	err := srvc.Delete(data)

	assert.Error(t, err)
}

func TestService_GetAll(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	dataLp := []*domain.LoginPass{
		{ID: 1, Login: "login1", Password: "password1", MetaInfo: "meta"},
		{ID: 2, Login: "login2", Password: "password2", MetaInfo: "meta"},
	}
	dataText := []*domain.Text{
		{ID: 1, Text: "text1", MetaInfo: "meta"},
		{ID: 2, Text: "text2", MetaInfo: "meta"},
	}
	dataBinary := []*domain.Binary{
		{ID: 1, Binary: "binary1", MetaInfo: "meta"},
		{ID: 2, Binary: "binary2", MetaInfo: "meta"},
	}
	dataBankCard := []*domain.BankCard{
		{ID: 1, Number: "5555", MetaInfo: "meta"},
		{ID: 2, Number: "6666", MetaInfo: "meta"},
	}

	for _, v := range dataLp {
		cache.Save(v)
	}
	for _, v := range dataText {
		cache.Save(v)
	}
	for _, v := range dataBinary {
		cache.Save(v)
	}
	for _, v := range dataBankCard {
		cache.Save(v)
	}

	dataLpResp, errLp := srvc.GetAll(dataType.LP)
	dataTextResp, errText := srvc.GetAll(dataType.TEXT)
	dataBinaryResp, errBinary := srvc.GetAll(dataType.BINARY)
	dataBankCardResp, errBankCard := srvc.GetAll(dataType.BANKCARD)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)

	assert.EqualValues(t, dataLp, dataLpResp)
	assert.EqualValues(t, dataText, dataTextResp)
	assert.EqualValues(t, dataBinary, dataBinaryResp)
	assert.EqualValues(t, dataBankCard, dataBankCardResp)
}

func TestService_GetCurrentUserID(t *testing.T) {
	userID := 10
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)
	cache.SetUserParams(userID, "")
	userIDResp := srvc.GetCurrentUserID()

	assert.Equal(t, userID, userIDResp)
}

func TestService_GetToken(t *testing.T) {
	token := "sdsdsdsd3242342"
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)
	cache.SetUserParams(1, token)
	tokenResp := srvc.GetToken()

	assert.Equal(t, token, tokenResp)
}

func TestService_GetVersionCache(t *testing.T) {
	version := 10
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)
	cache.SetVersion(version)
	versionResp := srvc.GetVersionCache()

	assert.Equal(t, version, versionResp)
}

func TestService_SetUserParams(t *testing.T) {
	userID := 10
	token := "sdsdsdsd3242342"
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)
	err := srvc.SetUserParams(userID, token)
	assert.Equal(t, nil, err)

	tokenResp := cache.GetToken()
	userIDResp := cache.GetCurrentUserID()

	assert.Equal(t, token, tokenResp)
	assert.Equal(t, userID, userIDResp)
}

func TestService_SetUserParamsFail(t *testing.T) {
	userID := -10
	token := "345345"
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)

	err := srvc.SetUserParams(userID, token)
	assert.Error(t, err)

	userID = 10
	token = ""

	err = srvc.SetUserParams(userID, token)
	assert.Error(t, err)
}

func TestService_SetVersionCache(t *testing.T) {
	ver := 10

	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := NewService(cache, nil, log)
	srvc.SetVersionCache(ver)

	verResp := cache.GetVersion()
	assert.Equal(t, ver, verResp)
}
