package model

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

const (
	STARTSEMESTER uint = 1
	ENDSEMESTER   uint = 14
)

type FileStatus string

const (
	PENDING  FileStatus = "pending"
	APPROVE  FileStatus = "approve"
	REJECTED FileStatus = "rejected"
)

type TAFile struct {
	Id          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FileName    string     `json:"file_name"`
	StudentName string     `json:"nama_mahasiswa" binding:"required"`
	NRP         string     `json:"nrp" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	Semester    uint       `json:"semester" binding:"required"`
	DosPem      string     `json:"dosen_pembimbing" binding:"required"`
	Status      FileStatus `json:"status"`
}

func NewFileTA(id string, filename string, studentName string, nrp string, email string, semester uint, dosPem string) (TAFile, error) {

	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return TAFile{}, errors.New("invalid email")
	}

	if semester < STARTSEMESTER && semester > ENDSEMESTER {
		return TAFile{}, errors.New("invalid semester")
	}

	return TAFile{
		Id:          uuid.MustParse(id),
		FileName:    filename,
		StudentName: studentName,
		NRP:         nrp,
		Email:       email,
		Semester:    semester,
		DosPem:      dosPem,
		Status:      PENDING,
	}, nil
}
