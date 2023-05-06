package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // autoload .env file
)

var (
	PORT              = os.Getenv("PORT")
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)
