package http

import (
	"base-setup/internal/configs"
	"base-setup/internal/utils"
	"base-setup/internal/validation"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	httpConfig configs.HTTP
	logger     *zap.Logger
	handler    Handler
}

func NewServer(
	httpConfig configs.HTTP,
	handler Handler,
	logger *zap.Logger,
) Server {
	return &server{
		httpConfig: httpConfig,
		handler:    handler,
		logger:     logger,
	}
}


func (s server) Start(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, s.logger)

	e := echo.New()
	e.Validator = validation.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	s.handler.RegisterRoutes(e)

	address := fmt.Sprintf("%s:%s", s.httpConfig.Host, s.httpConfig.Port)
	logger.With(zap.String("address", address)).Info("Starting http server")

	server := http.Server{
		Addr:    address,
		Handler: e,
	}

	return server.ListenAndServe()
}
