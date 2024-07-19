package utils

import (
	"base-setup/internal/common"
	"base-setup/internal/configs"
	"context"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger interface {
	
}

func getLoggerLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "panic":
		return zerolog.PanicLevel
	}
	return zerolog.InfoLevel
}

func InitializeLogger(logConfig configs.Log) (*zerolog.Logger, func(), error) {
	logDir := "logs"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			return nil, nil, err
		}
	}

	logPath := filepath.Join(logDir, "app.log")

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = file.Close()
	}

	multi := zerolog.MultiLevelWriter(file, zerolog.ConsoleWriter{Out: os.Stderr})
	logger := zerolog.New(multi).With().Timestamp().Stack().Logger()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(getLoggerLevel(logConfig.Level))

	return &logger, cleanup, nil
}

func LoggerWithContext(ctx context.Context, logger *zerolog.Logger) *zerolog.Logger {
	requestID, ok := ctx.Value(common.RequestIDContext).(string)
	if ok {
		newLogger := logger.With().Str("request_id", requestID).Logger()
		return &newLogger
	}

	return logger
}
