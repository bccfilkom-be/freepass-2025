package jwt

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(data jwt.MapClaims) (string, error) {
	config := config.GetConfig()
	key := []byte(config.GetString("secret.jwt-key"))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := t.SignedString(key)

	return tokenString, err
}

func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	config := config.GetConfig()
	key := []byte(config.GetString("secret.jwt-key"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return jwt.MapClaims{}, err
	}
}
