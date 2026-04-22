package response

import (
	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)


func FromConfirmationToGRPC(conf *yoopayment.Confirmer)*pb.CreatePaymentResponse{
	return &pb.CreatePaymentResponse{
		PaymentUrl: conf.,
	}
}