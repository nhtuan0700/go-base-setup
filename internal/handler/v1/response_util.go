package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

func Set400Response(c *gin.Context, err any) {
	res := ErrorResponse{
		Message: "invalid data",
	}
	switch e := err.(type) {
	case map[string]string:
		res.Details = e
	case error:
		res.Details = e.Error()
	}
	c.JSON(http.StatusBadRequest, res)
}

func Set200Response(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Set404Response(c *gin.Context) {
	c.JSON(http.StatusNotFound, "Model not found")
}

func Set500Response(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
}

