package repository

import (
	"errors"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"gorm.io/gorm"
)

type OAuthTokenRepository interface {
	Save(refreshToken string) error
	Get() (string, error)
}

type oauthTokenRepo struct {
	db *gorm.DB
}

func NewOAuthTokenRepo(db *gorm.DB) OAuthTokenRepository {
	return &oauthTokenRepo{db}
}

func (r *oauthTokenRepo) Save(refreshToken string) error {
	var token model.OAuthToken

	// cek apakah sudah ada (karena hanya butuh 1 token global)
	if err := r.db.First(&token).Error; err != nil {
		// jika belum ada, buat baru
		if errors.Is(err, gorm.ErrRecordNotFound) {
			token.RefreshToken = refreshToken
			return r.db.Create(&token).Error
		}
		// error lain
		return err
	}

	// kalau sudah ada, update
	token.RefreshToken = refreshToken
	return r.db.Save(&token).Error
}

func (r *oauthTokenRepo) Get() (string, error) {
	var token model.OAuthToken
	if err := r.db.First(&token).Error; err != nil {
		return "", err
	}
	return token.RefreshToken, nil
}
