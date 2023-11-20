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

type FlashCardCreateResponseData struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardCreateResponse struct {
	Message string
	Data    FlashCardListResponseData
}

type FlashCardReadResponseData struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardReadResponse struct {
	Message string
	Data    FlashCardReadResponseData
}

type FlashCardListResponseData struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardListResponse struct {
	Message string
	Data    FlashCardListResponseData
}

type FlashCardUpdateResponseData struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardUpdateResponse struct {
	Message string
	Data    FlashCardUpdateResponseData
}

type FlashCardDeleteResponseData struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardDeleteResponse struct {
	Message string
	Data    FlashCardDeleteResponseData
}

type FlashCardErrorResponse struct {
	Message string
	Code    uint
}
