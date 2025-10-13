package service

import (
	"context"
	"errors"
	"fmt"
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
	GetFileName(ctx context.Context, reqId string, fileId string, userNRP string) (string, error)
	GetAllFiles(ctx context.Context) ([]dto.GetAllFiles, error)
	ChangeFileStatus(ctx context.Context, req dto.ChangeFileStatus) error
}

type fileTAService struct {
	fileTARepo      repository.FileTARepository
	userFileReqRepo repository.UserFileReqRepository
}

func NewFileTAService(fileTARepo repository.FileTARepository, userFileReqRepo repository.UserFileReqRepository) FileTAService {
	return &fileTAService{fileTARepo: fileTARepo, userFileReqRepo: userFileReqRepo}
}

func (s *fileTAService) CreateFileTA(ctx context.Context, req dto.CreateFileTA, filename string, file multipart.File) error {
	err := os.MkdirAll("./file-ta", os.ModePerm)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

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
		return fileErr
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
	for index := range files {
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

	if status == string(model.REJECTED) {
		path := "./file-ta/" + filename
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *fileTAService) GetFileName(ctx context.Context, reqId string, fileId string, userNRP string) (string, error) {
	fmt.Println(userNRP)
	nrp, err := s.userFileReqRepo.GetUserNRPReq(ctx, nil, uuid.MustParse(reqId))
	if err != nil {
		return "", err
	}
	if userNRP != nrp {
		return "", errors.New("invalid NRP")
	}
	filename, err := s.fileTARepo.GetFileName(ctx, nil, uuid.MustParse(fileId))
	if err != nil {
		return "", err
	}
	return filename, nil
}
