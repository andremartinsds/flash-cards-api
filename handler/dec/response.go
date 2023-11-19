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

type CreateDecResponseData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type CreateDecResponse struct {
	Message string
	Data    CreateDecResponseData
}

type DecReadResponseData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type ReadDecResponse struct {
	Message string
	Data    DecReadResponseData
}

type ListDecResponseData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type ListDecResponse struct {
	Message string
	Data    []ListDecResponseData
}

type DecUpdateResponseData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type UpdateDecResponse struct {
	Message string
	Data    DecUpdateResponseData
}

type DeleteDecResponseData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type DeleteDecResponse struct {
	Message string
	Data    DeleteDecResponseData
}

type ErrorDecResponse struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}
