package start

import (
	"fmt"
	"log"
	"os"
	"time"

	"net"

	"github.com/SH1roV12/balance/math/internal/gen/pb"
	"github.com/SH1roV12/balance/math/internal/service"
	"google.golang.org/grpc"
)


func Start() {
	log.Println("Starting grpc server...")
	time.Sleep(time.Second * 3)
	port := os.Getenv("GRPC_PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Println(addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	srvc := &service.MathService{}

	pb.RegisterMathServiceServer(grpcServer, srvc)

	log.Printf("gRPC server started on :%s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}