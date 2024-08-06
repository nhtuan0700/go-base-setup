package handler

import (
	"base-setup/internal/handler/dto"
	"base-setup/internal/logic"

	"github.com/gin-gonic/gin"
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

func (h AuthHandler) SetHandler(rg *gin.RouterGroup) {
	g := rg.Group("auth")
	g.POST("/login", h.login)
}

func (h AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

	}

	
}
