package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckHealthHandler struct {
}

func NewCheckHealthHandler() CheckHealthHandler {
	return CheckHealthHandler{}
}

func (h CheckHealthHandler) SetHandler(rg *gin.RouterGroup) {
	rg.GET("/health", h.get)
}

func (h CheckHealthHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is ok"})
}
