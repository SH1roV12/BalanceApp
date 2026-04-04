package service

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/domain/user"
	"github.com/SH1roV12/balance/internal/pkg/id_generator/uuid"
	"go.uber.org/zap"

	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
)

type UserService struct{
	repo user.Repository
	sugar *zap.SugaredLogger
}

func NewUserService(repo user.Repository,sugar *zap.SugaredLogger)*UserService{
	return &UserService{repo: repo, sugar: sugar}
}


func(s *UserService) NewUser(ctx context.Context, req *request.RegisterUser)(*entity.User,error){
	id := uuid.GetID()
	user := entity.NewUser(id,req.FirstName,req.LastName,req.Username,0.00)
	return  s.repo.Create(ctx,user)
}

func(s *UserService) GetAll(ctx context.Context)([]*entity.User,error){
	return s.repo.GetAll(ctx)
}

func(s * UserService) GetById(ctx context.Context, user_id string)(*entity.User, error){
	return s.repo.GetByID(ctx,user_id)
}

