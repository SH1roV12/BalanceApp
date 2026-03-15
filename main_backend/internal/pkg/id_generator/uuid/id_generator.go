package uuid

import "github.com/google/uuid"

func GetID()string{
	id := uuid.NewString()
	return id
}