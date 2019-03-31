package main

import (
	"gapp/models"
	"gapp/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lestrrat/go-file-rotatelogs"
	"log"
	"time"
)

func main() {
	godotenv.Load(".env")

	// 设置分割日志
	//logWriterFormat := "var/log/gin.%Y%m%d.%H%M.log"
	logWriterFormat := "var/log/gin.%Y%m%d.log"
	logWriter, err := rotatelogs.New(
		logWriterFormat,
		rotatelogs.WithLinkName("var/log/access_log.log"),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}


	gin.DefaultWriter = logWriter
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/themes/***/***/*")
	// 初始化路由
	routers.InitRouter(engine)
	// 初始化模型
	models.ConnectDB()
	engine.Run(":8080")
}
