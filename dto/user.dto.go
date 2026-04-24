package dto

import "errors"



type UserGetByEmailReq struct {
	Email string `json:"email" binding:"required, email"`
}

type UserGetByEmailRes struct {
	Email           string  `json:"email"`
	DepartementName *string `json:"departement_name,omitempty"`
}

type UserLoginReq struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRefreshReq struct {
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}

type UserRefreshRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserMeRes struct {
	Email           string `json:"email"`
	DepartementName string `json:"departement_name,omitempty"`
}

const (
	MSG_USER_REGISTER_SUCCESS = "user registered successfully"
	MSG_USER_LOGIN_SUCCESS    = "user login successfully"
	MSG_USER_REFRESH_SUCCESS  = "user refresh token success"
	MSG_AUTH_SUCCESS          = "authorized"

	MSG_USER_REGISTER_FAILED       = "user registration failed"
	MSG_USER_LOGIN_FAILED          = "user login failed"
	MSG_USER_REFRESH_FAILED        = "user refresh token failed"
	MSG_AUTH_FAILED                = "unauthorized"
	MSG_INVALID_TOKEN_FAILED       = "invalid token"
	MSG_METHOD_TOKEN_FAILED        = "unexpected signing method"
	MSG_ACCESS_TOKEN_CREATE_FAILED = "failed to create access token"
	MSG_USER_FORBIDDEN             = "forbidden"

	MSG_USER_FOUND     = "user found"
	MSG_USER_NOT_FOUND = "user not found"
)

var (
	ErrNRPAlreadyExists        = errors.New("nrp already exists")
	ErrUserNotFound            = errors.New("user not found")
	ErrAccessTokenCreateFailed = errors.New("failed to create access token")
	ErrInvalidUserID           = errors.New("invalid user ID in token")
)
