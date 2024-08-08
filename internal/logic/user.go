package logic

import (
	"base-setup/internal/dataacess/database"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserOutput struct {
	ID    uint64
	Email string
	Name  string
}

type CreateUserParams struct {
	Email string
	Name  string
}

type UpdateUserParams struct {
	ID   uint64
	Name string
}

type UserLogic interface {
	GetUserByID(ctx context.Context, id uint64) (UserOutput, error)
	CreateUser(ctx context.Context, params CreateUserParams) (UserOutput, error)
	UpdateUser(ctx context.Context, params UpdateUserParams) (UserOutput, error)
	DeleteUser(ctx context.Context, id uint64) error
}

type userLogic struct {
	userDataAccessor database.UserDataAccessor
	db               *gorm.DB
	logger           *zap.Logger
}

func NewUserLogic(
	userDataAccessor database.UserDataAccessor,
	db *gorm.DB,
	logger *zap.Logger,
) UserLogic {
	return userLogic{
		userDataAccessor: userDataAccessor,
		db:               db,
		logger:           logger,
	}
}

func (u userLogic) GetUserByID(ctx context.Context, id uint64) (UserOutput, error) {
	data, ec := u.userDataAccessor.GetByID(ctx, id)
	if ec != database.DBOK {
		if ec == database.DBDataNotFound {
			return UserOutput{}, ErrNotFound
		}
		return UserOutput{}, ErrInternal
	}

	return UserOutput{
		ID:    uint64(data.ID),
		Email: data.Email,
		Name:  data.Name,
	}, nil
}

func (u userLogic) CreateUser(ctx context.Context, params CreateUserParams) (UserOutput, error) {
	id, ec := u.userDataAccessor.Create(ctx, database.User{
		Email: params.Email,
		Name:  params.Name,
	})
	if ec != database.DBOK {
		return UserOutput{}, ErrInternal
	}

	return UserOutput{
		ID:    id,
		Email: params.Email,
		Name:  params.Name,
	}, nil
}

func (u userLogic) UpdateUser(ctx context.Context, params UpdateUserParams) (UserOutput, error) {
	user, ec := u.userDataAccessor.GetByID(ctx, params.ID)
	if ec != database.DBOK {
		if ec == database.DBDataNotFound {
			return UserOutput{}, ErrNotFound
		}
		return UserOutput{}, ErrInternal
	}

	user.Name = params.Name
	ec = u.userDataAccessor.Update(ctx, user)
	if ec != database.DBOK {
		return UserOutput{}, ErrInternal
	}

	return UserOutput{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (u userLogic) DeleteUser(ctx context.Context, id uint64) error {
	user, ec := u.userDataAccessor.GetByID(ctx, uint64(id))
	if ec != database.DBOK {
		if ec == database.DBDataNotFound {
			return ErrNotFound
		}
		return ErrInternal
	}

	ec = u.userDataAccessor.Delete(ctx, user.ID)
	if ec != database.DBDeleteFailed {
		return ErrInternal
	}

	return nil
}
