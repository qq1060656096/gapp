package main

import (
	"gapp/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/qq1060656096/go-common"
	"strings"
	"time"
)

func main() {
	godotenv.Load(".app.env", ".db.env", ".redis.env")
	logDir := common.OsEnvManager.Get("LOG_DIR")
	logDir = strings.TrimRight(logDir, "/")
	// 设置分割日志
	logWriterFormat := logDir + "/gapp.%Y%m%d.log"
	logWriter, err := rotatelogs.New(
		logWriterFormat,
		rotatelogs.WithLinkName(logDir + "/gapp.access.log"),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		logrus.Errorf("gapp.rotatelogs.New.failed: %s", err)
		return
	}
	// 初始化数据库和redis连接
	common.DbConnInit()
	common.RedisConnInit()

	logrus.SetOutput(logWriter)
	gin.DefaultWriter = logWriter
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/themes/***/***/*")
	// 初始化路由
	routers.InitRouter(engine)
	addr := common.OsEnvManager.Get("ADDR")
	engine.Run(addr)
}
