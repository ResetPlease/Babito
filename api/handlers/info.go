package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InfoHanlder(c *gin.Context) {
	// It is a test variant of info handler for check database
	id := c.Query("id")
	UID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err,
		})
		return
	}
	user, err := h.db.GetUserDataByUserID(uint64(UID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
