package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"pert5/app/models"
	"pert5/app/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileService struct {
	repo       repository.IFileRepository
	uploadPath string
}

func NewFileService(repo repository.IFileRepository, uploadPath string) *FileService {
	return &FileService{
		repo:       repo,
		uploadPath: uploadPath,
	}
}

func (s *FileService) UploadFile(c *fiber.Ctx) error {
	return s.uploadFile(c, "")
}

// UploadFileAdmin godoc
// @Summary Upload file untuk alumni tertentu (admin only)
// @Description Admin mengunggah file untuk alumni tertentu berdasarkan ID
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param alumni_id path string true "Alumni ID"
// @Param file formData file true "File yang akan diunggah (PDF/JPG/PNG)"
// @Success 201 {object} models.FileResponse "File uploaded successfully"
// @Failure 400 {object} map[string]interface{} "Invalid alumni ID or file"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /pert5/files/upload/{alumni_id} [post]
func (s *FileService) UploadFileAdmin(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	if alumniID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Alumni ID is required",
		})
	}
	return s.uploadFile(c, alumniID)
}

// GetAllFiles godoc
// @Summary Mendapatkan semua file
// @Description Mengambil semua file dengan filter, pencarian, dan pagination
// @Tags Files
// @Produce json
// @Param search query string false "Pencarian berdasarkan nama file"
// @Param sortBy query string false "Kolom untuk pengurutan (default: uploadedAt)"
// @Param order query string false "Urutan asc/desc (default: desc)"
// @Param limit query int false "Jumlah data per halaman (default: 10)"
// @Param offset query int false "Offset data (default: 0)"
// @Success 200 {array} models.FileResponse "List of files"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /file [get]
func (s *FileService) uploadFile(c *fiber.Ctx, alumniIDStr string) error {
	userID := c.Locals("user_id").(primitive.ObjectID)
	var alumniID primitive.ObjectID
	var err error

	if alumniIDStr == "" {
		alumniID = userID
	} else {
		alumniID, err = primitive.ObjectIDFromHex(alumniIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "invalid alumni_id",
			})
		}
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "No file uploaded"})
	}

	// buka file dulu
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to open uploaded file"})
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	file.Seek(0, 0) // reset pointer
	detectedType := http.DetectContentType(buf[:n])

	var fileType string
	var maxSize int64

	switch detectedType {
	case "application/pdf":
		fileType = "document"
		maxSize = 2 * 1024 * 1024
	case "image/jpeg", "image/jpg", "image/png":
		fileType = "image"
		maxSize = 1 * 1024 * 1024
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "File type not allowed"})
	}

	if fileHeader.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": fmt.Sprintf("%s exceeds max size", fileType)})
	}

	ext := filepath.Ext(fileHeader.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(s.uploadPath, newFileName)

	if err := os.MkdirAll(s.uploadPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create upload directory"})
	}

	out, err := os.Create(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create file on server"})
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to write file to server"})
	}

	fileModel := &models.File{
		FileName:     newFileName,
		OriginalName: fileHeader.Filename,
		FilePath:     filePath,
		FileSize:     fileHeader.Size,
		FileType:     detectedType,
		UploadedAt:   time.Now(),
		AlumniID:     alumniID,
	}

	savedFile, err := s.repo.Create(context.Background(), *fileModel)
	if err != nil {
		os.Remove(filePath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to save file metadata"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("%s uploaded successfully", fileType),
		"data":    s.toFileResponse(savedFile),
	})
}

// GetAllFiles godoc
// @Summary Mendapatkan semua file
// @Description Mengambil semua file dengan filter, pencarian, dan pagination
// @Tags Files
// @Produce json
// @Param search query string false "Pencarian berdasarkan nama file"
// @Param sortBy query string false "Kolom untuk pengurutan (default: uploadedAt)"
// @Param order query string false "Urutan asc/desc (default: desc)"
// @Param limit query int false "Jumlah data per halaman (default: 10)"
// @Param offset query int false "Offset data (default: 0)"
// @Success 200 {array} models.FileResponse "List of files"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /file [get]
// Get all files
func (s *FileService) GetAllFiles(c *fiber.Ctx) error {
	// Get query parameters
	search := c.Query("search", "")
	sortBy := c.Query("sortBy", "uploadedAt")
	order := c.Query("order", "desc")
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	files, err := s.repo.GetAll(context.Background(), search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get files",
			"error":   err.Error(),
		})
	}

	var responses []models.FileResponse
	for _, f := range files {
		responses = append(responses, *s.toFileResponse(&f))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Files retrieved successfully",
		"data":    responses,
	})
}

// GetFileByID godoc
// @Summary Mendapatkan file berdasarkan ID
// @Description Mengambil file tunggal berdasarkan ID file
// @Tags Files
// @Produce json
// @Param id path string true "ID File"
// @Success 200 {object} models.FileResponse "File retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 404 {object} map[string]interface{} "File not found"
// @Security BearerAuth
// @Router /file/{id} [get]
// Get file by ID
func (s *FileService) GetFileByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
	}

	file, err := s.repo.GetByID(context.Background(), objectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File retrieved successfully",
		"data":    s.toFileResponse(file),
	})
}

// GetFilesByAlumniID godoc
// @Summary Mendapatkan file berdasarkan Alumni ID
// @Description Mengambil semua file yang dimiliki oleh alumni tertentu
// @Tags Files
// @Produce json
// @Param alumniID path string true "Alumni ID"
// @Success 200 {array} models.FileResponse "Files retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /file/alumni/{alumniID} [get]
func (s *FileService) GetFilesByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumniID")
	files, err := s.repo.GetByAlumniID(context.Background(), alumniID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get files",
			"error":   err.Error(),
		})
	}

	var responses []models.FileResponse
	for _, f := range files {
		responses = append(responses, *s.toFileResponse(&f))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Files retrieved successfully",
		"data":    responses,
	})
}

// UpdateFile godoc
// @Summary Update metadata file
// @Description Memperbarui informasi file seperti nama, deskripsi, atau jenis file
// @Tags Files
// @Accept json
// @Produce json
// @Param id path string true "ID File"
// @Param body body models.File true "Data file yang akan diperbarui"
// @Success 200 {object} models.FileResponse "File updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "File not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /file/{id} [put]
func (s *FileService) UpdateFile(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
	}
	file, err := s.repo.GetByID(context.Background(), objectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
			"error":   err.Error(),
		})
	}

	if err := c.BodyParser(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := s.repo.Update(context.Background(), objectID, *file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update file",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File updated successfully",
		"data":    s.toFileResponse(file),
	})
}

// DeleteFile godoc
// @Summary Menghapus file
// @Description Menghapus file dari server dan database
// @Tags Files
// @Produce json
// @Param id path string true "ID File"
// @Success 200 {object} map[string]interface{} "File deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 404 {object} map[string]interface{} "File not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete file"
// @Router /file/{id} [delete]
// Delete file (hard delete)
func (s *FileService) DeleteFile(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
			"error":   err.Error(),
		})
	}
	file, err := s.repo.GetByID(context.Background(), objectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
			"error":   err.Error(),
		})
	}

	// Hapus file dari storage
	if err := os.Remove(file.FilePath); err != nil {
		fmt.Println("Warning: Failed to delete file from storage:", err)
	}

	// Hapus dari database
	if err := s.repo.Delete(context.Background(), objectID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete file",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File deleted successfully",
	})
}

// Helper to convert to response
func (s *FileService) toFileResponse(file *models.File) *models.FileResponse {
	return &models.FileResponse{
		ID:           file.ID.Hex(),
		AlumniID:     file.AlumniID.Hex(),
		FileName:     file.FileName,
		OriginalName: file.OriginalName,
		FilePath:     file.FilePath,
		FileSize:     file.FileSize,
		FileType:     file.FileType,
		UploadedAt:   file.UploadedAt,
	}
}
