package database

import (
    // "database/sql"
    // "fmt"
    "log"
    "os"
		"time"
		"context"

    _ "github.com/lib/pq"
		"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// var DB *sql.DB

var DB *mongo.Database
var Client *mongo.Client

func ConnectDB() {
    // dsn := fmt.Sprintf(
    //     "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    //     os.Getenv("DB_HOST"),
    //     os.Getenv("DB_USER"),
    //     os.Getenv("DB_PASS"),
    //     os.Getenv("DB_NAME"),
    //     os.Getenv("DB_PORT"),
    // )

		    // var err error
    // DB, err = sql.Open("postgres", dsn)
    // if err != nil {
    //     log.Fatal("Error opening DB:", err)
    // }

    // if err = DB.Ping(); err != nil {
    //     log.Fatal("Error connecting to DB:", err)
    // }

    // var version string
    // if err := DB.QueryRow("SELECT version()").Scan(&version); err != nil {
    //     log.Fatal("Error checking DB version:", err)
    // }

    // log.Println("Connected database :", version)

		mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017" 
    }

		dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "alumni_db" 
    }

		clientOptions := options.Client().ApplyURI(mongoURI)
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }

    if err = Client.Ping(ctx, nil); err != nil {
        log.Fatal("Error pinging MongoDB:", err)
    }

    DB = Client.Database(dbName)

    log.Println("Connected to MongoDB database:", dbName)
}

func DBClose() {
	if Client != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			
			if err := Client.Disconnect(ctx); err != nil {
					log.Println("Error disconnecting from MongoDB:", err)
			} else {
					log.Println("Disconnected from MongoDB")
			}
	}
}