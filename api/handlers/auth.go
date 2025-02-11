package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthHandler(c *gin.Context) {
	h.logger.Error("Got message: auth")
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}
