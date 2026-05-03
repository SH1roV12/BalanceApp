package paymentservice

import (
	"context"

	"github.com/SH1roV12/balance/internal/gen/pb"
	"google.golang.org/grpc"
)


type PaymentService struct{
	client pb.UKassaPaymentClient
}

func NewPaymentService(client pb.UKassaPaymentClient)PaymentService{
	return PaymentService{client: client}
}

func (p *PaymentService) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest, opts ...grpc.CallOption) (*pb.CreatePaymentResponse, error){
	res,err := p.client.CreatePayment(ctx,in,opts...)
	if err != nil{
		return nil,err
	}
	return res,nil
}