package service

import (
	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
)

type UserService interface {
	Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error) {
	isUsername, err := us.userRepo.IsUsernameExist(userReq.Username)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}
	if isUsername {
		return dto.UserRegisterRes{}, dto.ErrUsernameAlreadyExists
	}

	isEmail, err := us.userRepo.IsEmailExist(userReq.Email)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	if isEmail {
		return dto.UserRegisterRes{}, dto.ErrEmailAlreadyExists
	}

	password, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	user := model.User{
		Email:    userReq.Email,
		Username: userReq.Username,
		Password: password,
	}

	usr, err := us.userRepo.Register(user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	return dto.UserRegisterRes{
		Email:    usr.Email,
		Username: usr.Username,
	}, nil
}
