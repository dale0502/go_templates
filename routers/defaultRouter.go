package routers

import (
	"demo_gin/controllers/web"
	"demo_gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouter := r.Group("/")
	defaultRouter.Use(middleware.InitMiddleware)
	{
		defaultRouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "todolist.html", gin.H{
				"title": "todolist表單",
			})
		})

		defaultRouter.GET("/user", web.WebController{}.Index)

		defaultRouter.POST("/user/add", web.WebController{}.Add)
	}
}
