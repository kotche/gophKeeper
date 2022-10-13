package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBinary creates a binary data
func (h *Handler) CreateBinary(ctx context.Context, r *pb.BinaryRequest) (*pb.BinaryResponse, error) {
	Binary := domain.Binary{
		UserID:   int(r.UserId),
		Binary:   r.Binary,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Create(ctx, &Binary)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler createBinary error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BinaryResponse{Id: int64(Binary.ID)}
	return &response, nil
}

// UpdateBinary updates a binary data
func (h *Handler) UpdateBinary(ctx context.Context, r *pb.BinaryUpdateRequest) (*pb.BinaryUpdateResponse, error) {
	Binary := domain.Binary{
		ID:       int(r.Id),
		UserID:   int(r.UserId),
		Binary:   r.Binary,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Update(ctx, &Binary)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler updateBinary error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BinaryUpdateResponse{}
	return &response, nil
}

// DeleteBinary deletes a binary data
func (h *Handler) DeleteBinary(ctx context.Context, r *pb.BinaryDeleteRequest) (*pb.BinaryDeleteResponse, error) {
	Binary := domain.Binary{
		ID:     int(r.Id),
		UserID: int(r.UserId),
	}

	err := h.Service.Data.Delete(ctx, &Binary)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler deleteBinary error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.BinaryDeleteResponse{}
	return &response, nil
}

// GetAllBinary returns all binary data by user id
func (h *Handler) GetAllBinary(ctx context.Context, r *pb.BinaryGetAllRequest) (*pb.BinaryGetAllResponse, error) {
	userID := int(r.UserId)
	data, err := h.Service.Data.GetAll(ctx, userID, dataType.BINARY)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler getAllBinary error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.BinaryGetAllResponse{}
	for _, v := range data.([]domain.Binary) {
		response.Binary = append(response.Binary, &pb.GetAllBinaryResponse{
			Id:       int64(v.ID),
			Binary:   v.Binary,
			MetaInfo: v.MetaInfo,
		})
	}

	return &response, nil
}
