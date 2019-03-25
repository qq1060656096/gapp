package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Get 简单请求显示不同主题的模板
// http://localhost:8080/app/v1/demo/simple-html?theme=default_mobile
func Get(c *gin.Context) {
	data := gin.H{
		"title": "gapp演示",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "GET请求演示",
	}
	theme := c.Query("theme")
	if theme == "" {
		theme = "default"
	}
	c.HTML(http.StatusOK, theme + "/simple/get.tmpl", data)
}