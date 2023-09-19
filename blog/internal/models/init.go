package models

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine
var Rdb *redis.Client

func init() {

	engine, err := xorm.NewEngine("mysql", "root:admin123@/blog?charset=utf8")
	if err != nil {
		log.Println("连接mysql失败！")
		return
	}
	Engine = engine
	log.Println("连接mysql成功！")

}
func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "admin123", // no password set
		DB:       0,          // uses default DB
		PoolSize: 1000,
	})
	ctx := context.Background()

	ping := client.Ping(ctx)
	if ping.String() == "ping: PONG" {
		log.Println("连接redis 成功!")
		Rdb = client
		return
	}
	log.Println("连接redis失败")

}
