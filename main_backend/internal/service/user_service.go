package service

import (
	"context"
	"errors"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/domain/user"
	"github.com/SH1roV12/balance/internal/pkg/id_generator/uuid"
	passwordHash "github.com/SH1roV12/balance/internal/pkg/passwordHasher/bcrypt"
	"github.com/SH1roV12/balance/internal/transport/dto/request"
	"go.uber.org/zap"
)

type UserService struct{
	repo user.Repository
	sugar *zap.SugaredLogger
}

func NewUserService(repo user.Repository,sugar *zap.SugaredLogger)*UserService{
	return &UserService{repo: repo, sugar: sugar}
}


func(s *UserService) NewUser(ctx context.Context, req *request.RegisterUser)(error){
	id := uuid.GetID()
	hashedPassword,err := passwordHash.GeneratePasswordHash(req.Password)
	if err != nil{
		s.sugar.Errorw("service", "hash error", err.Error())
		return err
	}
	user := entity.NewUser(id,req.FirstName,req.LastName,req.Username,0.00, req.Email, hashedPassword)
	return  s.repo.CreateUser(ctx,user)
}

func(s *UserService) GetAll(ctx context.Context)([]*entity.User,error){
	return s.repo.GetAllUsers(ctx)
}

func(s * UserService) GetById(ctx context.Context, user_id string)(*entity.User, error){
	return s.repo.GetUserByID(ctx,user_id)
}

func(s *UserService) GetByEmail(ctx context.Context, email,password string)(*entity.User, error){
	user,err := s.repo.GetUserByEmail(ctx,email)
	if err != nil{
		return nil,err
	}
	if !passwordHash.CheckHashPassword(password,user.Password){
		return nil,errors.New("wrong password")
	}
	return user,nil
}

func(s *UserService)SetBalance(ctx context.Context, req *request.Replenishment)error{
	err := s.repo.AddBalance(ctx,req.UserId,float64(req.Amount))
	if err != nil{
		return err
	}
	return nil
}