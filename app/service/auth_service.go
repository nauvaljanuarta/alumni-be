package service

import (
	"pert5/app/models"
	"pert5/app/repository"
	"pert5/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	repo *repository.AlumniRepository
}

func NewAuthService(repo *repository.AlumniRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
	}
	// cari alumni berdasarkan email mengecek apakah emailada atau tidak
	alumni, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "invalid email or password"})
	}
	// cek password
	if !utils.CheckPasswordHash(req.Password, alumni.Password) {
		return c.Status(401).JSON(fiber.Map{"message": "invalid email or password"})
	}
	// generate JWT
	token, err := utils.GenerateToken(models.Alumni{
		ID:       alumni.ID,
		Nama:     alumni.Nama,
		Role:     alumni.Role,
		Email:    alumni.Email,
		Fakultas: alumni.Fakultas,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":       alumni.ID,
			"nama":     alumni.Nama,
			"email":    alumni.Email,
			"role":     alumni.Role,
			"fakultas": alumni.Fakultas,
		},
		"token": token,
	})
}
