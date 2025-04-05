package config

import (
	"log/slog"
	"os"
	"strings"
)

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: getLogLevel(),
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(logger)

	return logger
}

func getLogLevel() *slog.LevelVar {
	lvl := new(slog.LevelVar)

	switch strings.ToLower(Env.Log.Level) {
	case "error":
		lvl.Set(slog.LevelError)
	case "warn":
		lvl.Set(slog.LevelWarn)
	case "debug":
		lvl.Set(slog.LevelDebug)
	case "info":
		lvl.Set(slog.LevelInfo)
	default:
		lvl.Set(slog.LevelInfo)
	}

	return lvl
}
