package router

import (
	"github.com/ResetPlease/Babito/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/info", handlers.InfoHanlder)
	r.GET("/api/buy/:item", handlers.BuyItemHandler)

	r.POST("/api/sendCoin", handlers.SendCoinHandler)
	r.POST("/api/auth", handlers.AuthHandler)

	return r
}
