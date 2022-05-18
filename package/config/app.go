package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func connect() {
	d, err := gorm.Open("mysql", "arsalan:arsalan/bookstore?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("connection to database fails.")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
