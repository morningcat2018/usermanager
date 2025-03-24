package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponseDTO struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

type JsonErrorDTO struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func Success(c *gin.Context, data interface{}, count int64) {
	json := &JsonResponseDTO{Code: 0, Message: "success", Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func Fail(c *gin.Context, msg interface{}) {
	json := &JsonErrorDTO{Code: -1, Message: msg}
	c.JSON(http.StatusOK, json)
}

func ErrorRecover(c *gin.Context) {
	if r := recover(); r != nil {
		Fail(c, r.(error).Error())
	}
}
