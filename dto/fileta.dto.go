package dto

import "github.com/HMTCITS/hmtc-backend-2025/model"

const (
	FileNotCreate = "error cannot create file"
)

type CreateFileTA struct {
	StudentName string `json:"student_name" form:"student_name" binding:"required"`
	NRP         string `json:"nrp" form:"nrp" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	Semester    uint   `json:"semester" form:"semester" binding:"required"`
	DosPem      string `json:"dosen_pembimbing" form:"dosen_pembimbing" binding:"required"`
}

type ChangeFileStatus struct {
	FileId string           `json:"file_id" form:"file_id" binding:"required"`
	Status model.FileStatus `json:"file_status" form:"file_status" binding:"required"`
}

type GetAllFiles struct {
	Id          string `json:"id"`
	FileName    string `json:"file_name"`
	StudentName string `json:"nama_mahasiswa"`
	NRP         string `json:"nrp"`
	Email       string `json:"email"`
	Semester    uint   `json:"semester"`
	DosPem      string `json:"dosen_pembimbing"`
}
