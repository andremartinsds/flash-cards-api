package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListUserHandler(ctx *gin.Context) {

	var user []schemas.User
	if err := db.Find(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "the user is not found")
		return
	}

	response := fromListUserToResponse(user)

	sendSuccess(ctx, http.StatusOK, "read user", response)

}
