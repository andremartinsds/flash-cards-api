package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateFlashCardHandler(ctx *gin.Context) {
	request := CreateFlashCardRequest{}

	ctx.BindJSON(&request)

	if err := request.FlashCardCreateValidate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var dec schemas.Dec
	if err := handler.DB.First(&dec, "id = ?", request.DecId).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "The dec does not exists")
		return
	}

	err := flashCardAlreadExists(request)

	if err != nil {
		sendError(ctx, http.StatusConflict, err.Error())
		return
	}

	flashCard := fromRequestFlashCardToFlashCard(request)

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

func flashCardAlreadExists(request CreateFlashCardRequest) error {
	var flashCard schemas.FlashCards
	handler.DB.First(&flashCard, "front = ? AND dec_id = ?", request.Front, request.DecId)

	if flashCard.FlashCardExists() {
		return errors.New("The flash card already exists")
	}

	return nil
}
