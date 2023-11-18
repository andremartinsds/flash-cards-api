package router

import (
	"github.com/andremartinsds/flash-cards-api/handler"
	hUser "github.com/andremartinsds/flash-cards-api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initilizeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	uriPath := "/api/"

	router.Use(cors.Default())
	api := router.Group(uriPath)
	{

		api.POST("user", hUser.CreateUserHandler)
		api.GET("user", hUser.ReadUserHandler)
		api.GET("users", hUser.ListUserHandler)
		api.DELETE("user", hUser.DeleteUserHandler)
		api.PUT("user", hUser.UpdateUserHandler)
	}

}
