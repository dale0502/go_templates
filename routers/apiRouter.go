package routers

import "github.com/gin-gonic/gin"

func ApiRouterInit(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api路由",
			})
		})
	}
}
