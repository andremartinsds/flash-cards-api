package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List dec
// @Description List dec
// @Tags Dec
// @Accept json
// @Produce json
// @Success 200 {object} ListDecResponse
// @Failure 400 {object} ErrorDecResponse
// @Failure 404 {object} ErrorDecResponse
// @Router /decs [get]
func ListDecHandler(ctx *gin.Context) {

	var dec []schemas.Dec
	if err := handler.DB.Find(&dec).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the dec is not found")
		return
	}

	response := fromListDecToResponse(dec)

	sendSuccess(ctx, http.StatusOK, "read dec", response)

}
