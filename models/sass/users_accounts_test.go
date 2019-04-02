package sass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStruct = map[string]UsersAccountsModel{
	"test1": {
		Uid: 1,
		AccountName: "test1@qq.com",
		AccountType: ACCOUNT_TYPE_EMAIL,
	},
}

// TestUserBindAccount 测试绑定账户
func TestUserBindAccount(t *testing.T) {
	ua := testStruct["test1"]
	UserUnbindAccount(&ua)
	b := UserBindAccount(&ua)
	assert.Equal(t, true, b, "sass.TestGetUserBindAccount test fail")


}

// TestGetUserBindAccount 测试获取绑定账户
func TestGetUserBindAccount(t *testing.T) {
	ua := testStruct["test1"]
	fua := GetUserBindAccount(&ua)
	assert.Equal(t, ua.Uid, fua.Uid, "sass.GetUserBindAccount test fail")
	assert.Equal(t, "test1@qq.com", fua.AccountName, "sass.GetUserBindAccount test fail")
}

