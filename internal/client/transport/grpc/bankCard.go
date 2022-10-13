package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateBankCard(number, meta string) error {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("createBankCard conn close error")
		}
	}()

	c := pb.NewBankCardServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return err
	}
	r := &pb.BankCardRequest{UserId: int64(userID), Number: number, MetaInfo: meta}

	ctx := context.Background()
	resp, err := c.CreateBankCard(ctx, r)
	if err != nil {
		return err
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

	if err = s.Service.Storage.IncVersion(); err != nil {
		s.Log.Debug().Msgf("createBankCard inc version error: %w", err)
	}

	return nil
}

func (s *Sender) UpdateBankCard(id int, number, meta string) error {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("updateBankCard conn close error")
		}
	}()

	c := pb.NewBankCardServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return err
	}
	r := &pb.BankCardUpdateRequest{Id: int64(id), UserId: int64(userID), Number: number, MetaInfo: meta}

	ctx := context.Background()
	_, err = c.UpdateBankCard(ctx, r)
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

	if err = s.Service.Storage.IncVersion(); err != nil {
		s.Log.Debug().Msgf("updateBankCard inc version error: %w", err)
	}

	return nil
}

func (s *Sender) DeleteBankCard(id int) error {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("deleteBankCard conn close error")
		}
	}()

	c := pb.NewBankCardServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return err
	}
	r := &pb.BankCardDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err = c.DeleteBankCard(ctx, r)
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

	if err = s.Service.Storage.IncVersion(); err != nil {
		s.Log.Debug().Msgf("deleteBankCard inc version error: %w", err)
	}

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
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return nil, fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("getAllBankCard conn close error")
		}
	}()

	c := pb.NewBankCardServiceClient(conn)

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return nil, err
	}
	r := &pb.BankCardGetAllRequest{UserId: int64(userID)}

	resp, err := c.GetAllBankCard(ctx, r)
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
