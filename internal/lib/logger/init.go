package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(path string) *slog.Logger {

	var (
		level  = slog.LevelInfo
		writer io.Writer
	)

	if path == "" {
		writer = io.MultiWriter(
			os.Stdout,
		)
	} else {
		writer = io.MultiWriter(
			os.Stdout,
			createFileWriter(path),
		)
	}

	logger := slog.New(
		slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: level,
		}),
	)

	return logger
}

func createFileWriter(path string) io.Writer {

	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}

	return &lumberjack.Logger{
		Filename:   path + "/app.log",
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}
}
