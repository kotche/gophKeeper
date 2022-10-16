package grpc

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kotche/gophKeeper/config/server"
	domainClient "github.com/kotche/gophKeeper/internal/client/domain"
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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type testParamsLoginPass struct {
	client pb.LoginPassServiceClient
	conn   *grpc.ClientConn
	lis    *bufconn.Listener
}

func NewTestParamsLoginPass(ctx context.Context, t *testing.T) *testParamsLoginPass {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsLoginPass{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewLoginPassServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.client = client

	return testParams
}

func (t *testParamsLoginPass) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService, dataService service_server.IData) {
	serverService := service_server.NewService(authService, dataService, nil)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsLoginPass) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func (t *testParamsLoginPass) GetCode(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.OK
}

func TestSender_CreateLoginPass(t *testing.T) {
	userID := 10
	username := "login"
	password := "password"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsLoginPass(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.LoginPass{
		UserID:   userID,
		Login:    username,
		Password: password,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(repo, nil, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.LoginPassRequest{UserId: int64(userID), Username: username, Password: password, MetaInfo: meta}
	_, actualError := testParams.client.CreateLoginPass(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_UpdateLoginPass(t *testing.T) {
	userID := 10
	ItemID := 55
	username := "login"
	password := "password"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsLoginPass(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.LoginPass{
		ID:       ItemID,
		UserID:   userID,
		Login:    username,
		Password: password,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(repo, nil, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.LoginPassUpdateRequest{Id: int64(ItemID), UserId: int64(userID), Username: username, Password: password, MetaInfo: meta}
	_, actualError := testParams.client.UpdateLoginPass(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_DeleteLoginPass(t *testing.T) {
	userID := 10
	ItemID := 55

	ctx := context.Background()

	testParams := NewTestParamsLoginPass(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.LoginPass{
		ID:     ItemID,
		UserID: userID,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(repo, nil, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.LoginPassDeleteRequest{Id: int64(ItemID), UserId: int64(userID)}
	_, actualError := testParams.client.DeleteLoginPass(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_GetAllLoginPass(t *testing.T) {
	userID := 10

	ctx := context.Background()

	testParams := NewTestParamsLoginPass(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := []domain.LoginPass{
		{ID: 1, UserID: userID, Login: "l1", Password: "p1", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Login: "l2", Password: "p2"},
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockILoginPassRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(repo, nil, nil, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.LoginPassGetAllRequest{UserId: int64(userID)}
	resp, actualError := testParams.client.GetAllLoginPass(ctx, req)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.LoginPass, 0, len(resp.LoginPassPairs))
	for _, v := range resp.LoginPassPairs {
		dataResp = append(dataResp, domain.LoginPass{
			ID:       int(v.Id),
			UserID:   userID,
			Login:    v.Login,
			Password: v.Password,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}

func TestSender_ReadLoginPassCache(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := service.NewService(cache, nil, log)
	sender := NewSender(srvc, nil, nil, log)

	data := []*domainClient.LoginPass{
		{ID: 1, Login: "l1", Password: "p1", MetaInfo: "777"},
		{ID: 2, Login: "l2", Password: "p2"},
		{ID: 3, Login: "l3", Password: "p3"},
	}
	for _, v := range data {
		cache.Save(v)
	}
	dataResp, actualError := sender.ReadLoginPassCache()
	assert.Equal(t, nil, actualError)
	assert.EqualValues(t, data, dataResp)
}
