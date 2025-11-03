package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(id uint, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	exp := time.Now().Add(time.Hour * 24).Unix() // 24 jam

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
