package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SimplePost(c *gin.Context) {
	c.String(http.StatusOK, "ddd")
}