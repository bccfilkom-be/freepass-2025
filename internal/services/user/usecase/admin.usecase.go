package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"

	"gorm.io/gorm"
)

type AdminUsecase struct {
	userRepository user.UserRepository
}

func NewAdminUsecase(userRepository user.UserRepository) user.AdminUsecase {
	return &AdminUsecase{
		userRepository,
	}
}

func (v *AdminUsecase) DeleteUser(userId uint) error {
	err := v.userRepository.Delete(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found!")
		} else {
			return err
		}
	}

	return nil
}

func (v *AdminUsecase) UpdateRole(userId uint, req *dto.UpdateUserRoleRequest) error {
	user, err := v.userRepository.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found!")
		} else {
			return err
		}
	}

	return v.userRepository.Update(domain.User{
		ID:   user.ID,
		Role: req.Role,
	})
}
