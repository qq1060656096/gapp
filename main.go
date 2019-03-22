package main

import (
	"gapp/models"
	"gapp/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	engine := gin.New()
	// 初始化路由
	routers.InitRouter(engine)
	// 初始化模型
	models.InitModel()

	engine.Run(":8080")
}
