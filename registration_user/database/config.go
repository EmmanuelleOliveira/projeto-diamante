package database

import (
	"log"
	"os"

	"github.com/Valgard/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &Config{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Name:     dbDatabase,
	}
}
