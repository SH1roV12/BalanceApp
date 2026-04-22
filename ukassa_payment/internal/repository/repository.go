package repository

import "github.com/SH1roV12/balance/ukassa/internal/infra/postgres"



type Repository struct{
	DB *postgres.Postgres
}

func NewRepository(db *postgres.Postgres)*Repository {
	return  &Repository{DB: db}
}