package repository

import (
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"gorm.io/gorm"
)

type ShortLinkRepository interface {
	GenerateShortLink(link model.LinkShortener) (model.LinkShortener, error)
	FindByShortUrl(link string) (model.LinkShortener, error)
}

type shortLinkRepository struct {
	DB *gorm.DB
}

func NewShortLinkRepository(db *gorm.DB) ShortLinkRepository {
	return &shortLinkRepository{DB: db}
}

func (r *shortLinkRepository) GenerateShortLink(link model.LinkShortener) (model.LinkShortener, error) {
	if err := r.DB.Create(&link).Error; err != nil {
		return model.LinkShortener{}, err
	}
	return link, nil
}

func (r *shortLinkRepository) FindByShortUrl(shorturl string) (model.LinkShortener, error) {
	var link model.LinkShortener

	if err := r.DB.Where("shorturl = ?", shorturl).First(&link).Error; err != nil {
		return link, err
	}

	//Update click-nya di goroutine
	go func(shorturl string) {
		_ = r.DB.Model(&model.LinkShortener{}).
			Where("shorturl = ?", shorturl).
			Update("click", gorm.Expr("click + ?", 1)).Error
	}(shorturl)

	return link, nil
}
