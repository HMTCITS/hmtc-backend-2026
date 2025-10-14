package repository

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFileReqRepository interface {
	NewUserFileReq(ctx context.Context, tx *gorm.DB, userReq model.UserFileReq) error
	UserFileStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (model.Status, error)
	GetFileIdByReqId(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error)
	GetUserNRPReq(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error)
	GetAllFileReq(ctx context.Context, tx *gorm.DB) ([]model.UserFileReq, error)
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

func (r *userFileReqRepository) UserFileStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (model.Status, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "req_id = ?", reqId).Error; err != nil {
		return "file not found", err
	}

	return file.Status, nil
}

func (r *userFileReqRepository) GetFileIdByReqId(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "req_id = ?", reqId).Error; err != nil {
		return "file not found", err
	}

	return file.FileId.String(), nil
}

func (r *userFileReqRepository) GetUserNRPReq(ctx context.Context, tx *gorm.DB, reqId uuid.UUID) (string, error) {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "req_id = ?", reqId).Error; err != nil {
		return "file not found", err
	}

	return file.NRP, nil
}

func (r *userFileReqRepository) ChangeReqStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID, status model.Status) error {
	if tx == nil {
		tx = r.db
	}

	var file model.UserFileReq
	if err := tx.WithContext(ctx).Take(&file, "req_id = ?", reqId).Error; err != nil {
		return err
	}

	file.Status = status

	if err := tx.WithContext(ctx).Save(&file).Error; err != nil {
		return err
	}

	return nil
}

func (r *userFileReqRepository) GetAllFileReq(ctx context.Context, tx *gorm.DB) ([]model.UserFileReq, error) {
	if tx == nil {
		tx = r.db
	}

	var fileReq []model.UserFileReq
	if err := tx.WithContext(ctx).Find(&fileReq).Error; err != nil {
		return []model.UserFileReq{}, err
	}

	return fileReq, nil
}
