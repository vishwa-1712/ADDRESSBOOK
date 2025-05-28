package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error in loading .env file : %v", err)
	}

	// "&Config" --> Creates a pointer of struct
	// This means config is of type *Config
	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_DBNAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}
	return config
}
