package repository

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFileReqRepository interface {
	NewUserFileReq(ctx context.Context, tx *gorm.DB, userReq model.UserFileReq) error
	UserFileStatus(ctx context.Context, tx *gorm.DB, reqId uuid.UUID)
}

type userFileReqRepository struct {
	db *gorm.DB
}

func NewUserFileRepository(db *gorm.DB) UserFileReqRepository {
	return &userFileReqRepository{db: db}
}
