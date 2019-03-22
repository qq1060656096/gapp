package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"strconv"
)

var Db gorm.DB;
// InitModel 初始化模型
func InitModel() {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"),
	)
	Db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatalf("models.InitModel err: %v", err)
	}
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	Db.DB().SetMaxOpenConns(maxConnections)
}

