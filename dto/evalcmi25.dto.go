package dto

type EvalCmi25Req struct {
	Nama       string `form:"nama" json:"nama" binding:"required"`
	Departemen string `form:"departemen" json:"departemen" binding:"required"`

	KejelasanKomunikasi int `form:"kejelasan_komunikasi" json:"kejelasan_komunikasi"`
	Responsivitas       int `form:"responsivitas" json:"responsivitas"`
	KoordinasiKegiatan  int `form:"koordinasi_kegiatan" json:"koordinasi_kegiatan"`
	Profesionalisme     int `form:"profesionalisme" json:"profesionalisme"`
	KeterbukaanFeedback int `form:"keterbukaan_feedback" json:"keterbukaan_feedback"`

	KualitasDukungan   int `form:"kualitas_dukung" json:"kualitas_dukung"`
	KeterlibatanAktif  int `form:"keterlibatan_aktif" json:"keterlibatan_aktif"`
	InovasiKreativitas int `form:"inovasi_kreativitas" json:"inovasi_kreativitas"`
	PemahamanTugas     int `form:"pemahaman_tugas" json:"pemahaman_tugas"`
	KepatuhanDeadline  int `form:"kepatuhan_deadline" json:"kepatuhan_deadline"`

	CdKonsistensiVisual int `form:"cd_konsistensi_visual" json:"cd_konsistensi_visual"`
	CdKesesuaianBrief   int `form:"cd_kesesuaian_brief" json:"cd_kesesuaian_brief"`
	CdEstetika          int `form:"cd_estetika" json:"cd_estetika"`
	CdKecepatanRevisi   int `form:"cd_kecepatan_revisi" json:"cd_kecepatan_revisi"`

	SmsStrategiKonten int `form:"sms_strategi_konten" json:"sms_strategi_konten"`
	SmsAudien         int `form:"sms_audien" json:"sms_audien"`
	SmsCaption        int `form:"sms_caption" json:"sms_caption"`
	SmsAnalitik       int `form:"sms_analitik" json:"sms_analitik"`

	MpKualitasProduksi int `form:"mp_kualitas_produksi" json:"mp_kualitas_produksi"`
	MpKonsep           int `form:"mp_konsep" json:"mp_konsep"`
	MpInovasi          int `form:"mp_inovasi" json:"mp_inovasi"`
	MpDokumentasi      int `form:"mp_dokumentasi" json:"mp_dokumentasi"`

	ItStabilitas int `form:"it_stabilitas" json:"it_stabilitas"`
	ItTeknis     int `form:"it_teknis" json:"it_teknis"`
	ItKeamanan   int `form:"it_keamanan" json:"it_keamanan"`
	ItUx         int `form:"it_ux" json:"it_ux"`

	UmpanBalikUmum   string `form:"umpan_balik_umum" json:"umpan_balik_umum"`
	SaranPerbaikan   string `form:"saran_perbaikan" json:"saran_perbaikan"`
	KomentarTambahan string `form:"komentar_tambahan" json:"komentar_tambahan"`
}
