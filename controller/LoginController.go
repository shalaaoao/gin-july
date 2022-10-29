package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	
	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}
