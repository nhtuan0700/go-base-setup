package logic

import (
	"base-setup/internal/dataacess/database"
	"context"
)

type CreateSessionParams struct {
	Email    string
	Password string
}

type RegisterAccountParams struct {
	Email    string
	Password string
}

type AuthLogic interface {
	CreateSession(ctx context.Context, params CreateSessionParams) (string, error)
	RegisterAccount(ctx context.Context, params RegisterAccountParams) error
}

type authLogic struct {
	userDataAccessor database.UserDataAccessor
}

func NewAuthLogic(userDataAccessor database.UserDataAccessor) AuthLogic {
	return authLogic{
		userDataAccessor: userDataAccessor,
	}
}

func (a authLogic) CreateSession(ctx context.Context, params CreateSessionParams) (string, error) {
	return "", nil
}

func (a authLogic) RegisterAccount(ctx context.Context, params RegisterAccountParams) error {
	// user, err := a.userDataAccessor.GetByEmail(ctx, req.Email)
	// if err == database.DBOK {
	// 	return false,
	// }
	// if err != nil && errors.Is(err) {
	// return false, err
	// }

	return nil
}
