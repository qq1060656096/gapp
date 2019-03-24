package demo

import (
	"gapp/models"
	"gapp/pkg/util"
	"github.com/gin-gonic/gin"
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
	var gormDemoUser GormDemoUser
	c.ShouldBind(&gormDemoUser)
	// 保存数据到数据库
	models.DB.Create(&gormDemoUser)
	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", gormDemoUser))
}

// UserDel删除记录
//
// curl -X DELETE "http://localhost:8080/api/v1/demo/gorm/user/1"
// {"code":"200","message":"success","data":1}%
func UserDel(c *gin.Context) {
	var gormDemoUser GormDemoUser
	id := c.Param("userId")
	models.DB.Where("id = ?", id).First(&gormDemoUser)
	if gormDemoUser.ID < 1{
		c.JSON(http.StatusOK, util.GetApiJsonResult("404", "not found", nil))
		return
	}
	gormDemoUser.ID, _ = strconv.Atoi(id)
	isDel := models.DB.Delete(gormDemoUser).RowsAffected
	if isDel < 1 {
		c.JSON(http.StatusOK, util.GetApiJsonResult("200", "fail", isDel))
	}
	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", isDel))
}

// UserUpdate 更新记录
//
// curl -X PUT "http://localhost:8080/api/v1/demo/gorm/user/1" -d "user=test.update&pass=123456.update"
// {"code":"200","message":"success","data":{"id":1,"user":"test.update","pass":"123456.update"}}
func UserUpdate(c *gin.Context) {
	var gormDemoUser, updateGormDemoUser GormDemoUser
	id := c.Param("userId")
	models.DB.Where("id = ?", id).First(&gormDemoUser)
	if gormDemoUser.ID < 1{
		c.JSON(http.StatusOK, util.GetApiJsonResult("404", "not found", gormDemoUser))
		return
	}
	c.ShouldBind(&updateGormDemoUser)
	models.DB.Model(&gormDemoUser).Update(updateGormDemoUser)
	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", gormDemoUser))
}

// UserQuery 分页查询数据
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm/user?page=1&pageSize=2"
// {"code":"200","message":"success","data":{"totalCount":3,"totalPageCount":1,"page":1,"pageSize":2,"lists":[{"id":1,"user":"test","pass":"123456"},{"id":2,"user":"test","pass":"123456"}]}}
func UserQuery(c *gin.Context) {
	var listGormDomeUser[] *GormDemoUser
	count :=0
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	offset := (page-1)*pageSize
	models.DB.Model(&GormDemoUser{}).Where("id > 0").Count(&count)
	models.DB.Offset(offset).Limit(pageSize).Find(&listGormDomeUser)
	c.JSON(http.StatusOK, util.GetApiJsonPagingResult("200", "success", listGormDomeUser, count, page, pageSize))
}

// UserDetail gormc查询单挑记录
//
// curl -X GET "http://localhost:8080/api/v1/demo/gorm/user/1"
// {"code":"200","message":"success","data":{"id":1,"user":"test","pass":"123456"}}
func UserDetail(c *gin.Context) {
	var gormDemoUser GormDemoUser
	id := c.Param("userId")
	models.DB.Where("id = ?", id).First(&gormDemoUser)
	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", gormDemoUser))
}