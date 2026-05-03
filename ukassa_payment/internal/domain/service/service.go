package service

import (
	"context"

	"github.com/SH1roV12/balance/ukassa/internal/transport/dto/request"
)

type Service interface{
	CreatePayment(ctx context.Context, dto request.CreatePayment)(string,error)
	ConfirmPayment(ctx context.Context, dto request.ConfirmPayment)error
}