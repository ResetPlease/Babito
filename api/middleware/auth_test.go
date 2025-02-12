package middleware_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ResetPlease/Babito/internal/models"
	testcore "github.com/ResetPlease/Babito/internal/test_core"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthMiddleware(t *testing.T) {
	testCore := testcore.NewTestCore()

	authCreds := models.AuthRequest{
		Username: "test",
		Password: "test",
	}
	authBody, err := json.Marshal(authCreds)
	assert.Equal(t, err, nil)

	router := gin.Default()
	authRouter := router.Group("/api")
	{
		authRouter.POST("/auth", testCore.Handler.AuthHandler)
	}
	secure := router.Use(testCore.Middleware.AuthMiddleware())
	{
		secure.GET("/auth/middleware", func(c *gin.Context) {
			// test for GetUserFromContext
			user, err := tools.GetUserFromContext(c)
			assert.Equal(t, err, nil)
			assert.Equal(t, user.Username, authCreds.Username)

			c.JSON(http.StatusOK, models.MessageOK)
		})
	}

	t.Run("test_authorization_ok", func(t *testing.T) {
		authRequest, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(string(authBody)))
		authResponse := httptest.NewRecorder()
		assert.Equal(t, err, nil)
		router.ServeHTTP(authResponse, authRequest)

		var responseToken models.AuthResponse
		err = json.Unmarshal(authResponse.Body.Bytes(), &responseToken)
		assert.Equal(t, err, nil)
		assert.NotEqual(t, responseToken.Token, nil)

		request, err := http.NewRequest(http.MethodGet, "/auth/middleware", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *responseToken.Token))
		assert.Equal(t, err, nil)

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, request)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		expected := `{"Message":"OK"}`
		assert.Equal(t, rr.Body.String(), expected)
	})

	t.Run("test_unauthorized_with_fake_token", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/auth/middleware", nil)
		request.Header.Set("Authorization", "Bearer jwt")
		assert.Equal(t, err, nil)

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, request)

		assert.Equal(t, rr.Result().StatusCode, http.StatusUnauthorized)
	})

	t.Run("test_unauthorized_expired_token", func(t *testing.T) {
		expired := time.Now().Add(-1 * time.Hour)
		claims := models.UserJWTClaims{
			UserID:   123,
			Username: "test",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expired),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "babito",
				Subject:   "user-auth",
				ID:        "token-id",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, err := token.SignedString([]byte(testCore.Middleware.Config.JWTSecret))
		assert.Equal(t, err, nil)

		request, err := http.NewRequest(http.MethodGet, "/auth/middleware", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signedToken))
		assert.Equal(t, err, nil)

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, request)

		assert.Equal(t, rr.Result().StatusCode, http.StatusUnauthorized)
	})

}
