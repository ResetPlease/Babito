package testcore

import (
	"log/slog"
	"os"

	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/ResetPlease/Babito/api/middleware"
	"github.com/ResetPlease/Babito/internal/db"
	"github.com/ResetPlease/Babito/internal/models"
)

type TestCore struct {
	DB         db.Database
	Handler    handlers.Handler
	Middleware middleware.Middleware
}

func GetTestDatabaseCreds() db.DatabaseCreds {
	var creds db.DatabaseCreds
	creds.Host = "localhost"
	creds.Port = "5050"
	creds.User = "postgres"
	creds.Password = "password"
	creds.DBname = "shop"
	return creds
}

func NewTestCore() *TestCore {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(logHandler)
	testConfig := models.NewConfig("jwtsecret", 1000)
	db := db.NewDatabaseController(GetTestDatabaseCreds, logger)
	return &TestCore{
		DB:         db,
		Handler:    *handlers.NewHandler(db, logger, *testConfig),
		Middleware: *middleware.NewMiddleware(db, logger, *testConfig),
	}
}
