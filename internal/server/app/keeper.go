package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kotche/gophKeeper/config/server"
	grpcServer "github.com/kotche/gophKeeper/internal/server/server/grpc"
	"github.com/kotche/gophKeeper/internal/server/service"
	"github.com/kotche/gophKeeper/internal/server/storage"
	"github.com/kotche/gophKeeper/internal/server/storage/postgres"
	grpcHandler "github.com/kotche/gophKeeper/internal/server/transport/grpc"
	"github.com/rs/zerolog"
)

type Keeper struct {
	Log *zerolog.Logger
	Cfg *server.Config
}

func NewKeeper(cfg *server.Config, log *zerolog.Logger) *Keeper {
	return &Keeper{Log: log, Cfg: cfg}
}

func (k *Keeper) Run() {
	pgx, err := postgres.NewPGX(k.Cfg.DSN)
	if err != nil {
		k.Log.Fatal().Err(err).Msg("db connection error")
	}
	authRepo := postgres.NewAuthPostgres(pgx.DB, k.Log)
	lpRepo := postgres.NewLoginPassPostgres(pgx.DB, k.Log)
	repo := storage.NewRepository(authRepo, lpRepo)

	jwt := service.NewJWTManager(k.Cfg.SecretKeyToken, k.Cfg.TokenDuration, k.Log)
	authService := service.NewAuthService(repo.Auth, k.Log, jwt, k.Cfg.SecretKeyPassword)
	lpService := service.NewLoginPassService(repo.LoginPass, k.Log)
	srvc := service.NewService(authService, lpService)

	handler := grpcHandler.NewHandler(srvc, k.Log, k.Cfg)
	grpcSrv := grpcServer.NewServer(k.Cfg, handler)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		k.Log.Info().Msg("Starting gRPC server")
		if err = grpcSrv.Run(); err != nil && err != http.ErrServerClosed {
			k.Log.Fatal().Err(err).Msg("gRCP server run error")
		}
	}()
	<-termChan

	grpcSrv.Stop()
	k.Log.Info().Msg("Shutdown gRPC server")
}
