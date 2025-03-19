package model

import "github.com/google/uuid"

type Departement struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Name string    `gorm:"type:varchar(50);unique" json:"name" form:"name" binding:"required"`

	Users []User `gorm:"constraint:OnDelete:SET NULL;" json:"users"`

	Timestamp
}
