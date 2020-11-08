package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:50;type:varchar(50)"`
	Sex  int    `gorm:"type:int(2)"`
}
