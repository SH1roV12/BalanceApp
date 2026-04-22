package grpc

import (
	"context"

	"github.com/SH1roV12/balance/ukassa/internal/domain/service"
	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	"github.com/SH1roV12/balance/ukassa/internal/transport/dto/request"
)

type UKassaPayment struct{
	service service.Service
	pb.UnimplementedUKassaPaymentServer
}

func (u *UKassaPayment) CreatePayment(ctx context.Context,in *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error){
	url,err := u.service.CreatePayment(ctx,request.ToDomain(in))
	if err != nil{
		return nil,err
	}
	return &pb.CreatePaymentResponse{PaymentUrl: url},err
}