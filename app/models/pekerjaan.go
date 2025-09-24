package models

import "time"

type Pekerjaan struct {
	ID                  int        `json:"id"`
	AlumniID            int        `json:"alumni_id"`
	NamaPerusahaan      string     `json:"nama_perusahaan"`
	PosisiJabatan       string     `json:"posisi_jabatan"`
	BidangIndustri      string     `json:"bidang_industri"`
	LokasiKerja         string     `json:"lokasi_kerja"`
	GajiRange           *string    `json:"gaji_range,omitempty"`
	TanggalMulaiKerja   time.Time  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string     `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string    `json:"deskripsi_pekerjaan,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	IsDeleted           bool       `json:"isDeleted"`
}

type CreatePekerjaan struct {
	AlumniID            int     `json:"alumni_id"`
	NamaPerusahaan      string  `json:"nama_perusahaan"`
	PosisiJabatan       string  `json:"posisi_jabatan"`
	BidangIndustri      string  `json:"bidang_industri"`
	LokasiKerja         string  `json:"lokasi_kerja"`
	GajiRange           *string `json:"gaji_range,omitempty"`
	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string  `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan,omitempty"`
	IsDeleted           bool    `json:"isDeleted"`
}

type UpdatePekerjaan struct {
	NamaPerusahaan      string     `json:"nama_perusahaan"`
	PosisiJabatan       string     `json:"posisi_jabatan"`
	BidangIndustri      string     `json:"bidang_industri"`
	LokasiKerja         string     `json:"lokasi_kerja"`
	GajiRange           *string    `json:"gaji_range,omitempty"`
	TanggalMulaiKerja   time.Time  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string     `json:"status_pekerjaan" validate:"oneof=aktif tidak_aktif resign"`
	DeskripsiPekerjaan  *string    `json:"deskripsi_pekerjaan,omitempty"`
	UpdatedAt           time.Time  `json:"updated_at"`
	IsDeleted           bool       `json:"isDeleted"`
}
