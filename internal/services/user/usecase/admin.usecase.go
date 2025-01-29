package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"

	"github.com/gin-gonic/gin"
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

func (v *AdminUsecase) DeleteUser(ctx *gin.Context, userId uint) error {
	loggedInUserId := ctx.GetUint("userId")

	if loggedInUserId == userId {
		return errors.New("You can't delete yourself!")
	}

	user, err := v.userRepository.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found!")
		} else {
			return err
		}
	}

	if user.Role == constant.ROLE_ADMIN {
		return errors.New("You can't delete other admin!")
	}

	err = v.userRepository.Delete(userId)
	if err != nil {
		return err
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
