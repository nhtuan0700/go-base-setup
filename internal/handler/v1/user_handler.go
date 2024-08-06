package handler

import (
	"base-setup/internal/handler/dto"
	"base-setup/internal/logic"
	"base-setup/internal/validation"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userLogic logic.UserLogic
	logger    *zap.Logger
}

func NewUserHandler(userLogic logic.UserLogic, logger *zap.Logger) UserHandler {
	return UserHandler{
		userLogic: userLogic,
		logger:    logger,
	}
}

func (h UserHandler) SetHandler(rg *gin.RouterGroup) {
	g := rg.Group("users")
	g.GET(":id", h.get)
	g.POST("", h.create)
	g.PUT(":id", h.update)
	g.DELETE(":id", h.delete)
}

func (h UserHandler) get(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Set400Response(c, errors.New("failed to parse id"))
		return
	}

	data, err := h.userLogic.GetUserByID(c, uint64(userID))
	if err != nil {
		Set404Response(c)
		return
	}

	Set200Response(c, data)
}

func (h UserHandler) create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ves := validation.GetValidationErrors(req, err)
		if ves != nil {
			Set400Response(c, ves)
			return
		}
		Set400Response(c, err)
		return
	}

	data, err := h.userLogic.CreateUser(c, req)
	if err != nil {
		Set500Response(c, err)
		return
	}

	Set200Response(c, data)
}

func (h UserHandler) update(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Set400Response(c, errors.New("failed to parse id"))
		return
	}
	
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ves := validation.GetValidationErrors(req, err)
		if ves != nil {
			Set400Response(c, ves)
			return
		}
		Set400Response(c, err)
		return
	}

	data, err := h.userLogic.UpdateUser(c, uint64(userID), req)
	if err != nil {
		Set500Response(c, err)
		return
	}

	Set200Response(c, data)
}

func (h UserHandler) delete(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Set400Response(c, errors.New("failed to parse id"))
		return
	}

	data, err := h.userLogic.DeleteUser(c, dto.DeleteUserRequest(userID))
	if err != nil {
		Set500Response(c, err)
		return
	}

	Set200Response(c, data)
}
