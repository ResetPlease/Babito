package tools

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetenvWithPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("empty value for key=%s", key))
	}
	return value
}

func GenerateJWTToken(userID uint64, username string, secretKey string) (string, error) {
	claims := models.UserJWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "babito",
			Subject:   "user-auth",
			ID:        "token-id",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GetUserFromContext(c *gin.Context) (*models.ContextUser, error) {
	rawUser, ok := c.Get(models.UserContextKey)
	if !ok {
		return nil, errors.New("user not exist in context")
	}
	user, ok := rawUser.(models.ContextUser)
	if !ok {
		return nil, errors.New("wrong user structure in context")
	}
	return &user, nil
}
