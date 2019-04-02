package sass

import (
	"gapp/models"
	"time"
)

const (
	ACCOUNT_TYPE = iota
	// 手机号码
	ACCOUNT_TYPE_MOBILE
	// 邮箱
	ACCOUNT_TYPE_EMAIL
	// 微信
	ACCOUNT_TYPE_WECHAT
)


// UsersAccountsModel 用户账户关联表
type UsersAccountsModel struct {
	models.BaseModel
	Uid uint `json:"uid" xml:"uid"`
	AccountName string `form:"account_name" json:"account_name" xml:"account_name" binding:"required"`
	AccountType int `form:"account_type" json:"account_type" xml:"account_type" binding:"required"`
}

// TableName 表名
func (ua *UsersAccountsModel) TableName() string {
	return "users_accounts"
}

// UserBindAccount 定用户绑账户
func UserBindAccount(ua *UsersAccountsModel) (bool) {
	rows := models.DB.Create(ua).RowsAffected
	if (rows < 1) {
		return false
	}
	return true
}

// UserUnbindAccount 用户解绑账户
func UserUnbindAccount(ua *UsersAccountsModel) bool {
	uaNew := UsersAccountsModel{
		Uid: ua.Uid,
		AccountName: ua.AccountName,
		AccountType: ua.AccountType,
	}
	uaNew.DeletedAt = time.Now().Unix()
	rows := models.DB.Table(ua.TableName()).
		Where("uid = ?", ua.Uid).
		Where("account_name = ?", ua.AccountName).
		Where("account_type = ?", ua.AccountType).
		Where("deleted_at = ?", 0).
		Updates(uaNew).RowsAffected
	if (rows < 1) {
		return false
	}
	return true
}

// GetUserBindAccount 获取用户绑定指定类型账户
func GetUserBindAccount(ua *UsersAccountsModel) *UsersAccountsModel {
	uaNew := &UsersAccountsModel{
		Uid: 0,
		AccountName: "",
	}
	models.DB.Table(ua.TableName()).
		Where("uid = ?", ua.Uid).
		Where("account_name = ?", ua.AccountName).
		Where("account_type = ?", ua.AccountType).
		Where("deleted_at = ?", 0).
		Unscoped().
		First(&uaNew)
	return uaNew
}

// GetUserBindAllAccounts 获取用户绑定的账户列表
func GetUserBindAllAccounts(uid uint) []*UsersAccountsModel {
	ua := []*UsersAccountsModel{}
	models.DB.Where("uid = ?", uid).
		Where("deleted_at = ?", 0).
		Unscoped().
		Find(&ua)
	return ua
}