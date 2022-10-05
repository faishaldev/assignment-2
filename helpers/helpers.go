package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvChecking() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
