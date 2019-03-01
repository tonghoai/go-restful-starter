package db

import (
	conf "restful-starter/config"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	Conn *gorm.DB
}

var instance *Connection
var once sync.Once

func GetMysql() *Connection {
	once.Do(func() {
		DB, err := gorm.Open("mysql", conf.GetConfig().MYSQL_DB)

		if err != nil {
			panic(err)
		}

		DB.LogMode(true)
		instance = &Connection{Conn: DB}
	})
	return instance
}
