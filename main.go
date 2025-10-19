package main

import (
	"context"
	"log"
	"time"

	// "os"

	"pert5/route"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// config.LoadEnv()
	// database.ConnectDB()
	// defer database.DB.Close()
	// log.Println("db connected")

	// app := fiber.New()
	// route.SetupRoutes(app, database.DB)

	// port := config.GetEnv("APP_PORT")
	// log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

	// defer jangan lupa

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("alumni") 

	app := fiber.New()

	route.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))

}
