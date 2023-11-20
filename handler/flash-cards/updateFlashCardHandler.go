package handler

import (
	"net/http"
	"time"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update Flash-Card
// @Description Update Flash-Card
// @Tags FlashCard
// @Accept json
// @Produce json
// @Param request body UpdateFlashCardRequest true "Request body"
// @Param id query string true "Flash-Card identification"
// @Success 200 {object} FlashCardUpdateResponse
// @Failure 400 {object} FlashCardErrorResponse
// @Failure 404 {object} FlashCardErrorResponse
// @Router /flash-card [put]
func UpdateFlashCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An id is required")
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
		sendError(ctx, http.StatusNotFound, "Flash card not found.")
		return
	}

	fromUpdateRequestFlashCardToFlashCardModel(updateFlachCardRequest, &flashCard)

	// TODO: add rules to handle with Last and Next revision
	flashCard.LastRevision = time.Now()
	flashCard.NextRevision = time.Now()

	if err = handler.DB.Save(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Error on updating flash card.")
		return
	}

	response := fromReadFlashCardToResponse(&flashCard)

	sendSuccess(ctx, http.StatusOK, "read flash card", response)

}
