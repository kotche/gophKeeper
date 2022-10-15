package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/client/domain"
	"github.com/kotche/gophKeeper/internal/client/domain/dataType"
	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) CreateBinary(binary, meta string) (int, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BinaryRequest{UserId: int64(userID), Binary: binary, MetaInfo: meta}

	ctx := context.Background()
	resp, err := s.ClientConn.Binary.CreateBinary(ctx, r)
	if err != nil {
		return -1, err
	}

	s.Log.Debug().Msgf("type binary create, userID %d, id: %d", userID, resp.Id)

	data := &domain.Binary{
		ID:       int(resp.Id),
		Binary:   binary,
		MetaInfo: meta,
	}

	if err = s.Service.Save(data); err != nil {
		s.Log.Err(err).Msgf("createBinary add to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return data.ID, nil
}

func (s *Sender) UpdateBinary(id int, binary, meta string) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BinaryUpdateRequest{Id: int64(id), UserId: int64(userID), Binary: binary, MetaInfo: meta}

	ctx := context.Background()
	_, err := s.ClientConn.Binary.UpdateBinary(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("type lp update, userID %d, id: %d", userID, id)

	data := &domain.Binary{
		ID:       id,
		Binary:   binary,
		MetaInfo: meta,
	}

	if err = s.Service.Update(data); err != nil {
		s.Log.Err(err).Msgf("updateBinary update binary to cache '%+v' error: %w", data, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) DeleteBinary(id int) error {
	userID := s.Service.GetCurrentUserID()
	r := &pb.BinaryDeleteRequest{Id: int64(id), UserId: int64(userID)}

	ctx := context.Background()
	_, err := s.ClientConn.Binary.DeleteBinary(ctx, r)
	if err != nil {
		return err
	}

	s.Log.Debug().Msgf("binary delete, userID %d, id: %d", userID, id)

	data := &domain.Binary{
		ID: id,
	}

	if err = s.Service.Delete(data); err != nil {
		s.Log.Err(err).Msgf("deleteBinary delete binary to cache '%d' error: %w", id, err)
	}

	s.Service.Storage.IncVersion()

	return nil
}

func (s *Sender) ReadBinaryCache() ([]*domain.Binary, error) {
	data, err := s.Service.GetAll(dataType.BINARY)
	if err != nil {
		return nil, err
	}
	return data.([]*domain.Binary), nil
}

func (s *Sender) GetAllBinary(ctx context.Context) ([]*domain.Binary, error) {
	userID := s.Service.GetCurrentUserID()

	r := &pb.BinaryGetAllRequest{UserId: int64(userID)}

	resp, err := s.ClientConn.Binary.GetAllBinary(ctx, r)
	if err != nil {
		return nil, err
	}

	data := make([]*domain.Binary, 0, len(resp.Binary))
	for _, v := range resp.Binary {
		data = append(data, &domain.Binary{
			ID:       int(v.Id),
			Binary:   v.Binary,
			MetaInfo: v.MetaInfo,
		})
	}

	return data, nil
}
