package handlers

import (
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

type Database interface {
	GetUserDataByUserID(id uint64) (*models.User, error)
}

type Handler struct {
	db     Database
	logger *slog.Logger
}

func NewHandler(db Database, logger *slog.Logger) *Handler {
	return &Handler{
		db:     db,
		logger: logger,
	}
}
