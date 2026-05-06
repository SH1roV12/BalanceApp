package handler

import (
	"github.com/SH1roV12/balance/internal/domain/payment"
	"github.com/SH1roV12/balance/internal/domain/user"

	"go.uber.org/zap"
)


type Handlers struct{
	userService user.Service
	paymentService payment.Service
	sugar *zap.SugaredLogger
}

func NewHandlers(service user.Service,payment payment.Service,sugar *zap.SugaredLogger)*Handlers{
	return &Handlers{userService: service,paymentService: payment, sugar: sugar}
}