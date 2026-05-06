package user

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/transport/dto/request"
)

type Service interface{
	NewUser(ctx context.Context, req *request.RegisterUser)(error)
	GetAll(ctx context.Context)([]*entity.User,error)
	GetById(ctx context.Context, user_id string)(*entity.User, error)
	GetByEmail(ctx context.Context, email,password string)(*entity.User, error)
	SetBalance(ctx context.Context, req *request.Replenishment)error
}