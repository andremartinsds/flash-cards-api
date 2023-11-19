package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An ID is required.")
		return
	}

	var user schemas.User
	if err := handler.DB.Preload("Dec").First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "The user was not found.")
		return
	}

	if userHasADec(user) {
		sendError(ctx, http.StatusBadRequest, "This user has a 'dec,' so you can't delete him/her.")
		return
	}

	if err := handler.DB.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "We encountered a problem while trying to remove it. Please try again.")
		return
	}

	response := fromDeleteUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "deleted", response)

}

func userHasADec(user schemas.User) bool {
	return len(user.Dec) > 0
}
