package service

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
)

type UserFileReqService interface {
	NewUserFileReq(ctx context.Context, req dto.UserFileReqDto) error
	UserFileStatus(ctx context.Context) (string, error)
}

type userFileReqService struct {
	userFileRepo repository.UserFileReqRepository
}

func NewUserFileService(userFileRepo repository.UserFileReqRepository) UserFileReqService {
	return &userFileReqService{userFileRepo: userFileRepo}
}

func (s *userFileReqService) NewUserFileReq(ctx context.Context, req dto.UserFileReqDto) error {

	userReq, err := model.NewUserFileReq(req.Name, req.NRP, req.Email, req.AlasanReq)
	if err != nil {
		return err
	}
	er := s.userFileRepo.NewUserFileReq(ctx, nil, userReq)
	if er != nil {
		return er
	}
	return nil
}

func (s *userFileReqService) UserFileStatus(ctx context.Context) (string, error) {

}
