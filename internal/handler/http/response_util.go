package http

import (
	"base-setup/internal/logic"
	"base-setup/internal/validation"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error400Response struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func Set400Response(c echo.Context, err error) error {
	res := Error400Response{
		Message: "Bad request",
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

func SetErrorResponse(c echo.Context, err error) error {
	if errors.Is(err, logic.ErrNotFound) {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Resource not found"})
	}
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
}
