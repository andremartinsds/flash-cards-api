package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Create a Flash-Card
// @Description Create a new Flash-Card
// @Tags FlashCard
// @Accept json
// @Produce json
// @Param request body CreateFlashCardRequest true "Request body"
// @Success 201 {object} FlashCardCreateResponse
// @Failure 400 {object} FlashCardErrorResponse
// @Failure 500 {object} FlashCardErrorResponse
// @Router /flash-card [post]
func CreateFlashCardHandler(ctx *gin.Context) {
	request := CreateFlashCardRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.FlashCardCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var dec schemas.Dec
	if err := handler.DB.First(&dec, "id = ?", request.DecId).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "The dec does not exists")
		return
	}

	err = flashCardAlreadyExists(request)

	if err != nil {
		sendError(ctx, http.StatusConflict, err.Error())
		return
	}

	flashCard := fromRequestFlashCardToFlashCardModel(request)

	// TODO: add rules to handle with Last and Next revision
	flashCard.LastRevision = time.Now()
	flashCard.NextRevision = time.Now()

	if err := handler.DB.Create(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := fromSaveFlashCardToResponse(flashCard)

	sendSuccess(ctx, http.StatusCreated, "create-flash-card", response)

}

func flashCardAlreadyExists(request CreateFlashCardRequest) error {
	var flashCard schemas.FlashCards
	handler.DB.First(&flashCard, "front = ? AND dec_id = ?", request.Front, request.DecId)

	if flashCard.FlashCardExists() {
		return errors.New("The flash card already exists")
	}

	return nil
}
