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
	PaymentPort string
	PaymentHost string
	ConfirmPort string
	
}

type Api struct{
	Port string
}
func GetConfig(sugar *zap.SugaredLogger)*Config{
	// err := godotenv.Load()
	// if err != nil{
	// 	sugar.Infow("Failed to load .env")
	// }
	sugar.Infow("env file","DB_HOST", os.Getenv("DB_HOST") )
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
		},
		GRPC: &GRPC{
			PaymentPort: os.Getenv("GRPC_PORT_PAYMENT"),
			ConfirmPort: os.Getenv("GRPC_PORT_CONFIRM"),
		},
	}
}