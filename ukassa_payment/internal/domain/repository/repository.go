package repository

import (
	"context"

	"github.com/SH1roV12/balance/ukassa/internal/infra/postgres"
)


type Repository interface{
	CreateTransaction(ctx context.Context,user_id,transaction_id,status string,price float32)(error)
	UpdateTransaction(ctx context.Context, status,id,external_id string)(error)
	UpdateTransactionByExternalID(ctx context.Context, status,external_id string)(error)
	GetUserIDByExternalID(ctx context.Context,external_id string)(string,error)
	GetStatusByExternalID(ctx context.Context,external_id string)(string,error)
	Transaction(ctx context.Context,fn func(tx any)error)(error)
	WithTx(tx any)*postgres.Repository
}