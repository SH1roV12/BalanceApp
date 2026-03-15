package response

import "github.com/SH1roV12/balance/internal/domain/entity"


type User struct{
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Balance   float64 `json:"balance"`
}

func FromEntityToDTO(user *entity.User)*User{
	return &User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
		Balance: user.Balance,
	}
}

func FromEntitysToDTOs(users []*entity.User)[]*User{
	dtoUsers := make([]*User, len(users))
	for i := range users{
		dtoUsers[i] = FromEntityToDTO(users[i])
	}
	return dtoUsers
}