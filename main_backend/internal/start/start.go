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
	grpc "github.com/SH1roV12/balance/internal/pkg/gRPC"
	"github.com/SH1roV12/balance/internal/service"
	paymentservice "github.com/SH1roV12/balance/internal/service/payment"

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
	paymentClient,conn := grpc.GetPaymentClient(config.GRPC.PaymentPort,sugar)
	
	defer conn.Close()
	db := postgres.StartDB(config.DB, sugar)
	userRepository := postgres.NewUserRepository(db.DB,sugar)
	paymentService := paymentservice.NewPaymentService(paymentClient)
	userService := service.NewUserService(userRepository,sugar)
	ctx,signal := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer signal()

	g,gCtx := errgroup.WithContext(ctx)
	g.Go(func()error{
		return StartApi(config,userService,paymentService,gCtx,sugar)
	})
	g.Go(func()error{
		return StartGRPC(userService,*config.GRPC,sugar,gCtx)
	})
	if err := g.Wait(); err != nil{
		sugar.Errorf("App error occurred %s", err.Error())
	}
}