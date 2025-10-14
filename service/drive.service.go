package service

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

type DriveService interface {
	UploadFileToDrive(fileBytes []byte, filename, folderID string) (string, error)
}

type driveService struct{}

func NewDriveService() DriveService {
	return &driveService{}
}

var (
	clientID     = "506226776429-7mmjoitqt3jr1g68ett370nmakh4lhka.apps.googleusercontent.com"
	clientSecret = "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"
)

func getDriveClient() (*drive.Service, error) {
	refreshToken, err := repository.LoadRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("refresh token kosong: %v", err)
	}

	token := &oauth2.Token{RefreshToken: refreshToken}
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{drive.DriveFileScope},
	}

	client := config.Client(context.Background(), token)
	return drive.New(client)
}

func (ds *driveService) UploadFileToDrive(fileBytes []byte, filename, folderID string) (string, error) {
	srv, err := getDriveClient()
	if err != nil {
		return "", err
	}

	driveFile := &drive.File{
		Name:    filename,
		Parents: []string{folderID},
	}

	uploadedFile, err := srv.Files.Create(driveFile).Media(bytes.NewReader(fileBytes)).Do()
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://drive.google.com/uc?id=%s", uploadedFile.Id)
	log.Println("Upload berhasil:", filename, fileURL)
	return fileURL, nil
}
