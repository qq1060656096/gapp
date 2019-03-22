package main

import (
	"github.com/gin-gonic/gin"
	"gin/routers"
)

func main() {
	engine := gin.New()
	// 初始化路由
	routers.InitRouter(engine)
	// 初始化模型
	//models.InitModel()

	engine.Run(":8080")
}
