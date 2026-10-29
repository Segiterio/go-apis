package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	Env   string
	DBUrl string
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found; using environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		panic("PORT is required")
	}

	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV is required")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		panic("DB_URL is required")
	}

	return Config{
		Port:  port,
		Env:   env,
		DBUrl: dbUrl,
	}

}
