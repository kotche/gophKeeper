package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateText creates a text data
func (h *Handler) CreateText(ctx context.Context, r *pb.TextRequest) (*pb.TextResponse, error) {
	Text := domain.Text{
		UserID:   int(r.UserId),
		Text:     r.Text,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Create(ctx, &Text)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler createText error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.TextResponse{Id: int64(Text.ID)}
	return &response, nil
}

// UpdateText updates a text data
func (h *Handler) UpdateText(ctx context.Context, r *pb.TextUpdateRequest) (*pb.TextUpdateResponse, error) {
	Text := domain.Text{
		ID:       int(r.Id),
		UserID:   int(r.UserId),
		Text:     r.Text,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Update(ctx, &Text)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler updateText error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.TextUpdateResponse{}
	return &response, nil
}

// DeleteText deletes a text data
func (h *Handler) DeleteText(ctx context.Context, r *pb.TextDeleteRequest) (*pb.TextDeleteResponse, error) {
	Text := domain.Text{
		ID:     int(r.Id),
		UserID: int(r.UserId),
	}

	err := h.Service.Data.Delete(ctx, &Text)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler deleteText error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.TextDeleteResponse{}
	return &response, nil
}

// GetAllText gets all text data by user id
func (h *Handler) GetAllText(ctx context.Context, r *pb.TextGetAllRequest) (*pb.TextGetAllResponse, error) {
	userID := int(r.UserId)
	data, err := h.Service.Data.GetAll(ctx, userID, dataType.TEXT)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler getAllText error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.TextGetAllResponse{}
	for _, v := range data.([]domain.Text) {
		response.Text = append(response.Text, &pb.GetAllTextResponse{
			Id:       int64(v.ID),
			Text:     v.Text,
			MetaInfo: v.MetaInfo,
		})
	}

	return &response, nil
}
