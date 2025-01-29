package lib

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// GenerateVerificationToken creates a random UUID token for email verification
func GenerateVerificationToken() (uuid.UUID, error) {
	return uuid.New(), nil
}

// ParseUUID parses a string into a UUID, returning ErrInvalidToken if invalid
func ParseUUID(token string) (uuid.UUID, error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return uuid.Nil, errors.New("invalid or expired token")
	}
	return id, nil
}
