package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	jwt_helper "jevvonn/bcc-be-freepass-2025/internal/helper/jwt"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/auth"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (v *AuthUsecase) SignIn(req *dto.SignInRequest) (dto.SignInResponse, error) {
	user, _ := v.userRepository.GetByEmail(req.Email)

	if user.ID == 0 {
		return dto.SignInResponse{}, errors.New("Invalid Email or Password!")
	}

	if !helper.ComparePassword(user.Password, req.Password) {
		return dto.SignInResponse{}, errors.New("Invalid Email or Password!")
	}

	tokenString, err := jwt_helper.CreateJWTToken(jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 48).Unix(),
	})

	if err != nil {
		log.Fatal(err)
		return dto.SignInResponse{}, err
	}

	return dto.SignInResponse{
		Token: tokenString,
	}, nil
}
