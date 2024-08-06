package middleware

import (
	"base-setup/internal/common"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := uuid.New().String()

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, common.RequestIDContext, requestID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
