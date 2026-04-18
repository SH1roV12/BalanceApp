package mathservice

import (
	"context"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"google.golang.org/grpc"
)

type MathService struct{
	client pb.MathServiceClient
}

func NewMathService(client pb.MathServiceClient)MathService{
	return MathService{client: client}
}

func(m MathService) Sum(ctx context.Context, in *pb.MathRequest, opts ...grpc.CallOption) (*pb.MathResponse, error){
	res,err := m.client.Sum(ctx,in,opts...)
	if err != nil{
		return nil,err
	}
	return res,nil
}