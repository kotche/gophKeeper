package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateLoginPass creates a login password pair
func (h *Handler) CreateLoginPass(ctx context.Context, r *pb.LoginPassRequest) (*pb.LoginPassResponse, error) {
	loginPass := domain.LoginPass{
		UserID:   int(r.UserId),
		Login:    r.Username,
		Password: r.Password,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Create(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler createLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassResponse{Id: int64(loginPass.ID)}
	return &response, nil
}

// UpdateLoginPass updates a login password pair
func (h *Handler) UpdateLoginPass(ctx context.Context, r *pb.LoginPassUpdateRequest) (*pb.LoginPassUpdateResponse, error) {
	loginPass := domain.LoginPass{
		ID:       int(r.Id),
		UserID:   int(r.UserId),
		Login:    r.Username,
		Password: r.Password,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.Data.Update(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler updateLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassUpdateResponse{}
	return &response, nil
}

// DeleteLoginPass deletes a login password pair
func (h *Handler) DeleteLoginPass(ctx context.Context, r *pb.LoginPassDeleteRequest) (*pb.LoginPassDeleteResponse, error) {
	loginPass := domain.LoginPass{
		ID:     int(r.Id),
		UserID: int(r.UserId),
	}

	err := h.Service.Data.Delete(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler deleteLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassDeleteResponse{}
	return &response, nil
}

// GetAllLoginPass returns all login password pairs by user id
func (h *Handler) GetAllLoginPass(ctx context.Context, r *pb.LoginPassGetAllRequest) (*pb.LoginPassGetAllResponse, error) {
	userID := int(r.UserId)
	data, err := h.Service.Data.GetAll(ctx, userID, dataType.LP)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler getAllLoginPass error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.LoginPassGetAllResponse{}
	for _, v := range data.([]domain.LoginPass) {
		response.LoginPassPairs = append(response.LoginPassPairs, &pb.GetAllLoginPassResponse{
			Id:       int64(v.ID),
			Login:    v.Login,
			Password: v.Password,
			MetaInfo: v.MetaInfo,
		})
	}

	return &response, nil
}
