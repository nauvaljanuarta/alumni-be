package route

import (
	"database/sql"
	"pert5/app/repository"
	"pert5/app/service"
	"pert5/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/pert5")

	// auth login 
	alumniRepo := repository.NewAlumniRepository(db)
	authService := service.NewAuthService(alumniRepo) 
	api.Post("/login", authService.Login)

	// alumnni routes
	alumniService := service.NewAlumniService(alumniRepo)
	alumni := api.Group("/alumni")
	// Admin & User (cukup RequireAuth)
	alumni.Get("/", middleware.RequireAuth, alumniService.GetAlumni)
	alumni.Get("/:id", middleware.RequireAuth, alumniService.GetByID)
	alumni.Get("/fakultas/:fakultas", middleware.RequireAuth, alumniService.GetByFakultas)
	// admin only
	alumni.Post("/", middleware.RequireAuth, middleware.AdminOnly(), alumniService.Create)
	alumni.Put("/:id", middleware.RequireAuth, middleware.AdminOnly(), alumniService.Update)
	alumni.Delete("/:id", middleware.RequireAuth, middleware.AdminOnly(), alumniService.Delete)
	
	// route pekerjaan
	pekerjaanRepo := repository.NewPekerjaanRepository(db)
	pekerjaanService := service.NewPekerjaanService(pekerjaanRepo)
	pekerjaan := api.Group("/pekerjaan")
	// Admin & User
	pekerjaan.Get("/", middleware.RequireAuth, pekerjaanService.GetAll)
	pekerjaan.Get("/:id", middleware.RequireAuth, pekerjaanService.GetByID)
	pekerjaan.Put("/softdelete/:id", middleware.RequireAuth, pekerjaanService.SoftDelete)
	// Hanya Admin
	pekerjaan.Get("/alumni/:alumni_id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.GetByAlumniID)
	pekerjaan.Post("/", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Create)
	pekerjaan.Put("/:id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Update)
	pekerjaan.Delete("/:id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Delete)
	pekerjaan.Put("/softdeletebulk", middleware.RequireAuth,middleware.AdminOnly(), pekerjaanService.SoftDeleteBulk)
}

// http://localhost:3000/pert5/pekerjaan/softdeleted/20