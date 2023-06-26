package app

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	config "github.com/golang_tasks/internal"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type App struct {
	cfg *config.Config
	logger *zap.Logger
}

func NewApp(config *config.Config, logger *zap.Logger) (App, error) {
	logger.Info("router initializing")

	router := httprouter.New()

	logger.Info("swagger docs initializing")
	router.Handler(
		http.MethodGet,
		"/swagger",
		http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently),
	)
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	app := App{
		cfg: config,
		logger: logger,
	}

	return app, nil
}
