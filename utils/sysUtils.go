package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
