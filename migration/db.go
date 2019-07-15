package migration

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
}
func GetDb() *gorm.DB {
	return Db
}
