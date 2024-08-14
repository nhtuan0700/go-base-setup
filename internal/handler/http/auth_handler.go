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

func (h AuthHandler) SetHandler(rg *echo.Group) {
	g := rg.Group("/auth")
	g.POST("/register", h.register)
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h AuthHandler) register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return Set400Response(c, err)
	}

	if e := c.Validate(req); e != nil {
		return Set400Response(c, e)
	}
	err := h.authLogic.RegisterAccount(c.Request().Context(), logic.RegisterAccountParams{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return SetErrorResponse(c, err)
	}

	return Set200Response(c, true)
}
