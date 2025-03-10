package dto

import "errors"

type UserRegisterReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

type UserRegisterRes struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

const (
	MSG_USER_REGISTER_SUCCESS = "user registered successfully"

	MSG_USER_REGISTER_FAILED = "user registration failed"

	MSG_USER_NOT_FOUND = "user not found"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrLogin                 = errors.New("invalid email or password")
	ErrUserNotFound          = errors.New("user not found")
)