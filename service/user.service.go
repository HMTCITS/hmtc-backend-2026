package service

import (
	"errors"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	myjwt "github.com/HMTCITS/hmtc-backend-2025/pkg/jwt"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error)
	Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error)
	RefreshToken(userReq dto.UserRefreshReq) (dto.UserRefreshRes, error)
	GetUserByNRP(userReq dto.UserGetByNRPReq) (dto.UserGetByNRPRes, error)
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
	isNRP, err := us.userRepo.IsNRPExist(userReq.NRP)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}
	if isNRP {
		return dto.UserRegisterRes{}, dto.ErrNRPAlreadyExists
	}

	var departementId *uuid.UUID
	if userReq.DepartementName != nil && *userReq.DepartementName != "" {
		var departement *model.Departement
		departement, err = us.userRepo.FindDepartementByName(*userReq.DepartementName)
		if err != nil {
			return dto.UserRegisterRes{}, dto.ErrDepartementNotFound
		}
		departementId = &departement.Id
	}

	user := model.User{
		NRP:           userReq.NRP,
		DepartementId: departementId,
	}

	usr, err := us.userRepo.Register(user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	return dto.UserRegisterRes{
		NRP:             usr.NRP,
		DepartementName: userReq.DepartementName,
	}, nil
}

func (us *userService) Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error) {
	user, err := us.userRepo.FindUserByNRP(userReq.NRP)
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	payload := make(map[string]string)
	payload["id"] = user.Id.String()
	payload["nrp"] = user.NRP
	payload["role"] = string(user.Role)
	payload["departemen_id"] = user.DepartementId.String()
	payload["departemen_name"] = user.Departement.Name

	accessTokentoken, err := myjwt.GenerateToken(payload, "Access")
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	refreshToken, err := myjwt.GenerateToken(payload, "Refresh")
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	return dto.UserLoginRes{AccessToken: accessTokentoken, RefreshToken: refreshToken}, nil
}

func (us *userService) RefreshToken(userReq dto.UserRefreshReq) (dto.UserRefreshRes, error) {
	payload, err := myjwt.GetPayload(userReq.RefreshToken)
	if err != nil {
		return dto.UserRefreshRes{}, err
	}

	newPayload := make(map[string]string)
	newPayload["id"] = payload["id"]
	newPayload["nrp"] = payload["nrp"]
	newPayload["role"] = payload["role"]
	newPayload["departemen_id"] = payload["departemen_id"]
	newPayload["departemen_name"] = payload["departemen_name"]

	newAccessToken, err := myjwt.GenerateToken(newPayload, "Access")
	if err != nil {
		return dto.UserRefreshRes{}, err
	}

	newRefreshToken, err := myjwt.GenerateToken(newPayload, "Refresh")
	if err != nil {
		return dto.UserRefreshRes{}, err
	}

	return dto.UserRefreshRes{AccessToken: newAccessToken, RefreshToken: newRefreshToken}, nil
}

func (us *userService) GetUserByNRP(userReq dto.UserGetByNRPReq) (dto.UserGetByNRPRes, error) {
	user, err := us.userRepo.FindUserByNRP(userReq.NRP)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserGetByNRPRes{}, dto.ErrUserNotFound
		}
		return dto.UserGetByNRPRes{}, err
	}

	var departementName *string
	if user.Departement != nil {
		departementName = &user.Departement.Name
	}

	return dto.UserGetByNRPRes{
		NRP:             user.NRP,
		DepartementName: departementName,
	}, nil
}
