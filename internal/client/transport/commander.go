package transport

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/rs/zerolog"
)

// ISender api transport client
type ISender interface {
	Login(username, password string) error
	Authentication(login, password string) error

	CreateLoginPass(login, password, meta string) (int, error)
	UpdateLoginPass(id int, login, password, meta string) error
	DeleteLoginPass(id int) error
	ReadLoginPassCache() ([]*domain.LoginPass, error)

	CreateText(text, meta string) (int, error)
	UpdateText(id int, text, meta string) error
	DeleteText(id int) error
	ReadTextCache() ([]*domain.Text, error)

	CreateBinary(binary, meta string) (int, error)
	UpdateBinary(id int, binary, meta string) error
	DeleteBinary(id int) error
	ReadBinaryCache() ([]*domain.Binary, error)

	CreateBankCard(number, meta string) (int, error)
	UpdateBankCard(id int, number, meta string) error
	DeleteBankCard(id int) error
	ReadBankCardCache() ([]*domain.BankCard, error)
}

// Commander manages commands
type Commander struct {
	Sender ISender
	Conf   *client.Config
	Log    *zerolog.Logger
}

func NewCommander(sender ISender, conf *client.Config, log *zerolog.Logger) *Commander {
	return &Commander{
		Sender: sender,
		Conf:   conf,
		Log:    log,
	}
}
