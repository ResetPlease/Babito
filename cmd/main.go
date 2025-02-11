package main

import (
	"log/slog"

	"github.com/ResetPlease/Babito/api/router"
	"github.com/ResetPlease/Babito/internal/db"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
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

	database := db.NewDatabaseController(logger)
	defer database.Close()

	r := router.SetupRouter(database, logger)

	err := r.Run(":8080")
	if err != nil {
		logger.Error("Server error: ", slog.Any("error", err))
	}
}
