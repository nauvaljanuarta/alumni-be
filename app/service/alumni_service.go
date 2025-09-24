package service

import (
	"pert5/app/models"
	"pert5/app/repository"
	"pert5/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AlumniService struct {
	repo *repository.AlumniRepository
}

func NewAlumniService(repo *repository.AlumniRepository) *AlumniService {
	return &AlumniService{repo: repo}
}

func (s *AlumniService) GetAlumni(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	offset := (page - 1) * limit

	sortByWhitelist := map[string]bool{
		"id": true, "nim": true, "nama": true, "jurusan": true,
		"angkatan": true, "tahun_lulus": true, "email": true,
		"fakultas": true, "role": true, "created_at": true,
	}
	if !sortByWhitelist[sortBy] {
		sortBy = "id" 
	}

	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	alumni, err := s.repo.GetAlumni(search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch alumni"})
	}

	total, err := s.repo.Count(search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to count alumni"})
	}

	response := models.AlumniResponse{
		Data: alumni,
		Meta: models.MetaInfo{
			Page:   page,
			Limit:  limit,
			Total:  total,
			Pages:  (total + limit - 1) / limit, 
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}

	return c.JSON(response)
}

func (s *AlumniService) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := s.repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"id": id, "message": "alumni doesnt exist"})
	}
	return c.JSON(data)
}

func (s *AlumniService) GetByFakultas(c *fiber.Ctx) error {
	fakultas := c.Params("fakultas")
	data, err := s.repo.GetByFakultas(fakultas)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"fakultas": fakultas, "message": "alumni doesnt exist"})
	}
	return c.JSON(data)
}

func (s *AlumniService) Create(c *fiber.Ctx) error {
	var reqs []models.CreateAlumni
	if err := c.BodyParser(&reqs); err != nil {
			// fallback: kalau body bukan array, coba parse tunggal
			var single models.CreateAlumni
			if err := c.BodyParser(&single); err != nil {
					return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
			}
			reqs = append(reqs, single)
	}

	var ids []int
	for _, req := range reqs {
			// default role = user
			if req.Role == "" {
					req.Role = "user"
			}

			hashedPassword, err := utils.HashPassword(req.Password)
			if err != nil {
					return c.Status(500).JSON(fiber.Map{"message": "failed to hash password"})
			}
			req.Password = hashedPassword

			id, err := s.repo.Create(req)
			if err != nil {
					return c.Status(500).JSON(fiber.Map{"message": err.Error()})
			}
			ids = append(ids, id)
	}

	return c.Status(201).JSON(fiber.Map{
			"message":      "alumni created",
			"inserted_ids": ids,
			"count":        len(ids),
	})
}


func (s *AlumniService) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req models.UpdateAlumni
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
	}
	if err := s.repo.Update(id, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "alumni updated"})
}

func (s *AlumniService) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := s.repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "alumni deleted"})
}
