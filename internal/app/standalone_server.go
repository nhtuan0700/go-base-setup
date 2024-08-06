package app

import (
	"base-setup/internal/handler/http"
	"base-setup/internal/utils"
	"context"
	"syscall"

	"go.uber.org/zap"
)

type Server struct {
	httpServer http.Server
	logger     *zap.Logger
}

func NewStandaloneServer(
	httpServer http.Server,
	logger *zap.Logger,
) *Server {
	return &Server{
		httpServer: httpServer,
		logger:     logger,
	}
}

func (s Server) Start() error {
	go func() {
		err := s.httpServer.Start(context.Background())
		s.logger.With(zap.Error(err)).Info("http server stopped")
	}()

	utils.BlockUntilSignal(syscall.SIGINT, syscall.SIGTERM)
	return nil
}
