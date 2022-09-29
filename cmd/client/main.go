package main

import (
	"github.com/kotche/gophKeeper/internal/client/app"
	"github.com/kotche/gophKeeper/internal/client/config"
	"github.com/kotche/gophKeeper/pkg/logger"
)

func main() {
	cfg, err := config.NewConfig()
	log := logger.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("client configuration error")
		return
	}

	client := app.NewApp(cfg, log)
	client.Run()
}
