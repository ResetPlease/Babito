package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyItemHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}
