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

type UserGetByNRPReq struct {
	NRP string `json:"nrp"`
}

type UserGetByNRPRes struct {
	NRP             string  `json:"nrp"`
	DepartementName *string `json:"departement_name,omitempty"`
}

type UserLoginReq struct {
	NRP string `json:"nrp" form:"nrp" binding:"required"`
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
	NRP             string `json:"nrp"`
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
