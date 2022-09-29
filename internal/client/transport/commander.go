package transport

import (
	"github.com/kotche/gophKeeper/internal/client/config"
	"github.com/rs/zerolog"
)

const (
	registration     = "reg"
	registrationDesc = "new user registration format: reg login password"
	authentication   = "auth"
	exit             = "exit"
)

type ISender interface {
	Registration(login, password string) error
}

type Commander struct {
	Sender ISender
	Conf   *config.Config
	Log    *zerolog.Logger
}

func NewCommander(sender ISender, conf *config.Config, log *zerolog.Logger) *Commander {
	return &Commander{
		Sender: sender,
		Conf:   conf,
		Log:    log,
	}
}
