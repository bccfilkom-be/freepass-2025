package repository

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepository{db}
}

func (v *UserRepository) Create(user domain.User) error {
	return v.db.Create(&user).Error
}

func (v *UserRepository) GetById(userId uint) (*domain.User, error) {
	var user *domain.User
	err := v.db.First(&user, userId).Error
	return user, err
}

func (v *UserRepository) GetByEmail(userEmail string) (*domain.User, error) {
	var user *domain.User
	err := v.db.Where("email = ?", userEmail).First(&user).Error
	return user, err
}

func (v *UserRepository) Update(data domain.User) error {
	if data.ID == 0 {
		return errors.New("User id not found!")
	}

	err := v.db.Updates(data).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found!")
		} else {
			return err
		}
	}

	return nil
}
