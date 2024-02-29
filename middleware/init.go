package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判斷用戶是否登入
	fmt.Println(time.Now().UnixMicro())
	c.Set("username", "Dale")
}
