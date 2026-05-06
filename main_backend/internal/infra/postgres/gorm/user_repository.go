package postgres

import (
	"context"
	"errors"

	"github.com/SH1roV12/balance/internal/domain/entity"
	customErrors "github.com/SH1roV12/balance/internal/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct{
	ID string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Username string `gorm:"unique;not null"`
	Balance float64 
	Email   string  `gorm:"not null;unique"`
	Password string  `gorm:"not null"`
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
		Email: user.Email,
		Password: user.Password,
	}
}

func GormToEntity(user *User)*entity.User{
	return &entity.User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
		Balance: user.Balance,
		Email: user.Email,
		Password: user.Password,
	}
}

func GormsToEntitys(users []*User)[]*entity.User{
	entityUsers := make([]*entity.User, len(users))
	for i, _ := range users{
		entityUsers[i] = GormToEntity(users[i])
	}
	return entityUsers
}

func(repo *UserRepository) CreateUser(ctx context.Context,user *entity.User)(error){
	gormUser := EntityToGorm(user)
	err := repo.db.WithContext(ctx).Create(&gormUser).Error
	if err !=nil{
		if errors.Is(err, gorm.ErrDuplicatedKey){
			return customErrors.NewRepoAppError(customErrors.Repo.User.AlreadyExist,err, "create")
		}
		return customErrors.NewRepoAppError(customErrors.Repo.User.CannotCreate,err, "create")
	}
	return nil
}

func(repo *UserRepository) GetAllUsers(ctx context.Context)([]*entity.User,error){
	var users []*User
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil{
		return nil,customErrors.NewRepoAppError(customErrors.Repo.User.CannotGetAll,err, "get all users")
	}
	return GormsToEntitys(users), nil
}


func(repo *UserRepository) GetUserByID(ctx context.Context, user_id string)(*entity.User,error){
	var user *User
	err := repo.db.WithContext(ctx).Where("id = ?", user_id).First(&user).Error
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,customErrors.NewRepoAppError(customErrors.Repo.User.NotFound,err, "get user by id")
		}
		return nil,customErrors.NewRepoAppError(customErrors.Repo.User.CannotGetById,err, "get user by id")
	}
	return GormToEntity(user),nil
}

func(repo *UserRepository) GetUserByEmail(ctx context.Context, email string)(*entity.User, error){
	var user *User
	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,customErrors.NewRepoAppError(customErrors.Repo.User.NotFound,err, "get user by email")
		}
		return nil,customErrors.NewRepoAppError(customErrors.Repo.User.CannotGetByEmail,err, "get user by email")
	}
	return GormToEntity(user),nil
}


func(repo *UserRepository) AddBalance(ctx context.Context, user_id string,amount float64)(error){
	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&User{}).Where("id = ?",user_id).Update("balance",gorm.Expr("balance + ?",amount))
		if result.Error != nil{
			return result.Error
		}

		if result.RowsAffected == 0{
			return customErrors.NewRepoAppError(customErrors.Repo.User.NotFound,nil,"add balance")
		}
		return nil
	})
}