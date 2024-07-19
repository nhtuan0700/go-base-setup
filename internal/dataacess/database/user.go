package database

import (
	"base-setup/internal/utils"
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataAccessor interface {
	GetByID(context.Context, uint64) (User, ErrorCode)
	GetByEmail(context.Context, string) (User, ErrorCode)
	Create(context.Context, User) (uint64, ErrorCode)
	Update(context.Context, User) ErrorCode
	Delete(context.Context, uint64) ErrorCode
	WithDB(*gorm.DB) UserDataAccessor
}

type userDataAccessor struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewUserDataAccessor(
	db *gorm.DB,
	logger *zerolog.Logger,
) UserDataAccessor {
	return &userDataAccessor{
		db:     db,
		logger: logger,
	}
}

func (u userDataAccessor) WithDB(db *gorm.DB) UserDataAccessor {
	return userDataAccessor{
		logger: u.logger,
		db:     db,
	}
}

func (u userDataAccessor) GetByID(ctx context.Context, id uint64) (User, ErrorCode) {
	logger := utils.LoggerWithContext(ctx, u.logger).With().Uint64("id", id).Logger()

	user := User{}
	result := u.db.WithContext(ctx).
		Where(map[string]interface{}{"id": id}).
		Take(&user)

	if err := result.Error; err != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, DBDataNotFound
		}

		logger.Error().Stack().Err(errors.WithStack(err)).Msg("failed to get user")
		return user, DBGetFailed
	}

	return user, DBOK
}

func (u userDataAccessor) Create(ctx context.Context, user User) (uint64, ErrorCode) {
	logger := utils.LoggerWithContext(ctx, u.logger).With().Any("user", user).Logger()

	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		logger.Error().Err(errors.WithStack(err)).Msg("failed to create user")
		return 0, DBInsertFailed
	}

	return uint64(user.ID), DBOK
}

func (u userDataAccessor) Update(ctx context.Context, user User) ErrorCode {
	logger := utils.LoggerWithContext(ctx, u.logger).With().Any("user", user).Logger()

	if err := u.db.WithContext(ctx).Save(&user).Error; err != nil {
		logger.Error().Err(errors.WithStack(err)).Msg("failed to update user")
		return DBUpdateFailed
	}

	return DBOK
}

func (u userDataAccessor) Delete(ctx context.Context, id uint64) ErrorCode {
	logger := utils.LoggerWithContext(ctx, u.logger).With().Any("id", id).Logger()

	if err := u.db.WithContext(ctx).Delete(&User{ID: id}).Error; err != nil {
		logger.Error().Stack().Err(errors.WithStack(err)).Msg("failed to delete user")
		return DBDeleteFailed
	}

	return DBDataNotFound
}

func (u userDataAccessor) GetByEmail(ctx context.Context, email string) (User, ErrorCode) {
	logger := utils.LoggerWithContext(ctx, u.logger).With().Str("email", email).Logger()

	user := User{}
	result := u.db.WithContext(ctx).
		Where(map[string]interface{}{"email": email}).
		Take(&user)

	if err := result.Error; err != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, DBDataNotFound
		}

		logger.Error().Stack().Err(errors.WithStack(err)).Msg("failed to get user by email")
		return user, DBGetFailed
	}

	return user, DBOK
}
