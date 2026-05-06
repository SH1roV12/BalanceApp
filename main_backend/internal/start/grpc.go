package start

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/SH1roV12/balance/internal/domain/user"
	"github.com/SH1roV12/balance/internal/gen/pb"
	"github.com/SH1roV12/balance/internal/pkg/config"
	handler "github.com/SH1roV12/balance/internal/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)



func StartGRPC(service user.Service, config config.GRPC, sugar *zap.SugaredLogger,ctx context.Context)error{
	sugar.Infow("grpc","starting..." )
	gRPCHandler := handler.NewYooKassaConfirmHandler(service)
	errChan := make(chan error, 1)
	grpcServer := grpc.NewServer()
	go func() {
		sugar.Infow("grpc", "opening tcp")
		lis,err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s",config.ConfirmPort))
		if err != nil{
			errChan <- err
		}
		
		pb.RegisterUKassaPaymentConfirmationServer(grpcServer,&gRPCHandler)
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