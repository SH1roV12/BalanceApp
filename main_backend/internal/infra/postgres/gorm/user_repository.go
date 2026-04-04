package postgres

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/infra/errorsrepo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct{
	ID string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Username string `gorm:"unique;not null"`
	Balance float64 
}

type UserRepository struct{
	db *gorm.DB
	sugar *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB,sugar *zap.SugaredLogger)*UserRepository{
	return &UserRepository{db: db, sugar: sugar}
}

func EntityToGorm(user *entity.User)*User{
	return &User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
		Balance: user.Balance,
	}
}

func GormToEntity(user *User)*entity.User{
	return &entity.User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
		Balance: user.Balance,
	}
}

func GormsToEntitys(users []*User)[]*entity.User{
	entityUsers := make([]*entity.User, len(users))
	for i, _ := range users{
		entityUsers[i] = GormToEntity(users[i])
	}
	return entityUsers
}

func(repo *UserRepository) Create(ctx context.Context,user *entity.User)(*entity.User,error){
	gormUser := EntityToGorm(user)
	err := repo.db.WithContext(ctx).Create(&gormUser).Error
	if err !=nil{
		return nil,errorsrepo.ErrCannotCreateUser
	}
	return GormToEntity(gormUser),nil
}

func(repo *UserRepository) GetAll(ctx context.Context)([]*entity.User,error){
	var users []*User
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil{
		return nil,err
	}
	return GormsToEntitys(users), nil
}


func(repo *UserRepository) GetByID(ctx context.Context, user_id string)(*entity.User,error){
	var user *User
	err := repo.db.WithContext(ctx).First(&user).Where("id = ?", user_id).Error
	if err != nil{
		return nil,err
	}
	return GormToEntity(user),nil
}