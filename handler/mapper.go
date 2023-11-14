package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

type UserCreateResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func fromRequestUserToUser(u CreateUserRequest) *schemas.User {
	return &schemas.User{
		Name:  u.Name,
		Email: u.Email,
		Pass:  u.Pass,
	}
}

func fromUserSaveToResponse(u *schemas.User) UserCreateResponse {
	return UserCreateResponse{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
