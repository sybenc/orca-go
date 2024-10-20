package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orca/conf"
	"sync"
)

// Mysql Global database connection.
var Mysql *gorm.DB
var once sync.Once

func InitMysql() {
	once.Do(func() {
		dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
			conf.GetString("database.user"),
			conf.GetString("database.password"),
			conf.GetString("database.host"),
			conf.GetString("database.port"),
			conf.GetString("database.dbname"),
			true,
			"Local")

		var err error
		Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Failed to connect mysql database")
		}
	})
}
