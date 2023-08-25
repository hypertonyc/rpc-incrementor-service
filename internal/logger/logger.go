package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hypertonyc/rpc-incrementor-service/internal/config"
)

var Logger *zap.Logger

func InitLogger(config config.AppConfig) {
	// Initialize zap logger configuration
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	// Set log level from config
	logLevel := getLogLevel(config.LogLevel)
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	// Create the logger
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
