package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
)

type MagangController interface {
	GetToken(ctx *gin.Context)
	Callback(ctx *gin.Context)
	Upload(ctx *gin.Context)
}

type magangController struct {
	magangService  service.MagangService
	oauthTokenRepo repository.OAuthTokenRepository
}

func NewMagangController(ms service.MagangService, otr repository.OAuthTokenRepository) MagangController {
	return &magangController{
		magangService:  ms,
		oauthTokenRepo: otr,
	}
}

// GetToken godoc
// @Summary Generate login URL untuk developer
// @Description Mendapatkan URL untuk login Google OAuth dan mendapatkan refresh token
// @Tags Magang
// @Produce text/html
// @Success 200 {string} string "HTML link untuk login Google"
// @Router /magang/get-token [get]
func (mc *magangController) GetToken(ctx *gin.Context) {
	config := &oauth2.Config{
		ClientID:     config.AppConfig.OauthClientID,
		ClientSecret: config.AppConfig.OauthClientSecret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			drive.DriveFileScope,
			sheets.SpreadsheetsScope,
		},
		RedirectURL: config.AppConfig.RedirectURL,
	}
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(`<a href="%s" target="_blank">Login with Google to get refresh token</a>`, url)))
}

func (mc *magangController) Callback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No code"})
		return
	}

	cfg := &oauth2.Config{
		ClientID:     config.AppConfig.OauthClientID,
		ClientSecret: config.AppConfig.OauthClientSecret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			drive.DriveFileScope,
			sheets.SpreadsheetsScope,
		},
		RedirectURL: config.AppConfig.RedirectURL,
	}

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "exchange error: " + err.Error()})
		return
	}

	refresh := strings.TrimSpace(token.RefreshToken)
	if refresh == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "refresh token kosong dari provider"})
		return
	}

	if err := mc.oauthTokenRepo.Save(refresh); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal simpan: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"refresh_token": refresh})
}

// Upload godoc
// @Summary Upload dokumen magang
// @Description Upload satu file ZIP berisi CV, Brainmap, Portofolio beserta jawaban pertanyaan umum dan divisi
// @Tags Magang
// @Accept multipart/form-data
// @Produce json
// @Param nama formData string true "Nama mahasiswa"
// @Param nrp formData string true "NRP mahasiswa"
// @Param kelompok_kp formData string true "Kelompok KP"
// @Param pertanyaan_umum[q1] formData string true "Pertanyaan umum Q1"
// @Param pertanyaan_umum[q2] formData string true "Pertanyaan umum Q2"
// @Param pertanyaan_umum[q3] formData string true "Pertanyaan umum Q3"
// @Param divisi_dipilih formData []string true "Divisi yang dipilih (min 1, max 3)" collectionFormat(multi) Enums(Marketing,Finance,IT,HR,CMI)
// @Param pertanyaan_divisi[Marketing][q1] formData string false "Soal 1 Marketing"
// @Param pertanyaan_divisi[Marketing][q2] formData string false "Soal 2 Marketing"
// @Param pertanyaan_divisi[Marketing][q3] formData string false "Soal 3 Marketing"
// @Param pertanyaan_divisi[Marketing][q4] formData string false "Soal 4 Marketing"
// @Param pertanyaan_divisi[Marketing][q5] formData string false "Soal 5 Marketing"
// (Ulangi untuk setiap divisi yang mungkin dipilih)
// @Param file_zip formData file true "File ZIP berisi CV, Brainmap, Portofolio"
// @Success 200 {object} map[string]interface{} "File URL dan data form"
// @Failure 400 {object} map[string]string{error=string}
// @Failure 500 {object} map[string]string{error=string}
// @Router /magang/upload [post]
func (mc *magangController) Upload(ctx *gin.Context) {
	var form dto.UploadDTO

	form.Nama = ctx.PostForm("nama")
	form.NRP = ctx.PostForm("nrp")
	form.KelompokKP = ctx.PostForm("kelompok_kp")

	form.PertanyaanUmum = dto.PertanyaanUmum{
		Q1: ctx.PostForm("pertanyaan_umum[q1]"),
		Q2: ctx.PostForm("pertanyaan_umum[q2]"),
		Q3: ctx.PostForm("pertanyaan_umum[q3]"),
	}

	form.DivisiDipilih = ctx.PostFormArray("divisi_dipilih")
	if len(form.DivisiDipilih) < 1 || len(form.DivisiDipilih) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Harus pilih 1-3 divisi"})
		return
	}

	form.PertanyaanDiv = make(dto.PertanyaanDivisi)
	for _, div := range form.DivisiDipilih {
		ds := dto.DivisiSoal{
			Q1: ctx.PostForm(fmt.Sprintf("pertanyaan_divisi[%s][q1]", div)),
			Q2: ctx.PostForm(fmt.Sprintf("pertanyaan_divisi[%s][q2]", div)),
			Q3: ctx.PostForm(fmt.Sprintf("pertanyaan_divisi[%s][q3]", div)),
			Q4: ctx.PostForm(fmt.Sprintf("pertanyaan_divisi[%s][q4]", div)),
			Q5: ctx.PostForm(fmt.Sprintf("pertanyaan_divisi[%s][q5]", div)),
		}
		form.PertanyaanDiv[div] = ds
	}

	// File ZIP
	fileHeader, err := ctx.FormFile("file_zip")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file_zip wajib"})
		return
	}

	// Open and read file
	fileObj, _ := fileHeader.Open()
	defer fileObj.Close()
	fileBytes, _ := io.ReadAll(fileObj)

	// Upload ke Drive
	fileURL, err := mc.magangService.UploadFile(fileBytes, fileHeader.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Upload ke Google Sheets (pakai DTO langsung)
	if err := mc.magangService.UploadToSheet(form, fileURL); err != nil {
		log.Println("Gagal upload ke spreadsheet:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload berhasil",
		"fileUrl": fileURL,
		"data":    form,
	})
}
