package entity


type User struct{
	ID        string
	FirstName string 
	LastName  string 
	Username  string 
	Balance   float64
	Email     string
	Password  string
	Role      Role
}

func NewUser(id,first_name,last_name,username string,balance float64, email,password string)*User{
	return &User{
		ID: id,
		FirstName: first_name,
		LastName: last_name,
		Username: username,
		Balance: balance,
		Email: email,
		Password: password,
	}
}

type Role string

var UserRoleAdmin Role = "admin"
var UserRoleUser  Role = "user"