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

type UserCreateResponseData struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserCreateResponse struct {
	Message string
	Data    UserCreateResponseData
}

type UserReadResponseData struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserReadResponse struct {
	Message string
	Data    UserReadResponseData
}

type UserUpdateResponseData struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserUpdateResponse struct {
	Message string
	Data    UserUpdateResponseData
}

type UserListResponseData struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserListResponse struct {
	Message string
	Data    []UserListResponseData
}

type UserDeleteResponseData struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}
type UserDeleteResponse struct {
	Message string
	Data    UserDeleteResponseData
}

type UserErrorResponse struct {
	Message string
	Code    uint
}
