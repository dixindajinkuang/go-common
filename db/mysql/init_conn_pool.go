package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnPool(dsn string, maxOpen, maxIdle int) (__DB *sql.DB) {
	var err error
	__DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	__DB.SetMaxOpenConns(maxOpen)
	__DB.SetMaxIdleConns(maxIdle)
	__DB.SetConnMaxLifetime(time.Hour * 1)
	if err = __DB.Ping(); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if __DB != nil {
				_ = __DB.Close()
			}
			panic(err)
		}
	}()
	return
}
