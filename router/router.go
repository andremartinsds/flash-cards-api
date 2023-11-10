package router

import "github.com/gin-gonic/gin"

func Initilize() {
	router := gin.Default()

	initilizeRoutes(router)

	//TODO: get por from .env
	router.Run(":8081")
}
