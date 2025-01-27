package jwt

import (
	"jevvonn/bcc-be-freepass-2025/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(data jwt.MapClaims) (string, error) {
	config := config.GetConfig()

	var key = []byte(config.GetString("secret.jwt-key"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := t.SignedString(key)

	return tokenString, err
}
