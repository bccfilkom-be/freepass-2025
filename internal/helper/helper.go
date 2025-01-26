package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}
