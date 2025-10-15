package model

import "github.com/google/uuid"

type TokenFile struct {
	RefreshToken string `json:"refresh_token"`
}

type OAuthToken struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	RefreshToken string    `gorm:"type:text"`

	Timestamp
}
