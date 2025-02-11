package main

import (
	"fmt"
	"log/slog"

	"github.com/ResetPlease/Babito/api/router"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("Babito service init")

	logWriter := &lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	loggerHandler := slog.NewTextHandler(logWriter, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(loggerHandler)

	r := router.SetupRouter(logger)

	err := r.Run(":8080")
	if err != nil {
		logger.Error("Server error: ", slog.Any("error", err))
	}
}
