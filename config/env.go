package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file tidak ditemukan, pakai environment default")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
