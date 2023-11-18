package handler

import (
	"net/http"
	"time"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateFlashCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var updateFlachCardRequest UpdateFlashCardRequest
	ctx.BindJSON(&updateFlachCardRequest)

	err := updateFlachCardRequest.FlashCardUpdateValidate()

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var flashCard schemas.FlashCards
	if err := handler.DB.First(&flashCard, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "flash card does not found")
		return
	}

	fromUpdateRequestFlashCardToFlashCard(updateFlachCardRequest, &flashCard)

	// TODO: add rules to handle with Last and Next revision
	flashCard.LastRevision = time.Now()
	flashCard.NextRevision = time.Now()

	if err = handler.DB.Save(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error on update update flash card")
		return
	}

	response := fromReadFlashCardToResponse(&flashCard)

	sendSuccess(ctx, http.StatusOK, "read flash card", response)

}
