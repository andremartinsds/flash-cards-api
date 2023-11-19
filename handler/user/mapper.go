package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

type UserCreateResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserReadResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserListResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDelete struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func fromRequestUserToUserModel(u CreateUserRequest) *schemas.User {
	return &schemas.User{
		Name:  u.Name,
		Email: u.Email,
		Pass:  u.Pass,
	}
}
func fromUpdateRequestUserToUserModel(u UpdateUserRequest, user *schemas.User) *schemas.User {

	if u.Email != "" {
		user.Name = u.Name
	}

	if u.Name != "" {
		user.Email = u.Email
	}

	if u.Pass != "" {
		user.Pass = u.Pass
	}

	return user
}

func fromSaveUserToResponse(u *schemas.User) UserCreateResponse {
	return UserCreateResponse{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func fromListUserToResponse(u []schemas.User) []UserListResponse {
	var userListResponseCollection []UserListResponse
	for i := 0; i < len(u); i++ {
		userListResponse := UserListResponse{
			Id:    u[i].ID,
			Name:  u[i].Name,
			Email: u[i].Email,
		}
		userListResponseCollection = append(userListResponseCollection, userListResponse)
	}
	return userListResponseCollection
}

func fromReadUserToResponse(u *schemas.User) UserReadResponse {
	return UserReadResponse{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func fromDeleteUserToResponse(u *schemas.User) UserDelete {
	return UserDelete{
		Id:    u.ID,
		Email: u.Email,
	}
}
