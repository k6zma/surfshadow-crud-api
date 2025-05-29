package logger

import (
	"log/slog"
	"os"
	"strings"
	"sync"
)

var once sync.Once

func InitLogger(logLevel string) {
	once.Do(func() {
		var handler slog.Handler

		level := strToSlogLevel(logLevel)
		handler = newJSONLogHandler(os.Stdout, level)

		logger := slog.New(handler)
		slog.SetDefault(logger)
	})
}

func strToSlogLevel(level string) slog.Level {
	level = strings.ToLower(level)

	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
