package controller

import (
	"gin-july/response"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, gin.H{"message": "hello world~"})
	response.Success(ctx, nil)
}
