package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"pert5/app/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock repository sesuai IPekerjaanRepository
type MockPekerjaanRepo struct {
	MockGetAll      func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error)
	MockCount       func(ctx context.Context, search string) (int64, error)
	MockGetByID     func(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error)
	MockGetByAlumni func(ctx context.Context, alumniID primitive.ObjectID) ([]models.Pekerjaan, error)
	MockCreate      func(ctx context.Context, req models.CreatePekerjaan) (*mongo.InsertOneResult, error)
	MockUpdate      func(ctx context.Context, id primitive.ObjectID, req models.UpdatePekerjaan) error
	MockDelete      func(ctx context.Context, id primitive.ObjectID) error
	MockSoftDelete  func(ctx context.Context, id primitive.ObjectID) error
	MockSoftBulk    func(ctx context.Context) error
	MockRestore     func(ctx context.Context, id primitive.ObjectID) error
	MockGetTrash    func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error)
	MockCountTrash  func(ctx context.Context, search string) (int64, error)
	MockDeleteTrash func(ctx context.Context, id primitive.ObjectID) error
}

func (m *MockPekerjaanRepo) GetAll(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
	return m.MockGetAll(ctx, search, sortBy, order, limit, offset)
}
func (m *MockPekerjaanRepo) Count(ctx context.Context, search string) (int64, error) {
	return m.MockCount(ctx, search)
}
func (m *MockPekerjaanRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error) {
	return m.MockGetByID(ctx, id)
}
func (m *MockPekerjaanRepo) GetByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]models.Pekerjaan, error) {
	return m.MockGetByAlumni(ctx, alumniID)
}
func (m *MockPekerjaanRepo) Create(ctx context.Context, req models.CreatePekerjaan) (*mongo.InsertOneResult, error) {
	return m.MockCreate(ctx, req)
}
func (m *MockPekerjaanRepo) Update(ctx context.Context, id primitive.ObjectID, req models.UpdatePekerjaan) error {
	return m.MockUpdate(ctx, id, req)
}
func (m *MockPekerjaanRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	return m.MockDelete(ctx, id)
}
func (m *MockPekerjaanRepo) SoftDelete(ctx context.Context, id primitive.ObjectID) error {
	return m.MockSoftDelete(ctx, id)
}
func (m *MockPekerjaanRepo) SoftDeleteBulk(ctx context.Context) error {
	return m.MockSoftBulk(ctx)
}
func (m *MockPekerjaanRepo) Restore(ctx context.Context, id primitive.ObjectID) error {
	return m.MockRestore(ctx, id)
}
func (m *MockPekerjaanRepo) GetTrash(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
	return m.MockGetTrash(ctx, search, sortBy, order, limit, offset)
}
func (m *MockPekerjaanRepo) CountTrash(ctx context.Context, search string) (int64, error) {
	return m.MockCountTrash(ctx, search)
}
func (m *MockPekerjaanRepo) DeleteTrash(ctx context.Context, id primitive.ObjectID) error {
	return m.MockDeleteTrash(ctx, id)
}

// Helper to build a Fiber app with single route invoking the handler
func makeApp(handler fiber.Handler) *fiber.App {
	app := fiber.New()
	// Add a wrapper so handler gets its own route; tests set path/params accordingly
	app.All("/*", handler)
	return app
}

func TestGetAllPekerjaan_Success(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}
	app := makeApp(svc.GetAll)

	// sample data
	p1 := models.Pekerjaan{ID: primitive.NewObjectID(), NamaPerusahaan: "A"}
	p2 := models.Pekerjaan{ID: primitive.NewObjectID(), NamaPerusahaan: "B"}
	mock.MockGetAll = func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
		return []models.Pekerjaan{p1, p2}, nil
	}
	mock.MockCount = func(ctx context.Context, search string) (int64, error) {
		return int64(2), nil
	}

	req := httptest.NewRequest(http.MethodGet, "/?page=1&limit=10&sortBy=created_at&order=desc", nil)
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}

	var body struct {
		Data []models.Pekerjaan `json:"data"`
		Meta struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Total int `json:"total"`
		} `json:"meta"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if len(body.Data) != 2 {
		t.Fatalf("expected 2 items got %d", len(body.Data))
	}
	if body.Meta.Total != 2 {
		t.Fatalf("expected total 2 got %d", body.Meta.Total)
	}
}

func TestGetByID_Pekerjaan_NotFound(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}
	app := makeApp(svc.GetByID)

	// repo returns error -> handler should 404
	mock.MockGetByID = func(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error) {
		return nil, errors.New("not found")
	}

	id := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodGet, "/"+id, nil)
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != 404 {
		t.Fatalf("expected 404 got %d", resp.StatusCode)
	}
}

func TestCreatePekerjaan_Success(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}
	app := makeApp(svc.Create)

	input := models.CreatePekerjaan{
		NamaPerusahaan: "Acme",
		PosisiJabatan:  "Dev",
	}
	mock.MockCreate = func(ctx context.Context, req models.CreatePekerjaan) (*mongo.InsertOneResult, error) {
		return &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil
	}

	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != 201 {
		t.Fatalf("expected 201 got %d", resp.StatusCode)
	}
	var body map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&body)
	if body["message"] != "pekerjaan created" {
		t.Fatalf("unexpected response message: %v", body["message"])
	}
}

func TestUpdatePekerjaan_Success(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}
	app := makeApp(svc.Update)

	mock.MockUpdate = func(ctx context.Context, id primitive.ObjectID, req models.UpdatePekerjaan) error {
		return nil
	}

	id := primitive.NewObjectID().Hex()
	updateReq := models.UpdatePekerjaan{PosisiJabatan: "Lead"}
	b, _ := json.Marshal(updateReq)
	req := httptest.NewRequest(http.MethodPut, "/"+id, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
}

func TestDeletePekerjaan_Success(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}
	app := makeApp(svc.Delete)

	mock.MockDelete = func(ctx context.Context, id primitive.ObjectID) error {
		return nil
	}

	id := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodDelete, "/"+id, nil)
	resp, _ := app.Test(req, -1)
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
}

// SoftDelete: test forbidden when not owner & not admin, and success when owner
func TestSoftDeletePekerjaan_ForbiddenAndSuccess(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}

	// existing pekerjaan owned by some other user
	ownerID := primitive.NewObjectID()
	existing := &models.Pekerjaan{ID: primitive.NewObjectID(), AlumniID: ownerID}

	mock.MockGetByID = func(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error) {
		return existing, nil
	}
	mock.MockSoftDelete = func(ctx context.Context, id primitive.ObjectID) error {
		return nil
	}

	// CASE 1: requester is different user (forbidden)
	app1 := fiber.New()
	app1.Delete("/:id", func(c *fiber.Ctx) error {
		// set locals to a different user id
		c.Locals("user_id", primitive.NewObjectID())
		c.Locals("role", "user")
		return svc.SoftDelete(c)
	})
	id := existing.ID.Hex()
	req1 := httptest.NewRequest(http.MethodDelete, "/"+id, nil)
	resp1, _ := app1.Test(req1, -1)
	if resp1.StatusCode != 403 {
		t.Fatalf("expected 403 got %d", resp1.StatusCode)
	}

	// CASE 2: requester is owner -> success
	app2 := fiber.New()
	app2.Delete("/:id", func(c *fiber.Ctx) error {
		c.Locals("user_id", ownerID) // set owner (primitive.ObjectID)
		c.Locals("role", "user")
		return svc.SoftDelete(c)
	})
	req2 := httptest.NewRequest(http.MethodDelete, "/"+id, nil)
	resp2, _ := app2.Test(req2, -1)
	if resp2.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp2.StatusCode)
	}
}

// SoftDeleteBulk: only admin allowed
func TestSoftDeleteBulk_UnauthorizedAndSuccess(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}

	mock.MockSoftBulk = func(ctx context.Context) error { return nil }

	// unauthorized
	app1 := fiber.New()
	app1.Delete("/", func(c *fiber.Ctx) error {
		c.Locals("role", "user")
		return svc.SoftDeleteBulk(c)
	})
	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	resp1, _ := app1.Test(req1, -1)
	if resp1.StatusCode != 403 {
		t.Fatalf("expected 403 got %d", resp1.StatusCode)
	}

	// authorized admin
	app2 := fiber.New()
	app2.Delete("/", func(c *fiber.Ctx) error {
		c.Locals("role", "admin")
		return svc.SoftDeleteBulk(c)
	})
	req2 := httptest.NewRequest(http.MethodDelete, "/", nil)
	resp2, _ := app2.Test(req2, -1)
	if resp2.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp2.StatusCode)
	}
}

func TestRestorePekerjaan_Permissions(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}

	ownerID := primitive.NewObjectID()
	existing := &models.Pekerjaan{ID: primitive.NewObjectID(), AlumniID: ownerID}

	mock.MockGetByID = func(ctx context.Context, id primitive.ObjectID) (*models.Pekerjaan, error) {
		return existing, nil
	}
	mock.MockRestore = func(ctx context.Context, id primitive.ObjectID) error { return nil }

	// non-owner -> forbidden
	app1 := fiber.New()
	app1.Put("/:id", func(c *fiber.Ctx) error {
		c.Locals("user_id", primitive.NewObjectID())
		c.Locals("role", "user")
		return svc.Restore(c)
	})
	req1 := httptest.NewRequest(http.MethodPut, "/"+existing.ID.Hex(), nil)
	resp1, _ := app1.Test(req1, -1)
	if resp1.StatusCode != 403 {
		t.Fatalf("expected 403 got %d", resp1.StatusCode)
	}

	// owner -> success
	app2 := fiber.New()
	app2.Put("/:id", func(c *fiber.Ctx) error {
		c.Locals("user_id", ownerID)
		c.Locals("role", "user")
		return svc.Restore(c)
	})
	req2 := httptest.NewRequest(http.MethodPut, "/"+existing.ID.Hex(), nil)
	resp2, _ := app2.Test(req2, -1)
	if resp2.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp2.StatusCode)
	}
}

func TestGetTrashAndDeleteTrash(t *testing.T) {
	mock := &MockPekerjaanRepo{}
	svc := PekerjaanService{repo: mock}

	p := models.Pekerjaan{ID: primitive.NewObjectID(), NamaPerusahaan: "TrashCo"}
	mock.MockGetTrash = func(ctx context.Context, search, sortBy, order string, limit, offset int) ([]models.Pekerjaan, error) {
		return []models.Pekerjaan{p}, nil
	}
	mock.MockCountTrash = func(ctx context.Context, search string) (int64, error) {
		return int64(1), nil
	}
	mock.MockDeleteTrash = func(ctx context.Context, id primitive.ObjectID) error {
		return nil
	}

	// GetTrash
	app1 := fiber.New()
	app1.Get("/", func(c *fiber.Ctx) error {
		return svc.GetTrash(c)
	})
	req1 := httptest.NewRequest(http.MethodGet, "/?page=1&limit=10", nil)
	resp1, _ := app1.Test(req1, -1)
	if resp1.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp1.StatusCode)
	}
	var body struct {
		Data []models.Pekerjaan `json:"data"`
		Meta struct {
			Total int `json:"total"`
		} `json:"meta"`
	}
	_ = json.NewDecoder(resp1.Body).Decode(&body)
	if len(body.Data) != 1 || body.Meta.Total != 1 {
		t.Fatalf("unexpected trash response")
	}

	// DeleteTrash
	app2 := fiber.New()
	app2.Delete("/:id", func(c *fiber.Ctx) error {
		return svc.DeleteTrash(c)
	})
	req2 := httptest.NewRequest(http.MethodDelete, "/"+p.ID.Hex(), nil)
	resp2, _ := app2.Test(req2, -1)
	if resp2.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp2.StatusCode)
	}
}
