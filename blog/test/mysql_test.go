package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"xorm.io/xorm"
)

func TestMysql(t *testing.T) {
	_, err := xorm.NewEngine("mysql", "root:admin123@/blog?charset=utf8")
	if err != nil {
		log.Println("连接mysql失败！")
		return
	}
	log.Println("连接mysql成功！")
}
