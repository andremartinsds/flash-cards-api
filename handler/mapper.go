package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

func fromRequestUserToUser(u CreateUserRequest) *schemas.User {
	return &schemas.User{
		Name:  u.Name,
		Email: u.Email,
		Pass:  u.Pass,
	}
}
