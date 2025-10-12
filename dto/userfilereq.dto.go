package dto

type UserFileReqDto struct {
	Name      string `json:"name" form:"name" binding:"required"`
	NRP       string `json:"nrp" form:"nrp" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	AlasanReq string `json:"alasan" form:"alasan" binding:"required"`
}

type UserFileResDto struct {
}
