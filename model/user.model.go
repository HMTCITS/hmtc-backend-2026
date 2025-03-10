package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Username string    `gorm:"type:varchar(255);unique" json:"username" form:"username" binding:"required"`
	Email    string    `gorm:"type:varchar(255);unique" json:"email" form:"email" binding:"required"`
	Password string    `gorm:"type:varchar(255)" json:"password" form:"password" binding:"required"`
}