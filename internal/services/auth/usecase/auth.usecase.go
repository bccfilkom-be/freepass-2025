package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/auth"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"
)

type AuthUsecase struct {
	userRepository user.UserRepository
}

func NewAuthUsecase(userRepository user.UserRepository) auth.AuthUsecase {
	return &AuthUsecase{userRepository}
}

func (v *AuthUsecase) SignUp(req *dto.SignUpRequest) error {
	result, _ := v.userRepository.GetByEmail(req.Email)

	if result.ID != 0 {
		return errors.New("User already exist!")
	}

	return v.userRepository.Create(domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.CreatePassword(req.Password),
	})
}
