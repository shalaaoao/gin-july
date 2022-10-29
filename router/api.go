package router

import (
	"gin-july/controller"
	"gin-july/controller/star"
	"github.com/gin-gonic/gin"
)

func CollectApiRoute(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", controller.Hello)

	starGroup := r.Group("/star")
	{
		starGroup.POST("/create", star.CreateLog)
		starGroup.GET("/lists", star.ListLog)
		starGroup.DELETE("/delete/:id", star.DeleteLog)
		starGroup.PUT("/edit/:id", star.EditLog)
	}

	return r
}
