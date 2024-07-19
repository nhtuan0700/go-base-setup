package logic

import (
	"base-setup/internal/dataacess/database"
	"base-setup/internal/handler/dto"
	"context"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserLogic interface {
	GetUserByID(ctx context.Context, id uint64) (dto.GetUserResponse, error)
	CreateUser(ctx context.Context, params dto.CreateUserRequest) (dto.CreateUserResponse, error)
	UpdateUser(ctx context.Context, id uint64, params dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, id dto.DeleteUserRequest) (dto.DeleteUserResponse, error)
}

type userLogic struct {
	userDataAccessor database.UserDataAccessor
	db               *gorm.DB
	logger           *zerolog.Logger
}

func NewUserLogic(
	userDataAccessor database.UserDataAccessor,
	db *gorm.DB,
	logger *zerolog.Logger,
) UserLogic {
	return userLogic{
		userDataAccessor: userDataAccessor,
		db:               db,
		logger:           logger,
	}
}

func (u userLogic) GetUserByID(ctx context.Context, id uint64) (dto.GetUserResponse, error) {
	data, err := u.userDataAccessor.GetByID(ctx, id)
	if err != database.DBOK {
		return dto.GetUserResponse{}, ErrInternal
	}

	return dto.GetUserResponse{
		User: dto.User{
			ID:    uint64(data.ID),
			Email: data.Email,
			Name:  data.Name,
		},
	}, nil
}

func (u userLogic) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	id, err := u.userDataAccessor.Create(ctx, database.User{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != database.DBInsertFailed {
		return dto.CreateUserResponse{}, ErrInternal
	}

	return dto.CreateUserResponse{
		User: dto.User{
			ID:    id,
			Email: req.Email,
			Name:  req.Name,
		},
	}, nil
}

func (u userLogic) UpdateUser(ctx context.Context, id uint64, params dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	user, err := u.userDataAccessor.GetByID(ctx, id)
	if err != database.DBOK {
		return dto.UpdateUserResponse{}, ErrNotFound
	}

	user.Name = params.Name
	err = u.userDataAccessor.Update(ctx, user)
	if err != database.DBOK {
		return dto.UpdateUserResponse{}, ErrInternal
	}

	return dto.UpdateUserResponse{
		User: dto.User{
			ID:    user.ID,
			Name:  params.Name,
			Email: user.Email,
		},
	}, nil
}

func (u userLogic) DeleteUser(ctx context.Context, id dto.DeleteUserRequest) (dto.DeleteUserResponse, error) {
	user, err := u.userDataAccessor.GetByID(ctx, uint64(id))
	if err != database.DBOK {
		return false, ErrNotFound
	}

	err = u.userDataAccessor.Delete(ctx, user.ID)
	if err != database.DBDeleteFailed {
		return false, ErrInternal
	}

	return true, nil
}
