package middleware

import (
	"base-setup/internal/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestMiddleware(c *gin.Context) {
	requestID := uuid.New().String()

	c.Set(common.RequestIDContext, requestID)
	c.Next()
}
