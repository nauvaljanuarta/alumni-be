package service

import (
	"context"
	"errors"
	"pert5/app/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockFileRepository struct {
	MockGetAll        func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.File, error)
	MockCount         func(ctx context.Context, search string) (int64, error)
	MockGetByID       func(ctx context.Context, id primitive.ObjectID) (*models.File, error)
	MockGetByAlumniID func(ctx context.Context, alumniID string) ([]models.File, error)
	MockCreate        func(ctx context.Context, file models.File) (*models.File, error)
	MockUpdate        func(ctx context.Context, id primitive.ObjectID, file models.File) error
	MockDelete        func(ctx context.Context, id primitive.ObjectID) error
}

func (m *MockFileRepository) GetAll(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.File, error) {
	return m.MockGetAll(ctx, search, sortBy, order, limit, offset)
}

func (m *MockFileRepository) Count(ctx context.Context, search string) (int64, error) {
	return m.MockCount(ctx, search)
}

func (m *MockFileRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.File, error) {
	return m.MockGetByID(ctx, id)
}

func (m *MockFileRepository) GetByAlumniID(ctx context.Context, alumniID string) ([]models.File, error) {
	return m.MockGetByAlumniID(ctx, alumniID)
}

func (m *MockFileRepository) Create(ctx context.Context, file models.File) (*models.File, error) {
	return m.MockCreate(ctx, file)
}

func (m *MockFileRepository) Update(ctx context.Context, id primitive.ObjectID, file models.File) error {
	return m.MockUpdate(ctx, id, file)
}

func (m *MockFileRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	return m.MockDelete(ctx, id)
}

func TestGetAllFiles(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	mockData := []models.File{
		{ID: primitive.NewObjectID(), FileName: "file1.jpg"},
		{ID: primitive.NewObjectID(), FileName: "file2.png"},
	}

	mockRepo.MockGetAll = func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.File, error) {
		return mockData, nil
	}

	files, err := svc.repo.GetAll(ctx, "", "created_at", "asc", 10, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestGetFileByID(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	expected := &models.File{ID: primitive.NewObjectID(), FileName: "test.pdf"}

	mockRepo.MockGetByID = func(ctx context.Context, id primitive.ObjectID) (*models.File, error) {
		return expected, nil
	}

	result, err := svc.repo.GetByID(ctx, primitive.NewObjectID())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.FileName != expected.FileName {
		t.Errorf("Expected %v, got %v", expected.FileName, result.FileName)
	}
}

func TestCreateFile(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	input := models.File{FileName: "upload.png"}

	mockRepo.MockCreate = func(ctx context.Context, f models.File) (*models.File, error) {
		f.ID = primitive.NewObjectID()
		return &f, nil
	}

	result, err := svc.repo.Create(ctx, input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.ID.IsZero() {
		t.Errorf("Expected non-zero ID after create")
	}
}

func TestUpdateFile(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockUpdate = func(ctx context.Context, id primitive.ObjectID, f models.File) error {
		return nil
	}

	err := svc.repo.Update(ctx, primitive.NewObjectID(), models.File{FileName: "updated.png"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestDeleteFile(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockDelete = func(ctx context.Context, id primitive.ObjectID) error {
		return nil
	}

	err := svc.repo.Delete(ctx, primitive.NewObjectID())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// ========== NEGATIVE TESTS ==========

func TestGetFileByID_NotFound(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockGetByID = func(ctx context.Context, id primitive.ObjectID) (*models.File, error) {
		return nil, errors.New("not found")
	}

	_, err := svc.repo.GetByID(ctx, primitive.NewObjectID())
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestCreateFile_Error(t *testing.T) {
	mockRepo := &MockFileRepository{}
	svc := FileService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockCreate = func(ctx context.Context, f models.File) (*models.File, error) {
		return nil, errors.New("insert error")
	}

	_, err := svc.repo.Create(ctx, models.File{})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
