package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteFlashCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var flashCard schemas.FlashCards
	if err := handler.DB.First(&flashCard, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "flashCard does not found")
		return
	}

	if err := handler.DB.Delete(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "we have a problem to remove it, please try again")
		return
	}

	response := fromDeleteFlashCardToResponse(&flashCard)

	sendSuccess(ctx, http.StatusOK, "deleted", response)

}
