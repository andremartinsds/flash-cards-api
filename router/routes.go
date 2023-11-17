package router

import (
	"github.com/andremartinsds/flash-cards-api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initilizeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	uriPath := "/api/"

	router.Use(cors.Default())
	api := router.Group(uriPath)
	{

		api.POST("user", handler.CreateUserHandler)
		api.GET("user", handler.ReadUserHandler)
		api.DELETE("user", handler.DeleteUserHandler)
		api.PUT("user", handler.UpdateUserHandler)
	}

}
