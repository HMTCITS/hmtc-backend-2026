package migration

import (
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return err
	}

	err := db.AutoMigrate(
		&model.User{},
		&model.OAuthToken{},
		&model.Departement{},
		&model.LinkShortener{})
	if err != nil {
		return err
	}

	return nil
}
