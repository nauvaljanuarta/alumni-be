package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Pekerjaan struct {
// 	ID                  int        `json:"id"`
// 	AlumniID            int        `json:"alumni_id"`
// 	NamaPerusahaan      string     `json:"nama_perusahaan"`
// 	PosisiJabatan       string     `json:"posisi_jabatan"`
// 	BidangIndustri      string     `json:"bidang_industri"`
// 	LokasiKerja         string     `json:"lokasi_kerja"`
// 	GajiRange           *string    `json:"gaji_range,omitempty"`
// 	TanggalMulaiKerja   time.Time  `json:"tanggal_mulai_kerja"`
// 	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja,omitempty"`
// 	StatusPekerjaan     string     `json:"status_pekerjaan"`
// 	DeskripsiPekerjaan  *string    `json:"deskripsi_pekerjaan,omitempty"`
// 	CreatedAt           time.Time  `json:"created_at"`
// 	UpdatedAt           time.Time  `json:"updated_at"`
// 	IsDeleted           bool       `json:"isdeleted"`
// }

// type CreatePekerjaan struct {
// 	AlumniID            int     `json:"alumni_id"`
// 	NamaPerusahaan      string  `json:"nama_perusahaan"`
// 	PosisiJabatan       string  `json:"posisi_jabatan"`
// 	BidangIndustri      string  `json:"bidang_industri"`
// 	LokasiKerja         string  `json:"lokasi_kerja"`
// 	GajiRange           *string `json:"gaji_range,omitempty"`
// 	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`
// 	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja,omitempty"`
// 	StatusPekerjaan     string  `json:"status_pekerjaan"`
// 	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan,omitempty"`
// 	IsDeleted           bool    `json:"isdeleted"`
// }

// type UpdatePekerjaan struct {
// 	NamaPerusahaan      string     `json:"nama_perusahaan"`
// 	PosisiJabatan       string     `json:"posisi_jabatan"`
// 	BidangIndustri      string     `json:"bidang_industri"`
// 	LokasiKerja         string     `json:"lokasi_kerja"`
// 	GajiRange           *string    `json:"gaji_range,omitempty"`
// 	TanggalMulaiKerja   time.Time  `json:"tanggal_mulai_kerja"`
// 	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja,omitempty"`
// 	StatusPekerjaan     string     `json:"status_pekerjaan" validate:"oneof=aktif tidak_aktif resign"`
// 	DeskripsiPekerjaan  *string    `json:"deskripsi_pekerjaan,omitempty"`
// 	UpdatedAt           time.Time  `json:"updated_at"`
// 	IsDeleted           bool       `json:"isdeleted"`
// }

type Pekerjaan struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AlumniID            primitive.ObjectID `bson:"alumni_id" json:"alumni_id"`
	NamaPerusahaan      string             `bson:"nama_perusahaan" json:"nama_perusahaan"`
	PosisiJabatan       string             `bson:"posisi_jabatan" json:"posisi_jabatan"`
	BidangIndustri      string             `bson:"bidang_industri" json:"bidang_industri"`
	LokasiKerja         string             `bson:"lokasi_kerja" json:"lokasi_kerja"`
	GajiRange           *string            `bson:"gaji_range,omitempty" json:"gaji_range,omitempty"`
	TanggalMulaiKerja   time.Time          `bson:"tanggal_mulai_kerja" json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time         `bson:"tanggal_selesai_kerja,omitempty" json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string             `bson:"status_pekerjaan" json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string            `bson:"deskripsi_pekerjaan,omitempty" json:"deskripsi_pekerjaan,omitempty"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updated_at"`
	IsDeleted           bool               `bson:"is_deleted" json:"is_deleted"`
}

type CreatePekerjaan struct {
	AlumniID            string    `bson:"alumni_id" json:"alumni_id"`
	NamaPerusahaan      string    `bson:"nama_perusahaan" json:"nama_perusahaan"`
	PosisiJabatan       string    `bson:"posisi_jabatan" json:"posisi_jabatan"`
	BidangIndustri      string    `bson:"bidang_industri" json:"bidang_industri"`
	LokasiKerja         string    `bson:"lokasi_kerja" json:"lokasi_kerja"`
	GajiRange           *string   `bson:"gaji_range,omitempty" json:"gaji_range,omitempty"`
	TanggalMulaiKerja   string    `bson:"tanggal_mulai_kerja" json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *string   `bson:"tanggal_selesai_kerja,omitempty" json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string    `bson:"status_pekerjaan" json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string   `bson:"deskripsi_pekerjaan,omitempty" json:"deskripsi_pekerjaan,omitempty"`
	IsDeleted           bool      `bson:"is_deleted" json:"is_deleted"`
	CreatedAt           time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at" json:"updated_at"`
}

type UpdatePekerjaan struct {
	NamaPerusahaan      string    `bson:"nama_perusahaan,omitempty" json:"nama_perusahaan,omitempty"`
	PosisiJabatan       string    `bson:"posisi_jabatan,omitempty" json:"posisi_jabatan,omitempty"`
	BidangIndustri      string    `bson:"bidang_industri,omitempty" json:"bidang_industri,omitempty"`
	LokasiKerja         string    `bson:"lokasi_kerja,omitempty" json:"lokasi_kerja,omitempty"`
	GajiRange           *string   `bson:"gaji_range,omitempty" json:"gaji_range,omitempty"`
	TanggalMulaiKerja   string    `bson:"tanggal_mulai_kerja,omitempty" json:"tanggal_mulai_kerja,omitempty"`
	TanggalSelesaiKerja *string   `bson:"tanggal_selesai_kerja,omitempty" json:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string    `bson:"status_pekerjaan,omitempty" json:"status_pekerjaan,omitempty"`
	DeskripsiPekerjaan  *string   `bson:"deskripsi_pekerjaan,omitempty" json:"deskripsi_pekerjaan,omitempty"`
	UpdatedAt           time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	IsDeleted           bool      `bson:"is_deleted" json:"is_deleted"`
	CreatedAt           time.Time `bson:"created_at" json:"created_at"`
}
