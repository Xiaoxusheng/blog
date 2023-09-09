package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {

	engine, err := xorm.NewEngine("mysql", "root:admin123@/blog?charset=utf8")
	if err != nil {
		log.Println("连接mysql失败！")
		return
	}
	Engine = engine
	log.Println("连接mysql成功！")

}
