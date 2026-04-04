package postgres

import (
	"fmt"
	"time"

	"github.com/SH1roV12/balance/internal/pkg/config"
	"go.uber.org/zap"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgres struct{
	*gorm.DB
}


func StartDB(cfg *config.Database,sugar *zap.SugaredLogger)*Postgres{
	var db *Postgres
	var err error

	for i := 0;i < 5; i++{
		sugar.Infow("Connecting to database")

		time.Sleep(time.Second * 2)
		db,err = setupDB(cfg)
		if err==nil{
			break
		}

		sugar.Warnf("Failed connect to DB, retry %d",i+1)
		time.Sleep(time.Second * 1)
	}

	if err != nil{
		sugar.Errorw("Failed connect to DB after 5 retry")
		return nil
	}
	sugar.Infow("Migrating DB tables....")
	err = migrate(db)
	time.Sleep(time.Second * 3)
	if err != nil{
		sugar.Errorw("cannot migrate db", err.Error())
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

func migrate(db *Postgres)error{
	return  db.AutoMigrate(&User{})
}