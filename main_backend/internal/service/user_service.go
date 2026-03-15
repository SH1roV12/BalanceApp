package service

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/domain/user"
	"github.com/SH1roV12/balance/internal/pkg/id_generator/uuid"

	"github.com/SH1roV12/balance/internal/transport/http/dto/request"
)

type UserService struct{
	repo user.Repository
}

func NewUserService(repo user.Repository)*UserService{
	return &UserService{repo: repo}
}


func(s *UserService) NewUser(ctx context.Context, req *request.RegisterUser)error{
	id := uuid.GetID()
	user := entity.NewUser(id,req.FirstName,req.LastName,req.Username,0.00)
	return  s.repo.Create(ctx,user)
}

func(s *UserService) GetAll(ctx context.Context)([]*entity.User,error){
	return s.repo.GetAll(ctx)
}