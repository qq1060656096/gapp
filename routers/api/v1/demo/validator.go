package demo

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-api"
	"log"
	"net/http"
)

type ValidatorRequest struct {
	Id int64 `json:"id" valid:"required~id必须是整数"`
	Name string `json:"name" valid:"required~名字不能为空,runelength(1|10)~名字在1-10个字符之间"`
}

// curl -X POST "http://localhost:8080/api/v1/demo/validator/post" -d "id=1&name=\"123456\""
/*
curl -X POST \
  http://localhost:8080/api/v1/demo/validator/post \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
  "id": 1,
  "name": "11"
}'

curl -X POST \
  http://localhost:8080/api/v1/demo/validator/post \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
  "id": "123",
  "name": "11"
}'
 */
func Validator(c *gin.Context) {
	var r ValidatorRequest
	c.ShouldBind(&r)
	log.Printf("failed to create rotatelogs: %#v", r)
	// 验证数据
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		c.JSON(http.StatusOK, api.NewResult("422", "govalidator fail").Simple(errMap))
		return
	}
	c.JSON(http.StatusOK, api.NewResult("200", "success"))
}