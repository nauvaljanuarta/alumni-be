package service


import (
	"pert5/app/models"
	"pert5/app/repository"
	"strconv"
	"strings"
	"time"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// type PekerjaanService struct {
// 	repo *repository.PekerjaanRepository
// }

// func NewPekerjaanService(repo *repository.PekerjaanRepository) *PekerjaanService {
// 	return &PekerjaanService{repo: repo}
// }

// func (s *PekerjaanService) GetAll(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit, _ := strconv.Atoi(c.Query("limit", "10"))
// 	sortBy := c.Query("sortBy", "created_at") 
// 	order := c.Query("order", "desc")
// 	search := c.Query("search", "")

// 	offset := (page - 1) * limit

// 	sortByWhitelist := map[string]bool{
// 		"id": true, "alumni_id": true, "nama_perusahaan": true,
// 		"posisi_jabatan": true, "bidang_industri": true,
// 		"lokasi_kerja": true, "status_pekerjaan": true,
// 		"created_at": true,
// 	}
// 	if !sortByWhitelist[sortBy] {
// 		sortBy = "created_at"
// 	}

// 	if strings.ToLower(order) != "desc" {
// 		order = "asc"
// 	}

// 	data, err := s.repo.GetAllPekerjaan(search, sortBy, order, limit, offset)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "pekerjaan not found"})
// 	}

// 	total, err := s.repo.Count(search)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "gagal menghitung pekerjaan"})
// 	}

// 	response := models.PekerjaanResponse{
// 		Data: data,
// 		Meta: models.MetaInfo{
// 			Page:   page,
// 			Limit:  limit,
// 			Total:  total,
// 			Pages:  (total + limit - 1) / limit,
// 			SortBy: sortBy,
// 			Order:  order,
// 			Search: search,
// 		},
// 	}

// 	return c.JSON(response)
// }

// func (s *PekerjaanService) GetByID(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	data, err := s.repo.GetByID(id)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
// 	}
// 	return c.JSON(data)
// }

// func (s *PekerjaanService) GetByAlumniID(c *fiber.Ctx) error {
// 	alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
// 	data, err := s.repo.GetByAlumniID(alumniID)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(data)
// }

// func (s *PekerjaanService) Create(c *fiber.Ctx) error {
// 	var req models.CreatePekerjaan
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
// 	}
// 	id, err := s.repo.Create(req)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.Status(201).JSON(fiber.Map{"id": id, "message": "pekerjaan created"})
// }

// func (s *PekerjaanService) Update(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	var req models.UpdatePekerjaan
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
// 	}
// 	if err := s.repo.Update(id, req); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "pekerjaan updated"})
// }

// func (s *PekerjaanService) Delete(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	if err := s.repo.Delete(id); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "pekerjaan deleted"})
// }

// func (s *PekerjaanService) SoftDelete(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
	
// 	userID := c.Locals("user_id")
// 	role := c.Locals("role").(string)
// 	isAdmin := role == "admin"
	
// 	existingData, err := s.repo.GetByID(id)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
// 	}
	
// 	if existingData.AlumniID != userID && !isAdmin {
// 		return c.Status(403).JSON(fiber.Map{"message": "bukan pekerjaanmu dan bukan admin"})
// 	}
	
// 	var updateReq models.UpdatePekerjaan
// 	if err := s.repo.SoftDelete(id, updateReq); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "pekerjaan soft deleted"})
// }

// func (s *PekerjaanService) SoftDeleteBulk(c *fiber.Ctx) error {

// 	role := c.Locals("role").(string)
// 	if role != "admin" {
// 		return c.Status(403).JSON(fiber.Map{"message": "unauthorized: admin access required"})
// 	}
	
// 	if err := s.repo.SoftDeleteBulk(); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "all pekerjaan soft deleted"})
// }

// func (s *PekerjaanService) Restore(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
	
// 	userID := c.Locals("user_id")
// 	role := c.Locals("role").(string)
// 	isAdmin := role == "admin"
	
// 	existingData, err := s.repo.GetByID(id)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
// 	}
	
// 	if existingData.AlumniID != userID && !isAdmin {
// 		return c.Status(403).JSON(fiber.Map{"message": "bukan pekerjaanmu dan bukan admin"})
// 	}
	
	
// 	var updateReq models.UpdatePekerjaan
// 	if err := s.repo.Restore(id, updateReq); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "pekerjaan restored"})
// }

// func (s *PekerjaanService) GetTrash(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit, _ := strconv.Atoi(c.Query("limit", "10"))
// 	sortBy := c.Query("sortBy", "created_at") 
// 	order := c.Query("order", "desc")
// 	search := c.Query("search", "")

// 	offset := (page - 1) * limit

// 	sortByWhitelist := map[string]bool{
// 		"id": true, "alumni_id": true, "nama_perusahaan": true,
// 		"posisi_jabatan": true, "bidang_industri": true,
// 		"lokasi_kerja": true, "status_pekerjaan": true,
// 		"created_at": true,
// 	}
// 	if !sortByWhitelist[sortBy] {
// 		sortBy = "created_at"
// 	}

// 	if strings.ToLower(order) != "desc" {
// 		order = "asc"
// 	}

// 	data, err := s.repo.GetTrash(search, sortBy, order, limit, offset)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "pekerjaan trash not found"})
// 	}

// 	total, err := s.repo.Count(search)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "gagal menghitung pekerjaan"})
// 	}

// 	response := models.PekerjaanResponse{
// 		Data: data,
// 		Meta: models.MetaInfo{
// 			Page:   page,
// 			Limit:  limit,
// 			Total:  total,
// 			Pages:  (total + limit - 1) / limit,
// 			SortBy: sortBy,
// 			Order:  order,
// 			Search: search,
// 		},
// 	}

// 	return c.JSON(response)
// }

// func (s *PekerjaanService) DeleteTrash(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	if err := s.repo.DeleteTrash(id); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "pekerjaan trash deleted hardly"})
// }



type PekerjaanService struct {
	repo repository.IPekerjaanRepository //  ubah ke interface
}

func NewPekerjaanService(repo repository.IPekerjaanRepository) *PekerjaanService {  //ubah ke interface
	return &PekerjaanService{repo: repo}
}

// GetAll godoc
// @Summary Get all pekerjaan
// @Description Mendapatkan daftar pekerjaan dengan pagination, sorting, dan pencarian
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param page query int false "Nomor halaman"
// @Param limit query int false "Jumlah data per halaman"
// @Param sortBy query string false "Kolom untuk sorting (default: created_at)"
// @Param order query string false "Urutan sort (asc/desc)"
// @Param search query string false "Kata kunci pencarian"
// @Success 200 {object} models.PekerjaanResponse
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan [get]
func (s *PekerjaanService) GetAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "created_at") 
	order := c.Query("order", "desc")
	search := c.Query("search", "")

	offset := (page - 1) * limit

	sortByWhitelist := map[string]bool{
		"_id": true, "alumni_id": true, "nama_perusahaan": true,
		"posisi_jabatan": true, "bidang_industri": true,
		"lokasi_kerja": true, "status_pekerjaan": true,
		"created_at": true, "updated_at": true,
	}
	if !sortByWhitelist[sortBy] {
		sortBy = "created_at"
	}

	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	data, err := s.repo.GetAll(ctx, search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "pekerjaan not found"})
	}

	total, err := s.repo.Count(ctx, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal menghitung pekerjaan"})
	}

	response := models.PekerjaanResponse{
		Data: data,
		Meta: models.MetaInfo{
			Page:   page,
			Limit:  limit,
			Total:  int(total), // ← Convert int64 ke int
			Pages:  (int(total) + limit - 1) / limit,
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}

	return c.JSON(response)
}

// GetByID godoc
// @Summary Get pekerjaan by ID
// @Description Mendapatkan detail pekerjaan berdasarkan ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Success 200 {object} models.Pekerjaan
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/{id} [get]
func (s *PekerjaanService) GetByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}

	data, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	if data == nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	
	return c.JSON(data)
}

// GetByAlumniID godoc
// @Summary Get pekerjaan by Alumni ID
// @Description Mendapatkan daftar pekerjaan berdasarkan alumni ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param alumni_id path string true "ID Alumni"
// @Success 200 {array} models.Pekerjaan
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/alumni/{alumni_id} [get]
func (s *PekerjaanService) GetByAlumniID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	alumniIDStr := c.Params("alumni_id")
	alumniID, err := primitive.ObjectIDFromHex(alumniIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid alumni ID format"})
	}

	data, err := s.repo.GetByAlumniID(ctx, alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(data)
}

// Create godoc
// @Summary Create pekerjaan baru
// @Description Menambahkan data pekerjaan baru
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param request body models.CreatePekerjaan true "Data pekerjaan baru"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan [post]
func (s *PekerjaanService) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var req models.CreatePekerjaan
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
	}

	result, err := s.repo.Create(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":      result.InsertedID,
		"message": "pekerjaan created",
	})
}

// Update godoc
// @Summary Update pekerjaan
// @Description Memperbarui data pekerjaan berdasarkan ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Param request body models.UpdatePekerjaan true "Data pekerjaan yang diperbarui"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/{id} [put]
func (s *PekerjaanService) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}

	var req models.UpdatePekerjaan
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
	}

	if err := s.repo.Update(ctx, objID, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "pekerjaan updated"})
}

// Delete godoc
// @Summary Hapus pekerjaan permanen
// @Description Menghapus pekerjaan dari database secara permanen
// @Tags Pekerjaan
// @Param id path string true "ID Pekerjaan"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/{id} [delete]
func (s *PekerjaanService) Delete(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}

	if err := s.repo.Delete(ctx, objID); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "pekerjaan deleted"})
}

// SoftDelete godoc
// @Summary Soft delete pekerjaan
// @Description Menghapus pekerjaan tanpa menghapus data di database
// @Tags Pekerjaan
// @Param id path string true "ID Pekerjaan"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/soft/{id} [delete]
func (s *PekerjaanService) SoftDelete(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}
	
	userID := c.Locals("user_id")
	role := c.Locals("role").(string)
	isAdmin := role == "admin"
	
	existingData, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	if existingData == nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	
	userObjID, ok := userID.(primitive.ObjectID)
	if !ok {
		if userIDStr, ok := userID.(string); ok {
			userObjID, err = primitive.ObjectIDFromHex(userIDStr) // ini convert format mongo ke string
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"message": "invalid user ID"})
			}
		} else {
			return c.Status(400).JSON(fiber.Map{"message": "invalid user ID format"})
		}
	}
	
	if existingData.AlumniID != userObjID && !isAdmin {
		return c.Status(403).JSON(fiber.Map{"message": "bukan pekerjaanmu dan bukan admin"})
	}
	
	if err := s.repo.SoftDelete(ctx, objID); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "pekerjaan soft deleted"})
}

// SoftDeleteBulk godoc
// @Summary Soft delete semua pekerjaan
// @Description Menghapus semua pekerjaan (hanya untuk admin)
// @Tags Pekerjaan
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/soft/all [delete]
func (s *PekerjaanService) SoftDeleteBulk(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 30*time.Second) 
	defer cancel()

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"message": "unauthorized: admin access required"})
	}
	
	if err := s.repo.SoftDeleteBulk(ctx); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "all pekerjaan soft deleted"})
}

// Restore godoc
// @Summary Restore pekerjaan
// @Description Mengembalikan pekerjaan yang dihapus (soft delete)
// @Tags Pekerjaan
// @Param id path string true "ID Pekerjaan"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/restore/{id} [put]
func (s *PekerjaanService) Restore(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}
	
	userID := c.Locals("user_id")
	role := c.Locals("role").(string)
	isAdmin := role == "admin"
	
	existingData, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	if existingData == nil {
		return c.Status(404).JSON(fiber.Map{"message": "pekerjaan not found"})
	}
	
	userObjID, ok := userID.(primitive.ObjectID)
	if !ok {
		if userIDStr, ok := userID.(string); ok {
			userObjID, err = primitive.ObjectIDFromHex(userIDStr)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"message": "invalid user ID"})
			}
		} else {
			return c.Status(400).JSON(fiber.Map{"message": "invalid user ID format"})
		}
	}
	
	if existingData.AlumniID != userObjID && !isAdmin {
		return c.Status(403).JSON(fiber.Map{"message": "bukan pekerjaanmu dan bukan admin"})
	}
	
	if err := s.repo.Restore(ctx, objID); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "pekerjaan restored"})
}

// GetTrash godoc
// @Summary Get semua pekerjaan yang dihapus (trash)
// @Description Mendapatkan daftar pekerjaan yang sudah dihapus secara soft delete
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param page query int false "Nomor halaman"
// @Param limit query int false "Jumlah data per halaman"
// @Success 200 {object} models.PekerjaanResponse
// @Security BearerAuth
// @Router /pekerjaan/trash [get]
func (s *PekerjaanService) GetTrash(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "updated_at") 
	order := c.Query("order", "desc")
	search := c.Query("search", "")

	offset := (page - 1) * limit

	sortByWhitelist := map[string]bool{
		"_id": true, "alumni_id": true, "nama_perusahaan": true,
		"posisi_jabatan": true, "bidang_industri": true,
		"lokasi_kerja": true, "status_pekerjaan": true,
		"created_at": true, "updated_at": true,
	}
	if !sortByWhitelist[sortBy] {
		sortBy = "updated_at"
	}

	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	data, err := s.repo.GetTrash(ctx, search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "pekerjaan trash not found"})
	}

	total, err := s.repo.CountTrash(ctx, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal menghitung pekerjaan trash"})
	}

	response := models.PekerjaanResponse{
		Data: data,
		Meta: models.MetaInfo{
			Page:   page,
			Limit:  limit,
			Total:  int(total), // ← Convert int64 ke int
			Pages:  (int(total) + limit - 1) / limit,
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}

	return c.JSON(response)
}

// DeleteTrash godoc
// @Summary Hapus pekerjaan dari trash (permanen)
// @Description Menghapus pekerjaan yang sudah dihapus (trash) secara permanen
// @Tags Pekerjaan
// @Param id path string true "ID Pekerjaan"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /pekerjaan/trash/{id} [delete]
func (s *PekerjaanService) DeleteTrash(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid ID format"})
	}

	if err := s.repo.DeleteTrash(ctx, objID); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.JSON(fiber.Map{"message": "pekerjaan trash deleted permanently"})
}
