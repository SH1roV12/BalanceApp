package handler

import "github.com/SH1roV12/balance/internal/domain/user"


type Handlers struct{
	userService user.Service
}

func NewHandlers(service user.Service)*Handlers{
	return &Handlers{userService: service}
}