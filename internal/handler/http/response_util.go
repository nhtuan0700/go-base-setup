package http

import (
	"base-setup/internal/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

func Set400Response(c echo.Context, err error) error {
	res := ErrorResponse{
		Message: "invalid data",
	}

	switch e := err.(type) {
	case validation.Validation:
		res.Details = e.Details
	case error:
		res.Details = e.Error()
	}
	return c.JSON(http.StatusBadRequest, res)
}

func Set200Response(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

func Set404Response(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Model not found"})
}

func Set500Response(c echo.Context, err error) error{
	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
}

