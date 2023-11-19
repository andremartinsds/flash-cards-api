package handler

import (
	"errors"
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Create Dec
// @Description Create a new dec
// @Tags Dec
// @Accept json
// @Produce json
// @Param request body CreateDecRequest true "Request body"
// @Success 201 {object} CreateDecResponse
// @Failure 400 {object} ErrorDecResponse
// @Failure 500 {object} ErrorDecResponse
// @Router /dec [post]
func CreateDecHandler(ctx *gin.Context) {
	request := CreateDecRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "We have a problem with that, please try again")
		return
	}

	if err := request.DecCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var user schemas.User
	if err := handler.DB.First(&user, "id = ?", request.UserID).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "The user does not exists")
		return
	}

	err = decAlreadyExists(request)

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

func decAlreadyExists(request CreateDecRequest) error {
	var decFound schemas.Dec
	handler.DB.First(&decFound, "name = ? AND user_id = ?", request.Name, request.UserID)

	if decFound.DecExists() {
		return errors.New("The dec already exists")
	}

	return nil
}
