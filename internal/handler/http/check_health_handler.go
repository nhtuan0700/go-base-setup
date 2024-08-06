package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckHealthHandler struct {
}

func NewCheckHealthHandler() CheckHealthHandler {
	return CheckHealthHandler{}
}

func (h CheckHealthHandler) SetHandler(rg *echo.Group) {
	rg.GET("/health", h.get)
}

func (h CheckHealthHandler) get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "API is ok"})
}
