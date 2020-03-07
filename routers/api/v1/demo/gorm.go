package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-common"
	"github.com/qq1060656096/go-api"
	"net/http"
	"strconv"
)

type GormDemoUser struct {
	ID int `json:"id"`
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Pass string `form:"pass" json:"pass" xml:"user" binding:"required"`
}




func (g GormDemoUser) TableName() string {
	return "demo_user"
}

// UserAdd gorm新增操作
//
// curl -X POST "http://localhost:8080/api/v1/demo/gorm/user" -d "user=test&pass=123456"
// {"code":"200","message":"success","data":{"id":5,"user":"test","pass":"123456"}}
func UserAdd(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var gormDemoUser GormDemoUser
	c.ShouldBind(&gormDemoUser)
	// 保存数据到数据库
	db.Create(&gormDemoUser)
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(gormDemoUser))
}

// UserDel删除记录
//
// curl -X DELETE "http://localhost:8080/api/v1/demo/gorm/user/1"
// {"code":"200","message":"success","data":1}%
func UserDel(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var gormDemoUser GormDemoUser
	id := c.Param("userId")
	db.Where("id = ?", id).First(&gormDemoUser)
	if gormDemoUser.ID < 1{
		c.JSON(http.StatusOK, api.NewResult("404", "not found").Simple(nil))
		return
	}
	gormDemoUser.ID, _ = strconv.Atoi(id)
	isDel := db.Delete(gormDemoUser).RowsAffected
	if isDel < 1 {
		c.JSON(http.StatusOK, api.NewResult("200", "fail").Simple(isDel))
	}
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(isDel))
}

// UserUpdate 更新记录
//
// curl -X PUT "http://localhost:8080/api/v1/demo/gorm/user/1" -d "user=test.update&pass=123456.update"
// {"code":"200","message":"success","data":{"id":1,"user":"test.update","pass":"123456.update"}}
func UserUpdate(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var gormDemoUser, updateGormDemoUser GormDemoUser
	id := c.Param("userId")
	db.Where("id = ?", id).First(&gormDemoUser)
	if gormDemoUser.ID < 1{
		c.JSON(http.StatusOK, api.NewResult("404", "not found").Simple(gormDemoUser))
		return
	}
	c.ShouldBind(&updateGormDemoUser)
	db.Model(&gormDemoUser).Update(updateGormDemoUser)
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(gormDemoUser))
}

// UserQuery 分页查询数据
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm/user?page=1&pageSize=2"
// {"code":"200","message":"success","data":{"totalCount":3,"totalPageCount":1,"page":1,"pageSize":2,"lists":[{"id":1,"user":"test","pass":"123456"},{"id":2,"user":"test","pass":"123456"}]}}
func UserQuery(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var listGormDomeUser[] *GormDemoUser
	count :=0
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	offset := (page-1)*pageSize
	db.Model(&GormDemoUser{}).Where("id > 0").Count(&count)
	db.Offset(offset).Limit(pageSize).Find(&listGormDomeUser)
	c.JSON(http.StatusOK, api.NewResult("200", "success").Paging(listGormDomeUser, count, page, pageSize))
}

// UserDetail gormc查询单挑记录
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm/user/1"
// {"code":"200","message":"success","data":{"id":1,"user":"test","pass":"123456"}}
func UserDetail(c *gin.Context) {
	db, _ := common.GetDefaultDbConn()
	var gormDemoUser GormDemoUser
	id := c.Param("userId")
	db.Where("id = ?", id).First(&gormDemoUser)
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(gormDemoUser))
}