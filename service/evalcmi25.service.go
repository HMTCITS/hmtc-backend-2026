package service

import (
	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/dto"
)

type EvalCmi25Service interface {
	UploadToSheet(req dto.EvalCmi25Req) error
}

type evalCmi25Service struct {
	sheetsService SheetsService
}

func NewEvalCmi25Service(ss SheetsService) EvalCmi25Service {
	return &evalCmi25Service{
		sheetsService: ss,
	}
}

func (es *evalCmi25Service) UploadToSheet(req dto.EvalCmi25Req) error {
	spreadsheetID := config.AppConfig.SheetsIDEvalCmi25
	sheetName := config.AppConfig.SheetsNameEvalCmi25

	// ===========================
	// URUTAN NILAI UNTUK SHEETS
	// ===========================
	values := []interface{}{
		// Identitas
		req.Nama,
		req.Departemen,

		// General Evaluation
		req.KejelasanKomunikasi,
		req.Responsivitas,
		req.KoordinasiKegiatan,
		req.Profesionalisme,
		req.KeterbukaanFeedback,

		req.KualitasDukungan,
		req.KeterlibatanAktif,
		req.InovasiKreativitas,
		req.PemahamanTugas,
		req.KepatuhanDeadline,

		// Creative Design
		req.CdKonsistensiVisual,
		req.CdKesesuaianBrief,
		req.CdEstetika,
		req.CdKecepatanRevisi,

		// Social Media
		req.SmsStrategiKonten,
		req.SmsAudien,
		req.SmsCaption,
		req.SmsAnalitik,

		// Media Production
		req.MpKualitasProduksi,
		req.MpKonsep,
		req.MpInovasi,
		req.MpDokumentasi,

		// IT Development
		req.ItStabilitas,
		req.ItTeknis,
		req.ItKeamanan,
		req.ItUx,

		// Essay
		req.UmpanBalikUmum,
		req.SaranPerbaikan,
		req.KomentarTambahan,
	}

	// ===========================
	// PUSH DATA KE GOOGLE SHEETS
	// ===========================
	return es.sheetsService.AppendRow(spreadsheetID, sheetName, values)
}
