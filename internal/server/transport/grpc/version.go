package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetVersion gets the current version of the data
func (h *Handler) GetVersion(ctx context.Context, r *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	version, err := h.Service.Version.GetVersion(ctx, int(r.UserId))
	if err != nil {
		h.Log.Err(err).Msg("handler getVersion error")
		return nil, status.Error(codes.Internal, "internal error")
	}
	response := pb.GetVersionResponse{Version: uint64(version)}
	return &response, nil
}
