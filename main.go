package main

import (
	"demo_gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//配置模板文件
	r.LoadHTMLGlob("templates/*")

	routers.DefaultRouterInit(r)
	routers.ApiRouterInit(r)

	r.Run(":8080") // 在 0.0.0.0:8080 上监听
}
