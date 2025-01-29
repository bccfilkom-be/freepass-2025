package user

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetUserDetail(userId uint) (dto.GetUserDetailResponse, error)
	GetUserProfile(userId uint) (dto.GetUserProfileResponse, error)
	UpdateUserProfile(userId uint, data *dto.UpdateUserProfileRequest) error
}

type AdminUsecase interface {
	DeleteUser(ctx *gin.Context, userId uint) error
	UpdateRole(userId uint, data *dto.UpdateUserRoleRequest) error
}
