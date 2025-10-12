package repository

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFileReqRepository interface {
	NewUserFileReq(ctx context.Context, tx *gorm.DB, userReq model.UserFileReq) error
	UserFileStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error)
	ChangeReqStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID, status model.Status) error
}

type userFileReqRepository struct {
	db *gorm.DB
}

func NewUserFileRepository(db *gorm.DB) UserFileReqRepository {
	return &userFileReqRepository{db: db}
}

func (r *userFileReqRepository) NewUserFileReq(ctx context.Context, tx *gorm.DB, userReq model.UserFileReq) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&userReq).Error; err != nil {
		return err
	}

	return nil
}

func (r *userFileReqRepository) UserFileStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "id = ?", reqId).Error; err != nil {
		return "file not found", err
	}

	return string(file.Status), nil
}

func (r *userFileReqRepository) ChangeReqStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID, status model.Status) error {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "id = ?", reqId).Error; err != nil {
		return err
	}

	file.Status = status

	if err := tx.WithContext(ctx).Save(&file).Error; err != nil {
		return err
	}

	return nil
}
