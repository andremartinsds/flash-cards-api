package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListDecHandler(ctx *gin.Context) {

	var dec []schemas.Dec
	if err := handler.DB.Find(&dec).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the dec is not found")
		return
	}

	response := fromListDecToResponse(dec)

	sendSuccess(ctx, http.StatusOK, "read dec", response)

}
