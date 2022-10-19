package service

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/rs/zerolog"
)

// ICache api client local repository
type ICache interface {
	SetUserParams(userID int, token string)
	GetCurrentUserID() int
	GetToken() string
	GetVersion() int
	SetVersion(version int)
	IncVersion()

	Save(data any) error
	Update(data any) error
	Delete(data any) error
	GetAll(dt dataType.DataType) (any, error)
	UpdateAll(data any) error
}

// Service client manager service
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

// SetUserParams records local repository data for an authorized user
func (s *Service) SetUserParams(userID int, token string) error {
	if userID < 0 {
		s.Log.Debug().Msgf("service setUserParams: userID %d less zero", userID)
		return errors.New("internal error")
	}

	if token == "" {
		s.Log.Debug().Msg("service setUserParams: token is empty")
		return errors.New("internal error")
	}

	s.Storage.SetUserParams(userID, token)
	return nil
}

// GetCurrentUserID gets the current user id
func (s *Service) GetCurrentUserID() int {
	return s.Storage.GetCurrentUserID()
}

// GetToken gets a token for authorization
func (s *Service) GetToken() string {
	return s.Storage.GetToken()
}

// GetVersionCache gets the current version data local repository
func (s *Service) GetVersionCache() int {
	return s.Storage.GetVersion()
}

// SetVersionCache records the version data
func (s *Service) SetVersionCache(version int) {
	s.Storage.SetVersion(version)
}

// Save writes data
func (s *Service) Save(data any) error {
	return s.Storage.Save(data)
}

// Update updates data
func (s *Service) Update(data any) error {
	return s.Storage.Update(data)
}

// Delete deletes data
func (s *Service) Delete(data any) error {
	return s.Storage.Delete(data)
}

// GetAll gets data by data type
func (s *Service) GetAll(dt dataType.DataType) (any, error) {
	data, err := s.Storage.GetAll(dt)
	if err != nil {
		return nil, err
	}

	switch d := data.(type) {
	case []*domain.LoginPass:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.Text:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.Binary:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	case []*domain.BankCard:
		sort.Slice(d, func(i, j int) bool {
			return d[i].ID < d[j].ID
		})
		return d, nil
	default:
		err = fmt.Errorf("unsupported type '%v'", reflect.TypeOf(data))
		s.Log.Err(err).Msg("service getAll error")
		return nil, err
	}
}

// UpdateAll updates data
func (s *Service) UpdateAll(data any) error {
	return s.Storage.UpdateAll(data)
}
