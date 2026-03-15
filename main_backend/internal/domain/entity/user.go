package entity


type User struct{
	ID        string
	FirstName string 
	LastName  string 
	Username  string 
	Balance   float64
}

func NewUser(id,first_name,last_name,username string,balance float64 )*User{
	return &User{
		ID: id,
		FirstName: first_name,
		LastName: last_name,
		Username: username,
		Balance: balance,
	}
}