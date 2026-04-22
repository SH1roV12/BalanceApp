package request

import (
	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
)


type CreatePayment struct{
	UserId        string                 `json:"user_id,omitempty"`
    TransactionId string                 `json:"transaction_id,omitempty"`
    Price         float32                `json:"price,omitempty"`
}

func ToDomain(grpc *pb.CreatePaymentRequest)CreatePayment{
	return CreatePayment{
		UserId: grpc.UserId,
		TransactionId: grpc.TransactionId,
		Price: grpc.Price,
	}
}