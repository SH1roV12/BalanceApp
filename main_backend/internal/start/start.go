package start

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	postgres "github.com/SH1roV12/balance/internal/infra/postgres/gorm"
	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/pkg/logger/zap"
	"github.com/SH1roV12/balance/internal/service"
	"golang.org/x/sync/errgroup"
)

func Start(){
	logger := zap.InitLogger()
	logger.Info("Starting users balance app...")
	config := config.GetConfig()
	db := postgres.StartDB(config.DB)
	userRepository := postgres.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	ctx,signal := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer signal()

	g,gCtx := errgroup.WithContext(ctx)
	g.Go(func()error{
		return StartApi(config,userService,gCtx)
	})
	if err := g.Wait(); err != nil{
		logger.Error("App error occurred", err.Error())
	}
}