package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

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

func fromSaveUserToResponse(u *schemas.User) UserCreateResponseData {
	return UserCreateResponseData{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func fromListUserToResponse(u []schemas.User) []UserListResponseData {
	var userListResponseCollection []UserListResponseData
	for i := 0; i < len(u); i++ {
		userListResponse := UserListResponseData{
			Id:    u[i].ID,
			Name:  u[i].Name,
			Email: u[i].Email,
		}
		userListResponseCollection = append(userListResponseCollection, userListResponse)
	}
	return userListResponseCollection
}

func fromReadUserToResponse(u *schemas.User) UserReadResponseData {
	return UserReadResponseData{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func fromDeleteUserToResponse(u *schemas.User) UserDeleteResponseData {
	return UserDeleteResponseData{
		Id:    u.ID,
		Email: u.Email,
	}
}
