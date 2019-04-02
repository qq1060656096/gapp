package sass

import "gapp/models"

// UsersModel 用户模型
type UsersModel struct {
	models.BaseModel
	Uid uint `gorm:"primary_key" json:"uid"`
	Pass string `json:"pass" form:"pass" xml:"pass" binding:"required"`
}

// TableName 表名
func (u *UsersModel) TableName() string {
	return "users"
}

// UsersCreate 创建用户
func UserCreate(u *UsersModel) (bool) {
	rows := models.DB.Create(u).RowsAffected
	if (rows < 1) {
		return false
	}
	return true
}

// 用户注册
func UserRegister(u *UsersModel, ua *UsersAccountsModel) bool {
	// 1. 创建账户
	// 2. 绑定用户账户
	b := UserCreate(u)
	if !b {
		return b
	}
	return UserBindAccount(ua)
}