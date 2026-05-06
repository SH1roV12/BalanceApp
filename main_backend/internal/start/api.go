package start

import (
	"context"
	"time"

	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/service"
	paymentservice "github.com/SH1roV12/balance/internal/service/payment"

	http "github.com/SH1roV12/balance/internal/transport/http/fiber"
	"github.com/SH1roV12/balance/internal/transport/http/fiber/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)



func StartApi(config *config.Config, userService *service.UserService,paymentService paymentservice.PaymentService,ctx context.Context,sugar *zap.SugaredLogger)error{
	app := fiber.New()
	handlers := handler.NewHandlers(userService,&paymentService,sugar)
	http.SetupRoutes(app,handlers,sugar)


	errChan := make(chan error,1)
	go func(){
		errChan <- app.Listen(":8080")
	}()

	select {
		case err := <- errChan:
			sugar.Errorf("api error: %s", err.Error())
			return err
		case <- ctx.Done():
			shutdownCtx,cancel := context.WithTimeout(ctx,time.Second * 7)
			defer cancel()
			if err := app.ShutdownWithContext(shutdownCtx);err != nil{
				sugar.Errorf("failed to graceful shutdown",err.Error())
				return err
			}
			return nil
	}
}