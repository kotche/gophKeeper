package grpc

import (
	"context"

	"github.com/kotche/gophKeeper/internal/pb"
)

// GetVersionCache gets the data version of the local repository
func (s *Sender) GetVersionCache() int {
	return s.Service.Storage.GetVersion()
}

// GetVersionServer gets the data version of the db server
func (s *Sender) GetVersionServer(ctx context.Context) (int, error) {
	userID := s.Service.GetCurrentUserID()
	r := &pb.GetVersionRequest{UserId: int64(userID)}
	resp, err := s.ClientConn.Version.GetVersion(ctx, r)
	if err != nil {
		return 0, err
	}
	return int(resp.Version), nil
}
