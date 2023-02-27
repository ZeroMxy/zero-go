package database

import (
	"sync"

	"zero-go/config"
	"zero-go/framework/log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Database struct {}

var (
	xorm_conn	*xorm.Engine
	xorm_once	sync.Once
	xorm_err	error
)

func initialize () {

	xorm_once.Do(func () {

		var drive = config.Database["drive"]

		// 驱动选择
		switch (drive) {
			case "mysql":
				xorm_conn, xorm_err = mysql_connection(drive)

			case "postgres":
				xorm_conn, xorm_err = postgres_connection(drive)

			default:
				xorm_conn, xorm_err = mysql_connection(drive)

		}

		if (xorm_err != nil) {
			log.Error(xorm_err)
			return
		}

		// 支持struct为驼峰式命名
		// 表结构为下划线命名之间的转换
		// 但是对于特定词支持更好
		// 比如ID会翻译成id而不是i_d。
		xorm_conn.SetMapper(names.GonicMapper {})

	})
	
	return

}

// 返回实例
func Conn () *xorm.Engine {

	initialize()

	return xorm_conn

}

// mysql 数据库连接
func mysql_connection (drive string) (*xorm.Engine, error) {

	var host     	= config.Database["host"]
	var port     	= config.Database["port"]
	var name 		= config.Database["name"]
	var username 	= config.Database["username"]
	var password 	= config.Database["password"]

	var dsn = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return xorm.NewEngine(drive, dsn)

}

// postgres 数据库连接
func postgres_connection (drive string) (*xorm.Engine, error) {

	var host     	= config.Database["host"]
	var port     	= config.Database["port"]
	var name 		= config.Database["name"]
	var username 	= config.Database["username"]
	var password 	= config.Database["password"]

	var dsn = "host=" + host + " user=" + username + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	return xorm.NewEngine(drive, dsn)

}