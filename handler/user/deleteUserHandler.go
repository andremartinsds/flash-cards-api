package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Delete User
// @Description User
// @Tags User
// @Accept json
// @Produce json
// @Param id query string true "User identification"
// @Success 200 {object} UserDeleteResponse
// @Failure 400 {object} UserErrorResponse
// @Failure 404 {object} UserErrorResponse
// @Router /user [delete]
func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An ID is required.")
		return
	}

	var user schemas.User
	if err := handler.DB.Preload("Dec").First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "User was not found.")
		return
	}

	if userHasADec(user) {
		sendError(ctx, http.StatusBadRequest, "This user has a 'dec,' so deletion is not allowed.")
		return
	}

	if err := handler.DB.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "We encountered an issue while attempting to remove it. Kindly try again.")
		return
	}

	response := fromDeleteUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "deleted", response)
}

func userHasADec(user schemas.User) bool {
	return len(user.Dec) > 0
}
