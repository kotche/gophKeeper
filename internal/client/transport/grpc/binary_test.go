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

type testParamsBinary struct {
	client pb.BinaryServiceClient
	conn   *grpc.ClientConn
	lis    *bufconn.Listener
}

func NewTestParamsBinary(ctx context.Context, t *testing.T) *testParamsBinary {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	testParams := &testParamsBinary{}
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(testParams.bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewBinaryServiceClient(conn)

	testParams.conn = conn
	testParams.lis = lis
	testParams.client = client

	return testParams
}

func (t *testParamsBinary) initServer(cfgServer *server.Config, log *zerolog.Logger, authService service_server.IAuthService, dataService service_server.IData) {
	serverService := service_server.NewService(authService, dataService, nil)
	handler := grpcHandler.NewHandler(serverService, log, cfgServer)
	grpcSrv := grpcServer.NewServer(cfgServer, handler)

	go func() {
		if err := grpcSrv.Run(t.lis); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("gRCP test server run error")
		}
	}()
}

func (t *testParamsBinary) bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return t.lis.Dial()
}

func (t *testParamsBinary) GetCode(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.OK
}

func TestSender_CreateBinary(t *testing.T) {
	userID := 10
	binary := "dgdgdfgdfg435345345"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsBinary(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Binary{
		UserID:   userID,
		Binary:   binary,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Create(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, repo, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BinaryRequest{UserId: int64(userID), Binary: binary, MetaInfo: meta}
	_, actualError := testParams.client.CreateBinary(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_UpdateBinary(t *testing.T) {
	userID := 10
	ItemID := 55
	binary := "dgdgdfgdfg435345345"
	meta := "meta info"

	ctx := context.Background()

	testParams := NewTestParamsBinary(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Binary{
		ID:       ItemID,
		UserID:   userID,
		Binary:   binary,
		MetaInfo: meta,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Update(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, repo, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BinaryUpdateRequest{Id: int64(ItemID), UserId: int64(userID), Binary: binary, MetaInfo: meta}
	_, actualError := testParams.client.UpdateBinary(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_DeleteBinary(t *testing.T) {
	userID := 10
	ItemID := 55

	ctx := context.Background()

	testParams := NewTestParamsBinary(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := &domain.Binary{
		ID:     ItemID,
		UserID: userID,
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().Delete(gomock.Any(), data).Return(nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, repo, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BinaryDeleteRequest{Id: int64(ItemID), UserId: int64(userID)}
	_, actualError := testParams.client.DeleteBinary(ctx, req)
	assert.Equal(t, nil, actualError)
}

func TestSender_GetAllBinary(t *testing.T) {
	userID := 10

	ctx := context.Background()

	testParams := NewTestParamsBinary(ctx, t)
	defer testParams.conn.Close()

	log := logger.Init("")
	cfgServer := &server.Config{
		Security: server.Security{
			SecretKeyPassword: "1",
		},
	}

	data := []domain.Binary{
		{ID: 1, UserID: userID, Binary: "5555dfgdfgdfg", MetaInfo: "meta"},
		{ID: 2, UserID: userID, Binary: "6666dfgdgdfgdfgdgdfg"},
	}

	user := &domain.User{
		ID: userID,
	}

	control := gomock.NewController(t)
	repo := mock_server.NewMockIBinaryRepo(control)
	repo.EXPECT().GetAll(gomock.Any(), userID).Return(data, nil).Times(1)

	jwt := service_server.NewJWTManager(cfgServer.SecretKeyToken, cfgServer.TokenDuration, log)
	pe := service_server.NewPasswordEncryptor(cfgServer.SecretKeyPassword)
	authService := service_server.NewAuthService(nil, jwt, pe, log)

	token, _ := authService.GenerateToken(user)
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	dataService := service_server.NewDataService(nil, nil, repo, nil, log)
	testParams.initServer(cfgServer, log, authService, dataService)

	req := &pb.BinaryGetAllRequest{UserId: int64(userID)}
	resp, actualError := testParams.client.GetAllBinary(ctx, req)
	assert.Equal(t, nil, actualError)

	dataResp := make([]domain.Binary, 0, len(resp.Binaries))
	for _, v := range resp.Binaries {
		dataResp = append(dataResp, domain.Binary{
			ID:       int(v.Id),
			UserID:   userID,
			Binary:   v.Binary,
			MetaInfo: v.MetaInfo,
		})
	}

	assert.EqualValues(t, data, dataResp)
}

func TestSender_ReadBinaryCache(t *testing.T) {
	log := logger.Init("")
	cache := storage.NewCache(log)
	srvc := service.NewService(cache, nil, log)
	sender := NewSender(srvc, nil, nil, log)

	data := []*domainClient.Binary{
		{ID: 1, Binary: "555dfgdf", MetaInfo: "777"},
		{ID: 2, Binary: "666dfgdf"},
		{ID: 3, Binary: "888dfgdfgdfg"},
	}
	for _, v := range data {
		cache.Save(v)
	}
	dataResp, actualError := sender.ReadBinaryCache()
	assert.Equal(t, nil, actualError)
	assert.EqualValues(t, data, dataResp)
}
