package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCoinHandler(c *gin.Context) {
	userData, err := tools.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
		return
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, models.ErrorEmptyRequestBody)
		return
	}

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error("Got error while read request body", slog.Any("error", err), slog.Any("body", c.Request.Body))
		c.JSON(http.StatusBadRequest, models.ErrorBadRequest)
		return
	}

	var requestData models.SendCoinRequest
	err = json.Unmarshal(data, &requestData)
	if err != nil {
		h.logger.Error("Got error while unmarshal body to SendCoinRequest", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, models.ErrorWrongDataFormat)
		return
	}

	if userData.Username == requestData.ToUser {
		h.logger.Error("Replenishment of your account is prohibited", slog.Any("user", userData))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Errors: "Replenishment of your account is prohibited",
		})
		return
	}

	if requestData.Amount <= 0 {
		h.logger.Error("Amount less or equal 0", slog.Any("user", userData))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Errors: "Amount should be more than zero",
		})
		return
	}

	err = h.db.SendCoinByUsername(userData.ID, requestData.ToUser, int64(requestData.Amount))
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			h.logger.Error("DB error: user not exist", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, models.ErrorUserNotExist)
			return
		} else if errors.Is(err, models.ErrNotEnoughtFunds) {
			h.logger.Error("DB error: NotEnoughtFunds", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, models.ErrorNotEnoughtFunds)
			return
		}
		h.logger.Error("DB error", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}
}
