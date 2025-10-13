package service

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
)

type UserFileReqService interface {
	NewUserFileReq(ctx context.Context, req dto.UserFileReqDto) error
	ChangeUserFileReqStatus(ctx context.Context, req dto.ChangeUserFileReqDto) error
	UserFileReqStatus(ctx context.Context, reqId string) (dto.UserFileResDto, error)
}

type userFileReqService struct {
	userFileRepo repository.UserFileReqRepository
}

func NewUserFileService(userFileRepo repository.UserFileReqRepository) UserFileReqService {
	return &userFileReqService{userFileRepo: userFileRepo}
}

func (s *userFileReqService) NewUserFileReq(ctx context.Context, req dto.UserFileReqDto) error {

	userReq, err := model.NewUserFileReq(req.Name, req.NRP, req.Email, req.AlasanReq, uuid.MustParse(req.FileId))
	if err != nil {
		return err
	}
	er := s.userFileRepo.NewUserFileReq(ctx, nil, userReq)
	if er != nil {
		return er
	}
	return nil
}

func (s *userFileReqService) ChangeUserFileReqStatus(ctx context.Context, req dto.ChangeUserFileReqDto) error {
	err := s.userFileRepo.ChangeReqStatus(ctx, nil, uuid.MustParse(req.ReqId), req.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *userFileReqService) UserFileReqStatus(ctx context.Context, reqId string) (dto.UserFileResDto, error) {

	status, err := s.userFileRepo.UserFileStatus(ctx, nil, uuid.MustParse(reqId))
	if err != nil {
		return dto.UserFileResDto{}, err
	}

	var link *string
	if status == model.StatusApproved {
		fileId, err := s.userFileRepo.GetFileIdByReqId(ctx, nil, uuid.MustParse(reqId))
		if err != nil {
			return dto.UserFileResDto{}, err
		}
		tmp := "http://localhost:8080/api/file-ta/download/" + reqId + "/" + fileId
		link = &tmp
	}

	return dto.UserFileResDto{
		Status: string(status),
		Link:   link,
	}, nil
}
