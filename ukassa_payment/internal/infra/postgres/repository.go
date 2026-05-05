package postgres

import (
	"context"

	"gorm.io/gorm"
)



type Repository struct{
	DB *gorm.DB
}

func NewRepository(db *gorm.DB)*Repository {
	return  &Repository{DB: db}
}



func(r *Repository)Transaction(ctx context.Context,fn func(tx any)error)(error){
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func(r *Repository)WithTx(tx any)*Repository{
	return &Repository{
		DB: tx.(*gorm.DB),
	}
}