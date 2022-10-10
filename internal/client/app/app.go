package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/c-bata/go-prompt"
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/service"
	"github.com/kotche/gophKeeper/internal/client/storage"
	"github.com/kotche/gophKeeper/internal/client/transport"
	grpcTransport "github.com/kotche/gophKeeper/internal/client/transport/grpc"
	"github.com/kotche/gophKeeper/internal/client/updater"
	"github.com/rs/zerolog"
)

type App struct {
	Conf *client.Config
	Log  *zerolog.Logger
}

func NewApp(conf *client.Config, log *zerolog.Logger) *App {
	return &App{Conf: conf, Log: log}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	clientConn := grpcTransport.Connection{}
	cache := storage.NewCache(a.Log)
	srvc := service.NewService(cache, a.Conf, a.Log)
	sender := grpcTransport.NewSender(srvc, clientConn, a.Conf, a.Log)
	commander := transport.NewCommander(sender, a.Conf, a.Log)

	updater := updater.NewUpdater(sender, srvc, a.Conf, a.Log)
	go updater.Run(ctx)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	fmt.Println("GophKeeper start")

	p := prompt.New(
		commander.Executor,
		commander.Completer,
		prompt.OptionTitle("menu"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Green),
	)
	go p.Run()

	<-termChan
	cancel()
	fmt.Println("GophKeeper stop")
}
