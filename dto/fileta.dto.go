package dto

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
