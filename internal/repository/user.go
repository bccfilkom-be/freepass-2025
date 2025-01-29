package repository

import (
	"freepass-bcc/entity"
	"freepass-bcc/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateProfile(id string, request *model.UpdateProfile) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	AddNewEC(id string) error
	RemoveRole(id string) error
	GetUser(param model.UserParam) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) UpdateProfile(id string, request *model.UpdateProfile) (*entity.User, error) {
	tx := u.db.Begin()
	var user entity.User

	userID, err := uuid.Parse(id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user.UserID = userID

	profileParse := *parseUpdateProfile(request, &user)
	err = tx.Debug().Updates(&profileParse).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil

}

func (u *UserRepository) GetUserByID(id string) (*entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) AddNewEC(id string) error {
	tx := u.db.Begin()
	var user entity.User

	err := tx.Debug().Model(&user).Where("user_id = ?", id).Update("role_id", 2).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil

}

func (u *UserRepository) RemoveRole(id string) error {
	tx := u.db.Begin()
	var user entity.User

	err := tx.Debug().Model(&user).Where("user_id = ?", id).Update("role_id", 4).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func parseUpdateProfile(model *model.UpdateProfile, user *entity.User) *entity.User {
	if model.Name != "" {
		user.Name = model.Name
	}

	if model.Email != "" {
		user.Email = model.Email
	}

	if model.Address != "" {
		user.Address = model.Address
	}

	return user
}
