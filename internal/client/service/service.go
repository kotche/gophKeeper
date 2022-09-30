package service

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/rs/zerolog"
)

type CacheContract interface {
}

type Service struct {
	storage CacheContract
	Conf    *client.Config
	Log     *zerolog.Logger
}

func NewService(conf *client.Config, log *zerolog.Logger) *Service {
	return &Service{Conf: conf, Log: log}
}
