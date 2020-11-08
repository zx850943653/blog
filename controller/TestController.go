package controller

import (
	"github.com/gin-gonic/gin"
	"zx/database"
	"zx/model"
)

var db = database.Mysql()

func Index(*gin.Context) {

	user := model.User{
		Name: "zx",
		Sex:  1,
	}
	db.Create(&user)
}
