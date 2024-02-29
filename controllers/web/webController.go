package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	BaseController
}

func (con WebController) Index(c *gin.Context) {
	userName, _ := c.Get("username")
	con.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "这是用户列表",
		"user":    userName,
	})
}

func (con WebController) Add(c *gin.Context) {
	con.MessageResponse(c, "新增用户成功")
}

func (con WebController) Edit(c *gin.Context) {
	con.SuccessResponse(c, http.StatusAccepted, gin.H{"message": "用户信息已更新"})
}
