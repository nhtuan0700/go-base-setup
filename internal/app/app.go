package app

import (
	"base-setup/internal/configs"
	"base-setup/internal/handler/v1"
	"base-setup/internal/utils"
	"fmt"
	"log"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	logger     *zerolog.Logger
	httpConfig configs.HTTP
	handler    handler.Handler
}

func NewServer(
	logger *zerolog.Logger,
	httpConfig configs.HTTP,
	handler handler.Handler,
) *Server {
	return &Server{
		logger:     logger,
		httpConfig: httpConfig,
		handler:    handler,
	}
}

func (s Server) Start() error {
	go func() {
		server := gin.Default()

		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		server.Use(cors.New(config))

		s.handler.RegisterRoutes(server)

		address := fmt.Sprintf("%s:%s", s.httpConfig.Address, s.httpConfig.Port)
		log.Println("Starting on: " + address)
		s.logger.Info().Msg("Starting on: " + address)
		_ = server.Run()

		// s.logger.Error().Stack().Err(errors.WithStack(err)).Send()
	}()

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)

	return nil
}
