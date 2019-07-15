package main

import (
	. "urlshorner/migration"
	"urlshorner/routes"
)

//var Db *gorm.DB
func main() {
	//var db=GetDb()
	Init()

	routes.AllRoutes()
}
