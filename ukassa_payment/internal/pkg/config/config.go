package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct{
	DB *Database
	YooKassa *YooKassa
}

type Database struct{
	Host string
	User string
	Password string
	DBName  string
	Port string
	SSlMode string
}

type YooKassa struct{
	ID string
	SecretKey string
	ReturnURL string
}

func GetConfig(sugar *zap.SugaredLogger)*Config{
	godotenv.Load()
	
	sugar.Infow("config", "getConfig", "getting config...")
	return &Config{
		DB: &Database{
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
			Port: os.Getenv("DB_PORT"),
			SSlMode: os.Getenv("DB_SSL"),
		},
		YooKassa: &YooKassa{
			ID: os.Getenv("YOOKASSA_SHOP_ID"),
			SecretKey: os.Getenv("YOOKASSA_STORE_SECRET_KEY"),
		},
	}
}