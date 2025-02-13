package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get(models.AuthHeaderKey)
		bearerToken = strings.TrimSpace(bearerToken)
		splitToken := strings.Split(bearerToken, " ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
			c.Abort()
			return
		}
		token := splitToken[1]

		user, err := tools.ParseJWTToken(token, m.Config)
		if err != nil {
			m.logger.Error("Got error while check JWT", slog.Any("error", err))
			c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
			c.Abort()
			return
		}

		c.Set(models.UserContextKey, *user)
		c.Next()
	}
}
