package handler

import (
	"net/http"

	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/andremartinsds/flash-cards-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param request body UpdateUserRequest true "Request body"
// @Param id query string true "User identification"
// @Success 200 {object} UserUpdateResponse
// @Failure 400 {object} UserErrorResponse
// @Failure 404 {object} UserErrorResponse
// @Router /user [put]
func UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "An ID is required")
		return
	}

	var updateUserRequest UpdateUserRequest
	ctx.BindJSON(&updateUserRequest)

	err := updateUserRequest.UpdateUserValidate()

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var user schemas.User
	if err := handler.DB.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "User was not found")
		return
	}

	fromUpdateRequestUserToUserModel(updateUserRequest, &user)

	if err = handler.DB.Save(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error on update user")
		return
	}

	response := fromReadUserToResponse(&user)

	sendSuccess(ctx, http.StatusOK, "read user", response)

}
