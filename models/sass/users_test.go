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

var userTestStruct = map[string]UsersAccountsModel{
	"test2User": {
		Uid: 2,
		AccountName: "test2@qq.com",
		AccountType: ACCOUNT_TYPE_EMAIL,
	},
	"test2UserAccount": {
		Uid: 2,
		AccountName: "15400000012",
		AccountType: ACCOUNT_TYPE_MOBILE,
	},
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

func TestUsersRegister(t *testing.T) {
	u := &UsersModel{
		Uid: 2,
		Pass: "p123456",
	}
	ua := &UsersAccountsModel{
		Uid: 2,
		AccountName: "test2@qq.com",
		AccountType: ACCOUNT_TYPE_EMAIL,
	}
	models.DB.Table(u.TableName()).Where("uid = 2").Unscoped().Delete(&UsersModel{})
	models.DB.Table(ua.TableName()).Where("uid = 2").Unscoped().Delete(&UsersAccountsModel{})
	b := UserRegister(u, ua)
	assert.Equal(t, true, b, "sass.TestUsersRegister test fail")
}