package http

import (
	"base-setup/internal/logic"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authLogic logic.AuthLogic
	logger    *zap.Logger
}

func NewAuthHandler(authLogic logic.AuthLogic, logger *zap.Logger) AuthHandler {
	return AuthHandler{
		authLogic: authLogic,
		logger:    logger,
	}
}

func (h AuthHandler) SetHandler(rg echo.Group) {
	g := rg.Group("/auth")
	g.POST("/login", h.login)
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (h AuthHandler) login(c echo.Context) error {
	return nil
}
