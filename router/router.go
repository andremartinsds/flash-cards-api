package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initilize() {
	router := gin.Default()

	initilizeRoutes(router)

	port := ":" + os.Getenv("APPLICATION_PORT")
	router.Run(port)
}
