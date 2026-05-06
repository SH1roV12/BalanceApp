package start

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/SH1roV12/balance/ukassa/internal/infra/postgres"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	getclient "github.com/SH1roV12/balance/ukassa/internal/pkg/gRPC"
	logger "github.com/SH1roV12/balance/ukassa/internal/pkg/logger/zap"

	"github.com/SH1roV12/balance/ukassa/internal/pkg/yookassa"
	"github.com/SH1roV12/balance/ukassa/internal/service"
	"golang.org/x/sync/errgroup"
)



func Start(){
	logger := logger.GetLogger()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Payment module", "Starting")
	config := config.GetConfig(sugar)
	sugar.Infow("start", "id", config.YooKassa.ID, config.YooKassa.SecretKey)
	paymentHandler := yookassa.GetYooKassaPayment(config.YooKassa)
	confirmPaymentClient,conn := getclient.GetConfirmClient(config.GRPC.ConfirmPort,sugar)
	defer conn.Close()
	db := postgres.StartDB(config.DB,sugar)
	repo := postgres.NewRepository(db.DB)
	service := service.NewService(repo,paymentHandler,config.YooKassa,sugar,confirmPaymentClient)
	ctx,cancel := signal.NotifyContext(context.Background(), os.Interrupt,syscall.SIGTERM)
	defer cancel()
	g,gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		sugar.Infow("main start", "starting grpc goroutine")
		return StartGRPC(service,*config.GRPC,sugar,gCtx)
	})
	g.Go(func() error {
		sugar.Infow("main start", "starting http goroutine")
		return StartHTTP(service,*config.HTTP,sugar,gCtx)
	})

	if err := g.Wait(); err != nil || !(errors.Is(err,context.Canceled)){
		sugar.Fatalw("main start", "error", err.Error())
	}
}