package service

import (
	"context"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
)

type GalleriesService interface {
	CreateGallery(ctx context.Context, req dto.GalleryCreateReq) (model.Gallery, error)
	GetGalleries(ctx context.Context, userRole, userDepartement string) ([]model.Gallery, error)
	GetGalleryByID(ctx context.Context, id uuid.UUID) (model.Gallery, error)
	UpdateGallery(ctx context.Context, id uuid.UUID, req dto.GalleryUpdateReq) (model.Gallery, error)
	DeleteGallery(ctx context.Context, id uuid.UUID) error
}

type galleriesService struct {
	repo repository.GalleriesRepository
}

func NewGalleriesService(repo repository.GalleriesRepository) GalleriesService {
	return &galleriesService{repo: repo}
}

func (s *galleriesService) CreateGallery(ctx context.Context, req dto.GalleryCreateReq) (model.Gallery, error) {
	gallery := model.Gallery{
		Id:           uuid.New(),
		Title:        req.Title,
		Description:  req.Description,
		GDriveLink:   req.GDriveLink,
		ThumbnailUrl: req.ThumbnailUrl,
	}

	if req.EventDate != "" {
		parsed, err := time.Parse("2006-01-02", req.EventDate)
		if err != nil {
			parsed, err = time.Parse(time.RFC3339, req.EventDate)
			if err != nil {
				return model.Gallery{}, err
			}
		}
		gallery.EventDate = model.Timestamp{
			CreatedAt: parsed,
			UpdatedAt: parsed,
		}
	}

	return s.repo.CreateGallery(ctx, gallery)
}

func (s *galleriesService) GetGalleries(ctx context.Context, userRole, userDepartement string) ([]model.Gallery, error) {
	return s.repo.GetGalleries(ctx)
}

func (s *galleriesService) GetGalleryByID(ctx context.Context, id uuid.UUID) (model.Gallery, error) {
	return s.repo.GetGalleryByID(ctx, id)
}

func (s *galleriesService) UpdateGallery(ctx context.Context, id uuid.UUID, req dto.GalleryUpdateReq) (model.Gallery, error) {
	updates := model.Gallery{}
	hasUpdate := false

	if req.Title != nil {
		updates.Title = *req.Title
		hasUpdate = true
	}
	if req.Description != nil {
		updates.Description = *req.Description
		hasUpdate = true
	}
	if req.GDriveLink != nil {
		updates.GDriveLink = *req.GDriveLink
		hasUpdate = true
	}
	if req.ThumbnailUrl != nil {
		updates.ThumbnailUrl = *req.ThumbnailUrl
		hasUpdate = true
	}
	if req.EventDate != nil && *req.EventDate != "" {
		parsed, err := time.Parse("2006-01-02", *req.EventDate)
		if err != nil {
			parsed, err = time.Parse(time.RFC3339, *req.EventDate)
			if err != nil {
				return model.Gallery{}, err
			}
		}
		updates.EventDate = model.Timestamp{
			CreatedAt: parsed,
			UpdatedAt: parsed,
		}
		hasUpdate = true
	}

	if !hasUpdate {
		return s.repo.GetGalleryByID(ctx, id)
	}

	return s.repo.UpdateGallery(ctx, id, updates)
}

func (s *galleriesService) DeleteGallery(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteGallery(ctx, id)
}
