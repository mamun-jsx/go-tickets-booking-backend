package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dns  string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error : can not load .env file")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Dns:  os.Getenv("DNS"),
	}
}

