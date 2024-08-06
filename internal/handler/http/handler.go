package http

import (
	"base-setup/internal/handler/http/middleware"

	"github.com/labstack/echo/v4"
)


type Handler struct {
	CheckHealthHandler CheckHealthHandler
	UserHandler        UserHandler
	// PostHandler        PostHandler
}

func (h Handler) RegisterRoutes(r *echo.Echo) {
	r.Use(middleware.RequestMiddleware())
	rg := r.Group("/api/v1")

	h.CheckHealthHandler.SetHandler(rg)
	h.UserHandler.SetHandler(rg)
	// h.PostHandler.SetHandler(rg)
}
