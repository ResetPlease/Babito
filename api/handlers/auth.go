package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	models "github.com/ResetPlease/Babito/internal/models"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthHandler(c *gin.Context) {
	var creds models.AuthRequest

	if c.Request.Body == nil {
		h.logger.Error("Got empty request body")
		c.JSON(http.StatusBadRequest, models.ErrorEmptyRequestBody)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error("Got error while read request body", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, models.ErrorEmptyRequestBody)
		return
	}

	err = json.Unmarshal(body, &creds)
	if err != nil {
		h.logger.Error("Got error while unmarshal user data", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, models.ErrorWrongDataFormat)
		return
	}

	if len(creds.Username) == 0 || len(creds.Password) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrorMissingRequiredField)
		return
	}

	hashedPassword := tools.GenerateHash(creds.Password)
	// try to create new user else do nothing and get userID + username
	userData, err := h.db.CreateNewUser(creds.Username, hashedPassword, h.config.DefaultUserBalance)
	if err != nil {
		h.logger.Error("failed to create new user", slog.Any("error", err), slog.Any("username", creds.Username))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}

	if hashedPassword != userData.HashedPassword {
		c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
		return
	}

	token, err := tools.GenerateJWTToken(userData.ID, userData.Username, h.config.JWTSecret)
	if err != nil {
		h.logger.Error("failed to create JWT token for user", slog.Any("error", err), slog.Any("username", userData.Username))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		Token: &token,
	})
}
