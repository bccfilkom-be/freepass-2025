package user

import "jevvonn/bcc-be-freepass-2025/internal/models/domain"

type UserRepository interface {
	Create(user domain.User) error
	GetById(userId uint) (*domain.User, error)
	GetByEmail(userEmail string) (*domain.User, error)
}
