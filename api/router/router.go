package router

import (
	"log/slog"

	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(logger *slog.Logger) *gin.Engine {
	r := gin.Default()
	handler := handlers.NewHandler(nil, logger)

	r.GET("/api/info", handler.InfoHanlder)
	r.GET("/api/buy/:item", handler.BuyItemHandler)

	r.POST("/api/sendCoin", handler.SendCoinHandler)
	r.POST("/api/auth", handler.AuthHandler)

	return r
}
