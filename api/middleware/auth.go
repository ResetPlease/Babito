package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const authHeaderKey = "Authorization"

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get(authHeaderKey)
		bearerToken = strings.TrimSpace(bearerToken)
		splitToken := strings.Split(bearerToken, " ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
			c.Abort()
			return
		}
		token := splitToken[1]

		claims := &models.UserJWTClaims{}
		tokenJWT, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.Config.JWTSecret), nil
		})

		if err != nil || !tokenJWT.Valid {
			m.logger.Error("Got error while check signing JWT", slog.Any("error", err))
			c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
			c.Abort()
			return
		}

		if claims.ExpiresAt == nil || claims.ExpiresAt.Unix() < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
			c.Abort()
			return
		}

		user := models.ContextUser{
			ID:       claims.UserID,
			Username: claims.Username,
		}
		c.Set(models.UserContextKey, user)

		c.Next()
	}
}
