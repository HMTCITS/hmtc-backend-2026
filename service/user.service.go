package service

import (
	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
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
