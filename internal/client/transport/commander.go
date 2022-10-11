package transport

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/rs/zerolog"
)

type ISender interface {
	Login(username, password string) error
	Authentication(login, password string) error

	CreateLoginPass(login, password, meta string) error
	UpdateLoginPass(id int, login, password, meta string) error
	DeleteLoginPass(id int) error
	ReadLoginPassCache() ([]*domain.LoginPass, error)
}

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
