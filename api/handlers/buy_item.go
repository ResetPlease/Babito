package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
)

func (h *Handler) BuyItemHandler(c *gin.Context) {
	user, err := tools.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
		return
	}

	itemName := c.Param(models.ParamItemName)

	err = h.db.BuyItemByName(user.ID, itemName)

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			h.logger.Error("DB error: user not exist", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, models.ErrorUserNotExist)
			return
		} else if errors.Is(err, models.ErrNotEnoughtFunds) {
			h.logger.Error("DB error: not enought funds", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, models.ErrorNotEnoughtFunds)
			return
		} else if errors.Is(err, models.ErrProductNotFound) {
			h.logger.Error("DB error: product not found", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, models.ErrorProductNotFound)
			return
		}
		h.logger.Error("DB error", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}
}
