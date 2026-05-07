package repository

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GalleriesRepository interface {
	CreateGallery(ctx context.Context, gallery model.Gallery) (model.Gallery, error)
	GetGalleries(ctx context.Context, visibilities []model.GalleryVisibility) ([]model.Gallery, error)
	GetGalleryByID(ctx context.Context, galleryId uuid.UUID) (model.Gallery, error)
	UpdateGallery(ctx context.Context, galleryId uuid.UUID, gallery model.Gallery) (model.Gallery, error)
	DeleteGallery(ctx context.Context, galleryId uuid.UUID) error
}

type galleriesRepository struct {
	db *gorm.DB
}

func NewGalleriesRepository(db *gorm.DB) GalleriesRepository {
	return &galleriesRepository{
		db: db,
	}
}

func (r *galleriesRepository) CreateGallery(ctx context.Context, gallery model.Gallery) (model.Gallery, error) {
	if err := r.db.WithContext(ctx).Create(&gallery).Error; err != nil {
		return model.Gallery{}, err
	}
	return gallery, nil
}

func (r *galleriesRepository) GetGalleries(ctx context.Context, visibilities []model.GalleryVisibility) ([]model.Gallery, error) {
	var galleries []model.Gallery

	query := r.db.WithContext(ctx)
	if len(visibilities) > 0 {
		query = query.Where("visibility IN ?", visibilities)
	}

	if err := query.Order("created_at DESC").Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}

func (r *galleriesRepository) GetGalleryByID(ctx context.Context, galleryId uuid.UUID) (model.Gallery, error) {
	var gallery model.Gallery
	if err := r.db.WithContext(ctx).Where("id = ?", galleryId).First(&gallery).Error; err != nil {
		return model.Gallery{}, err
	}
	return gallery, nil
}

func (r *galleriesRepository) UpdateGallery(ctx context.Context, galleryId uuid.UUID, gallery model.Gallery) (model.Gallery, error) {
	var existing model.Gallery
	if err := r.db.WithContext(ctx).Where("id = ?", galleryId).First(&existing).Error; err != nil {
		return model.Gallery{}, err
	}

	if err := r.db.WithContext(ctx).Model(&existing).Updates(gallery).Error; err != nil {
		return model.Gallery{}, err
	}

	if err := r.db.WithContext(ctx).Where("id = ?", galleryId).First(&existing).Error; err != nil {
		return model.Gallery{}, err
	}

	return existing, nil
}

func (r *galleriesRepository) DeleteGallery(ctx context.Context, galleryId uuid.UUID) error {
	result := r.db.WithContext(ctx).Where("id = ?", galleryId).Delete(&model.Gallery{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
