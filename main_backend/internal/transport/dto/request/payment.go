package request

import "github.com/SH1roV12/balance/internal/gen/pb"

type Payment struct{
	Price float32 `json:"price"`
	
}

type Replenishment struct{
	UserId string
	Amount float32
}

func ToDomain(payment Payment, userId string)*pb.CreatePaymentRequest{
	return &pb.CreatePaymentRequest{
		Price: payment.Price,
		UserId: userId,
	}
}

func ToDomainReplenishment(in *pb.ConfirmPaymentRequest)*Replenishment{
	return &Replenishment{
		UserId: in.UserId,
		Amount: in.Amount,
	}
}