package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.UserCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := fromRequestUserToUser(request)

	if err := db.Create(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-user", user)

}
