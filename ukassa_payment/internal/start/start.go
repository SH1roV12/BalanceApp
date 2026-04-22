package start

import (
	"github.com/SH1roV12/balance/ukassa/internal/infra/postgres"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	logger "github.com/SH1roV12/balance/ukassa/internal/pkg/logger/zap"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/yookassa"
	"github.com/SH1roV12/balance/ukassa/internal/repository"
	"github.com/SH1roV12/balance/ukassa/internal/service"
)



func Start(){
	logger := logger.GetLogger()
	defer logger.Sync()
	sugar := logger.Sugar()
	config := config.GetConfig(sugar)
	paymentHandler := yookassa.GetYooKassaPayment(config.YooKassa)
	db := postgres.StartDB(config.DB,sugar)
	repo := repository.NewRepository(db)
	service := service.NewService(repo,paymentHandler)
	
}