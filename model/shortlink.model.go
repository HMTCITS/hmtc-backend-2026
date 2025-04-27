package model

import (
	"fmt"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type LinkShortener struct {
	Fullurl  string `gorm:"type:varchar(255)" json:"link" form:"link" binding:"required,link"`
	Shorturl string `gorm:"type:varchar(255)" json:"shortlink"`
	Click    int8   `gorm:"type:int;default:0" json:"click"`
}

func (l *LinkShortener) BeforeCreate(tx *gorm.DB) (err error) {
	if l.Shorturl == "" {
		sid, _ := shortid.Generate()
		fmt.Print(sid)
		l.Shorturl = sid
	}
	return
}
