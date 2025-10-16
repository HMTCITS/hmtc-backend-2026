package service

import (
	"io"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/dto"
)

type MagangService interface {
	UploadFile(r io.Reader, filename string) (string, error)
	UploadToSheet(form dto.UploadDTO, fileURL string) error
}

type magangService struct {
	driveService  DriveService
	sheetsService SheetsService
}

func NewMagangService(ds DriveService, ss SheetsService) MagangService {
	return &magangService{
		driveService:  ds,
		sheetsService: ss,
	}
}

func (ms *magangService) UploadFile(r io.Reader, filename string) (string, error) {
	folderID := config.AppConfig.GDriveFolderID

	return ms.driveService.UploadFileToDrive(r, filename, folderID)
}

func (ms *magangService) UploadToSheet(form dto.UploadDTO, fileURL string) error {
	spreadsheetID := config.AppConfig.SheetsID
	sheetName := config.AppConfig.SheetsName

	// 1. Base fields
	values := []interface{}{
		// time.Now().Format("2006-01-02 15:04:05"),
		form.Nama,
		form.NRP,
		form.KelompokKP,
		fileURL,
		form.PertanyaanUmum.Q1,
		form.PertanyaanUmum.Q2,
		form.PertanyaanUmum.Q3,
	}

	// 2. Divisi dan jawabannya (maks 3 divisi × (nama + 5 jawaban))
	for i := 0; i < 3; i++ {
		if i < len(form.DivisiDipilih) {
			divName := form.DivisiDipilih[i]
			values = append(values, divName)

			if soal, ok := form.PertanyaanDiv[divName]; ok {
				values = append(values,
					soal.Q1,
					soal.Q2,
					soal.Q3,
					soal.Q4,
					soal.Q5,
				)
			} else {
				// Jika tidak ada jawaban divisinya (harusnya tidak kejadian)
				values = append(values, "", "", "", "", "")
			}
		} else {
			// Tidak ada divisi → isi kosong semua
			values = append(values, "", "", "", "", "", "")
		}
	}

	// 3. Kirim ke Google Sheets
	return ms.sheetsService.AppendRow(spreadsheetID, sheetName, values)
}
