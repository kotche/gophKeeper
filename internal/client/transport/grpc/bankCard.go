package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateBankCard(number, meta string) (int, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BankCardRequest{UserId: int64(userID), Number: number, MetaInfo: meta}

	ctx := context.Background()
	resp, err := s.ClientConn.BankCard.CreateBankCard(ctx, r)
	if err != nil {
		return -1, err
	}

	s.Log.Debug().Msgf("type bank card create, userID %d, id: %d", userID, resp.Id)

	data := &domain.BankCard{
		ID:       int(resp.Id),
		Number:   number,
		MetaInfo: meta,
	}

	if err = s.Service.Save(data); err != nil {
		s.Log.Err(err).Msgf("createBankCard add to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return data.ID, nil
}

func (s *Sender) UpdateBankCard(id int, number, meta string) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BankCardUpdateRequest{Id: int64(id), UserId: int64(userID), Number: number, MetaInfo: meta}

	ctx := context.Background()
	_, err := s.ClientConn.BankCard.UpdateBankCard(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp update, userID %d, id: %d", userID, id)

	data := &domain.BankCard{
		ID:       id,
		Number:   number,
		MetaInfo: meta,
	}

	if err = s.Service.Update(data); err != nil {
		s.Log.Err(err).Msgf("updateBankCard update bank card to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) DeleteBankCard(id int) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BankCardDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err := s.ClientConn.BankCard.DeleteBankCard(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("BankCard delete, userID %d, id: %d", userID, id)

	data := &domain.BankCard{
		ID: id,
	}

	if err = s.Service.Delete(data); err != nil {
		s.Log.Err(err).Msgf("deleteBankCard delete bank card to cache '%d' error: %w", id, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) ReadBankCardCache() ([]*domain.BankCard, error) {
	data, err := s.Service.GetAll(dataType.BANKCARD)
	if err != nil {
		return nil, err
	}
	return data.([]*domain.BankCard), nil
}

func (s *Sender) GetAllBankCard(ctx context.Context) ([]*domain.BankCard, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BankCardGetAllRequest{UserId: int64(userID)}

	resp, err := s.ClientConn.BankCard.GetAllBankCard(ctx, r)
	if err != nil {
		return nil, err
	}

	data := make([]*domain.BankCard, 0, len(resp.BankCard))
	for _, v := range resp.BankCard {
		data = append(data, &domain.BankCard{
			ID:       int(v.Id),
			Number:   v.Number,
			MetaInfo: v.MetaInfo,
		})
	}

	return data, nil
}
