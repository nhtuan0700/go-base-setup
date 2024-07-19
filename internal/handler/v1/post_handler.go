package handler

import (
	"base-setup/internal/handler/dto"
	"base-setup/internal/logic"
	"base-setup/internal/validation"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog"
)

type PostHandler struct {
	userLogic logic.UserLogic
	logger    *zerolog.Logger
}

func NewPostHandler(userLogic logic.UserLogic, logger *zerolog.Logger) PostHandler {
	return PostHandler{
		userLogic: userLogic,
		logger:    logger,
	}
}

func (h PostHandler) SetHandler(rg *gin.RouterGroup) {
	g := rg.Group("posts")
	g.GET(":id", h.get)
	g.POST("", h.create)
	g.PUT(":id", h.update)
	g.DELETE(":id", h.delete)
}

func (h PostHandler) get(c *gin.Context) {
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

func (h PostHandler) create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.MustBindWith(&req, binding.FormMultipart); err != nil {
		ves := validation.GetValidationErrors(req, err)
		log.Println(err)
		if ves != nil {
			Set400Response(c, ves)
			return
		}
		Set400Response(c, err)
		return
	}

	// data, err := h.userLogic.CreateUser(c, req)
	// if err != nil {
	// 	Set500Response(c, err)
	// 	return
	// }

	Set200Response(c, dto.CreatePostResponse{})
}

func (h PostHandler) update(c *gin.Context) {
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

func (h PostHandler) delete(c *gin.Context) {
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
