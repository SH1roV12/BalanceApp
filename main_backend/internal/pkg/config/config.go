package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	DB *Database
	Api *Api
}

type Database struct{
	Host string
	User string
	Password string
	DBName  string
	Port string
	SSlMode string
}

type Api struct{
	Port string
}
func GetConfig()*Config{
	err := godotenv.Load()
	if err != nil{
		log.Println("Failed to load .env")
	}

	return &Config{
		DB: &Database{
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
			Port: os.Getenv("DB_PORT"),
			SSlMode: os.Getenv("DB_SSL"),
		},
		Api: &Api{
			Port: os.Getenv("API_PORT"),
		},
	}
}