package logic

import (
	"base-setup/internal/dataacess/database"
	"base-setup/internal/utils"
	"context"

	"go.uber.org/zap"
)

type CreateSessionParams struct {
	Email    string
	Password string
}

type RegisterAccountParams struct {
	Name     string
	Email    string
	Password string
}

type AuthLogic interface {
	CreateSession(ctx context.Context, params CreateSessionParams) (string, error)
	RegisterAccount(ctx context.Context, params RegisterAccountParams) error
}

type authLogic struct {
	userDataAccessor database.UserDataAccessor
	logger           *zap.Logger
}

func NewAuthLogic(
	userDataAccessor database.UserDataAccessor,
	logger *zap.Logger,
) AuthLogic {
	return authLogic{
		userDataAccessor: userDataAccessor,
		logger:           logger,
	}
}

func (a authLogic) CreateSession(ctx context.Context, params CreateSessionParams) (string, error) {
	return "", nil
}

func (a authLogic) RegisterAccount(ctx context.Context, params RegisterAccountParams) error {
	_, ec := a.userDataAccessor.GetByEmail(ctx, params.Email)
	if ec == database.DBOK {
		return ErrDuplicateEmail
	}

	hashedPassword, err := utils.Hash(params.Password)
	if err != nil {
		return err
	}

	_, ec = a.userDataAccessor.Create(ctx, database.User{
		Email:    params.Email,
		Password: hashedPassword,
		Name:     params.Name,
	})

	if ec != database.DBOK {
		return ErrInternal
	}

	return nil
}
