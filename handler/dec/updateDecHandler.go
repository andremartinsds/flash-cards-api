package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateDecHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var updateDecRequest UpdateDecRequest
	ctx.BindJSON(&updateDecRequest)

	err := updateDecRequest.UpdateDecValidate()

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var dec schemas.Dec
	if err := handler.DB.First(&dec, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the dec is not found")
		return
	}

	fromUpdateRequestDecToDec(updateDecRequest, &dec)

	if err = handler.DB.Save(&dec).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error on update dec")
		return
	}

	response := fromReadDecToResponse(&dec)

	sendSuccess(ctx, http.StatusOK, "read dec", response)

}
