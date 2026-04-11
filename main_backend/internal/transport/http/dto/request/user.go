package request


type RegisterUser struct{
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct{
	Email     string `json:"email"`
	Password  string `json:"password"`
}