package yookassa

import (
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
)


func GetYooKassaPayment(cfg *config.YooKassa)*yookassa.PaymentHandler{
	yooclient := yookassa.NewClient(cfg.ID,cfg.SecretKey)
	payment := yookassa.NewPaymentHandler(yooclient)
	return payment
}