package service

import (
	"fmt"
	"log"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
)

type MagangService interface {
	UploadFile(fileBytes []byte, filename string) (string, error)
	UploadToSheet(form dto.UploadDTO, fileURL string) error
	UploadToSheetSA(data map[string]interface{}) error
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

func (ms *magangService) UploadFile(fileBytes []byte, filename string) (string, error) {
	// ID folder magang
	const folderID = "1IVZqff-KivZFFrBTMqZOcEn4JhMF5J_1"
	return ms.driveService.UploadFileToDrive(fileBytes, filename, folderID)
}
func (ms *magangService) UploadToSheet(form dto.UploadDTO, fileURL string) error {
	const spreadsheetID = "1188huzNIVi0vuxdvInPUtx1XeFhWuNhpKiVFyc7U51s"
	sheetName := "from web"

	// 1. Base fields
	values := []interface{}{
		time.Now().Format("2006-01-02 15:04:05"),
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

func (ms *magangService) UploadToSheetSA(data map[string]interface{}) error {
	const spreadsheetID = "1188huzNIVi0vuxdvInPUtx1XeFhWuNhpKiVFyc7U51s"
	sheetName := "from web"
	credsFile := "service-account.json"

	// ubah data menjadi array interface
	values := []interface{}{
		data["nama"],
		data["nrp"],
		data["kelompok_kp"],
		data["file_url"],
		// data["pertanyaan_umum"],
		// data["divisi_dipilih"],
		// data["pertanyaan_divisi"],
	}

	// panggil fungsi helper service account
	err := ms.sheetsService.AppendRowServiceAccount(credsFile, spreadsheetID, sheetName, values)
	if err != nil {
		log.Println("Gagal upload ke spreadsheet (SA):", err)
		return fmt.Errorf("upload spreadsheet error: %v", err)
	}
	return nil
}
