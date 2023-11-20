package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List User
// @Description List User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserListResponse
// @Failure 400 {object} UserErrorResponse
// @Failure 404 {object} UserErrorResponse
// @Router /users [get]
func ListUserHandler(ctx *gin.Context) {

	var user []schemas.User
	if err := handler.DB.Find(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "User was not found")
		return
	}

	response := fromListUserToResponse(user)

	sendSuccess(ctx, http.StatusOK, "read user", response)

}
