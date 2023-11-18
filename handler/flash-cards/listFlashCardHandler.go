package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListFlashCardHandler(ctx *gin.Context) {

	var flashCard []schemas.FlashCards
	if err := handler.DB.Find(&flashCard).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the flashCard collection does not found")
		return
	}

	response := fromListFlashCardToResponse(flashCard)

	sendSuccess(ctx, http.StatusOK, "read flash card", response)

}
