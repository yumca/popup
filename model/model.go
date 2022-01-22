package model

import (
	"database/sql"
	"log"
	"popup/library"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func NewConnection() *gorm.DB {
	conf, err := library.GetConf()
	if err != nil {
		log.Fatal("mainServer GetConfig Error:", err)
	}
	var dialector gorm.Dialector
	var config gorm.Config
	switch conf.Db.DbType {
	case "mysql":
		dialector = mysql.Open(conf.Db.DbUser + ":" + conf.Db.DbPwd + "@(" + conf.Db.DbHost + ":" + strconv.Itoa(conf.Db.DbPort) + ")/" + conf.Db.Database + "?charset=utf8mb4&parseTime=True&loc=Local")
		config = gorm.Config{}
	case "sqlite":
		dialector = sqlite.Open(library.ProgramDir() + conf.Db.Database + ".db?charset=utf8mb4&parseTime=True&loc=Local")
		config = gorm.Config{}
	default:
	}
	conn, err := gorm.Open(dialector, &config)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ = conn.DB()
	sqlDB.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDB.SetMaxOpenConns(30)                   //最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

	return conn
}

func GetDb() *gorm.DB {
	if db == nil {
		db = NewConnection()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		db = NewConnection()
	}
	return db
	// defer Db.Close()
}
