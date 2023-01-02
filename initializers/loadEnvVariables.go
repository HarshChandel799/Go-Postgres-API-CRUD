package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVvariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
