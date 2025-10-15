package model

type TokenFile struct {
	RefreshToken string `json:"refresh_token"`
}

type OAuthToken struct {
	ID           uint   `gorm:"primaryKey"`
	RefreshToken string `gorm:"type:text"`

	Timestamp
}
