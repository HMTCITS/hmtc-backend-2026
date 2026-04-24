package model

import "github.com/google/uuid"

type Gallery struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title        string    `gorm:"type:varchar(100);not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	EventDate    Timestamp `gorm:"embedded;embeddedPrefix:event_" json:"event_date"`
	GDriveLink   string    `gorm:"type:text" json:"gdrive_link"`
	ThumbnailUrl string    `gorm:"type:text" json:"thumbnail_url"`

	Timestamp
}
