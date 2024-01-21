package logger

import (
	"log/slog"
	"os"
	"time"
)

const TimeFormat = "15:04:05"
const DateFormat = "02 января 2006"

var ConfigureTime = slog.Group(
	"date",
	slog.String("", time.Now().Format(DateFormat)),
	slog.String("time", time.Now().Format(TimeFormat)),
)

func replace(groups []string, a slog.Attr) slog.Attr {
	if a.Key != slog.TimeKey || len(groups) != 0 {
		return a
	}
	return slog.Attr{}
}

func LogStart(level string) *slog.Logger {
	programLevel := new(slog.LevelVar)
	switch level {
	case "slog.LevelDebug":
		programLevel.Set(slog.LevelDebug)
	case "slog.LevelInfo":
		programLevel.Set(slog.LevelInfo)
	case "slog.LevelWarn":
		programLevel.Set(slog.LevelWarn)
	case "slog.LevelError":
		programLevel.Set(slog.LevelError)
	default:
		programLevel.Set(slog.LevelDebug)
	}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       programLevel.Level(),
		ReplaceAttr: replace,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	newLogger := logger.With(ConfigureTime)
	return newLogger
}

func WriteLogInTheFile(f *os.File) *slog.Logger {
	handler := slog.NewTextHandler(f, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
		ReplaceAttr: replace,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	newLogger := logger.With(ConfigureTime)
	return newLogger
}

func OpenLogFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
