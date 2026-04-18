package handler

import (
	"github.com/SH1roV12/balance/internal/domain/math"
	"github.com/SH1roV12/balance/internal/domain/user"
	"go.uber.org/zap"
)


type Handlers struct{
	userService user.Service
	mathService math.Service
	sugar *zap.SugaredLogger
}

func NewHandlers(service user.Service,math math.Service,sugar *zap.SugaredLogger)*Handlers{
	return &Handlers{userService: service,mathService: math, sugar: sugar}
}