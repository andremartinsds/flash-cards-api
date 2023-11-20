package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Read Flash-card
// @Description Read Flash-card
// @Tags FlashCard
// @Produce json
// @Param id query string true "Flash-card identification"
// @Success 200 {object} FlashCardReadResponse
// @Failure 400 {object} FlashCardErrorResponse
// @Failure 404 {object} FlashCardErrorResponse
// @Router /flash-card [get]
func ReadFlashCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An id is required")
		return
	}

	var flashCard schemas.FlashCards
	if err := handler.DB.First(&flashCard, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Flash card not found")
		return
	}

	response := fromReadFlashCardToResponse(&flashCard)

	sendSuccess(ctx, http.StatusOK, "read flash card", response)

}
