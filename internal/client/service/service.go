package service

import (
	"errors"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/rs/zerolog"
)

type ICache interface {
	SetUserParams(userID int, token string) error
	GetCurrentUserID() (int, error)
}

type Service struct {
	Storage ICache
	Conf    *client.Config
	Log     *zerolog.Logger
}

func NewService(storage ICache, conf *client.Config, log *zerolog.Logger) *Service {
	return &Service{
		Storage: storage,
		Conf:    conf,
		Log:     log,
	}
}

func (s *Service) SetUserParams(userID int, token string) error {
	if userID < 0 {
		s.Log.Debug().Msgf("service setUserParams: userID %d less zero", userID)
		return errors.New("internal error")
	}

	if token == "" {
		s.Log.Debug().Msg("service setUserParams: token is empty")
		return errors.New("internal error")
	}

	if err := s.Storage.SetUserParams(userID, token); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCurrentUserID() (int, error) {
	return s.Storage.GetCurrentUserID()
}
