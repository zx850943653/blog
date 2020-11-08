package web

import (
	"github.com/gin-gonic/gin"
	"zx/controller"
	"zx/database"
	"zx/migration"
)

var db = database.Mysql()

func Route(r *gin.Engine) {
	//数据库迁移
	migration.Init()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	})
	r.GET("/mysql", controller.Index)

}
