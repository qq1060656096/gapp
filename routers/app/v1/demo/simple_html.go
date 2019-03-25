package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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