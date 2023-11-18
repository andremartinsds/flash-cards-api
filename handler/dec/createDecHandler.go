package handler

import (
	"errors"
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateDecHandler(ctx *gin.Context) {
	request := CreateDecRequest{}

	ctx.BindJSON(&request)

	if err := request.DecCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var user schemas.User
	if err := handler.DB.First(&user, "id = ?", request.UserID).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "The user does not exists")
		return
	}

	err := decAlreadExists(request)

	if err != nil {
		sendError(ctx, http.StatusConflict, err.Error())
		return
	}

	dec := fromRequestDecToDec(request)

	if err := handler.DB.Create(&dec).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := fromSaveDecToResponse(dec)

	sendSuccess(ctx, http.StatusCreated, "create-user", response)

}

func decAlreadExists(request CreateDecRequest) error {
	var decFound schemas.Dec
	handler.DB.First(&decFound, "name = ? AND user_id = ?", request.Name, request.UserID)

	if decFound.DecExists() {
		return errors.New("The dec already exists")
	}

	return nil
}
