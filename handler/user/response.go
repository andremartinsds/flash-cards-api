package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {

	ctx.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-Type", "Application/json")

	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}
