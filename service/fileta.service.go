package service

import (
	"context"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
)

type FileTAService interface {
	CreateFileTA(ctx context.Context, req dto.CreateFileTA, filename string) error
}

type fileTAService struct {
	fileTARepo repository.FileTARepository
}

func NewFileTAService(fileTARepo repository.FileTARepository) FileTAService {
	return &fileTAService{fileTARepo: fileTARepo}
}

func (s *fileTAService) CreateFileTA(ctx context.Context, req dto.CreateFileTA, filename string) error {

	fileId := uuid.New().String()
	fileTA, err := model.NewFileTA(fileId, filename, req.StudentName, req.NRP, req.Email, req.Semester, req.DosPem)
	if err != nil {
		return err
	}

	fileErr := s.fileTARepo.CreateFileTA(ctx, nil, fileTA)
	if fileErr != nil {
		return err
	}

	return nil
}
