package config

import (
	"os"

	"go.uber.org/zap"
)

type Config struct{
	DB *Database
	Api *Api
	GRPC *GRPC
}

type Database struct{
	Host string
	User string
	Password string
	DBName  string
	Port string
	SSlMode string
}

type GRPC struct{
	YooKassaHost string
	PaymentPort string
	ConfirmPort string
}

type Api struct{
	Port string
	Access_Secret string
	Refresh_Secret string
}
func GetConfig(sugar *zap.SugaredLogger)*Config{
	
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
			Port: os.Getenv("APP_PORT"),
			Access_Secret: getEnv("JWT_ACCESS_SECRET",sugar),
			Refresh_Secret: getEnv("JWT_REFRESH_SECRET",sugar),
		},
		GRPC: &GRPC{
			PaymentPort: os.Getenv("GRPC_PORT_PAYMENT"),
			ConfirmPort: os.Getenv("GRPC_PORT_CONFIRM"),
			YooKassaHost: os.Getenv("GRPC_YOOKASSA_HOST"),
		},
	}
}

func getEnv(key string, sugar *zap.SugaredLogger)string{
	value := os.Getenv(key)
	if value == ""{
		sugar.Fatalln("Critical environment variable is missing", "variable", key)
	}
	return value
}