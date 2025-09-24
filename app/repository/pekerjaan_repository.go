package repository

import (
	"database/sql"
	"fmt"
	"pert5/app/models"
)

type PekerjaanRepository struct {
	DB *sql.DB
}

func NewPekerjaanRepository(db *sql.DB) *PekerjaanRepository {
	return &PekerjaanRepository{DB: db}
}
func (r *PekerjaanRepository) GetAllPekerjaan(search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
	query := fmt.Sprintf(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
				OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1
		ORDER BY %s %s`, sortBy, order)

	var rows *sql.Rows
	var err error
	if limit > 0 {
		rows, err = r.DB.Query(query+" LIMIT $2 OFFSET $3", "%"+search+"%", limit, offset)
	} else {
		rows, err = r.DB.Query(query, "%"+search+"%")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Pekerjaan
	for rows.Next() {
		var p models.Pekerjaan
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func (r *PekerjaanRepository) Count(search string) (int, error) {
	var total int
	countQuery := `
		SELECT COUNT(*) FROM pekerjaan_alumni 
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
			OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1
	`
	err := r.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}

func (r *PekerjaanRepository) GetByID(id int) (*models.Pekerjaan, error) {
	var p models.Pekerjaan
	err := r.DB.QueryRow(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
		FROM pekerjaan_alumni WHERE id=$1`, id).
		Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PekerjaanRepository) GetByAlumniID(alumniID int) ([]models.Pekerjaan, error) {
	rows, err := r.DB.Query(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
		FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY created_at DESC`, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Pekerjaan
	for rows.Next() {
		var p models.Pekerjaan
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func (r *PekerjaanRepository) Create(req models.CreatePekerjaan) (int, error) {
	var id int
	err := r.DB.QueryRow(`
		INSERT INTO pekerjaan_alumni (
			alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
			lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
			status_pekerjaan, deskripsi_pekerjaan
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
		req.AlumniID, req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri,
		req.LokasiKerja, req.GajiRange, req.TanggalMulaiKerja, req.TanggalSelesaiKerja,
		req.StatusPekerjaan, req.DeskripsiPekerjaan).Scan(&id)
	return id, err
}

func (r *PekerjaanRepository) Update(id int, req models.UpdatePekerjaan) error {
	_, err := r.DB.Exec(`
		UPDATE pekerjaan_alumni SET 
			nama_perusahaan=$1, 
			posisi_jabatan=$2, 
			bidang_industri=$3, 
			lokasi_kerja=$4, 
			gaji_range=$5, 
			tanggal_mulai_kerja=$6, 
			tanggal_selesai_kerja=$7, 
			status_pekerjaan=$8, 
			deskripsi_pekerjaan=$9, 
			updated_at=NOW() 
		WHERE id=$10`,
		req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri,
		req.LokasiKerja, req.GajiRange, req.TanggalMulaiKerja, req.TanggalSelesaiKerja,
		req.StatusPekerjaan, req.DeskripsiPekerjaan, id)
	return err
}

func (r *PekerjaanRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1`, id)
	return err
}

func (r *PekerjaanRepository) SoftDelete(id int,req models.UpdatePekerjaan) error {
	_, err := r.DB.Exec(`UPDATE pekerjaan_alumni SET isdelete=true, updated_at=NOW() WHERE id=$1`, id)
	return err
}

func (r *PekerjaanRepository) SoftDeleteBulk() error {
	query := `UPDATE pekerjaan_alumni SET isdelete=true, updated_at=NOW() WHERE isdelete=false`
	_, err := r.DB.Exec(query)
	return err
}
