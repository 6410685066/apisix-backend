package service

import (
	"apisix-backend/config"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint, username string) (string, error) {
	jwtSecret := config.GetJWTSecret()
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(30 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
