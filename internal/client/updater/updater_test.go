package updater

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/mocks/client/updater"
	"github.com/kotche/gophKeeper/logger"
	"github.com/stretchr/testify/assert"
)

func TestUpdater_updateData(t *testing.T) {
	lp := []*domain.LoginPass{
		{ID: 1, Login: "login1", Password: "password1", MetaInfo: "meta"},
		{ID: 1, Login: "login1", Password: "password2", MetaInfo: "meta"},
	}
	text := []*domain.Text{
		{ID: 1, Text: "text1", MetaInfo: "meta"},
		{ID: 2, Text: "text2"},
	}
	binary := []*domain.Binary{
		{ID: 1, Binary: "binary1", MetaInfo: "meta"},
		{ID: 2, Binary: "binary2"},
	}
	bank := []*domain.BankCard{
		{ID: 1, Number: "5555", MetaInfo: "meta"},
		{ID: 2, Number: "6666"},
	}

	control := gomock.NewController(t)
	sender := mock_updater.NewMockISender(control)
	sender.EXPECT().GetAllLoginPass(gomock.Any()).Return(lp, nil).Times(1)
	sender.EXPECT().GetAllText(gomock.Any()).Return(text, nil).Times(1)
	sender.EXPECT().GetAllBinary(gomock.Any()).Return(binary, nil).Times(1)
	sender.EXPECT().GetAllBankCard(gomock.Any()).Return(bank, nil).Times(1)

	srvc := mock_updater.NewMockIService(control)
	srvc.EXPECT().UpdateAll(gomock.Any()).Return(nil).Times(4)

	updater := NewUpdater(sender, srvc, nil, logger.Init(""))
	ctx := context.Background()
	err := updater.updateData(ctx)
	assert.Equal(t, nil, err)
}
