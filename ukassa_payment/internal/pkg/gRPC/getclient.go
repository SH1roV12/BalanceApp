package getclient

import (
	"fmt"
	"time"

	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.uber.org/zap"
)



func GetConfirmClient(addr string,sugar *zap.SugaredLogger)(pb.UKassaPaymentConfirmationClient,*grpc.ClientConn){
	var conn *grpc.ClientConn
	sugar.Infow(addr)
		sugar.Infoln("Connecting to grpc...")
		time.Sleep(time.Second * 1)
		
		conn,_ = grpc.NewClient(fmt.Sprintf("app:%s",addr),grpc.WithTransportCredentials(insecure.NewCredentials()))
		
		
		
		
	
	client := pb.NewUKassaPaymentConfirmationClient(conn)
	return client,conn
}