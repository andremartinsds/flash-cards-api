package handler

import (
	"errors"
	"net/http"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.UserCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := userAlreadExists(request.Email)

	if err != nil {
		sendError(ctx, http.StatusConflict, err.Error())
		return
	}

	user := fromRequestUserToUser(request)

	if err := db.Create(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := fromUserSaveToResponse(user)

	sendSuccess(ctx, http.StatusCreated, "create-user", response)

}

func userAlreadExists(email string) error {
	var userFound schemas.User
	db.First(&userFound, "email = ?", email)

	if userFound.UserExistis() {
		return errors.New("The user already exists")
	}

	return nil
}
