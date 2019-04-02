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
	"test2": {
		Uid: 1,
		AccountName: "15400000010",
		AccountType: ACCOUNT_TYPE_MOBILE,
	},
}

// TestUserBindAccount 测试绑定账户
func TestUserBindAccount(t *testing.T) {
	ua := testStruct["test1"]
	UserUnbindAccount(&ua)
	b := UserBindAccount(&ua)
	assert.Equal(t, true, b, "sass.TestGetUserBindAccount test fail")

	ua = testStruct["test2"]
	UserUnbindAccount(&ua)
	UserBindAccount(&ua)
}

// TestGetUserBindAccount 测试获取绑定账户
func TestGetUserBindAccount(t *testing.T) {
	ua := testStruct["test1"]
	fua := GetUserBindAccount(&ua)
	assert.Equal(t, ua.Uid, fua.Uid, "sass.GetUserBindAccount test fail")
	assert.Equal(t, "test1@qq.com", fua.AccountName, "sass.GetUserBindAccount test fail")
}


// TestGetUserBindAllAccounts 测试用户绑定了那些账户
func TestGetUserBindAllAccounts(t *testing.T) {
	uaList := GetUserBindAllAccounts(1)
	assert.Equal(t, 1, uaList[0].AccountType)
	assert.Equal(t, "15400000010", uaList[0].AccountName)
	assert.Equal(t, "test1@qq.com", uaList[1].AccountName)

}