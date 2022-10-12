package service

import (
	"sort"

	"github.com/kotche/gophKeeper/internal/client/domain"
)

func (s *Service) AddBinary(data *domain.Binary) error {
	return s.Storage.AddBinary(data)
}

func (s *Service) UpdateBinary(data *domain.Binary) error {
	return s.Storage.UpdateBinary(data)
}

func (s *Service) DeleteBinary(id int) error {
	return s.Storage.DeleteBinary(id)
}

func (s *Service) ReadAllBinaryCache() ([]*domain.Binary, error) {
	data, err := s.Storage.ReadAllBinary()
	if err != nil {
		return data, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	return data, err
}

func (s *Service) UpdateAllBinaryCache(data []*domain.Binary) error {
	return s.Storage.UpdateAllBinary(data)
}
