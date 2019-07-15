package controller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	. "urlshorner/model"
)

func InitDB(db *gorm.DB) *gorm.DB{
	fmt.Println("here")
	var err error
	db, err = gorm.Open("mysql", "root:root@/urlshort?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ")
	}
	//Migrate the Schema
	db.AutoMigrate(&UrlModel{})
	return db
	//fmt.Println("done")
}
