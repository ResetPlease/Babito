package models

import "github.com/golang-jwt/jwt/v5"

type UserJWTClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Config struct {
	JWTSecret          string
	DefaultUserBalance int64
}

func NewConfig(JWTSecret string, DefaultUserBalance int64) *Config {
	return &Config{
		JWTSecret:          JWTSecret,
		DefaultUserBalance: DefaultUserBalance,
	}
}
