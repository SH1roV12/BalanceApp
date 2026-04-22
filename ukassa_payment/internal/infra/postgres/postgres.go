package postgres

import (
	"fmt"
	"os"
	"time"

	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	"go.uber.org/zap"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgres struct{
	 *gorm.DB
}


func StartDB(cfg *config.Database, sugar *zap.SugaredLogger)*Postgres{
	var db *Postgres
	var err error
	sugar.Infow("Connecting to database")
	for i := 0; i < 5; i++{

		time.Sleep(time.Second * 2)
		db,err = setupDB(cfg)
		if err == nil{
			break
		}

		sugar.Warnw("postgres", "start db", "retry", i+1)
	}

	
	if err != nil{
		sugar.Warnw("postgres", "start db", "cannot connect to db")
		os.Exit(1)
	}
	sugar.Infow("postgres", "start db", "migrating...")
	err = migrate(db)

	if err != nil{
		sugar.Warnw("postgres", "start db", "cannot migrate db")
		os.Exit(1)
	}
	sugar.Infow("postgres", "start db", "successfully connected")
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

func migrate(db *Postgres)error{
	return  db.AutoMigrate()
}