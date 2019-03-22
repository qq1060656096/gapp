package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-gin-app/routers/api/v1/demo"
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


		//g.GET("/demo/simple/user", demo.UserQuery)
		g.POST("/demo/simple/post", demo.SimplePost)
		//g.PUT("/demo/simple/user/:userId", demo.UserUpdate)
		//g.DELETE("/demo/simple/user/:userId", demo.UserDel)
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