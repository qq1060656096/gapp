package demo

import (
	"gapp/models"
	"gapp/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DemoRawInsertSql 原生insert sql 插入数据
//
// curl -X POST "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":1}%
func DemoRawInsertSql(c *gin.Context) {
	sql := "insert demo_user(`user`, `pass`) values(?, ?);"
	rows := models.DB.Exec(sql, "test.DemoRawInsertSql", "123456.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,util.GetApiJsonResult("200", "fail", rows))
		return
	}
	c.JSON(http.StatusOK,util.GetApiJsonResult("200", "success", rows))
}

// DemoRawDeleteSql 原生 delete sql 删除数据
//
// curl -X DELETE "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":1}
func DemoRawDeleteSql(c *gin.Context) {
	sql := "delete from demo_user where `user` = ? limit 1;"
	rows := models.DB.Exec(sql, "test.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,util.GetApiJsonResult("200", "fail", rows))
		return
	}
	c.JSON(http.StatusOK,util.GetApiJsonResult("200", "success", rows))
}

// DemoRawUpdateSql 原生 update sql 更新数据
// curl -X PUT "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":3}
func DemoRawUpdateSql(c *gin.Context) {
	sql := "update demo_user set `user` = ? where `user` = ?"
	rows := models.DB.Exec(sql, "test.DemoRawUpdateSql", "test.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,util.GetApiJsonResult("200", "fail", rows))
		return
	}
	c.JSON(http.StatusOK,util.GetApiJsonResult("200", "success", rows))
}

// DemoRawSelecOnetSql 原生 select sql 查询单条数据
// curl -X GET "http://localhost:8080/api/v1/demo/gorm-raw-sql/user/first-detail"
//{"code":"200","message":"success","data":{"id":2,"user":"test","pass":"123456"}}
func DemoRawSelecOnetSql(c *gin.Context) {
	var gormDemoUser GormDemoUser
	sql := "select * from demo_user limit 1"
	models.DB.Raw(sql).Scan(&gormDemoUser)
	models.DB.Raw(sql).Scan(&gormDemoUser)
	c.JSON(http.StatusOK,util.GetApiJsonResult("200", "success", gormDemoUser))
}

// DemoRawSelectAllSql 原生 select sql 查询多条数据
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":[{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""}]}
func DemoRawSelectAllSql(c *gin.Context) {
	var ListsGormDemoUser[] *GormDemoUser
	var gormDemoUser GormDemoUser

	sql := "select * from demo_user limit 10"
	rows, err := models.DB.Raw(sql).Rows()
	if err != nil {
		c.JSON(http.StatusOK,util.GetApiJsonResult("400", "fail", err))
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&gormDemoUser)
		ListsGormDemoUser = append(ListsGormDemoUser, &gormDemoUser)
	}
	c.JSON(http.StatusOK,util.GetApiJsonResult("200", "success", ListsGormDemoUser))
}