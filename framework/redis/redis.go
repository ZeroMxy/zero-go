package redis

import (
	"zero-go/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

type Redis struct {}

var (
	rds_conn *redis.Client
	rds_once sync.Once
)

func initialize () {

	rds_once.Do(func () {

		var host 		= config.Redis["host"]
		var port 		= config.Redis["port"]
		var password 	= config.Redis["password"]

		rds_conn = redis.NewClient(&redis.Options {
			Addr: 		host + ":" + port,
			Password: 	password,
			DB: 		0,
		})
		
	})

	return

}

// 返回实例
func Conn () *redis.Client {

	initialize()

	return rds_conn

}