package main

import (
	"fmt"
	"time"

	"github.com/kotche/gophKeeper/config/client"
	"github.com/kotche/gophKeeper/internal/client/app"
	"github.com/kotche/gophKeeper/logger"
)

var (
	buildVersion string
	buildDate    time.Time
)

func main() {
	printBuildInfo()

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

// example: go run -ldflags "-X main.buildVersion=v1.0 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" main.go
func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
}
