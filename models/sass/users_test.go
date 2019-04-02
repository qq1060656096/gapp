package sass

import (
	"fmt"
	"gapp/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	dir, _ := os.Getwd()
	dir = (filepath.Dir(filepath.Dir(dir)))
	testEnv := dir + "/.env"
	fmt.Println(testEnv)
	godotenv.Load(testEnv)
	models.ConnectDB()
}

// TestUserCreate 测试用户创建
func TestUserCreate(t *testing.T) {
	var u *UsersModel
	u = &UsersModel{
		Pass: "123456",
	}
	b := UserCreate(u)
	assert.Equal(t, true, b, "sass.UserCreate Test fail")
	assert.Equal(t, true, u.Uid > 0, "sass.UserCreate Test fail")
}