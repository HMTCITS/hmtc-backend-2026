package model

import "github.com/google/uuid"

type UserRole string

const (
	Guest UserRole = "user"
	Admin UserRole = "admin"
)

type User struct {
	Id  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	NRP string    `gorm:"type:varchar(35);unique" json:"nrp" form:"nrp" binding:"required"`

	DepartementId *uuid.UUID   `gorm:"type:uuid" json:"departement_id"`
	Departement   *Departement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"departement"`

	Role UserRole `gorm:"type:varchar(10);default:'user'" json:"role" form:"role" binding:"required"`

	Timestamp
}
