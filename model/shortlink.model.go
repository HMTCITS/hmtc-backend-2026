package model

type LinkShortener struct {
	Fullurl  string `gorm:"type:varchar(255)" json:"link" form:"link" binding:"required,link"`
	Shorturl string `gorm:"type:varchar(255)" json:"shortlink"`
	Click    int8   `gorm:"type:int;default:0" json:"click"`
}
