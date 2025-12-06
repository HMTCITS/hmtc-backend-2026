package controller

import (
	"net/http"
	"strconv"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/gin-gonic/gin"
)

type EvalCmi25Controller interface {
	Upload(ctx *gin.Context)
}

type evalCmi25Controller struct {
	evalService service.EvalCmi25Service
}

func NewEvalCmi25Controller(s service.EvalCmi25Service) EvalCmi25Controller {
	return &evalCmi25Controller{
		evalService: s,
	}
}

// Helper convert string → int (default 0)
func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Upload godoc
// @Summary Mengunggah hasil evaluasi CMI 2025
// @Description Menerima data evaluasi dalam bentuk form-data dan menyimpannya ke Google Sheets.
// @Tags Evaluasi CMI 2025
// @Accept multipart/form-data
// @Produce json
//
// @Param nama formData string true "Nama penilai"
// @Param departemen formData string true "Departemen penilai"
//
// @Param kejelasan_komunikasi formData int false "Kejelasan komunikasi"
// @Param responsivitas formData int false "Responsivitas"
// @Param koordinasi_kegiatan formData int false "Koordinasi kegiatan"
// @Param profesionalisme formData int false "Profesionalisme"
// @Param keterbukaan_feedback formData int false "Keterbukaan terhadap feedback"
//
// @Param kualitas_dukung formData int false "Kualitas dukungan"
// @Param keterlibatan_aktif formData int false "Keterlibatan aktif"
// @Param inovasi_kreativitas formData int false "Inovasi & kreativitas"
// @Param pemahaman_tugas formData int false "Pemahaman tugas"
// @Param kepatuhan_deadline formData int false "Kepatuhan deadline"
//
// @Param cd_konsistensi_visual formData int false "Consistensi visual (CD)"
// @Param cd_kesesuaian_brief formData int false "Kesesuaian brief (CD)"
// @Param cd_estetika formData int false "Estetika (CD)"
// @Param cd_kecepatan_revisi formData int false "Kecepatan revisi (CD)"
//
// @Param sms_strategi_konten formData int false "Strategi konten (SM)"
// @Param sms_audien formData int false "Pemahaman audiens (SM)"
// @Param sms_caption formData int false "Copywriting/caption (SM)"
// @Param sms_analitik formData int false "Analitik konten (SM)"
//
// @Param mp_kualitas_produksi formData int false "Kualitas produksi (MP)"
// @Param mp_konsep formData int false "Konsep (MP)"
// @Param mp_inovasi formData int false "Inovasi (MP)"
// @Param mp_dokumentasi formData int false "Dokumentasi (MP)"
//
// @Param it_stabilitas formData int false "Stabilitas sistem (IT)"
// @Param it_teknis formData int false "Kemampuan teknis (IT)"
// @Param it_keamanan formData int false "Keamanan sistem (IT)"
// @Param it_ux formData int false "User experience (IT)"
//
// @Param umpan_balik_umum formData string false "Feedback umum"
// @Param saran_perbaikan formData string false "Saran perbaikan"
// @Param komentar_tambahan formData string false "Komentar tambahan"
//
// @Success 200 {object} map[string]interface{} "Evaluasi berhasil disimpan"
// @Failure 400 {object} map[string]string "Form tidak valid"
// @Failure 500 {object} map[string]string "Gagal menyimpan ke Google Sheets"
//
// @Router /evaluasi-cmi/submit [post]
func (ec *evalCmi25Controller) Upload(ctx *gin.Context) {
	var req dto.EvalCmi25Req

	// Bind form-data langsung ke struct
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Form tidak valid: " + err.Error(),
		})
		return
	}

	// Identitas
	req.Nama = ctx.PostForm("nama")
	req.Departemen = ctx.PostForm("departemen")

	req.KejelasanKomunikasi = toInt(ctx.PostForm("kejelasan_komunikasi"))
	req.Responsivitas = toInt(ctx.PostForm("responsivitas"))
	req.KoordinasiKegiatan = toInt(ctx.PostForm("koordinasi_kegiatan"))
	req.Profesionalisme = toInt(ctx.PostForm("profesionalisme"))
	req.KeterbukaanFeedback = toInt(ctx.PostForm("keterbukaan_feedback"))

	req.KualitasDukungan = toInt(ctx.PostForm("kualitas_dukung"))
	req.KeterlibatanAktif = toInt(ctx.PostForm("keterlibatan_aktif"))
	req.InovasiKreativitas = toInt(ctx.PostForm("inovasi_kreativitas"))
	req.PemahamanTugas = toInt(ctx.PostForm("pemahaman_tugas"))
	req.KepatuhanDeadline = toInt(ctx.PostForm("kepatuhan_deadline"))

	// Creative Design
	req.CdKonsistensiVisual = toInt(ctx.PostForm("cd_konsistensi_visual"))
	req.CdKesesuaianBrief = toInt(ctx.PostForm("cd_kesesuaian_brief"))
	req.CdEstetika = toInt(ctx.PostForm("cd_estetika"))
	req.CdKecepatanRevisi = toInt(ctx.PostForm("cd_kecepatan_revisi"))

	// Social Media
	req.SmsStrategiKonten = toInt(ctx.PostForm("sms_strategi_konten"))
	req.SmsAudien = toInt(ctx.PostForm("sms_audien"))
	req.SmsCaption = toInt(ctx.PostForm("sms_caption"))
	req.SmsAnalitik = toInt(ctx.PostForm("sms_analitik"))

	// Media Production
	req.MpKualitasProduksi = toInt(ctx.PostForm("mp_kualitas_produksi"))
	req.MpKonsep = toInt(ctx.PostForm("mp_konsep"))
	req.MpInovasi = toInt(ctx.PostForm("mp_inovasi"))
	req.MpDokumentasi = toInt(ctx.PostForm("mp_dokumentasi"))

	// IT Dev
	req.ItStabilitas = toInt(ctx.PostForm("it_stabilitas"))
	req.ItTeknis = toInt(ctx.PostForm("it_teknis"))
	req.ItKeamanan = toInt(ctx.PostForm("it_keamanan"))
	req.ItUx = toInt(ctx.PostForm("it_ux"))

	// Essay
	req.UmpanBalikUmum = ctx.PostForm("umpan_balik_umum")
	req.SaranPerbaikan = ctx.PostForm("saran_perbaikan")
	req.KomentarTambahan = ctx.PostForm("komentar_tambahan")

	if err := ec.evalService.UploadToSheet(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan evaluasi: " + err.Error(),
		})
		return
	}

	// ===========================
	// SUCCESS RESPONSE
	// ===========================
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Evaluasi berhasil disimpan!",
		"data":    req,
	})
}
