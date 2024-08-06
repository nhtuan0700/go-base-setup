package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

func Set400Response(c echo.Context, err any) error {
	res := ErrorResponse{
		Message: "invalid data",
	}
	switch e := err.(type) {
	case map[string]string:
		res.Details = e
	case error:
		res.Details = e.Error()
	}
	return c.JSON(http.StatusBadRequest, res)
}

func Set200Response(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

func Set404Response(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Model not found")
}

func Set500Response(c echo.Context, err error) error{
	return c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
}

