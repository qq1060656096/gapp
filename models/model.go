package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // indirect
	"log"
	"os"
	"strconv"
)

const (
	DRIVER_MY_SQL = "mysql"
	DRIVER_POSTGRE_SQL = "postgres"
	DRIVER_SQLITE3 = "sqlite3"
	DRIVER_SQL_SERVER = "mssql"
)

var DB *gorm.DB


// ConnectDB 连接数据库
func ConnectDB() {
	driver := os.Getenv("DB_DRIVER")
	switch driver {
	case DRIVER_MY_SQL:
		DB = ConnectDbMySQL(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_CHARSET"),
		)
	case DRIVER_POSTGRE_SQL:
		DB = ConnectDbPostgreSQL(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
		)
	case DRIVER_SQLITE3:
		DB = ConnectDbSqlite3(
			os.Getenv("DB_HOST"),
		)
	case DRIVER_SQL_SERVER:
		DB = ConnectDbMySQL(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_CHARSET"),
		)
	default:
		log.Fatalf("models.ConnectDB driver err: %s", driver)
	}
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	openConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	DB.DB().SetMaxOpenConns(maxConnections)
	DB.DB().SetMaxIdleConns(openConnections)
	dbLog := os.Getenv("DB_LOG")
	if dbLog == "true" {
		DB.LogMode(true)
	}
}

// ConnectDbMySQL 初始化Mysql db
func ConnectDbMySQL(host, port, database, user, pass , charset string) *gorm.DB {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(DRIVER_MY_SQL, dns)
	if err != nil {
		log.Fatalf("models.InitDbMySQL err: %v", err)
	}
	db.SingularTable(true)
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	openConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	db.DB().SetMaxOpenConns(maxConnections)
	db.DB().SetMaxIdleConns(openConnections)
	return db
}

// ConnectDbPostgreSQL 连接dbpostgresql数据库
func ConnectDbPostgreSQL(host, port, database, user, pass string) *gorm.DB {
	dns := fmt.Sprintf(
		"host=%s port=myport user=%s dbname=%s password=%s",
		host,
		port,
		user,
		database,
		pass,
	)

	db, err := gorm.Open(DRIVER_POSTGRE_SQL, dns)
	if err != nil {
		log.Fatalf("models.ConnectDbPostgreSQL err: %v", err)
	}
	db.SingularTable(true)
	return db
}

// ConnectDbSqlite3 连接sqlite3数据库
func ConnectDbSqlite3(host string) *gorm.DB {
	dns := fmt.Sprintf(
		"%s",
		host,
	)

	db, err := gorm.Open(DRIVER_SQLITE3, dns)
	if err != nil {
		log.Fatalf("models.ConnectDbSqlite3 err: %v", err)
	}
	db.SingularTable(true)
	return db
}

// ConnectDbSqlServer 连接 sql server数据库
func ConnectDbSqlServer(host, port, database, user, pass string) *gorm.DB {
	dns := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		user,
		pass,
		host,
		port,
		database,
	)

	db, err := gorm.Open(DRIVER_SQL_SERVER, dns)
	if err != nil {
		log.Fatalf("models.ConnectDbSqlServer err: %v", err)
	}
	db.SingularTable(true)
	return db
}