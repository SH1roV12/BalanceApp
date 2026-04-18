package service

import (
	"context"

	"github.com/SH1roV12/balance/math/internal/gen/pb"
)

type MathService struct{
	pb.UnimplementedMathServiceServer
}


func(m *MathService) Sum(ctx context.Context, in *pb.MathRequest)(*pb.MathResponse,error){
	return &pb.MathResponse{
		Result: in.Num1+in.Num2,
	},nil
}