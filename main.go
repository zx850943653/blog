package main

import (
	"github.com/gin-gonic/gin"
	"zx/database"
	"zx/web"
)

var db = database.Mysql()

func main() {
	rout := gin.Default()
	web.Route(rout)
	rout.Run(":9090")
	//程序停止时关闭数据库
	defer db.Close()
}
