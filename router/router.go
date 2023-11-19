package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	port := ":" + os.Getenv("APPLICATION_PORT")
	err := router.Run(port)
	if err != nil {
		//TODO: add logger
		return
	}
}
