package routes

import (
	"concurrency-server/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		queue := v1.Group("/queue")
		{
			queue.POST("/get", handler.QueueHandler.GetQueue)
			queue.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "test successful",
				})
			})
		}
	}
}
