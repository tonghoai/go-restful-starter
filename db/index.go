package db

import (
	"fmt"
	conf "restful-starter/config"
	shared "restful-starter/shared/model"
)

func Ping() {
	PingMysql()
}

func PingMysql() {
	err := GetMysql().Conn.DB().Ping()
	if err != nil {
		panic(err)
	}
	if conf.GetConfig().MYSQL_GENERATOR_DB == "1" {
		GetMysql().Conn.DropTableIfExists(shared.User{})
		GetMysql().Conn.CreateTable(shared.User{})
	}
	fmt.Println("mysql is well")
}
