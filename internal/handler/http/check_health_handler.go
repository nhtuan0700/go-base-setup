package http

import (
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

type CheckHealthResponse struct {
	Message string `json:"message" example:"API is ok"`
}

// get
// @Summary Check api health
// @Description Check api health
// @Accept json
// @Produce json
// @Success 200 {object} CheckHealthResponse
// @Router /health [get]
func (h CheckHealthHandler) get(c echo.Context) error {
	return Set200Response(c, CheckHealthResponse{
		Message: "API is ok",
	})
}
