package grpc

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/user"
	"github.com/SH1roV12/balance/internal/gen/pb"
	"github.com/SH1roV12/balance/internal/transport/dto/request"
	"google.golang.org/protobuf/types/known/emptypb"
)




type YooKassaConfirm struct{
	service user.Service
	pb.UnimplementedUKassaPaymentConfirmationServer
}

func NewYooKassaConfirmHandler(service user.Service)YooKassaConfirm{
	return YooKassaConfirm{service: service,UnimplementedUKassaPaymentConfirmationServer: pb.UnimplementedUKassaPaymentConfirmationServer{}}
}

func (u *YooKassaConfirm) ConfirmPayment(ctx context.Context, in *pb.ConfirmPaymentRequest) (*emptypb.Empty, error){
	err := u.service.SetBalance(ctx,request.ToDomainReplenishment(in))
	if err != nil{
		return nil,err
	}
	return &emptypb.Empty{},err
}