package http

import (
	"base-setup/internal/logic"
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
	g.POST("", h.create)
	g.PUT("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

type User struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GetUserResponse struct {
	User
}

// get
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} GetUserResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/get/:id [get]
func (h UserHandler) get(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
	}

	output, err := h.userLogic.GetUserByID(c.Request().Context(), uint64(userID))
	if err != nil {
		return SetErrorResponse(c, err)
	}

	return Set200Response(c, GetUserResponse{
		User{
			ID:    output.ID,
			Email: output.Email,
			Name:  output.Name,
		},
	})
}

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

type CreateUserResponse struct {
	User
}

// create
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param Request body CreateUserRequest true "CreateUserRequest"
// @Success 200 {object} CreateUserResponse
// @Success 500 {object} ErrorResponse
// @Router /users/ [post]
func (h UserHandler) create(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return Set400Response(c, err)
	}

	if e := c.Validate(req); e != nil {
		return Set400Response(c, e)
	}

	output, err := h.userLogic.CreateUser(c.Request().Context(), logic.CreateUserParams{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		return SetErrorResponse(c, err)
	}

	return Set200Response(c, CreateUserResponse{
		User{
			ID:    output.ID,
			Email: output.Email,
			Name:  output.Name,
		},
	})
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserResponse struct {
	User
}

// update
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param Request body UpdateUserRequest true "UpdateUserRequest"
// @Success 200 {object} UpdateUserResponse
// @Success 404 {object} ErrorResponse
// @Success 500 {object} ErrorResponse
// @Router /users/:id [put]
func (h UserHandler) update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
	}

	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return Set400Response(c, err)
	}

	if err := c.Validate(req); err != nil {
		return Set400Response(c, err)
	}

	output, err := h.userLogic.UpdateUser(c.Request().Context(), logic.UpdateUserParams{
		ID:   uint64(userID),
		Name: req.Name,
	})
	if err != nil {
		return SetErrorResponse(c, err)
	}

	return Set200Response(c, UpdateUserResponse{
		User{
			ID:    output.ID,
			Name:  output.Name,
			Email: output.Email,
		},
	})
}

// delete
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 boolean status
// @Success 404 {object} ErrorResponse
// @Success 500 {object} ErrorResponse
// @Router /users/:id [delete]
func (h UserHandler) delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Set400Response(c, errors.New("failed to parse id"))
	}

	err = h.userLogic.DeleteUser(c.Request().Context(), uint64(userID))
	if err != nil {
		return SetErrorResponse(c, err)
	}

	return Set200Response(c, true)
}
