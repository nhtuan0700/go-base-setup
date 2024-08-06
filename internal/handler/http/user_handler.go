package http

import (
	"base-setup/internal/dto"
	"base-setup/internal/logic"
	"base-setup/internal/validation"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	userLogic logic.UserLogic
	logger    *zap.Logger
}

func NewUserHandler(userLogic logic.UserLogic, logger *zap.Logger) UserHandler {
	return UserHandler{
		userLogic: userLogic,
		logger:    logger,
	}
}

func (h UserHandler) SetHandler(rg *echo.Group) {
	g := rg.Group("/users")
	g.GET("/:id", h.get)
	g.POST("/", h.create)
	g.PUT("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

func (h UserHandler) get(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
	}

	data, err := h.userLogic.GetUserByID(c.Request().Context(), uint64(userID))
	if err != nil {
		return Set404Response(c)
	}

	return Set200Response(c, data)
}

func (h UserHandler) create(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		ves := validation.GetValidationErrors(req, err)
		if ves != nil {
			return Set400Response(c, ves)
		}
		return Set400Response(c, err)
	}

	data, err := h.userLogic.CreateUser(c.Request().Context(), req)
	if err != nil {
		return Set500Response(c, err)
	}

	return Set200Response(c, data)
}

func (h UserHandler) update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
		
	}
	
	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		ves := validation.GetValidationErrors(req, err)
		if ves != nil {
			return Set400Response(c, ves)
		}
		return Set400Response(c, err)
	}

	data, err := h.userLogic.UpdateUser(c.Request().Context(), uint64(userID), req)
	if err != nil {
		return Set500Response(c, err)
	}

	return Set200Response(c, data)
}

func (h UserHandler) delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
	}

	data, err := h.userLogic.DeleteUser(c.Request().Context(), dto.DeleteUserRequest(userID))
	if err != nil {
		return Set500Response(c, err)
	}

	return Set200Response(c, data)
}
