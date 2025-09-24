package main

import (
	"log"
	"fmt"
	// "os"
	"pert5/config"
	"pert5/database"
	"pert5/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// config.LoadEnv()
	// db := database.ConnectDB()
	// app := config.NewApp(db)

	// port := os.Getenv("APP_PORT")
	// if port == "" {
	// 	port = "3000"
	// }

	// log.Fatal(app.Listen(":" + port))

	config.LoadEnv()
  database.ConnectDB()
  defer database.DB.Close()
  log.Println("db connected")

	app := fiber.New()
	route.SetupRoutes(app, database.DB)

	port := config.GetEnv("APP_PORT")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

	// defer jangan lupa
	
}
