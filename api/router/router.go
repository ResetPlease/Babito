package router

import (
	"log/slog"

	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/ResetPlease/Babito/internal/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter(database *db.DatabaseController, logger *slog.Logger) *gin.Engine {
	r := gin.Default()
	handler := handlers.NewHandler(database, logger)

	r.GET("/api/info", handler.InfoHanlder)
	r.GET("/api/buy/:item", handler.BuyItemHandler)

	r.POST("/api/sendCoin", handler.SendCoinHandler)
	r.POST("/api/auth", handler.AuthHandler)

	return r
}
