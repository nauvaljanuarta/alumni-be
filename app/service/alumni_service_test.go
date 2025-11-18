package service

import (
	"context"
	"errors"
	"pert5/app/models"
	"pert5/utils"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockAlumniRepository struct {
	MockGetAlumni     func(ctx context.Context, search string, limit, offset int, sortBy, order string) ([]models.Alumni, error)
	MockGetAlumniByID func(ctx context.Context, id string) (*models.Alumni, error)
	MockGetByEmail    func(ctx context.Context, email string) (*models.Alumni, error)
	MockCreateAlumni  func(ctx context.Context, req *models.Alumni) (*models.Alumni, error)
	MockUpdateAlumni  func(ctx context.Context, id string, req *models.Alumni) error
	MockDeleteAlumni  func(ctx context.Context, id string) error
	MockCount         func(ctx context.Context, search string) (int, error)
}

func (m *MockAlumniRepository) GetAlumni(ctx context.Context, search string, limit, offset int, sortBy, order string) ([]models.Alumni, error) {
	return m.MockGetAlumni(ctx, search, limit, offset, sortBy, order)
}

func (m *MockAlumniRepository) GetAlumniByID(ctx context.Context, id string) (*models.Alumni, error) {
	return m.MockGetAlumniByID(ctx, id)
}

func (m *MockAlumniRepository) GetByEmail(ctx context.Context, email string) (*models.Alumni, error) {
	return m.MockGetByEmail(ctx, email)
}

func (m *MockAlumniRepository) CreateAlumni(ctx context.Context, req *models.Alumni) (*models.Alumni, error) {
	return m.MockCreateAlumni(ctx, req)
}

func (m *MockAlumniRepository) UpdateAlumni(ctx context.Context, id string, req *models.Alumni) error {
	return m.MockUpdateAlumni(ctx, id, req)
}

func (m *MockAlumniRepository) DeleteAlumni(ctx context.Context, id string) error {
	return m.MockDeleteAlumni(ctx, id)
}

func (m *MockAlumniRepository) Count(ctx context.Context, search string) (int, error) {
	return m.MockCount(ctx, search)
}


func TestGetAlumni(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	mockData := []models.Alumni{
		{ID: primitive.NewObjectID(), Nama: "Janu"},
		{ID: primitive.NewObjectID(), Nama: "Dika"},
	}

	mockRepo.MockGetAlumni = func(ctx context.Context, search string, limit, offset int, sortBy, order string) ([]models.Alumni, error) {
		return mockData, nil
	}

	alumni, err := svc.repo.GetAlumni(ctx, "", 10, 0, "nama", "asc")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(alumni) != 2 {
		t.Errorf("Expected 2 alumni, got %d", len(alumni))
	}
}

func TestGetByID(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	expected := &models.Alumni{ID: primitive.NewObjectID(), Nama: "Janu"}

	mockRepo.MockGetAlumniByID = func(ctx context.Context, id string) (*models.Alumni, error) {
		return expected, nil
	}

	result, err := svc.repo.GetAlumniByID(ctx, "123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.Nama != expected.Nama {
		t.Errorf("Expected name %v, got %v", expected.Nama, result.Nama)
	}
}

func TestCreateAlumni(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	hashed, _ := utils.HashPassword("pass123")
	input := &models.Alumni{Nama: "Janu", Password: hashed}

	mockRepo.MockCreateAlumni = func(ctx context.Context, alumni *models.Alumni) (*models.Alumni, error) {
		alumni.ID = primitive.NewObjectID()
		return alumni, nil
	}

	result, err := svc.repo.CreateAlumni(ctx, input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.ID.IsZero() {
		t.Errorf("Expected created ID, got zero object ID")
	}
}

func TestUpdateAlumni(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockUpdateAlumni = func(ctx context.Context, id string, alumni *models.Alumni) error {
		return nil
	}

	err := svc.repo.UpdateAlumni(ctx, "123", &models.Alumni{Nama: "Updated"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestDeleteAlumni(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockDeleteAlumni = func(ctx context.Context, id string) error {
		return nil
	}

	err := svc.repo.DeleteAlumni(ctx, "123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// ========== NEGATIVE TESTS ==========

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockGetAlumniByID = func(ctx context.Context, id string) (*models.Alumni, error) {
		return nil, errors.New("not found")
	}

	_, err := svc.repo.GetAlumniByID(ctx, "missing-id")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestCreateAlumni_Error(t *testing.T) {
	mockRepo := &MockAlumniRepository{}
	svc := AlumniService{repo: mockRepo}
	ctx := context.Background()

	mockRepo.MockCreateAlumni = func(ctx context.Context, alumni *models.Alumni) (*models.Alumni, error) {
		return nil, errors.New("insert error")
	}

	_, err := svc.repo.CreateAlumni(ctx, &models.Alumni{})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
