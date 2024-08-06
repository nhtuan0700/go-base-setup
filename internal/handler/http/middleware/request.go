package middleware

import (
	"base-setup/internal/common"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := uuid.New().String()
			c.Set(common.RequestIDContext, requestID)
			return next(c)
		}
	}
}
