package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": data})
}

func Fail(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "message": message, "data": nil})
}
