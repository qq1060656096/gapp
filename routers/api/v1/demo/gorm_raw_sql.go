package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-api"
	"github.com/qq1060656096/go-common"
	"net/http"
)

// DemoRawInsertSql 原生insert sql 插入数据
//
// curl -X POST "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":1}%
func DemoRawInsertSql(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	sql := "insert demo_user(`user`, `pass`) values(?, ?);"
	rows := db.Exec(sql, "test.DemoRawInsertSql", "123456.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,api.NewResult("200", "fail").Simple(rows))
		return
	}
	c.JSON(http.StatusOK,api.NewResult("200", "success").Simple(rows))
}

// DemoRawDeleteSql 原生 delete sql 删除数据
//
// curl -X DELETE "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":1}
func DemoRawDeleteSql(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	sql := "delete from demo_user where `user` = ? limit 1;"
	rows := db.Exec(sql, "test.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,api.NewResult("200", "fail").Simple(rows))
		return
	}
	c.JSON(http.StatusOK,api.NewResult("200", "success").Simple(rows))
}

// DemoRawUpdateSql 原生 update sql 更新数据
// curl -X PUT "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":3}
func DemoRawUpdateSql(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	sql := "update demo_user set `user` = ? where `user` = ?"
	rows := db.Exec(sql, "test.DemoRawUpdateSql", "test.DemoRawInsertSql").RowsAffected
	if rows < 1 {
		c.JSON(http.StatusOK,api.NewResult("200", "fail").Simple(rows))
		return
	}
	c.JSON(http.StatusOK,api.NewResult("200", "success").Simple(rows))
}

// DemoRawSelecOnetSql 原生 select sql 查询单条数据
// curl -X GET "http://localhost:8080/api/v1/demo/gorm-raw-sql/user/first-detail"
//{"code":"200","message":"success","data":{"id":2,"user":"test","pass":"123456"}}
func DemoRawSelecOnetSql(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var gormDemoUser GormDemoUser
	sql := "select * from demo_user limit 1"
	db.Raw(sql).Scan(&gormDemoUser)
	db.Raw(sql).Scan(&gormDemoUser)
	c.JSON(http.StatusOK,api.NewResult("200", "success").Simple(gormDemoUser))
}

// DemoRawSelectAllSql 原生 select sql 查询多条数据
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm-raw-sql/user"
// {"code":"200","message":"success","data":[{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""},{"id":0,"user":"","pass":""}]}
func DemoRawSelectAllSql(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var ListsGormDemoUser[] *GormDemoUser
	var gormDemoUser GormDemoUser

	sql := "select * from demo_user limit 10"
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		c.JSON(http.StatusOK,api.NewResult("400", "fail").Simple(err))
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&gormDemoUser)
		ListsGormDemoUser = append(ListsGormDemoUser, &gormDemoUser)
	}
	c.JSON(http.StatusOK,api.NewResult("200", "success").Simple(ListsGormDemoUser))
}