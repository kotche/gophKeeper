package main

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/app"
	"github.com/kotche/gophKeeper/logger"
)

func main() {
	cfg, err := client.NewConfig()
	if err != nil {
		log := logger.Init("")
		log.Fatal().Err(err).Msg("client configuration error")
		return
	}
	log := logger.Init(cfg.LogLevel)

	client := app.NewApp(cfg, log)
	client.Run()
}
