package models

import (
	"time"
)

// BaseModel 基模型定义
type BaseModel struct {
	CreatedAt int64 `gorm:"type:int;"`
	UpdatedAt int64 `gorm:"type:int;"`
	DeletedAt int64 `gorm:"type:int;"`

}

func (bm *BaseModel) BeforeSave() (err error) {
	if bm.CreatedAt < 1 {
		bm.CreatedAt = time.Now().Unix()
	}
	return
}

func (bm *BaseModel) BeforeUpdate() (err error) {
	bm.UpdatedAt = time.Now().Unix()
	return
}

func (bm *BaseModel) BeforeDelete33() (err error) {
	bm.DeletedAt = time.Now().Unix()
	return
}

func (bm *BaseModel) AfterFind() (err error) {

	return
}