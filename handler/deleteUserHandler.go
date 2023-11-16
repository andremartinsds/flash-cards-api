package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var user schemas.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "does user not found")
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "we have a problem to remove it, please try again")
		return
	}

	response := fromDeleteUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "deleted", response)

}
