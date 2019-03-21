package routers

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) *gin.Engine {
	// api 接口路由配置
	ApiV1(engine)
	ApiV2(engine)
	ApiV3(engine)

	// app 应用路由配置
	AppV1(engine)
	AppV2(engine)
	AppV3(engine)
	return engine
}



