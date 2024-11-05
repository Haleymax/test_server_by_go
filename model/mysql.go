package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASSWORD  = "root"
	HOST      = "127.0.0.1"
	PORT      = "3306"
	DATABASE  = "webglhost"
	CHARSET   = "utf8"
)

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)

	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		log.Println("dbDSN:" + dbDSN)
		panic("数据源配置不正确 : " + MysqlDbErr.Error())
	}
	//最大连接数
	MysqlDb.SetMaxOpenConns(100)

	//闲置连接数
	MysqlDb.SetMaxIdleConns(20)

	//最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)
}
