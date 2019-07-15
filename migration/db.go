package migration

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//. "urlshorner"
	. "urlshorner/model"
)

var Db *gorm.DB
func Init() {
	fmt.Println("here")

	var err error
	Db, err = gorm.Open("mysql", "root:root@/urlshort?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ")
	}
	//Migrate the Schema
	Db.AutoMigrate(&UrlModel{})
	//return db
	//fmt.Println("done")
}
func GetDb() *gorm.DB {
	return Db
}
