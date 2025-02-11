package handlers

import (
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

type Database interface {
	GetUserDataByUserID(id uint64) (*models.User, error)
	GetUserDataByUsername(username string) (*models.User, error)
	CreateNewUser(username string, hashedPassword string, balance int64) (*models.User, error)
}

type Handler struct {
	db     Database
	logger *slog.Logger
	config models.Config
}

func NewHandler(db Database, logger *slog.Logger, config models.Config) *Handler {
	return &Handler{
		db:     db,
		logger: logger,
		config: config,
	}
}
