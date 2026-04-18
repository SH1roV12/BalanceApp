package math

import (
	"context"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"google.golang.org/grpc"
)


type Service interface{
	Sum(ctx context.Context, in *pb.MathRequest, opts ...grpc.CallOption) (*pb.MathResponse, error)
}