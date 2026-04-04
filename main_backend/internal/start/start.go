package start

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	postgres "github.com/SH1roV12/balance/internal/infra/postgres/gorm"
	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/service"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func Start(){
	logger,err := zap.NewProduction()
	if err != nil{
		log.Fatalf("cannot start logger, %s",err.Error())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Balance app is starting...")
	time.Sleep(time.Second * 3)
	config := config.GetConfig(sugar)
	db := postgres.StartDB(config.DB, sugar)
	userRepository := postgres.NewUserRepository(db.DB,sugar)
	userService := service.NewUserService(userRepository,sugar)
	ctx,signal := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer signal()

	g,gCtx := errgroup.WithContext(ctx)
	g.Go(func()error{
		return StartApi(config,userService,gCtx,sugar)
	})
	if err := g.Wait(); err != nil{
		sugar.Errorf("App error occurred %s", err.Error())
	}
}