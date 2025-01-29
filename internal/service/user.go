package service

import (
	"errors"
	"freepass-bcc/entity"
	"freepass-bcc/internal/repository"
	"freepass-bcc/model"
	"freepass-bcc/pkg/bcrypt"
	"freepass-bcc/pkg/jwt"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(param *model.UserRegister) error
	Login(param model.UserLogin) (model.LoginResponse, error)
	UpdateProfile(id string, request *model.UpdateProfile) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	AddNewEC(id string) error
	RemoveRole(id string) error
	GetUser(param model.UserParam) (entity.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		UserRepository: userRepository,
		bcrypt:         bcrypt,
		jwtAuth:        jwtAuth,
	}
}

func (us *UserService) Register(param *model.UserRegister) error {
	hash, err := us.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	user := &entity.User{
		UserID:   id,
		Name:     param.Name,
		Email:    param.Email,
		Password: hash,
		RoleID:   3,
	}

	_, err = us.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(param model.UserLogin) (model.LoginResponse, error) {
	var result model.LoginResponse

	user, err := us.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = us.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := us.jwtAuth.CreateJWTToken(user.UserID)
	if err != nil {
		return result, errors.New("failed to create jwt")
	}

	result.UserID = user.UserID
	result.Token = token
	result.RoleID = user.RoleID

	return result, nil
}

func (us *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return us.UserRepository.GetUser(param)
}

func (us *UserService) UpdateProfile(id string, request *model.UpdateProfile) (*entity.User, error) {
	user, err := us.UserRepository.UpdateProfile(id, request)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) AddNewEC(id string) error {
	err := us.UserRepository.AddNewEC(id)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) RemoveRole(id string) error {
	err := us.UserRepository.RemoveRole(id)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByID(id string) (*entity.User, error) {
	user, err := us.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
