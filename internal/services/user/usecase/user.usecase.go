package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"

	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) user.UserUsecase {
	return &UserUsecase{
		userRepository,
	}
}

func (v *UserUsecase) GetUserProfile(userId uint) (dto.GetUserProfileResponse, error) {
	user, err := v.userRepository.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.GetUserProfileResponse{}, errors.New("User not found")
		} else {
			return dto.GetUserProfileResponse{}, err
		}
	}

	return dto.GetUserProfileResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (v *UserUsecase) GetUserDetail(userId uint) (dto.GetUserDetailResponse, error) {
	user, err := v.userRepository.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.GetUserDetailResponse{}, errors.New("User not found")
		} else {
			return dto.GetUserDetailResponse{}, err
		}
	}

	return dto.GetUserDetailResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
