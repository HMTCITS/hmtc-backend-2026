package service

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GalleryService interface {
	CreateGallery(ctx context.Context, req dto.GalleryCreateReq, createdBy string) (dto.GalleryResponse, error)
	GetGalleries(ctx context.Context, userNRP string, userDeptName string, userRole string, filters dto.GalleryFilterParams) ([]dto.GalleryResponse, error)
	GetGalleryByID(ctx context.Context, id string) (dto.GalleryResponse, error)
	UpdateGallery(ctx context.Context, id string, req dto.GalleryUpdateReq) (dto.GalleryResponse, error)
	DeleteGallery(ctx context.Context, id string) error
}

type galleryService struct {
	repo repository.GalleriesRepository
}

func NewGalleryService(repo repository.GalleriesRepository) GalleryService {
	return &galleryService{repo: repo}
}

func validateGDriveLink(link string) error {
	parsed, err := url.ParseRequestURI(link)
	if err != nil {
		return dto.ErrGalleryInvalidGDriveLink
	}
	host := strings.ToLower(parsed.Host)
	if !strings.Contains(host, "drive.google.com") && !strings.Contains(host, "docs.google.com") {
		return dto.ErrGalleryInvalidGDriveLink
	}
	return nil
}

func galleryToResponse(g model.Gallery) dto.GalleryResponse {
	resp := dto.GalleryResponse{
		Id:           g.Id.String(),
		Title:        g.Title,
		Description:  g.Description,
		GDriveLink:   g.GDriveLink,
		ThumbnailUrl: g.ThumbnailUrl,
		CreatedAt:    g.Timestamp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    g.Timestamp.UpdatedAt.Format(time.RFC3339),
	}

	if !g.EventDate.CreatedAt.IsZero() {
		resp.EventDate = g.EventDate.CreatedAt.Format("2006-01-02")
	}

	return resp
}

func (s *galleryService) CreateGallery(ctx context.Context, req dto.GalleryCreateReq, createdBy string) (dto.GalleryResponse, error) {
	if err := validateGDriveLink(req.GDriveLink); err != nil {
		return dto.GalleryResponse{}, err
	}

	gallery := model.Gallery{
		Title:        req.Title,
		Description:  req.Description,
		GDriveLink:   req.GDriveLink,
		ThumbnailUrl: req.ThumbnailUrl,
	}

	if req.EventDate != "" {
		t, err := time.Parse("2006-01-02", req.EventDate)
		if err != nil {
			return dto.GalleryResponse{}, err
		}
		gallery.EventDate = model.Timestamp{CreatedAt: t, UpdatedAt: t}
	}

	created, err := s.repo.CreateGallery(ctx, gallery)
	if err != nil {
		return dto.GalleryResponse{}, err
	}

	return galleryToResponse(created), nil
}

func (s *galleryService) GetGalleries(ctx context.Context, userNRP string, userDeptName string, userRole string, filters dto.GalleryFilterParams) ([]dto.GalleryResponse, error) {
	var galleries []model.Gallery
	var err error

	galleries, err = s.repo.GetGalleries(ctx, filters)

	if err != nil {
		return nil, err
	}

	responses := make([]dto.GalleryResponse, len(galleries))
	for i, g := range galleries {
		responses[i] = galleryToResponse(g)
	}

	return responses, nil
}

func (s *galleryService) GetGalleryByID(ctx context.Context, id string) (dto.GalleryResponse, error) {
	galleryId, err := uuid.Parse(id)
	if err != nil {
		return dto.GalleryResponse{}, dto.ErrGalleryInvalidID
	}

	gallery, err := s.repo.GetGalleryByID(ctx, galleryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.GalleryResponse{}, dto.ErrGalleryNotFound
		}
		return dto.GalleryResponse{}, err
	}

	return galleryToResponse(gallery), nil
}

func (s *galleryService) UpdateGallery(ctx context.Context, id string, req dto.GalleryUpdateReq) (dto.GalleryResponse, error) {
	galleryId, err := uuid.Parse(id)
	if err != nil {
		return dto.GalleryResponse{}, dto.ErrGalleryInvalidID
	}

	// Build partial update model
	updateData := model.Gallery{}

	if req.Title != nil {
		updateData.Title = *req.Title
	}
	if req.Description != nil {
		updateData.Description = *req.Description
	}
	if req.GDriveLink != nil {
		if err := validateGDriveLink(*req.GDriveLink); err != nil {
			return dto.GalleryResponse{}, err
		}
		updateData.GDriveLink = *req.GDriveLink
	}
	if req.ThumbnailUrl != nil {
		if *req.ThumbnailUrl == "" {
			return dto.GalleryResponse{}, dto.ErrGalleryThumbnailRequired
		}
		updateData.ThumbnailUrl = *req.ThumbnailUrl
	}
	if req.EventDate != nil && *req.EventDate != "" {
		t, err := time.Parse("2006-01-02", *req.EventDate)
		if err != nil {
			return dto.GalleryResponse{}, err
		}
		updateData.EventDate = model.Timestamp{CreatedAt: t, UpdatedAt: t}
	}

	updated, err := s.repo.UpdateGallery(ctx, galleryId, updateData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.GalleryResponse{}, dto.ErrGalleryNotFound
		}
		return dto.GalleryResponse{}, err
	}

	return galleryToResponse(updated), nil
}

func (s *galleryService) DeleteGallery(ctx context.Context, id string) error {
	galleryId, err := uuid.Parse(id)
	if err != nil {
		return dto.ErrGalleryInvalidID
	}

	err = s.repo.DeleteGallery(ctx, galleryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.ErrGalleryNotFound
		}
		return err
	}

	return nil
}
