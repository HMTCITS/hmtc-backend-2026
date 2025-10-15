package dto

type PertanyaanUmum struct {
	Q1 string `json:"q1" form:"q1" binding:"required"`
	Q2 string `json:"q2" form:"q2" binding:"required"`
	Q3 string `json:"q3" form:"q3" binding:"required"`
}

type PertanyaanDivisi map[string]DivisiSoal

type DivisiSoal struct {
	Q1 string `json:"q1" form:"q1" binding:"required"`
	Q2 string `json:"q2" form:"q2" binding:"required"`
	Q3 string `json:"q3" form:"q3" binding:"required"`
	Q4 string `json:"q4" form:"q4" binding:"required"`
	Q5 string `json:"q5" form:"q5" binding:"required"`
}

type UploadDTO struct {
	Nama           string           `json:"nama" form:"nama" binding:"required"`
	NRP            string           `json:"nrp" form:"nrp" binding:"required"`
	KelompokKP     string           `json:"kelompok_kp" form:"kelompok_kp" binding:"required"`
	PertanyaanUmum PertanyaanUmum   `json:"pertanyaan_umum" form:"pertanyaan_umum" binding:"required"`
	DivisiDipilih  []string         `json:"divisi_dipilih" form:"divisi_dipilih" binding:"required,min=1,max=3"`
	PertanyaanDiv  PertanyaanDivisi `json:"pertanyaan_divisi" form:"pertanyaan_divisi" binding:"required"`
	// FileZIP        *multipart.FileHeader `json:"file_zip" form:"file_zip" binding:"required"`
}
