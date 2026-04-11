package user

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
)


type Repository interface{
	CreateUser(ctx context.Context,user *entity.User)(error)
	GetAllUsers(ctx context.Context)([]*entity.User,error)
	GetUserByID(ctx context.Context, user_id string)(*entity.User,error)
	GetUserByEmail(ctx context.Context, email string)(*entity.User, error)
}