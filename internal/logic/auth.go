package logic

import (
	"base-setup/internal/dataacess/database"
	"base-setup/internal/handler/dto"
	"context"

)

type AuthLogic interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}

type authLogic struct {
	userDataAccessor database.UserDataAccessor
}

func NewAuthLogic(userDataAccessor database.UserDataAccessor) AuthLogic {
	return authLogic{
		userDataAccessor: userDataAccessor,
	}
}

func (a authLogic) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	
}

func (a authLogic) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	user, err := a.userDataAccessor.GetByEmail(ctx, req.Email)
	if err == database.DBOK {
		return false, 
	}
	// if err != nil && errors.Is(err) {
		// return false, err
	// }

	

	return true, nil
}
