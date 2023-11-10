package router

import (
	handler "github.com/andremartinsds/flash-cards-api/handler/user"
	"github.com/gin-gonic/gin"
)

func initilizeRoutes(router *gin.Engine) {

	uriPath := "/api/"

	api := router.Group(uriPath)
	{
		api.POST("user", handler.CreateUserHandler)
	}

}
