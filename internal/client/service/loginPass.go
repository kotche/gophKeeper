package service

import (
	"sort"

	"github.com/kotche/gophKeeper/internal/client/domain"
)

func (s *Service) AddLoginPassword(data *domain.LoginPass) error {
	return s.Storage.AddLoginPassword(data)
}

func (s *Service) UpdateLoginPassword(data *domain.LoginPass) error {
	return s.Storage.UpdateLoginPassword(data)
}

func (s *Service) DeleteLoginPassword(id int) error {
	return s.Storage.DeleteLoginPassword(id)
}

func (s *Service) ReadAllLoginPasswordCache() ([]*domain.LoginPass, error) {
	lpPairs, err := s.Storage.ReadAllLoginPassword()
	if err != nil {
		return lpPairs, err
	}

	sort.Slice(lpPairs, func(i, j int) bool {
		return lpPairs[i].ID < lpPairs[j].ID
	})

	return lpPairs, err
}

func (s *Service) UpdateAllLoginPassCache(lpPairs []*domain.LoginPass) error {
	return s.Storage.UpdateAllLoginPass(lpPairs)
}
