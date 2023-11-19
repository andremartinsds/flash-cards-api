package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Delete dec
// @Description Delete dec
// @Tags Dec
// @Accept json
// @Produce json
// @Param id query string true "Dec identification"
// @Success 200 {object} DeleteDecResponse
// @Failure 400 {object} ErrorDecResponse
// @Failure 404 {object} ErrorDecResponse
// @Router /dec [delete]
func DeleteDecHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var dec schemas.Dec
	if err := handler.DB.First(&dec, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "dec does not found")
		return
	}

	if err := handler.DB.Delete(&dec).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "we have a problem to remove it, please try again")
		return
	}

	response := fromDeleteDecToResponse(&dec)

	sendSuccess(ctx, http.StatusOK, "deleted", response)

}
