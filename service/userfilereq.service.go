package service

import "github.com/HMTCITS/hmtc-backend-2025/repository"

type UserFileReqService interface {
	NewUserFileReq()
	UserFileStatus()
}

type userFileReqService struct {
	userFileRepo repository.UserFileReqRepository
}
