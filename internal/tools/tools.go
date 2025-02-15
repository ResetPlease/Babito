package tools

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func ParseJWTToken(token string, config models.Config) (*models.ContextUser, error) {
	claims := &models.UserJWTClaims{}
	tokenJWT, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTSecret), nil
	})

	if err != nil || !tokenJWT.Valid {
		return nil, errors.New("token invalid")
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	user := models.ContextUser{
		ID:       claims.UserID,
		Username: claims.Username,
	}
	return &user, nil
}

func GenerateHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
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
