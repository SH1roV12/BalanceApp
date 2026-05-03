package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SH1roV12/balance/ukassa/internal/domain/entity"
	"github.com/SH1roV12/balance/ukassa/internal/domain/repository"
	pb "github.com/SH1roV12/balance/ukassa/internal/gen"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/config"
	"github.com/SH1roV12/balance/ukassa/internal/pkg/id_gen/uuid"
	"go.uber.org/zap"

	"github.com/SH1roV12/balance/ukassa/internal/transport/dto/request"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)



type Service struct{
	repo repository.Repository
	sdk *yookassa.PaymentHandler
	cfg *config.YooKassa
	sugar *zap.SugaredLogger
	client pb.UKassaPaymentConfirmationClient
}

func NewService(repo repository.Repository,sdk *yookassa.PaymentHandler, cfg *config.YooKassa,sugar *zap.SugaredLogger,client pb.UKassaPaymentConfirmationClient )*Service{
	return &Service{
		repo: repo,
		sdk: sdk,
		cfg: cfg,
		sugar: sugar,
		client: client,
	}
}

func(s *Service)CreatePayment(ctx context.Context, dto request.CreatePayment)(string,error){
	
	transactionId := uuid.GetID()
	s.sugar.Infow("Payment","dto", dto)
	err := s.repo.CreateTransaction(ctx,dto.UserId,transactionId,entity.Created,dto.Price)
	if err != nil {
		return "",err
	}
	
	payment,err := s.sdk.CreatePayment(ctx, &yoopayment.Payment{
		ID: transactionId,
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

	err = s.repo.UpdateTransaction(ctx, entity.InProgress,transactionId,payment.ID)
	if err != nil {
    	return "", err
	}

	return paymentURL, nil
}


func(s *Service)ConfirmPayment(ctx context.Context, dto request.ConfirmPayment)error{
	if dto.Status != "succeeded"{
		s.sugar.Infoln("waiting for payment...")
		return nil
	}
	if dto.Status == "canceled"{
		s.sugar.Infoln("payment was not successful")
		err := s.repo.UpdateTransactionByExternalID(ctx,entity.Canceled,dto.ID)
		if err != nil{
			return err
		}
		return nil
	}
	err := s.repo.UpdateTransactionByExternalID(ctx,entity.Succeeded,dto.ID)
	userId,err := s.repo.GetUserIDByExternalID(ctx,dto.ID)
	if err != nil{
		return err
	}
	in := pb.ConfirmPaymentRequest{
		UserId: userId,
		Amount: dto.Amount,
	}
	_,err = s.client.ConfirmPayment(ctx,&in)
	if err != nil{
		return err
	}
	return  nil
}