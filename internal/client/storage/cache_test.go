package storage

import (
	"sync/atomic"
	"testing"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestCache_Save(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := cache.Save(dataLp)
	errText := cache.Save(dataText)
	errBinary := cache.Save(dataBinary)
	errBankCard := cache.Save(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestCache_SaveFail(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)
	data := "some data"
	err := cache.Save(data)
	assert.Error(t, err)
}

func TestCache_Update(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := cache.Update(dataLp)
	errText := cache.Update(dataText)
	errBinary := cache.Update(dataBinary)
	errBankCard := cache.Update(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestCache_UpdateFail(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)
	data := "some data"
	err := cache.Update(data)
	assert.Error(t, err)
}

func TestCache_UpdateAll(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)

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

	errLp := cache.UpdateAll(dataLp)
	errText := cache.UpdateAll(dataText)
	errBinary := cache.UpdateAll(dataBinary)
	errBankCard := cache.UpdateAll(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)

	dataLpResp, _ := cache.GetAll(dataType.LP)
	dataTextResp, _ := cache.GetAll(dataType.TEXT)
	dataBinaryResp, _ := cache.GetAll(dataType.BINARY)
	dataBankCardResp, _ := cache.GetAll(dataType.BANKCARD)

	assert.EqualValues(t, dataLp, dataLpResp)
	assert.EqualValues(t, dataText, dataTextResp)
	assert.EqualValues(t, dataBinary, dataBinaryResp)
	assert.EqualValues(t, dataBankCard, dataBankCardResp)
}

func TestCache_UpdateAllFail(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)
	data := "some data"
	err := cache.UpdateAll(data)
	assert.Error(t, err)
}

func TestCache_Delete(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)

	dataLp := &domain.LoginPass{ID: 1, Login: "login", Password: "password", MetaInfo: "777"}
	dataText := &domain.Text{ID: 1, Text: "555dfgdf", MetaInfo: "777"}
	dataBinary := &domain.Binary{ID: 1, Binary: "555dfgdf", MetaInfo: "777"}
	dataBankCard := &domain.BankCard{ID: 1, Number: "5555", MetaInfo: "777"}

	errLp := cache.Delete(dataLp)
	errText := cache.Delete(dataText)
	errBinary := cache.Delete(dataBinary)
	errBankCard := cache.Delete(dataBankCard)

	assert.Equal(t, nil, errLp)
	assert.Equal(t, nil, errText)
	assert.Equal(t, nil, errBinary)
	assert.Equal(t, nil, errBankCard)
}

func TestCache_DeleteFail(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)
	data := "some data"
	err := cache.Delete(data)
	assert.Error(t, err)
}

func TestCache_GetAll(t *testing.T) {
	log := logger.Init("")
	cache := NewCache(log)

	tests := []struct {
		name string
		dt   dataType.DataType
		data any
	}{
		{
			name: "loginPass",
			dt:   dataType.LP,
			data: []*domain.LoginPass{
				{ID: 1, Login: "login1", Password: "password1", MetaInfo: "meta"},
				{ID: 2, Login: "login2", Password: "password2", MetaInfo: "meta"},
			},
		},
		{
			name: "text",
			dt:   dataType.TEXT,
			data: []*domain.Text{
				{ID: 1, Text: "text1", MetaInfo: "meta"},
				{ID: 2, Text: "text2", MetaInfo: "meta"},
			},
		},
		{
			name: "binary",
			dt:   dataType.BINARY,
			data: []*domain.Binary{
				{ID: 1, Binary: "binary1", MetaInfo: "meta"},
				{ID: 2, Binary: "binary2", MetaInfo: "meta"},
			},
		},
		{
			name: "bankCard",
			dt:   dataType.BANKCARD,
			data: []*domain.BankCard{
				{ID: 1, Number: "5555", MetaInfo: "meta"},
				{ID: 2, Number: "6666", MetaInfo: "meta"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_ = cache.UpdateAll(tt.data)
			_, err := cache.GetAll(tt.dt)
			assert.Equal(t, nil, err)
		})
	}
}

func TestCache_GetCurrentUserID(t *testing.T) {
	userID := 10
	log := logger.Init("")
	cache := NewCache(log)
	cache.SetUserParams(userID, "")
	userIDResp := cache.GetCurrentUserID()
	assert.Equal(t, userID, userIDResp)
}

func TestCache_GetToken(t *testing.T) {
	token := "sdsdsdsd3242342"
	log := logger.Init("")
	cache := NewCache(log)
	cache.SetUserParams(1, token)
	tokenResp := cache.GetToken()
	assert.Equal(t, token, tokenResp)
}

func TestCache_GetVersion(t *testing.T) {
	version := 10
	log := logger.Init("")
	cache := NewCache(log)
	cache.SetVersion(version)
	versionResp := cache.GetVersion()
	assert.Equal(t, version, versionResp)
}

func TestCache_IncVersion(t *testing.T) {
	ver := 10
	var verCheck atomic.Uint64
	verCheck.Swap(uint64(ver))
	verCheck.Add(1)

	cache := NewCache(nil)
	cache.SetVersion(ver)
	cache.IncVersion()

	assert.Equal(t, verCheck, cache.version)
}

func TestCache_SetUserParams(t *testing.T) {
	userID := 10
	token := "sdsdsdsd3242342"
	log := logger.Init("")
	cache := NewCache(log)
	cache.SetUserParams(userID, token)

	tokenResp := cache.GetToken()
	userIDResp := cache.GetCurrentUserID()
	assert.Equal(t, token, tokenResp)
	assert.Equal(t, userID, userIDResp)
}

func TestCache_SetUserParamsFail(t *testing.T) {
	userID := -10
	token := "345345"
	log := logger.Init("")
	cache := NewCache(log)

	cache.SetUserParams(userID, token)
	assert.Equal(t, userID, cache.userID)
	assert.Equal(t, token, cache.token)
}

func TestCache_SetVersion(t *testing.T) {
	ver := 10
	var verCheck atomic.Uint64
	verCheck.Swap(uint64(ver))

	log := logger.Init("")
	cache := NewCache(log)
	cache.SetVersion(ver)
	assert.Equal(t, verCheck, cache.version)
}
