package database

import (
	"base-setup/internal/utils"
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint64 `gorm:"primarykey"`
	Title     string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostDataAccessor interface {
	Create(ctx context.Context, post Post) (uint64, error)
	WithDB(db *gorm.DB) PostDataAccessor
}

type postDataAccessor struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostDataAccessor(
	db *gorm.DB,
	logger *zap.Logger,
) PostDataAccessor {
	return &postDataAccessor{
		db:     db,
		logger: logger,
	}
}

func (u postDataAccessor) WithDB(db *gorm.DB) PostDataAccessor {
	return postDataAccessor{
		logger: u.logger,
		db:     db,
	}
}

func (u postDataAccessor) Create(ctx context.Context, post Post) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, u.logger).With(zap.Any("post", post))

	if err := u.db.WithContext(ctx).Create(&post).Error; err != nil {
		logger.With(zap.Error(err)).Error("failed to create post")
		return 0, err
	}

	return uint64(post.ID), nil
}
