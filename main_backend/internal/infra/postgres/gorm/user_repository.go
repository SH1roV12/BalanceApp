package postgres

import (
	"context"

	"github.com/SH1roV12/balance/internal/domain/entity"
	"github.com/SH1roV12/balance/internal/infra/errorsrepo"
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
}

func NewUserRepository(db *gorm.DB)*UserRepository{
	return &UserRepository{db: db}
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

func(repo *UserRepository) Create(ctx context.Context,user *entity.User)error{
	gormUser := EntityToGorm(user)
	err := repo.db.WithContext(ctx).Create(&gormUser).Error
	if err !=nil{
		return errorsrepo.ErrCannotCreateUser
	}
	return nil
}

func(repo *UserRepository) GetAll(ctx context.Context)([]*entity.User,error){
	var users []*User
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil{
		return nil,err
	}
	return GormsToEntitys(users), nil
}