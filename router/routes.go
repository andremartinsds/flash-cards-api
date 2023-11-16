package router

import (
	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/gin-gonic/gin"
)

func initilizeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	uriPath := "/api/"

	api := router.Group(uriPath)
	{
		api.POST("user", handler.CreateUserHandler)
		api.GET("user", handler.ReadUserHandler)
		api.DELETE("user", handler.DeleteUserHandler)
	}

}
