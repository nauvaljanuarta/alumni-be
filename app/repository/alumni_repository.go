package repository

import (
	"context"
	"pert5/app/models"
	"time"
	// "database/sql"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type AlumniRepository struct {
// 	DB *sql.DB
// }

// func NewAlumniRepository(db *sql.DB) *AlumniRepository {
// 	return &AlumniRepository{DB: db}
// }

// func (r *AlumniRepository) GetByEmail(email string) (*models.Alumni, error) {
// 	var a models.Alumni
// 	err := r.DB.QueryRow(`
// 		SELECT id, nim, nama, role, fakultas, jurusan, angkatan,
// 		       tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at
// 		FROM alumni WHERE email=$1`, email).
// 		Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
// 			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.Password,
// 			&a.CreatedAt, &a.UpdatedAt)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// func (r *AlumniRepository) GetAlumni(search, sortBy, order string, limit, offset int) ([]models.Alumni, error) {
// 	query := fmt.Sprintf(`
// 		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email,
// 		       no_telepon, alamat, fakultas, role, password, created_at, updated_at
// 		FROM alumni
// 		WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
// 		      OR nim ILIKE $1 OR fakultas ILIKE $1
// 		ORDER BY %s %s
// 	`, sortBy, order)

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

// 	var alumni []models.Alumni
// 	for rows.Next() {
// 		var a models.Alumni
// 		if err := rows.Scan(
// 			&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
// 			&a.Email, &a.NoTelepon, &a.Alamat, &a.Fakultas, &a.Role,
// 			&a.Password, &a.CreatedAt, &a.UpdatedAt,
// 		); err != nil {
// 			return nil, err
// 		}
// 		alumni = append(alumni, a)
// 	}
// 	return alumni, nil
// }

// func (r *AlumniRepository) Count(search string) (int, error) {
// 	var total int
// 	countQuery := `
// 		SELECT COUNT(*) FROM alumni
// 		WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
// 		      OR nim ILIKE $1 OR fakultas ILIKE $1
// 	`
// 	err := r.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
// 	if err != nil && err != sql.ErrNoRows {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *AlumniRepository) GetByID(id int) (*models.Alumni, error) {
// 	var a models.Alumni
// 	err := r.DB.QueryRow(`SELECT id, nim, nama, role, fakultas, jurusan, angkatan,
// 		tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at
// 		FROM alumni WHERE id=$1`, id).
// 		Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
// 			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.Password,
// 			&a.CreatedAt, &a.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// func (r *AlumniRepository) GetByFakultas(fakultas string) ([]models.Alumni, error) {
// 	rows, err := r.DB.Query(`SELECT id, nim, nama, role, fakultas, jurusan, angkatan,
// 		tahun_lulus, email, no_telepon, alamat, password, created_at, updated_at
// 		FROM alumni WHERE fakultas=$1`, fakultas)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var alumniList []models.Alumni
// 	for rows.Next() {
// 		var a models.Alumni
// 		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Role, &a.Fakultas, &a.Jurusan, &a.Angkatan,
// 			&a.TahunLulus, &a.Email, &a.NoTelepon,
// 			&a.Alamat, &a.Password, &a.CreatedAt, &a.UpdatedAt); err != nil {
// 			return nil, err
// 		}
// 		alumniList = append(alumniList, a)
// 	}
// 	return alumniList, nil
// }

// func (r *AlumniRepository) Create(req models.CreateAlumni) (int, error) {
// 	var id int
// 	err := r.DB.QueryRow(`
// 		INSERT INTO alumni (nim, nama, role, fakultas, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, password)
// 		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`,
// 		req.NIM, req.Nama, req.Role, req.Fakultas, req.Jurusan, req.Angkatan,
// 		req.TahunLulus, req.Email, req.NoTelepon, req.Alamat, req.Password).Scan(&id)
// 	return id, err
// }

// func (r *AlumniRepository) Update(id int, req models.UpdateAlumni) error {
// 	_, err := r.DB.Exec(`
// 		UPDATE alumni SET nama=$1, role=$2, fakultas=$3, jurusan=$4, angkatan=$5,
// 		tahun_lulus=$6, email=$7, no_telepon=$8, alamat=$9, password=$10, updated_at=NOW()
// 		WHERE id=$11`,
// 		req.Nama, req.Role, req.Fakultas, req.Jurusan, req.Angkatan,
// 		req.TahunLulus, req.Email, req.NoTelepon, req.Alamat, req.Password, id)
// 	return err
// }

// func (r *AlumniRepository) Delete(id int) error {
// 	_, err := r.DB.Exec(`DELETE FROM alumni WHERE id=$1`, id)
// 	return err
// }

// interface untuk kontrak repository alumni
type IAlumniRepository interface {
	GetAlumni(ctx context.Context, search string, limit, offset int, sortBy, order string) ([]models.Alumni, error)
	GetAlumniByID(ctx context.Context, id string) (*models.Alumni, error)
	GetByEmail(ctx context.Context, email string) (*models.Alumni, error)
	CreateAlumni(ctx context.Context, req *models.Alumni) (*models.Alumni, error)
	UpdateAlumni(ctx context.Context, id string, req *models.Alumni) error
	DeleteAlumni(ctx context.Context, id string) error
	Count(ctx context.Context, search string) (int, error)
}

type AlumniRepository struct {
	collection *mongo.Collection
}

func NewAlumniRepository(db *mongo.Database) IAlumniRepository {
	return &AlumniRepository{
		collection: db.Collection("alumni"),
	}
}

func (r *AlumniRepository) Count(ctx context.Context, search string) (int, error) {
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama": bson.M{"$regex": search, "$options": "i"}},
				{"nim": bson.M{"$regex": search, "$options": "i"}},
				{"email": bson.M{"$regex": search, "$options": "i"}},
				{"jurusan": bson.M{"$regex": search, "$options": "i"}},
				{"fakultas": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	
	return int(count), nil
}

func (r *AlumniRepository) GetAlumni(ctx context.Context, search string, limit, offset int, sortBy, order string) ([]models.Alumni, error) {
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"nama": bson.M{"$regex": search, "$options": "i"}},
				{"nim": bson.M{"$regex": search, "$options": "i"}},
				{"email": bson.M{"$regex": search, "$options": "i"}},
				{"jurusan": bson.M{"$regex": search, "$options": "i"}},
				{"fakultas": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	findOptions := mongoOptions(limit, offset, sortBy, order)

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alumniList []models.Alumni
	if err := cursor.All(ctx, &alumniList); err != nil {
		return nil, err
	}

	return alumniList, nil
}

func mongoOptions(limit, offset int, sortBy, order string) *options.FindOptions {
	opts := &options.FindOptions{}
	if limit > 0 {
		opts.SetLimit(int64(limit))
		opts.SetSkip(int64(offset))
	}
	if sortBy != "" {
		sortOrder := 1
		if order == "desc" {
			sortOrder = -1
		}
		opts.SetSort(bson.D{{Key: sortBy, Value: sortOrder}})
	}
	return opts
}

func (r *AlumniRepository) GetByEmail(ctx context.Context, email string) (*models.Alumni, error) {
	var alumni models.Alumni
	filter := bson.M{"email": email}

	err := r.collection.FindOne(ctx, filter).Decode(&alumni)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &alumni, nil
}

func (r *AlumniRepository) GetAlumniByID(ctx context.Context, id string) (*models.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var alumni models.Alumni
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&alumni)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &alumni, nil
}

func (r *AlumniRepository) CreateAlumni(ctx context.Context, req *models.Alumni) (*models.Alumni, error) {
	req.ID = primitive.NilObjectID
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	req.ID = result.InsertedID.(primitive.ObjectID)
	return req, nil
}

func (r *AlumniRepository) UpdateAlumni(ctx context.Context, id string, req *models.Alumni) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	req.UpdatedAt = time.Now()
	update := bson.M{"$set": req}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *AlumniRepository) DeleteAlumni(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
