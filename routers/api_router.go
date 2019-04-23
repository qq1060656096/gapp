package routers

import (
	"github.com/gin-gonic/gin"
	"gapp/routers/api/v1/demo"
)

// ApiV1 接口v1版本路由
func ApiV1(engine *gin.Engine) {
	g := engine.Group("/api/v1")
	{
		// grom 模型操作示例
		g.POST("/demo/gorm/user", demo.UserAdd)
		g.GET("/demo/gorm/user", demo.UserQuery)
		g.GET("/demo/gorm/user/:userId", demo.UserDetail)
		g.DELETE("/demo/gorm/user/:userId", demo.UserDel)
		g.PUT("/demo/gorm/user/:userId", demo.UserUpdate)

		// grom 原生sql操作示例
		g.POST("/demo/gorm-raw-sql/user", demo.DemoRawInsertSql)
		g.GET("/demo/gorm-raw-sql/user/first-detail", demo.DemoRawSelecOnetSql)
		g.GET("/demo/gorm-raw-sql/user", demo.DemoRawSelectAllSql)
		g.DELETE("/demo/gorm-raw-sql/user", demo.DemoRawDeleteSql)
		g.PUT("/demo/gorm-raw-sql/user", demo.DemoRawUpdateSql)

		// 简单 get post put delete请求示例
		g.GET("/demo/simple/get", demo.Get)
		g.POST("/demo/simple/post", demo.Post)
		g.PUT("/demo/simple/put", demo.Put)
		g.DELETE("/demo/simple/delete", demo.Delete)

		// 简单验证示例
		g.POST("/demo/validator/post", demo.Validator)
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