package user

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
)

type Service interface{
	NewUser(ctx context.Context, req *request.RegisterUser)(*entity.User,error)
	GetAll(ctx context.Context)([]*entity.User,error)
	GetById(ctx context.Context, user_id string)(*entity.User, error)
}