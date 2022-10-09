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
		return nil, status.Error(codes.Internal, "internal error")
	}

	response := pb.LoginPassResponse{Id: int64(loginPass.ID)}
	return &response, nil
}

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
