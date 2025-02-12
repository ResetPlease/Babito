package router

import (
	"log/slog"

	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/ResetPlease/Babito/internal/db"
	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config models.Config, database *db.DatabaseController, logger *slog.Logger) *gin.Engine {
	r := gin.Default()
	handler := handlers.NewHandler(database, logger, config)

	r.GET("/api/info", handler.InfoHanlder)
	r.GET("/api/buy/:item", handler.BuyItemHandler)

	r.POST("/api/sendCoin", handler.SendCoinHandler)
	r.POST("/api/auth", handler.AuthHandler)

	return r
}
