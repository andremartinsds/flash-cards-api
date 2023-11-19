package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update dec
// @Description Update dec
// @Tags Dec
// @Accept json
// @Produce json
// @Param request body UpdateDecRequest true "Request body"
// @Param id query string true "Dec identification"
// @Success 200 {object} UpdateDecResponse
// @Failure 400 {object} ErrorDecResponse
// @Failure 404 {object} ErrorDecResponse
// @Router /dec [put]
func UpdateDecHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var updateDecRequest UpdateDecRequest
	err := ctx.BindJSON(&updateDecRequest)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = updateDecRequest.UpdateDecValidate()

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

	response := fromUpdateDecToResponse(&dec)

	sendSuccess(ctx, http.StatusOK, "read dec", response)

}
