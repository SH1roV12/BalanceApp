package payment

import (
	"context"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"google.golang.org/grpc"
)


type Service interface{
	CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest, opts ...grpc.CallOption) (*pb.CreatePaymentResponse, error)
}