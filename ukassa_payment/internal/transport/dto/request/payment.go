package request

import (
	"strconv"

	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	yoowebhook "github.com/rvinnie/yookassa-sdk-go/yookassa/webhook"
)


type CreatePayment struct{
	UserId        string                 `json:"user_id,omitempty"`
    Price         float32                `json:"price,omitempty"`
}


type ConfirmPayment struct{
	ID string
	Status string
	Amount float32
}

func ToDomainCreatePayment(grpc *pb.CreatePaymentRequest)CreatePayment{
	return CreatePayment{
		UserId: grpc.UserId,
		Price: grpc.Price,
	}
}

func ToDomainConfirmPayment(in yoowebhook.WebhookEvent[yoopayment.Payment])ConfirmPayment{
	value64,_:= strconv.ParseFloat(in.Object.Amount.Value, 32)
	
	
	return ConfirmPayment{
		ID: in.Object.ID,
		Status: string(in.Object.Status),
		Amount: float32(value64),
	}
}