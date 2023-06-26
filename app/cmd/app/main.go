package main

import (
	"log"

	config "github.com/golang_tasks/internal"
	app "github.com/golang_tasks/internal/app"
	"github.com/golang_tasks/pkg/common/logging"
)

func main () {
	log.Print("Config initializing")
	cfg := config.GetConfig()

	log.Print("Logger initializing")
	logger := logging.GetLogger()

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
}