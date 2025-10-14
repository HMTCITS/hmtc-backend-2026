package dto

import "github.com/HMTCITS/hmtc-backend-2025/model"

type UserFileReqDto struct {
	FileId    string `json:"file_id" form:"file_id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	NRP       string `json:"nrp" form:"nrp" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	AlasanReq string `json:"alasan" form:"alasan" binding:"required"`
}

type ChangeUserFileReqDto struct {
	ReqId  string       `json:"req_id" form:"req_id" binding:"required"`
	Status model.Status `json:"status" form:"status" binding:"required"`
}

type UserFileResDto struct {
	Status string  `json:"status" form:"status" binding:"required"`
	Link   *string `json:"link"`
}
