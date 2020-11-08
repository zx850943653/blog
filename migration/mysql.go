package migration

import (
	"zx/database"
	"zx/model"
)

var db = database.Mysql()

func Init() {
	db.AutoMigrate(&model.User{})
}
