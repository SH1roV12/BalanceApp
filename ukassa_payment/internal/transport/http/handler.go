package http

import (
	"encoding/json"
	"net/http"

	"github.com/SH1roV12/balance/ukassa/internal/domain/service"
	"github.com/SH1roV12/balance/ukassa/internal/transport/dto/request"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	yoowebhook "github.com/rvinnie/yookassa-sdk-go/yookassa/webhook"
	"go.uber.org/zap"
)



type YooKassaPaymentConfirmHTTPHandlers struct{
	service service.Service
	sugar *zap.SugaredLogger
}


func NewYooKassaPaymentConfirmHTTPHandlers(service service.Service,sugar *zap.SugaredLogger)*YooKassaPaymentConfirmHTTPHandlers{
	return &YooKassaPaymentConfirmHTTPHandlers{
		service: service,
		sugar: sugar,
	}
}


func(h *YooKassaPaymentConfirmHTTPHandlers)HandleWebhook(w http.ResponseWriter, r *http.Request){
	var webhookEvent yoowebhook.WebhookEvent[yoopayment.Payment]
	err := json.NewDecoder(r.Body).Decode(&webhookEvent)
	if err != nil {
		http.Error(w, "Invalid webhook data", http.StatusBadRequest)
		return
	}
	
	err = h.service.ConfirmPayment(r.Context(),request.ToDomainConfirmPayment(webhookEvent))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.sugar.Infoln("Webhook обработан: %+v", webhookEvent)
	h.sugar.Infoln("Тип Webhook: %+v", webhookEvent.Type)
	h.sugar.Infoln("Событие: %+v", webhookEvent.Event)
	h.sugar.Infoln("Данные о платеже: %+v", webhookEvent.Object)

	
	w.WriteHeader(http.StatusOK)
}