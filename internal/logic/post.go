package logic

import (
	// "base-setup/internal/dataacess/database"
	// "base-setup/internal/handler/dto"
	// "context"

	// "github.com/rs/zerolog"
	// "gorm.io/gorm"
)

// type PostLogic interface {
// 	CreateUser(ctx context.Context, params dto.CreatePostRequest) (dto.CreatePostResponse, error)
// }

// type postLogic struct {
// 	postDataAccessor database.PostDataAccessor
// 	db               *gorm.DB
// 	logger           *zerolog.Logger
// }

// func NewPostLogic(
// 	userDataAccessor database.UserDataAccessor,
// 	db *gorm.DB,
// 	logger *zerolog.Logger,
// ) PostLogic {
// 	return postLogic{
// 		userDataAccessor: postDataAccessor,
// 		db:               db,
// 		logger:           logger,
// 	}
// }

// func (u userLogic) GetUserByID(ctx context.Context, id uint64) (dto.GetUserResponse, error) {
// 	data, err := u.userDataAccessor.GetByID(ctx, id)
// 	if err != nil {
// 		return dto.GetUserResponse{}, err
// 	}

// 	return dto.GetUserResponse{
// 		User: dto.User{
// 			ID:    uint64(data.ID),
// 			Email: data.Email,
// 			Name:  data.Name,
// 		},
// 	}, nil
// }

// func (u userLogic) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
// 	id, err := u.userDataAccessor.Create(ctx, database.User{
// 		Email: req.Email,
// 		Name:  req.Name,
// 	})
// 	if err != nil {
// 		return dto.CreateUserResponse{}, err
// 	}

// 	return dto.CreateUserResponse{
// 		User: dto.User{
// 			ID:    id,
// 			Email: req.Email,
// 			Name:  req.Name,
// 		},
// 	}, nil
// }
