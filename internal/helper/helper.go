package helper

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func ComparePassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

func StringISOToDateTime(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}
