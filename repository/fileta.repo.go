package repository

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileTARepository interface {
	CreateFileTA(ctx context.Context, tx *gorm.DB, fileTa model.TAFile) error
	GetAllFileTA(ctx context.Context, tx *gorm.DB) ([]model.TAFile, error)
	ChangeFileStatus(ctx context.Context, tx *gorm.DB, fileId uuid.UUID, status model.FileStatus) (string, string, error)
	GetFileStatus(ctx context.Context, tx *gorm.DB, fileId uuid.UUID) (string, error)
}

type fileTARepository struct {
	db *gorm.DB
}

func NewFileTARepository(db *gorm.DB) FileTARepository {
	return &fileTARepository{db: db}
}

func (r *fileTARepository) CreateFileTA(ctx context.Context, tx *gorm.DB, fileTa model.TAFile) error {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&fileTa).Error; err != nil {
		return err
	}

	return nil
}

func (r *fileTARepository) GetAllFileTA(ctx context.Context, tx *gorm.DB) ([]model.TAFile, error) {

	if tx == nil {
		tx = r.db
	}

	var files []model.TAFile
	if err := tx.WithContext(ctx).Find(&files).Error; err != nil {
		return []model.TAFile{}, err
	}

	return files, nil
}

func (r *fileTARepository) ChangeFileStatus(ctx context.Context, tx *gorm.DB, fileId uuid.UUID, status model.FileStatus) (string, string, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.TAFile
	if err := tx.Take(&file, "id = ?", fileId).Error; err != nil {
		return "", "", err
	}

	file.Status = status
	if err := tx.Save(&file).Error; err != nil {
		return "", "", err
	}

	return string(status), file.FileName, nil
}

func (r *fileTARepository) GetFileStatus(ctx context.Context, tx *gorm.DB, fileId uuid.UUID) (string, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.TAFile
	if err := tx.Take(&file, "id = ?", fileId).Error; err != nil {
		return "Fail", err
	}

	return string(file.Status), nil
}
