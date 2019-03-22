package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get get请求示例
//
// curl -X GET "http://localhost:8080/api/v1/demo/simple/get?user=test&pass=123456"
//
// url: "/api/v1/demo/simple/get?user=test&pass=123456"
// demoUser: demo.DemoUser{User:"test", Pass:"123456"}
// query.user: "test"
// query.pass: "123456"
func Get(c *gin.Context) {
	var demoUser DemoUser
	c.ShouldBindQuery(&demoUser)
	str := `
url: %#v
demoUser: %#v
query.user: %#v
query.pass: %#v

`
	str = fmt.Sprintf(
		str,
		c.Request.RequestURI,
		demoUser,
		c.Query("user"),
		c.Query("pass"),
	)
	c.String(http.StatusOK, str)
}

// Post post请求示例
//
// curl -X POST "http://localhost:8080/api/v1/demo/simple/post" -d "user=test&pass=123456"
//
// url: "/api/v1/demo/simple/post"
// demoUser: demo.DemoUser{User:"test", Pass:"123456"}
// formData.user: "test"
// formData.pass: "123456"
func Post(c *gin.Context) {
	var demoUser DemoUser
	c.ShouldBind(&demoUser)
	str := `
url: %#v
demoUser: %#v
formData.user: %#v
formData.pass: %#v

`
	str = fmt.Sprintf(
		str,
		c.Request.RequestURI,
		demoUser,
		c.PostForm("user"),
		c.PostForm("pass"),
	)
	c.String(http.StatusOK, str)
}

// Put put请求示例
//
// curl -X PUT "http://localhost:8080/api/v1/demo/simple/put" -d "user=test&pass=123456"
// url: "/api/v1/demo/simple/put"
// demoUser: demo.DemoUser{User:"test", Pass:"123456"}
// formData.user: "test"
// formData.pass: "123456"
func Put(c *gin.Context) {
	var demoUser DemoUser
	c.ShouldBind(&demoUser)
	str := `
url: %#v
demoUser: %#v
formData.user: %#v
formData.pass: %#v

`
	str = fmt.Sprintf(
		str,
		c.Request.RequestURI,
		demoUser,
		c.PostForm("user"),
		c.PostForm("pass"),
	)
	c.String(http.StatusOK, str)
}

// Delete delete请求示例
//
// curl -X DELETE "http://localhost:8080/api/v1/demo/simple/delete?user=test&pass=123456"
//
// url: "/api/v1/demo/simple/delete?user=test&pass=123456"
//  demoUser: demo.DemoUser{User:"test", Pass:"123456"}
// formData.user: "test"
// formData.pass: "123456"
func Delete(c *gin.Context) {
	var demoUser DemoUser
	c.ShouldBind(&demoUser)
	str := `
url: %#v
demoUser: %#v
formData.user: %#v
formData.pass: %#v

`
	str = fmt.Sprintf(
		str,
		c.Request.RequestURI,
		demoUser,
		c.Query("user"),
		c.Query("pass"),
	)
	c.String(http.StatusOK, str)
}
