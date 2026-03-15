package start

import (
	"context"
	"time"

	"github.com/SH1roV12/balance/internal/pkg/config"
	"github.com/SH1roV12/balance/internal/service"
	http "github.com/SH1roV12/balance/internal/transport/http/fiber"
	"github.com/SH1roV12/balance/internal/transport/http/fiber/handler"
	"github.com/gofiber/fiber/v2"
)



func StartApi(config *config.Config, userService *service.UserService,ctx context.Context)error{
	app := fiber.New()
	handlers := handler.NewHandlers(userService)
	http.SetupRoutes(app,handlers)


	errChan := make(chan error,1)
	go func(){
		errChan <- app.Listen(config.Api.Port)
	}()

	select {
		case err := <- errChan:
			return err
		case <- ctx.Done():
			shutdownCtx,cancel := context.WithTimeout(ctx,time.Second * 7)
			defer cancel()
			if err := app.ShutdownWithContext(shutdownCtx);err != nil{
				return err
			}
			return nil
	}
}