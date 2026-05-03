package start

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	middleware "github.com/SH1roV12/balance/ukassa/internal/pkg/middleware/http"
	"github.com/SH1roV12/balance/ukassa/internal/service"
	handler "github.com/SH1roV12/balance/ukassa/internal/transport/http"
	"go.uber.org/zap"
)


func StartHTTP(service *service.Service, config config.HTTP, sugar *zap.SugaredLogger,ctx context.Context)error{
	var allowedCIDRs = []string{
    "0.0.0.0/0",     
    "::/0",          
    "127.0.0.1/32",
}

	handlers := handler.NewYooKassaPaymentConfirmHTTPHandlers(service,sugar)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/webhooks", handlers.HandleWebhook)

	
	protectedMux := middleware.IPFilterMiddleware(mux,allowedCIDRs)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", config.Port),
		Handler: protectedMux,
	}
	
	sugar.Infoln("Starting server on :8080")
	errChan := make(chan error,1)
	go func(){
		err := srv.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	select{
		case err := <-errChan:
			sugar.Errorln("Server failed to start: %v", err)
			return err
		case <- ctx.Done():
			ctx,cancel := context.WithTimeout(context.Background(),time.Second * 7)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil{
				sugar.Errorln("Cannot stop server gracefully")
				return err
			}
			sugar.Fatalln("server gracefully stopped")
	}
	return nil
}