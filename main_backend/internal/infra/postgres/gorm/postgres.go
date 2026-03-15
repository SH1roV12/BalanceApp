package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/SH1roV12/balance/internal/pkg/config"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgres struct{
	*gorm.DB
}


func StartDB(cfg *config.Database)*Postgres{
	var db *Postgres
	var err error

	for i := 0;i < 5; i++{
		log.Println("Connecting to database")

		time.Sleep(time.Second * 2)
		db,err = setupDB(cfg)
		if err==nil{
			break
		}

		log.Printf("Failed connect to DB, retry %d",i+1)
		time.Sleep(time.Second * 1)
	}

	if err != nil{
		log.Fatalln("Failed connect to DB after 5 retry")
		return nil
	}

	return db
}

func setupDB(cfg *config.Database)(*Postgres,error){
	db,err := gorm.Open(driver.Open(DSN(cfg)),&gorm.Config{})
	if err != nil{
		return nil,err
	}
	return &Postgres{DB: db},nil
}


func DSN(cfg *config.Database)string{
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	cfg.Host,cfg.User, cfg.Password,cfg.DBName,cfg.Port,cfg.SSlMode)
}