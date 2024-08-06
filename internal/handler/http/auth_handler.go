package http

import (
	"base-setup/internal/logic"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authLogic logic.UserLogic
	logger    *zap.Logger
}

func NewAuthHandler(authLogic logic.UserLogic, logger *zap.Logger) AuthHandler {
	return AuthHandler{
		authLogic: authLogic,
		logger:    logger,
	}
}

func (h AuthHandler) SetHandler(rg echo.Group) {
	g := rg.Group("/auth")
	g.POST("/login", h.login)
}

func (h AuthHandler) login(c echo.Context) error {
	return nil
}
