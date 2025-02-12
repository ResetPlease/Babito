package middleware

import (
	"log/slog"

	"github.com/ResetPlease/Babito/internal/db"
	"github.com/ResetPlease/Babito/internal/models"
)

type Middleware struct {
	Config models.Config
	logger *slog.Logger
	db     db.Database
}

func NewMiddleware(db db.Database, logger *slog.Logger, config models.Config) *Middleware {
	return &Middleware{
		Config: config,
		logger: logger,
		db:     db,
	}
}
