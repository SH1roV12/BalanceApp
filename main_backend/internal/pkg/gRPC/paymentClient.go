package grpc

import (
	"fmt"
	"log"
	"time"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func GetPaymentClient(addr string,sugar *zap.SugaredLogger)(pb.UKassaPaymentClient,*grpc.ClientConn){
	var conn *grpc.ClientConn
	sugar.Infow(addr)
		log.Println("Connecting to grpc...")
		time.Sleep(time.Second * 1)
		
		conn,_ = grpc.NewClient(fmt.Sprintf("yookassa:%s",addr),grpc.WithTransportCredentials(insecure.NewCredentials()))
		
		
		
		
	
	client := pb.NewUKassaPaymentClient(conn)
	return client,conn
}