package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBankCard creates a bank card data
func (h *Handler) CreateBankCard(ctx context.Context, r *pb.BankCardRequest) (*pb.BankCardResponse, error) {
	BankCard := domain.BankCard{
		UserID:   int(r.UserId),
		Number:   r.Number,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Create(ctx, &BankCard)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler createBankCard error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BankCardResponse{Id: int64(BankCard.ID)}
	return &response, nil
}

// UpdateBankCard updates a bank card data
func (h *Handler) UpdateBankCard(ctx context.Context, r *pb.BankCardUpdateRequest) (*pb.BankCardUpdateResponse, error) {
	BankCard := domain.BankCard{
		ID:       int(r.Id),
		UserID:   int(r.UserId),
		Number:   r.Number,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Update(ctx, &BankCard)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler updateBankCard error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BankCardUpdateResponse{}
	return &response, nil
}

// DeleteBankCard deletes a bank card data
func (h *Handler) DeleteBankCard(ctx context.Context, r *pb.BankCardDeleteRequest) (*pb.BankCardDeleteResponse, error) {
	BankCard := domain.BankCard{
		ID:     int(r.Id),
		UserID: int(r.UserId),
	}

	err := h.Service.Data.Delete(ctx, &BankCard)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler deleteBankCard error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BankCardDeleteResponse{}
	return &response, nil
}

// GetAllBankCard gets all bank card by user id
func (h *Handler) GetAllBankCard(ctx context.Context, r *pb.BankCardGetAllRequest) (*pb.BankCardGetAllResponse, error) {
	userID := int(r.UserId)
	data, err := h.Service.Data.GetAll(ctx, userID, dataType.BANKCARD)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler getAllBankCard error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.BankCardGetAllResponse{}
	for _, v := range data.([]domain.BankCard) {
		response.BankCard = append(response.BankCard, &pb.GetAllBankCardResponse{
			Id:       int64(v.ID),
			Number:   v.Number,
			MetaInfo: v.MetaInfo,
		})
	}

	return &response, nil
}
