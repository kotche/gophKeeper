package main

import (
	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/app"
	"github.com/kotche/gophKeeper/logger"
)

func main() {
	cfg, err := client.NewConfig()
	log := logger.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("client configuration error")
		return
	}

	client := app.NewApp(cfg, log)
	client.Run()
}
