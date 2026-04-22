package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SH1roV12/balance/ukassa/internal/domain/entity"
	"github.com/SH1roV12/balance/ukassa/internal/domain/repository"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	"github.com/SH1roV12/balance/ukassa/internal/transport/dto/request"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)



type Service struct{
	repo repository.Repository
	sdk *yookassa.PaymentHandler
	cfg *config.YooKassa
}

func NewService(repo repository.Repository,sdk *yookassa.PaymentHandler, cfg *config.YooKassa )*Service{
	return &Service{
		repo: repo,
		sdk: sdk,
		cfg: cfg,
	}
}

func(s *Service)CreatePayment(ctx context.Context, dto request.CreatePayment)(string,error){
	
	err := s.repo.CreateTransaction(ctx,dto.UserId,dto.TransactionId,entity.Created,dto.Price)
	if err != nil {
		return "",err
	}
	
	payment,err := s.sdk.CreatePayment(ctx, &yoopayment.Payment{
		ID: dto.TransactionId,
		Amount: &yoocommon.Amount{
			Value: fmt.Sprintf("%.2f", dto.Price),
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
		Confirmation: yoopayment.Redirect{
			Type: yoopayment.TypeRedirect,
			ReturnURL: s.cfg.ReturnURL,
		},
		Description: "Test payment",
	})
	if err != nil || payment == nil{
		return "",err
	}
	


	var paymentURL string
	if payment.Confirmation != nil {
    	if conf, ok := payment.Confirmation.(map[string]interface{}); ok {
        	paymentURL, _ = conf["confirmation_url"].(string)
    	}
	}
	
	if paymentURL == ""{
		return "", errors.New("failed to get url")
	}

	err = s.repo.UpdateTransaction(ctx, entity.InProgress, dto.TransactionId,payment.ID)
	if err != nil {
    	return "", err
	}

	return paymentURL, nil
}