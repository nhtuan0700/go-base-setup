package database

import (
	"base-setup/internal/configs"
	"base-setup/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDSN(dbConfig configs.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
}

type gormZeroLog struct {
	logger *zerolog.Logger
}

func (la gormZeroLog) LogMode(_ logger.LogLevel) logger.Interface {
	return la
}

func (la gormZeroLog) Info(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).Info().Any("args", args)
}

func (la gormZeroLog) Warn(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).Warn().Any("args", args)
}

func (la gormZeroLog) Error(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).Error().Any("args", args)
}

func (la gormZeroLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()

	utils.LoggerWithContext(ctx, la.logger).Debug().
		Str("sql", sql).
		Dur("elapsed", time.Since(begin)).
		Int64("rows", rows).
		Err(errors.WithStack(err)).
		Send()
}

func InitializeDB(
	logger *zerolog.Logger,
	dbConfig configs.Database,
) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getDSN(dbConfig)), &gorm.Config{
		Logger: gormZeroLog{logger: logger},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
