package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func ReadUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "the id is required")
		return
	}

	var user schemas.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the user is not found")
		return
	}

	response := fromReadUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "read user", response)

}
