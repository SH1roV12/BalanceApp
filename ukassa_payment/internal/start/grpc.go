package start

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	"github.com/SH1roV12/balance/ukassa/internal/service"
	handler "github.com/SH1roV12/balance/ukassa/internal/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)


func StartGRPC(service *service.Service, config config.GRPC, sugar *zap.SugaredLogger,ctx context.Context)error{
	sugar.Infow("grpc","starting..." )
	gRPCHandler := handler.NewYooKassaPayment(service)
	errChan := make(chan error, 1)
	grpcServer := grpc.NewServer()
	go func() {
		sugar.Infow("grpc", "opening tcp")
		lis,err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s",config.PaymentPort))
		if err != nil{
			errChan <- err
		}
		
		pb.RegisterUKassaPaymentServer(grpcServer,&gRPCHandler)
		grpcServer.Serve(lis)
		
	}()
	select{
	case err := <- errChan:
		sugar.Errorw("grpc","start","error", err.Error() )
		return err
	case <- ctx.Done():
		sugar.Errorw("grpc","graceful shutdown")
		defer grpcServer.Stop()
		time.Sleep(time.Second*7)
		grpcServer.GracefulStop()
		sugar.Errorw("grpc","Server stopped gracefully.")
	}
	sugar.Infow("grpc","successfully starting!" )
	return nil
}