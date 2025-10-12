package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/google/uuid"
)

type FileTAService interface {
	CreateFileTA(ctx context.Context, req dto.CreateFileTA, filename string, file multipart.File) error
	GetFileStatus(ctx context.Context, fileId string) (string, error)
	GetAllFiles(ctx context.Context) ([]dto.GetAllFiles, error)
	ChangeFileStatus(ctx context.Context, req dto.ChangeFileStatus) error
}

type fileTAService struct {
	fileTARepo repository.FileTARepository
}

func NewFileTAService(fileTARepo repository.FileTARepository) FileTAService {
	return &fileTAService{fileTARepo: fileTARepo}
}

func (s *fileTAService) CreateFileTA(ctx context.Context, req dto.CreateFileTA, filename string, file multipart.File) error {

	out, err := os.Create("./file-ta/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

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

func (s *fileTAService) GetFileStatus(ctx context.Context, fileId string) (string, error) {

	status, err := s.fileTARepo.GetFileStatus(ctx, nil, uuid.MustParse(fileId))
	if err != nil {
		return "", err
	}

	return status, nil
}

func (s *fileTAService) GetAllFiles(ctx context.Context) ([]dto.GetAllFiles, error) {

	files, err := s.fileTARepo.GetAllFileTA(ctx, nil)
	if err != nil {
		return []dto.GetAllFiles{}, nil
	}

	var allFiles []dto.GetAllFiles
	for index, _ := range files {
		fileDetail := dto.GetAllFiles{
			Id:          files[index].Id.String(),
			FileName:    files[index].FileName,
			StudentName: files[index].StudentName,
			NRP:         files[index].NRP,
			Semester:    files[index].Semester,
			DosPem:      files[index].DosPem,
		}
		allFiles = append(allFiles, fileDetail)
	}

	return allFiles, nil
}

func (s *fileTAService) ChangeFileStatus(ctx context.Context, req dto.ChangeFileStatus) error {

	status, filename, err := s.fileTARepo.ChangeFileStatus(ctx, nil, uuid.MustParse(req.FileId), model.FileStatus(req.Status))
	if err != nil {
		return err
	}

	if status == string(model.StatusRejected) {
		path := "./file-ta/" + filename
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}
