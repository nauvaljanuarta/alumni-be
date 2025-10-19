package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
  "go.mongodb.org/mongo-driver/bson/primitive"
	)

// type Alumni struct {
// 	ID         int       `json:"id"`
// 	NIM        string    `json:"nim"`
// 	Nama       string    `json:"nama"`
// 	Jurusan    string    `json:"jurusan"`
// 	Angkatan   int       `json:"angkatan"`
// 	TahunLulus int       `json:"tahun_lulus"`
// 	Email      string    `json:"email"`
// 	NoTelepon  string    `json:"no_telepon"`
// 	Alamat     string    `json:"alamat"`
// 	Fakultas   string    `json:"fakultas"`
// 	Role       string    `json:"role"`
// 	Password   string    `json:"-"` // agar password tidak muncul pada saat di get
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }

// type CreateAlumni struct {
// 	NIM        string `json:"nim"`
// 	Nama       string `json:"nama"`
// 	Jurusan    string `json:"jurusan"`
// 	Angkatan   int    `json:"angkatan"`
// 	TahunLulus int    `json:"tahun_lulus"`
// 	Email      string `json:"email"`
// 	NoTelepon  string `json:"no_telepon"`
// 	Alamat     string `json:"alamat"`
// 	Fakultas   string `json:"fakultas"`
// 	Role       string `json:"role"`
// 	Password   string `json:"password"`
// }

// type UpdateAlumni struct {
// 	Nama       string `json:"nama"`
// 	Jurusan    string `json:"jurusan"`
// 	Angkatan   int    `json:"angkatan"`
// 	TahunLulus int    `json:"tahun_lulus"`
// 	Email      string `json:"email"`
// 	NoTelepon  string `json:"no_telepon"`
// 	Alamat     string `json:"alamat"`
// 	Role       string `json:"role"`
// 	Fakultas   string `json:"fakultas"`
// 	Password   string `json:"password"`
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // response setelah login
// type LoginResponse struct {
// 	Alumni Alumni `json:"alumni"`
// 	Token  string `json:"token"`
// }

// // JWT claims
// type JWTClaims struct {
// 	UserID int    `json:"user_id"`
// 	Email  string `json:"email"`
// 	Role   string `json:"role"`
// 	jwt.RegisteredClaims
// }

type Alumni struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NIM        string             `bson:"nim" json:"nim"`
	Nama       string             `bson:"nama" json:"nama"`
	Jurusan    string             `bson:"jurusan" json:"jurusan"`
	Angkatan   int                `bson:"angkatan" json:"angkatan"`
	TahunLulus int                `bson:"tahun_lulus" json:"tahun_lulus"`
	Email      string             `bson:"email" json:"email"`
	NoTelepon  string             `bson:"no_telepon" json:"no_telepon"`
	Alamat     string             `bson:"alamat" json:"alamat"`
	Fakultas   string             `bson:"fakultas" json:"fakultas"`
	Role       string             `bson:"role" json:"role"`
	Password   string             `bson:"password,omitempty" json:"-"` // disembunyikan saat response
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateAlumni struct {
	NIM        string `bson:"nim" json:"nim"`
	Nama       string `bson:"nama" json:"nama"`
	Jurusan    string `bson:"jurusan" json:"jurusan"`
	Angkatan   int    `bson:"angkatan" json:"angkatan"`
	TahunLulus int    `bson:"tahun_lulus" json:"tahun_lulus"`
	Email      string `bson:"email" json:"email"`
	NoTelepon  string `bson:"no_telepon" json:"no_telepon"`
	Alamat     string `bson:"alamat" json:"alamat"`
	Fakultas   string `bson:"fakultas" json:"fakultas"`
	Role       string `bson:"role" json:"role"`
	Password   string `bson:"password" json:"password"`
}

type UpdateAlumni struct {
	Nama       string `bson:"nama,omitempty" json:"nama"`
	Jurusan    string `bson:"jurusan,omitempty" json:"jurusan"`
	Angkatan   int    `bson:"angkatan,omitempty" json:"angkatan"`
	TahunLulus int    `bson:"tahun_lulus,omitempty" json:"tahun_lulus"`
	Email      string `bson:"email,omitempty" json:"email"`
	NoTelepon  string `bson:"no_telepon,omitempty" json:"no_telepon"`
	Alamat     string `bson:"alamat,omitempty" json:"alamat"`
	Fakultas   string `bson:"fakultas,omitempty" json:"fakultas"`
	Role       string `bson:"role,omitempty" json:"role"`
	Password   string `bson:"password,omitempty" json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Alumni Alumni `json:"alumni"`
	Token  string `json:"token"`
}

type JWTClaims struct {
	UserID primitive.ObjectID `json:"user_id"`
	Email  string             `json:"email"`
	Role   string             `json:"role"`
	jwt.RegisteredClaims
}