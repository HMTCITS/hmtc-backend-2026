package model

import "github.com/google/uuid"

type UserRole string

const (
	Guest UserRole = "user"
	Admin UserRole = "admin"
)

type User struct {
	Id  			uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Email          string    `gorm:"type:varchar(50);unique" json:"email" form:"email" binding:"required"`
	PasswordHash   string    `gorm:"type:text;not null" json:"-"`
	DepartmentName string    `gorm:"type:varchar(50)" json:"department_name"`

	Role UserRole `gorm:"type:varchar(10);default:'user'" json:"role" form:"role" binding:"required"`

	Timestamp
}
