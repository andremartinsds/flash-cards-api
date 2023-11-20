package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Delete Flash-Card
// @Description Flash-Card
// @Tags FlashCard
// @Accept json
// @Produce json
// @Param id query string true "Flash-Card identification"
// @Success 200 {object} FlashCardDeleteResponse
// @Failure 400 {object} FlashCardErrorResponse
// @Failure 404 {object} FlashCardErrorResponse
// @Router /flash-card [delete]
func DeleteFlashCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An id is requried")
		return
	}

	var flashCard schemas.FlashCards
	if err := handler.DB.First(&flashCard, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "FlashCard not found")
		return
	}

	if err := handler.DB.Delete(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "We are experiencing an issue with the removal process. Please try again.")
		return
	}

	response := fromDeleteFlashCardToResponse(&flashCard)

	sendSuccess(ctx, http.StatusOK, "deleted", response)

}
