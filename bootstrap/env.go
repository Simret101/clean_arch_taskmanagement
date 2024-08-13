package bootstrap

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURI  string
	DatabaseName string
	JWTSecretKey string
	TokenExpiry  time.Duration
	ServerPort   string
}

func LoadEnv() *Env {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		DatabaseURI:  os.Getenv("DATABASE_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		TokenExpiry:  time.Duration(60) * time.Minute,
		ServerPort:   os.Getenv("SERVER_PORT"),
	}
}
