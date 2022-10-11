package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateLoginPass creates a login password pair
func (h *Handler) CreateLoginPass(ctx context.Context, r *pb.LoginPassRequest) (*pb.LoginPassResponse, error) {
	loginPass := dataType.LoginPass{
		UserID:   int(r.UserId),
		Login:    r.Username,
		Password: r.Password,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.LoginPass.Create(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler createLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassResponse{Id: int64(loginPass.ID)}
	return &response, nil
}

// UpdateLoginPass updates a login password pair
func (h *Handler) UpdateLoginPass(ctx context.Context, r *pb.LoginPassUpdateRequest) (*pb.LoginPassUpdateResponse, error) {
	loginPass := dataType.LoginPass{
		ID:       int(r.Id),
		UserID:   int(r.UserId),
		Login:    r.Username,
		Password: r.Password,
		MetaInfo: r.MetaInfo,
	}

	err := h.Service.LoginPass.Update(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler updateLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassUpdateResponse{}
	return &response, nil
}

// DeleteLoginPass deletes a login password pair
func (h *Handler) DeleteLoginPass(ctx context.Context, r *pb.LoginPassDeleteRequest) (*pb.LoginPassDeleteResponse, error) {
	loginPass := dataType.LoginPass{
		ID:     int(r.Id),
		UserID: int(r.UserId),
	}

	err := h.Service.LoginPass.Delete(ctx, &loginPass)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler deleteLoginPass error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.LoginPassDeleteResponse{}
	return &response, nil
}

// GetAllLoginPass returns all login password pairs
func (h *Handler) GetAllLoginPass(ctx context.Context, r *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	userID := int(r.UserId)
	lpPairs, err := h.Service.LoginPass.GetAll(ctx, userID)
	if err != nil {
		h.Log.Error().Err(err).Msg("handler getAllLoginPass error")
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.GetAllResponse{}
	for _, lp := range lpPairs {
		response.LoginPassPairs = append(response.LoginPassPairs, &pb.GetAllLoginPassResponse{
			Id:       int64(lp.ID),
			Login:    lp.Login,
			Password: lp.Password,
			MetaInfo: lp.MetaInfo,
		})
	}

	return &response, nil
}
