package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ()

func CreateDb() {
	gorm.Open(mysql.Open("py:123@tcp(localhost:3306)/gorm?charset=utf8&TimeParsed=true&loc=Local"))
}
