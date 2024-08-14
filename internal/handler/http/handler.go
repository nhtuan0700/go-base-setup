package http

import (
	"base-setup/internal/handler/http/middleware"

	"github.com/labstack/echo/v4"
)


type Handler struct {
	CheckHealthHandler CheckHealthHandler
	UserHandler        UserHandler
	AuthHandler        AuthHandler
}

// @title           Go example
// @version         1.0
// @description     Go example API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      HOST_NAME
// @BasePath  /api/v1

func (h Handler) RegisterRoutes(r *echo.Echo) {
	r.Use(middleware.RequestMiddleware())
	rg := r.Group("/api/v1")

	h.CheckHealthHandler.SetHandler(rg)
	h.UserHandler.SetHandler(rg)
	h.AuthHandler.SetHandler(rg)
}
