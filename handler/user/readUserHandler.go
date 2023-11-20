package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Read User
// @Description Read User
// @Tags User
// @Produce json
// @Param id query string true "User identification"
// @Success 200 {object} UserReadResponse
// @Failure 400 {object} UserErrorResponse
// @Failure 404 {object} UserErrorResponse
// @Router /user [get]
func ReadUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var user schemas.User
	if err := handler.DB.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the user is not found")
		return
	}

	response := fromReadUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "read user", response)

}
