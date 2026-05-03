package postgres

import (
	"context"
	"errors"
	"time"
)


type Transaction struct{
	ID string `gorm:"unique;not null"`
	UserID string `gorm:"not null;column:user_id"`
	ExternalID string    `gorm:"column:external_id"`
	Status string `gorm:"not null"`
	Price float32 `gorm:"not null"`
	CreatedAt time.Time
}

func(r *Repository) CreateTransaction(ctx context.Context,user_id,transaction_id,status string,price float32)(error){
	var Transaction Transaction = Transaction{
		ID: transaction_id,
		UserID: user_id,
		Price: price,
		Status: status,
	} 
	err := r.DB.WithContext(ctx).Create(&Transaction).Error
	if err != nil{
		return err
	}
	return nil
}

func(r *Repository) UpdateTransaction(ctx context.Context, status,id,external_id string)(error){
	req := r.DB.WithContext(ctx).Model(&Transaction{}).Where("id = ?", id).Updates(map[string]interface{}{"status":status,"external_id":external_id})
	if req.Error != nil{
		return req.Error
	}
	if req.RowsAffected == 0 {
		return FailedUpdate 
	}
	return nil
}


func(r *Repository) UpdateTransactionByExternalID(ctx context.Context, status,external_id string)(error){
	req := r.DB.WithContext(ctx).Model(&Transaction{}).Where("external_id = ?", external_id).Updates(map[string]interface{}{"status":status})
	if req.Error != nil{
		return req.Error
	}
	if req.RowsAffected == 0 {
		return FailedUpdate 
	}
	return nil
}


func(r *Repository) GetUserIDByExternalID(ctx context.Context,external_id string)(string,error){
	var UserID string
	err := r.DB.WithContext(ctx).Model(&Transaction{}).Where("external_id = ?", external_id).Select("user_id").Scan(&UserID).Error
	if err != nil{
		return "",err
	}
	return UserID,nil
}


var FailedUpdate error = errors.New("failed to update transaction ")