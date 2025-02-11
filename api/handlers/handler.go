package handlers

import "log/slog"

type Database interface {
	TestFunction(id string) bool
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
