package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List Flash-Card
// @Description List Flash-Card
// @Tags FlashCard
// @Accept json
// @Produce json
// @Success 200 {object} FlashCardListResponse
// @Failure 400 {object} FlashCardErrorResponse
// @Failure 404 {object} FlashCardErrorResponse
// @Router /flash-card [get]
func ListFlashCardHandler(ctx *gin.Context) {

	var flashCard []schemas.FlashCards
	if err := handler.DB.Find(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "FlashCard collection not found")
		return
	}

	response := fromListFlashCardToResponse(flashCard)

	sendSuccess(ctx, http.StatusOK, "read flash card", response)

}
