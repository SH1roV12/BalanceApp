package user

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
)


type Repository interface{
	Create(ctx context.Context,user *entity.User)(*entity.User,error)
	GetAll(ctx context.Context)([]*entity.User,error)
	GetByID(ctx context.Context, user_id string)(*entity.User,error)
}