package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	SecretKey       string
	TokenExpiration time.Duration
)

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	SecretKey = os.Getenv("SECRET_KEY")
	if SecretKey == "" {
		log.Fatal("SECRET_KEY environment variable is required")
	}

	expirationMinutes := os.Getenv("TOKEN_EXPIRATION_MINUTES")
	if expirationMinutes == "" {
		TokenExpiration = 15 * time.Minute
	} else {
		duration, err := time.ParseDuration(expirationMinutes + "m")
		if err != nil {
			log.Fatal("Invalid TOKEN_EXPIRATION_MINUTES value")
		}
		TokenExpiration = duration
	}
}
