package repository

import (
	// "database/sql"
	// "fmt"
	"context"
	"time"
	"pert5/app/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type PekerjaanRepository struct {
// 	DB *sql.DB
// }

// func NewPekerjaanRepository(db *sql.DB) *PekerjaanRepository {
// 	return &PekerjaanRepository{DB: db}
// }

// func (r *PekerjaanRepository) GetAllPekerjaan(search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
// 	query := fmt.Sprintf(`
// 		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
// 				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
// 				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, isdeleted
// 		FROM pekerjaan_alumni
// 		WHERE isdeleted=false AND (nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
// 				OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1)
// 		ORDER BY %s %s`, sortBy, order)

// 	var rows *sql.Rows
// 	var err error
// 	if limit > 0 {
// 		rows, err = r.DB.Query(query+" LIMIT $2 OFFSET $3", "%"+search+"%", limit, offset)
// 	} else {
// 		rows, err = r.DB.Query(query, "%"+search+"%")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var list []models.Pekerjaan
// 	for rows.Next() {
// 		var p models.Pekerjaan
// 		if err := rows.Scan(
// 			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt, &p.IsDeleted); err != nil {
// 			return nil, err
// 		}
// 		list = append(list, p)
// 	}
// 	return list, nil
// }

// func (r *PekerjaanRepository) Count(search string) (int, error) {
// 	var total int
// 	countQuery := `
// 		SELECT COUNT(*) FROM pekerjaan_alumni 
// 		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
// 			OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1
// 	`
// 	err := r.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
// 	if err != nil && err != sql.ErrNoRows {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *PekerjaanRepository) GetByID(id int) (*models.Pekerjaan, error) {
// 	var p models.Pekerjaan
// 	err := r.DB.QueryRow(`
// 		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
// 				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
// 				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
// 		FROM pekerjaan_alumni WHERE id=$1`, id).
// 		Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }

// func (r *PekerjaanRepository) GetByAlumniID(alumniID int) ([]models.Pekerjaan, error) {
// 	rows, err := r.DB.Query(`
// 		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
// 				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
// 				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
// 		FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY created_at DESC`, alumniID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var list []models.Pekerjaan
// 	for rows.Next() {
// 		var p models.Pekerjaan
// 		if err := rows.Scan(
// 			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
// 			return nil, err
// 		}
// 		list = append(list, p)
// 	}
// 	return list, nil
// }

// func (r *PekerjaanRepository) Create(req models.CreatePekerjaan) (int, error) {
// 	var id int
// 	err := r.DB.QueryRow(`
// 		INSERT INTO pekerjaan_alumni (
// 			alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
// 			lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
// 			status_pekerjaan, deskripsi_pekerjaan
// 		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
// 		req.AlumniID, req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri,
// 		req.LokasiKerja, req.GajiRange, req.TanggalMulaiKerja, req.TanggalSelesaiKerja,
// 		req.StatusPekerjaan, req.DeskripsiPekerjaan).Scan(&id)
// 	return id, err
// }

// func (r *PekerjaanRepository) Update(id int, req models.UpdatePekerjaan) error {
// 	_, err := r.DB.Exec(`
// 		UPDATE pekerjaan_alumni SET 
// 			nama_perusahaan=$1, 
// 			posisi_jabatan=$2, 
// 			bidang_industri=$3, 
// 			lokasi_kerja=$4, 
// 			gaji_range=$5, 
// 			tanggal_mulai_kerja=$6, 
// 			tanggal_selesai_kerja=$7, 
// 			status_pekerjaan=$8, 
// 			deskripsi_pekerjaan=$9, 
// 			updated_at=NOW() 
// 		WHERE id=$10`,
// 		req.NamaPerusahaan, req.PosisiJabatan, req.BidangIndustri,
// 		req.LokasiKerja, req.GajiRange, req.TanggalMulaiKerja, req.TanggalSelesaiKerja,
// 		req.StatusPekerjaan, req.DeskripsiPekerjaan, id)
// 	return err
// }

// func (r *PekerjaanRepository) Delete(id int) error {
// 	_, err := r.DB.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1`, id)
// 	return err
// }

// func (r *PekerjaanRepository) SoftDelete(id int,req models.UpdatePekerjaan) error {
// 	_, err := r.DB.Exec(`UPDATE pekerjaan_alumni SET isdeleted=true, updated_at=NOW() WHERE id=$1`, id)
// 	return err
// }

// func (r *PekerjaanRepository) SoftDeleteBulk() error {
// 	query := `UPDATE pekerjaan_alumni SET isdeleted=true, updated_at=NOW() WHERE isdeleted=false`
// 	_, err := r.DB.Exec(query)
// 	return err
// }

// func (r *PekerjaanRepository) Restore(id int,req models.UpdatePekerjaan) error {
// 	_, err := r.DB.Exec(`UPDATE pekerjaan_alumni SET isdeleted=false, updated_at=NOW() WHERE id=$1`, id)
// 	return err
// }

// func (r *PekerjaanRepository) GetTrash(search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
// 	query := fmt.Sprintf(`
// 		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
// 				 lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
// 				 status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, isdeleted
// 		FROM pekerjaan_alumni
// 		WHERE isdeleted=true AND (nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
// 				OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1)
// 		ORDER BY %s %s`, sortBy, order)

// 	var rows *sql.Rows
// 	var err error
// 	if limit > 0 {
// 		rows, err = r.DB.Query(query+" LIMIT $2 OFFSET $3", "%"+search+"%", limit, offset)
// 	} else {
// 		rows, err = r.DB.Query(query, "%"+search+"%")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var list []models.Pekerjaan
// 	for rows.Next() {
// 		var p models.Pekerjaan
// 		if err := rows.Scan(
// 			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,&p.IsDeleted); err != nil {
// 			return nil, err
// 		}
// 		list = append(list, p)
// 	}
// 	return list, nil
// }

// func (r *PekerjaanRepository) CountTrash(search string) (int, error) {
// 	var total int
// 	countQuery := `
// 		SELECT COUNT(*) FROM pekerjaan_alumni 
// 		WHERE isdeleted=true AND (nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 
// 			OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1)
// 	`
// 	err := r.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
// 	if err != nil && err != sql.ErrNoRows {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *PekerjaanRepository) DeleteTrash(id int) error {
// 	_, err := r.DB.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1 AND isdeleted=true`, id)
// 	return err
// }

// interface kkontrak dengan 
type IPekerjaanRepository interface {
	GetAll(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error)
	Count(ctx context.Context, search string) (int64, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error)
	GetByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]models.Pekerjaan, error)
	Create(ctx context.Context, req models.CreatePekerjaan) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, id primitive.ObjectID, req models.UpdatePekerjaan) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	SoftDelete(ctx context.Context, id primitive.ObjectID) error
	SoftDeleteBulk(ctx context.Context) error
	Restore(ctx context.Context, id primitive.ObjectID) error
	GetTrash(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error)
	CountTrash(ctx context.Context, search string) (int64, error)
	DeleteTrash(ctx context.Context, id primitive.ObjectID) error
}

type pekerjaanRepository struct {
	collection *mongo.Collection
}

func NewPekerjaanRepository(db *mongo.Database) IPekerjaanRepository {
	return &pekerjaanRepository{
		collection: db.Collection("pekerjaan_alumni"),
	}
}

func (r *pekerjaanRepository) GetAll(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
	filter := bson.M{"is_deleted": false}

	if search != "" {
		filter["$or"] = []bson.M{
			{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
			{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
			{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
			{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	sortOrder := 1
	if order == "DESC" || order == "desc" {
		sortOrder = -1
	}
	
	if sortBy == "" {
		sortBy = "created_at"
	}

	opts := options.Find().
		SetSort(bson.M{sortBy: sortOrder}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	
	return list, nil
}

func (r *pekerjaanRepository) Count(ctx context.Context, search string) (int64, error) {
	filter := bson.M{"isdeleted": false}
	
	if search != "" {
		filter["$or"] = []bson.M{
			{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
			{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
			{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
			{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	
	return r.collection.CountDocuments(ctx, filter)
}

func (r *pekerjaanRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error) {
	var pekerjaan models.Pekerjaan
	filter := bson.M{"_id": id}
	
	err := r.collection.FindOne(ctx, filter).Decode(&pekerjaan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Tidak ditemukan
		}
		return nil, err
	}
	
	return &pekerjaan, nil
}

func (r *pekerjaanRepository) GetByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]models.Pekerjaan, error) {
	filter := bson.M{
		"alumni_id": alumniID,
		"is_deleted": false,
	}
	
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	
	return list, nil
}

func (r *pekerjaanRepository) Create(ctx context.Context, req models.CreatePekerjaan) (*mongo.InsertOneResult, error) {
	// Convert AlumniID ke ObjectID
	alumniID, err := primitive.ObjectIDFromHex(req.AlumniID)
	if err != nil {
			return nil, errors.New("invalid alumni_id")
	}

	// Parse tanggal mulai
	tMulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
	if err != nil {
			return nil, errors.New("invalid tanggal_mulai_kerja")
	}

	// Parse tanggal selesai jika ada
	var tSelesai *time.Time
	if req.TanggalSelesaiKerja != nil && *req.TanggalSelesaiKerja != "" {
			ts, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
			if err != nil {
					return nil, errors.New("invalid tanggal_selesai_kerja")
			}
			tSelesai = &ts
	}

	now := time.Now()
	doc := bson.M{
			"alumni_id": alumniID,
			"nama_perusahaan": req.NamaPerusahaan,
			"posisi_jabatan": req.PosisiJabatan,
			"bidang_industri": req.BidangIndustri,
			"lokasi_kerja": req.LokasiKerja,
			"gaji_range": req.GajiRange,
			"tanggal_mulai_kerja": tMulai,
			"tanggal_selesai_kerja": tSelesai,
			"status_pekerjaan": req.StatusPekerjaan,
			"deskripsi_pekerjaan": req.DeskripsiPekerjaan,
			"created_at": now,
			"updated_at": now,
			"is_deleted": false,
	}

	return r.collection.InsertOne(ctx, doc)
}

func (r *pekerjaanRepository) Update(ctx context.Context, id primitive.ObjectID, req models.UpdatePekerjaan) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_perusahaan":       req.NamaPerusahaan,
			"posisi_jabatan":        req.PosisiJabatan,
			"bidang_industri":       req.BidangIndustri,
			"lokasi_kerja":          req.LokasiKerja,
			"gaji_range":            req.GajiRange,
			"tanggal_mulai_kerja":   req.TanggalMulaiKerja,
			"tanggal_selesai_kerja": req.TanggalSelesaiKerja,
			"status_pekerjaan":      req.StatusPekerjaan,
			"deskripsi_pekerjaan":   req.DeskripsiPekerjaan,
			"updated_at":            time.Now(),
		},
	}
	
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}

func (r *pekerjaanRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}

func (r *pekerjaanRepository) SoftDelete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id":       id,
		"is_deleted": false, 
	}
	update := bson.M{
		"$set": bson.M{
			"is_deleted":  true,
			"updated_at": time.Now(),
		},
	}
	
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}

func (r *pekerjaanRepository) SoftDeleteBulk(ctx context.Context) error {
	filter := bson.M{"is_deleted": false}
	update := bson.M{
		"$set": bson.M{
			"is_deleted":  true,
			"updated_at": time.Now(),
		},
	}
	
	_, err := r.collection.UpdateMany(ctx, filter, update)
	return err
}

func (r *pekerjaanRepository) Restore(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id":       id,
		"is_deleted": true, 
	}
	update := bson.M{
		"$set": bson.M{
			"is_deleted":  false,
			"updated_at": time.Now(),
		},
	}
	
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}

func (r *pekerjaanRepository) GetTrash(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
	filter := bson.M{"is_deleted": true}
	
	// Tambahkan search filter 
	if search != "" {
		filter["$or"] = []bson.M{
			{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
			{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
			{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
			{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	sortOrder := 1
	if order == "DESC" || order == "desc" {
		sortOrder = -1
	}
	
	if sortBy == "" {
		sortBy = "updated_at" 
	}

	opts := options.Find().
		SetSort(bson.M{sortBy: sortOrder}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []models.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	
	return list, nil
}

// CountTrash - Menghitung total pekerjaan di trash
func (r *pekerjaanRepository) CountTrash(ctx context.Context, search string) (int64, error) {
	filter := bson.M{"is_deleted": true}
	
	if search != "" {
		filter["$or"] = []bson.M{
			{"nama_perusahaan": bson.M{"$regex": search, "$options": "i"}},
			{"posisi_jabatan": bson.M{"$regex": search, "$options": "i"}},
			{"bidang_industri": bson.M{"$regex": search, "$options": "i"}},
			{"lokasi_kerja": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	
	return r.collection.CountDocuments(ctx, filter)
}

// DeleteTrash - Menghapus pekerjaan dari trash secara permanen
func (r *pekerjaanRepository) DeleteTrash(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id":       id,
		"is_deleted": true, 
	}
	
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}