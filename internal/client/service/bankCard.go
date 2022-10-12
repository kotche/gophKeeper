package service

import (
	"sort"

	"github.com/kotche/gophKeeper/internal/client/domain"
)

func (s *Service) AddBankCard(data *domain.BankCard) error {
	return s.Storage.AddBankCard(data)
}

func (s *Service) UpdateBankCard(data *domain.BankCard) error {
	return s.Storage.UpdateBankCard(data)
}

func (s *Service) DeleteBankCard(id int) error {
	return s.Storage.DeleteBankCard(id)
}

func (s *Service) ReadAllBankCardCache() ([]*domain.BankCard, error) {
	data, err := s.Storage.ReadAllBankCard()
	if err != nil {
		return data, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	return data, err
}

func (s *Service) UpdateAllBankCardCache(data []*domain.BankCard) error {
	return s.Storage.UpdateAllBankCard(data)
}
