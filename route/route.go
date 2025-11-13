package route

import (
	// "database/sql"
	"pert5/app/repository"
	"pert5/app/service"
	"pert5/middleware"
	_ "pert5/docs" 
	"github.com/gofiber/swagger"


	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// func SetupRoutes(app *fiber.App, db *sql.DB) {
func SetupRoutes(app *fiber.App, db *mongo.Database) {
	api := app.Group("/pert5")
	api.Get("/swagger/*", swagger.HandlerDefault)

	// auth login 
	alumniRepo := repository.NewAlumniRepository(db)
	authService := service.NewAuthService(alumniRepo) 
	api.Post("/login", authService.Login)

	// alumnni routes
	alumniService := service.NewAlumniService(alumniRepo)
	alumni := api.Group("/alumni")
	
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

	// uts
	pekerjaan.Put("/restore/:id", middleware.RequireAuth, pekerjaanService.Restore)
	pekerjaan.Get("/trash", middleware.RequireAuth, pekerjaanService.GetTrash)
	pekerjaan.Delete("/deletetrash/:id", middleware.RequireAuth, pekerjaanService.DeleteTrash)
	
	// admin & user
	pekerjaan.Get("/", middleware.RequireAuth, pekerjaanService.GetAll)
	pekerjaan.Get("/:id", middleware.RequireAuth, pekerjaanService.GetByID)
	pekerjaan.Put("/softdelete/:id", middleware.RequireAuth, pekerjaanService.SoftDelete)

	// Hanya Admin
	pekerjaan.Get("/alumni/:alumni_id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.GetByAlumniID)
	pekerjaan.Post("/", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Create)
	pekerjaan.Put("/:id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Update)
	pekerjaan.Delete("/:id", middleware.RequireAuth, middleware.AdminOnly(), pekerjaanService.Delete)
	pekerjaan.Put("/softdeletebulk", middleware.RequireAuth,middleware.AdminOnly(), pekerjaanService.SoftDeleteBulk)

  // route file
fileRepo := repository.NewFileRepository(db)
fileService := service.NewFileService(fileRepo, "./uploads")
file := api.Group("/file")


// user bisa lihat file
file.Get("/", middleware.RequireAuth, fileService.GetAllFiles)
file.Get("/:id", middleware.RequireAuth, fileService.GetFileByID)
file.Get("/alumni/:alumni_id", middleware.RequireAuth, fileService.GetFilesByAlumniID)

// admin only upload/update/delete
file.Post("/", middleware.RequireAuth, fileService.UploadFile)              
file.Post("/:alumni_id", middleware.RequireAuth, middleware.AdminOnly(), fileService.UploadFileAdmin) // admin upload ke alumni lain
file.Put("/:id", middleware.RequireAuth, middleware.AdminOnly(), fileService.UpdateFile)
file.Delete("/:id", middleware.RequireAuth, middleware.AdminOnly(), fileService.DeleteFile)
}

