package orm

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var ctx = context.Background()
var GormDB *gorm.DB

type MysqlConf struct {
	DataSource   string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifeTime  int
	Debug        bool
}

func NewGorm(c *MysqlConf) *gorm.DB {
	if c == nil {
		panic("config cannot be nil")
	}
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(fmt.Sprintf("db connnet err: %s", err.Error()))
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
		if err = sqlDB.Ping(); err != nil {
			log.Panicf("db ping err: %s", err.Error())
		}
	} else {
		log.Panicf("db connnet err: %v", err)
	}
	if c.Debug {
		db.Debug()
	}
	return db
}
