package database

import (
	"base-setup/internal/configs"
	"base-setup/internal/utils"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDSN(dbConfig configs.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
}

type gormZeroLog struct {
	logger *zap.Logger
}

func (la gormZeroLog) LogMode(_ logger.LogLevel) logger.Interface {
	return la
}

func (la gormZeroLog) Info(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).With(zap.Any("args", args)).Info(s)
}

func (la gormZeroLog) Warn(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).With(zap.Any("args", args)).Warn(s)
}

func (la gormZeroLog) Error(ctx context.Context, s string, args ...interface{}) {
	utils.LoggerWithContext(ctx, la.logger).With(zap.Any("args", args)).Error(s)
}

func (la gormZeroLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()

	utils.LoggerWithContext(ctx, la.logger).Debug(
		"debug: ",
		zap.String("sql", sql),
		zap.Duration("elapsed", time.Since(begin)),
		zap.Int64("rows", rows),
		zap.Error(err),
	)	
}

func InitializeDB(
	logger *zap.Logger,
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
