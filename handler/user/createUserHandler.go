package handler

import (
	"errors"
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 201 {object} UserCreateResponse
// @Failure 400 {object} UserErrorResponse
// @Failure 500 {object} UserErrorResponse
// @Router /user [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.UserCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = userAlreadyExists(request.Email)

	if err != nil {
		sendError(ctx, http.StatusConflict, err.Error())
		return
	}

	user := fromRequestUserToUserModel(request)

	if err := handler.DB.Create(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := fromSaveUserToResponse(user)

	sendSuccess(ctx, http.StatusCreated, "create-user", response)

}

func userAlreadyExists(email string) error {
	var userFound schemas.User
	handler.DB.First(&userFound, "email = ?", email)

	if userFound.UserExists() {
		return errors.New("User already exists")
	}

	return nil
}
