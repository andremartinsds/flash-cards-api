package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func ReadDecHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var dec schemas.Dec
	if err := handler.DB.First(&dec, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the dec does not found")
		return
	}

	response := fromReadDecToResponse(&dec)

	sendSuccess(ctx, http.StatusOK, "read dec", response)

}
