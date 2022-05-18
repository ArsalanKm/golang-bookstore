package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "docker:password@tcp(http://localhost:3306)/godocker?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		fmt.Println("connection to database fails.")
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
