package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BuyItemHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}
