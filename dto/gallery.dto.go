package dto

import "errors"

type GalleryCreateReq struct {
	Title        string `json:"title" form:"title" binding:"required"`
	Description  string `json:"description" form:"description"`
	EventDate    string `json:"event_date" form:"event_date"`
	GDriveLink   string `json:"gdrive_link" form:"gdrive_link" binding:"required"`
	ThumbnailUrl string `json:"thumbnail_url" form:"thumbnail_url" binding:"required"`
}

type GalleryUpdateReq struct {
	Title        *string `json:"title" form:"title"`
	Description  *string `json:"description" form:"description"`
	EventDate    *string `json:"event_date" form:"event_date"`
	GDriveLink   *string `json:"gdrive_link" form:"gdrive_link"`
	ThumbnailUrl *string `json:"thumbnail_url" form:"thumbnail_url"`
}

type GalleryResponse struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	EventDate    string `json:"event_date,omitempty"`
	GDriveLink   string `json:"gdrive_link"`
	ThumbnailUrl string `json:"thumbnail_url"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GalleryFilterParams struct {
	EventDateFrom string `form:"event_date_from"`
	EventDateTo   string `form:"event_date_to"`
}

const (
	MSG_GALLERY_CREATE_SUCCESS = "gallery created successfully"
	MSG_GALLERY_GET_SUCCESS    = "galleries fetched successfully"
	MSG_GALLERY_DETAIL_SUCCESS = "gallery fetched successfully"
	MSG_GALLERY_UPDATE_SUCCESS = "gallery updated successfully"
	MSG_GALLERY_DELETE_SUCCESS = "gallery deleted successfully"

	MSG_GALLERY_CREATE_FAILED = "failed to create gallery"
	MSG_GALLERY_GET_FAILED    = "failed to fetch galleries"
	MSG_GALLERY_UPDATE_FAILED = "failed to update gallery"
	MSG_GALLERY_DELETE_FAILED = "failed to delete gallery"
	MSG_GALLERY_NOT_FOUND     = "gallery not found"
	MSG_GALLERY_INVALID_ID    = "invalid gallery id"
)

var (
	ErrGalleryNotFound          = errors.New("gallery not found")
	ErrGalleryInvalidID         = errors.New("invalid gallery id")
	ErrGalleryInvalidGDriveLink = errors.New("gdrive_link must be a valid Google Drive URL")
	ErrGalleryThumbnailRequired = errors.New("thumbnail_url is required")
)
