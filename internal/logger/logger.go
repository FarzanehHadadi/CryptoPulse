package logger

import (
	"context"
	"cryptoPulse/internal/config"
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger = slog.Default()
	once   sync.Once
)

func Init(cfg config.LoggerConfig) {
	once.Do(func() {
		var logLevel slog.Level
		level := cfg.LogLevel
		switch level {
		case "debug":
			logLevel = slog.LevelDebug
		case "info":
			logLevel = slog.LevelInfo
		case "warn":
			logLevel = slog.LevelWarn
		default:
			logLevel = slog.LevelInfo
		}
		opts := &slog.HandlerOptions{
			Level: logLevel,
		}
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
		logger = logger.With(
			"service", cfg.ServiceName,
			"env", cfg.Environment,
			"version", cfg.Version,
		)
	})
}

// Convenience methods
func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func WithContext(ctx context.Context) *slog.Logger {
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		return logger.With("trace_id", traceID)
	}
	return logger
}
