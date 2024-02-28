package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "todolist.html", gin.H{
				"title": "todolist表單",
			})
		})
	}
}
