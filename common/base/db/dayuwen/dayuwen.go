package dayuwen

import (
	"fmt"
	"time"

	// mysql import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DB dayuwen db
	DB *gorm.DB
)

// Init dayuwen db
func Init(user, password, addr, db string) {
	var err error

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", user, password, "tcp", addr, db)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	DB.DB().SetMaxOpenConns(150)
	DB.DB().SetMaxIdleConns(15)
	DB.SingularTable(true)
}
