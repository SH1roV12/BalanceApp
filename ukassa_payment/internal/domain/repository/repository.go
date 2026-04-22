package repository

import "context"


type Repository interface{
	CreateTransaction(ctx context.Context,user_id,transaction_id,status string,price float32)(error)
	UpdateTransaction(ctx context.Context, status,id,external_id string)(error)
}