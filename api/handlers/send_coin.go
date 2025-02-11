package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCoinHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}
