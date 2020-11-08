package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/zx?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print("数据库链接出错了，错误原因：%v", err)
		return nil
	}
	return db
}
