package routers

import (
	"github.com/gin-gonic/gin"
	"gapp/routers/api/v1/demo"
)

// ApiV1 接口v1版本路由
func ApiV1(engine *gin.Engine) {
	g := engine.Group("/api/v1")
	{
		g.POST("/demo/gorm/user", demo.UserAdd)
		g.GET("/demo/gorm/user", demo.UserQuery)
		g.GET("/demo/gorm/user/:userId", demo.UserDetail)
		g.DELETE("/demo/gorm/user/:userId", demo.UserDel)
		g.PUT("/demo/gorm/user/:userId", demo.UserUpdate)


		g.GET("/demo/simple/get", demo.Get)
		g.POST("/demo/simple/post", demo.Post)
		g.PUT("/demo/simple/put", demo.Put)
		g.DELETE("/demo/simple/delete", demo.Delete)
	}
}

// ApiV2 接口v2版本路由
func ApiV2(engine *gin.Engine) {
	//g := engine.Group("/api/v2")
	//{
	//	g.GET("/demo/json")
	//}
}

// ApiV3 接口v3版本路由
func ApiV3(engine *gin.Engine) {
	//g := engine.Group("/api/v3")
	//{
	//	g.GET("/demo/json")
	//}
}