package grpc

import (
	"fmt"
	"time"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"github.com/SH1roV12/balance/internal/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func GetPaymentClient(cfg config.GRPC,sugar *zap.SugaredLogger,)(pb.UKassaPaymentClient,*grpc.ClientConn){
	var conn *grpc.ClientConn
		sugar.Infoln("Connecting to grpc...")
		time.Sleep(time.Second * 1)
		
		conn,_ = grpc.NewClient(fmt.Sprintf("%s:%s",cfg.YooKassaHost,cfg.PaymentPort),grpc.WithTransportCredentials(insecure.NewCredentials()))
		
		
		
		
	
	client := pb.NewUKassaPaymentClient(conn)
	sugar.Infoln("Successfully connected!")
	return client,conn
}