package handler

import (
	"github.com/SH1roV12/balance/internal/domain/user"
	"go.uber.org/zap"
)


type Handlers struct{
	userService user.Service
	sugar *zap.SugaredLogger
}

func NewHandlers(service user.Service,sugar *zap.SugaredLogger)*Handlers{
	return &Handlers{userService: service, sugar: sugar}
}