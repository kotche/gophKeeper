package grpc

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/kotche/gophKeeper/internal/client/storage"
	mock_server "github.com/kotche/gophKeeper/internal/mocks/server"
	"github.com/kotche/gophKeeper/internal/pb"
	"github.com/kotche/gophKeeper/internal/server/domain"
	grpcServer "github.com/kotche/gophKeeper/internal/server/server/grpc"
	service_server "github.com/kotche/gophKeeper/internal/server/service"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"github.com/kotche/gophKeeper/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
)

type testParamsVersion struct {
	client pb.VersionServiceClient
	conn   *grpc.ClientConn
	lis    *bufconn.Listener
}

func NewTestParamsVersion(ctx context.Context, t *testing.T) *testParamsVersion {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsVersion{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewVersionServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.client = client

	return testParams
}

func (t *testParamsVersion) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService, ver service_server.IVersion) {
	serverService := service_server.NewService(authService, nil, ver)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsVersion) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func TestSender_GetVersionCache(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := service.NewService(cache, nil, log)
	sender := NewSender(srvc, nil, nil, log)

	verTest := 5
	cache.SetVersion(verTest)
	verCache := sender.GetVersionCache()
	assert.Equal(t, verTest, verCache)
}

func TestSender_GetVersionServer(t *testing.T) {
	userID := 10
	var ver uint
	ver = 2567

	ctx := context.Background()
	control := gomock.NewController(t)
	repo := mock_server.NewMockIVersionRepo(control)
	repo.EXPECT().GetVersion(gomock.Any(), userID).Return(ver, nil).Times(1)

	log := logger.Init("")
	cfgServer := &server.Config{}
	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	user := &domain.User{
		ID: userID,
	}

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	verService := service_server.NewVersionService(repo, log)
	testParams := NewTestParamsVersion(ctx, t)
	testParams.initServer(cfgServer, log, authService, verService)

	req := &pb.GetVersionRequest{UserId: int64(userID)}
	resp, actualError := testParams.client.GetVersion(ctx, req)
	assert.Equal(t, nil, actualError)
	assert.Equal(t, uint64(ver), resp.Version)
}
