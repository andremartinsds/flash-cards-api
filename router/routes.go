package router

import (
	"github.com/andremartinsds/flash-cards-api/docs"
	"github.com/andremartinsds/flash-cards-api/handler"
	hDec "github.com/andremartinsds/flash-cards-api/handler/dec"
	hFlashCard "github.com/andremartinsds/flash-cards-api/handler/flash-cards"
	hUser "github.com/andremartinsds/flash-cards-api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	uriPath := "/api/"

	docs.SwaggerInfo.BasePath = uriPath

	router.Use(cors.Default())

	api := router.Group(uriPath)
	{
		api.POST("user", hUser.CreateUserHandler)
		api.GET("user", hUser.ReadUserHandler)
		api.GET("users", hUser.ListUserHandler)
		api.DELETE("user", hUser.DeleteUserHandler)
		api.PUT("user", hUser.UpdateUserHandler)

		api.POST("dec", hDec.CreateDecHandler)
		api.GET("dec", hDec.ReadDecHandler)
		api.GET("decs", hDec.ListDecHandler)
		api.DELETE("dec", hDec.DeleteDecHandler)
		api.PUT("dec", hDec.UpdateDecHandler)

		api.POST("flash-card", hFlashCard.CreateFlashCardHandler)
		api.GET("flash-card", hFlashCard.ReadFlashCardHandler)
		api.GET("flash-cards", hFlashCard.ListFlashCardHandler)
		api.DELETE("flash-card", hFlashCard.DeleteFlashCardHandler)
		api.PUT("flash-card", hFlashCard.UpdateFlashCardHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
