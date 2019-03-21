package routers

import "github.com/gin-gonic/gin"

// AppV1 应用v1版本路由
func AppV1(engine *gin.Engine) {
	g := engine.Group("/app/v1")
	{
		g.GET("/demo/json")
	}
}

// AppV2 应用v2版本路由
func AppV2(engine *gin.Engine) {
	g := engine.Group("/app/v2")
	{
		g.GET("/demo/json")
	}
}

// AppV3 应用v3版本路由
func AppV3(engine *gin.Engine) {
	g := engine.Group("/app/v3")
	{
		g.GET("/demo/json")
	}
}