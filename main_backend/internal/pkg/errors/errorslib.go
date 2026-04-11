package customErrors

import "errors"

//Repo errors
var repoUserAlreadyExist error = errors.New("user already exist")
var repoCannotCreateUser error = errors.New("cannot create new user")
var repoCannotGetAllUsers error = errors.New("cannot get all users")
var repoNotFound error = errors.New("user not found")
var repoCannotGetUserById error = errors.New("cannot get user by id")
var repoCannotGetUserByEmail error = errors.New("cannot get user by email")

type repoUser struct{
	AlreadyExist error
	CannotCreate error
	CannotGetAll error
	CannotGetById error
	NotFound error
	CannotGetByEmail error
}


var Repo  =  struct{
	User repoUser
}{
	User: repoUser{
		AlreadyExist: repoUserAlreadyExist,
		CannotCreate: repoCannotCreateUser,
		CannotGetAll: repoCannotGetAllUsers,
		CannotGetById: repoCannotGetUserById,
		NotFound: repoNotFound,
		CannotGetByEmail: repoCannotGetUserByEmail,
	},
}



