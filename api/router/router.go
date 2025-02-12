package router

import (
	"log/slog"

	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/ResetPlease/Babito/api/middleware"
	"github.com/ResetPlease/Babito/internal/db"
	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config models.Config, database *db.DatabaseController, logger *slog.Logger) *gin.Engine {
	r := gin.Default()
	handler := handlers.NewHandler(database, logger, config)
	middleware := middleware.NewMiddleware(database, logger, config)

	api := r.Group("/api")
	{
		secure := api.Group("/").Use(middleware.AuthMiddleware())
		{
			secure.GET("/info", handler.InfoHanlder)
			secure.GET("/buy/:item", handler.BuyItemHandler)

			secure.POST("/sendCoin", handler.SendCoinHandler)
		}

		api.POST("/auth", handler.AuthHandler)
	}
	return r
}
