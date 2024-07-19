package handler

import (
	"base-setup/internal/handler/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	CheckHealthHandler CheckHealthHandler
	UserHandler        UserHandler
	PostHandler        PostHandler
}

func (h Handler) RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.RequestMiddleware)
	rg := r.Group("api/v1")

	h.CheckHealthHandler.SetHandler(rg)
	h.UserHandler.SetHandler(rg)
	h.PostHandler.SetHandler(rg)
}
