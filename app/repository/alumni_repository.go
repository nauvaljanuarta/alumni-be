package repository

import (
	"database/sql"
	"pert5/app/models"
	"fmt"
)
type AlumniRepository struct {
	DB *sql.DB
}

func NewAlumniRepository(db *sql.DB) *AlumniRepository {
	return &AlumniRepository{DB: db}
}

func (r *AlumniRepository) GetByEmail(email string) (*models.Alumni, error) {
	var a models.Alumni
	err := r.DB.QueryRow(`
		SELECT id, nim, nama, role, fakultas, jurusan, angkatan, 
		       tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at
		FROM alumni WHERE email=$1`, email).
		Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.Password,
			&a.CreatedAt, &a.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AlumniRepository) GetAlumni(search, sortBy, order string, limit, offset int) ([]models.Alumni, error) {
	query := fmt.Sprintf(`
		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email,
		       no_telepon, alamat, fakultas, role, password, created_at, updated_at
		FROM alumni
		WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
		      OR nim ILIKE $1 OR fakultas ILIKE $1
		ORDER BY %s %s
	`, sortBy, order)

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

	var alumni []models.Alumni
	for rows.Next() {
		var a models.Alumni
		if err := rows.Scan(
			&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
			&a.Email, &a.NoTelepon, &a.Alamat, &a.Fakultas, &a.Role,
			&a.Password, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		alumni = append(alumni, a)
	}
	return alumni, nil
}

func (r *AlumniRepository) Count(search string) (int, error) {
	var total int
	countQuery := `
		SELECT COUNT(*) FROM alumni
		WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
		      OR nim ILIKE $1 OR fakultas ILIKE $1
	`
	err := r.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}


func (r *AlumniRepository) GetByID(id int) (*models.Alumni, error) {
	var a models.Alumni
	err := r.DB.QueryRow(`SELECT id, nim, nama, role, fakultas, jurusan, angkatan, 
		tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at 
		FROM alumni WHERE id=$1`, id).
		Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.Password,
			&a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AlumniRepository) GetByFakultas(fakultas string) ([]models.Alumni, error) {
	rows, err := r.DB.Query(`SELECT id, nim, nama, role, fakultas, jurusan, angkatan, 
		tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at 
		FROM alumni WHERE fakultas=$1`, fakultas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alumniList []models.Alumni
	for rows.Next() {
		var a models.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon,
			&a.Alamat, &a.Password, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}
	return alumniList, nil
}

func (r *AlumniRepository) Create(req models.CreateAlumni) (int, error) {
	var id int
	err := r.DB.QueryRow(`
		INSERT INTO alumni (nim, nama, role, fakultas, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, password)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`,
		req.NIM, req.Nama, req.Role, req.Fakultas, req.Jurusan, req.Angkatan, 
		req.TahunLulus, req.Email, req.NoTelepon, req.Alamat, req.Password).Scan(&id)
	return id, err
}

func (r *AlumniRepository) Update(id int, req models.UpdateAlumni) error {
	_, err := r.DB.Exec(`
		UPDATE alumni SET nama=$1, role=$2, fakultas=$3, jurusan=$4, angkatan=$5, 
		tahun_lulus=$6, email=$7, no_telepon=$8, alamat=$9, password=$10, updated_at=NOW() 
		WHERE id=$11`,
		req.Nama, req.Role, req.Fakultas, req.Jurusan, req.Angkatan,
		req.TahunLulus, req.Email, req.NoTelepon, req.Alamat, req.Password, id)
	return err
}

func (r *AlumniRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM alumni WHERE id=$1`, id)
	return err
}
