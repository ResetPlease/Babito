package handlers

import (
	"log/slog"

	"github.com/ResetPlease/Babito/internal/db"
	"github.com/ResetPlease/Babito/internal/models"
)

type Handler struct {
	db     db.Database
	logger *slog.Logger
	config models.Config
}

func NewHandler(db db.Database, logger *slog.Logger, config models.Config) *Handler {
	return &Handler{
		db:     db,
		logger: logger,
		config: config,
	}
}
