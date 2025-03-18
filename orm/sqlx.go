package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var SqlxDB *sqlx.DB

type SqlxConf struct {
	DataSource   string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifeTime  int
}

func NewSqlx(c *SqlxConf) *sqlx.DB {
	if c == nil {
		panic("config cannot be nil")
	}

	db, err := sqlx.Connect("mysql", c.DataSource)
	if err != nil {
		panic(fmt.Sprintf("db connnet err: %s", err.Error()))
	}
	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("db ping err: %s", err.Error()))
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
	fmt.Println("db connect success......")
	return db
}
