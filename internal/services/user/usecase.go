package user

import "jevvonn/bcc-be-freepass-2025/internal/models/dto"

type UserUsecase interface {
	GetUserDetail(userId uint) (dto.GetUserDetailResponse, error)
	GetUserProfile(userId uint) (dto.GetUserProfileResponse, error)
}
