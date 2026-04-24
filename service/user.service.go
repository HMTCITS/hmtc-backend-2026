package service

import (
	"errors"
	"strings"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"gorm.io/gorm"
)

type UserService interface {
	// Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error)
	Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error)
	GetUserByEmail(userReq dto.UserGetByEmailReq) (dto.UserGetByEmailRes, error)
	Me(userId string) (dto.UserMeRes, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

// func (us *userService) Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error) {
// 	isEmail, err := us.userRepo.IsEmailExist(userReq.NRP)
// 	if err != nil {
// 		return dto.UserRegisterRes{}, err
// 	}
// 	if isEmail {
// 		return dto.UserRegisterRes{}, dto.ErrNRPAlreadyExists
// 	}

// 	var departementId *uuid.UUID
// 	if userReq.DepartementName != nil && *userReq.DepartementName != "" {
// 		var departement *model.Departement
// 		departement, err = us.userRepo.FindDepartementByName(*userReq.DepartementName)
// 		if err != nil {
// 			return dto.UserRegisterRes{}, dto.ErrDepartementNotFound
// 		}
// 		departementId = &departement.Id
// 	}

// 	user := model.User{
// 		NRP:           userReq.NRP,
// 		DepartementId: departementId,
// 	}

// 	usr, err := us.userRepo.Register(user)
// 	if err != nil {
// 		return dto.UserRegisterRes{}, err
// 	}

// 	return dto.UserRegisterRes{
// 		NRP:             usr.NRP,
// 		DepartementName: userReq.DepartementName,
// 	}, nil
// }

func (us *userService) Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error) {
	if !strings.HasSuffix(userReq.Email, "@hmtc-its.com") {
		return dto.UserLoginRes{}, dto.ErrInvalidUserID
	}

	isUser, err := us.userRepo.FindUserByEmail(userReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserLoginRes{}, dto.ErrUserNotFound
		}
		return dto.UserLoginRes{}, err
	}

	if !utils.CheckPasswordHash(userReq.Password, isUser.PasswordHash) {
		return dto.UserLoginRes{}, dto.ErrUserNotFound
	}

	accessToken, err := utils.GenerateToken(isUser.Id, string(isUser.Role), isUser.DepartmentName)
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	refreshToken, err := utils.GenerateRefreshToken(isUser.Id, string(isUser.Role), isUser.DepartmentName)
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	user := dto.UserLoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return user, nil
}

func (us *userService) GetUserByEmail(userReq dto.UserGetByEmailReq) (dto.UserGetByEmailRes, error) {
	user, err := us.userRepo.FindUserByEmail(userReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserGetByEmailRes{}, dto.ErrUserNotFound
		}
		return dto.UserGetByEmailRes{}, err
	}

	var departementName *string
	if user.DepartmentName != "" {
		departementName = &user.DepartmentName
	}

	return dto.UserGetByEmailRes{
		Email:           user.Email,
		DepartementName: departementName,
	}, nil
}

func (us *userService) Me(userId string) (dto.UserMeRes, error) {
	user, err := us.userRepo.FindUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserMeRes{}, dto.ErrUserNotFound
		}
		return dto.UserMeRes{}, err
	}

	departementName := user.DepartmentName

	return dto.UserMeRes{
		Email:           user.Email,
		DepartementName: departementName,
	}, nil
}
