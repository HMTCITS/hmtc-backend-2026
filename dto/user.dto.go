package dto

import "errors"

type UserRegisterReq struct {
	NRP             string  `json:"nrp" form:"nrp" binding:"required"`
	DepartementName *string `json:"departement_name,omitempty" form:"departement_name,omitempty"`
}

type UserRegisterRes struct {
	NRP             string  `json:"nrp"`
	DepartementName *string `json:"departement_name,omitempty"`
}

const (
	MSG_USER_REGISTER_SUCCESS = "user registered successfully"

	MSG_USER_REGISTER_FAILED = "user registration failed"

	MSG_USER_NOT_FOUND = "user not found"
)

var (
	ErrNRPAlreadyExists = errors.New("nrp already exists")
	ErrUserNotFound     = errors.New("user not found")
)
