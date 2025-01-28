package helper

import (
	"strconv"
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

func StringToUint(value string) (uint, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(result), nil
}
