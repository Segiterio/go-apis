package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error while loading .env")
	}

	port := os.Getenv("PORT")

	if port == "" {
		panic("PORT is required")
	}

	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV is required")
	}

	return Config{
		Port: port,
		Env:  env,
	}

}
