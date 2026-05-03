package postgres



type Repository struct{
	DB *Postgres
}

func NewRepository(db *Postgres)*Repository {
	return  &Repository{DB: db}
}