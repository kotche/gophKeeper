package grpc

import (
	"context"
	"fmt"

	"github.com/kotche/gophKeeper/internal/pb"
)

func (s *Sender) GetVersionCache() (int, error) {
	return s.Service.Storage.GetVersion()
}

func (s *Sender) GetVersionServer(ctx context.Context) (int, error) {
	portTCP := fmt.Sprintf(":%s", s.Conf.Port)
	conn, err := s.ClientConn.GetClientConn(portTCP, s.Log, s.getInterceptors())
	if err != nil {
		return 0, fmt.Errorf("server is not available: %s", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			s.Log.Err(err).Msg("getVersionServer conn close error")
		}
	}()

	userID, err := s.Service.GetCurrentUserID()
	if err != nil {
		return 0, err
	}

	c := pb.NewVersionServiceClient(conn)
	r := &pb.GetVersionRequest{UserId: int64(userID)}
	resp, err := c.GetVersion(ctx, r)
	if err != nil {
		return 0, err
	}

	return int(resp.Version), nil
}
