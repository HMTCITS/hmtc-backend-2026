package controller

import "github.com/HMTCITS/hmtc-backend-2025/service"

type UserFileReqController interface {
	NewUserFileReq()
	UserFileStatus()
}

type userFileReqController struct {
	userFileService service.UserFileReqService
}
