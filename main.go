package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB

func main()  {
	router:=gin.Default()
	v1:=router.Group("/urlshort"){
		v1.POST("/", createUrlShorter)
	//	//v1.GET("/", fetchAllUrlShorter)
	//	//v1.GET("/:id", fetchSingleUrlShorter)
	//	//v1.PUT("/:id", updateUrlShorter)
	//	//v1.DELETE("/:id", deleteUrlShorter)
	}

}
func init()  {
	var err error
	db,err=gorm.Open("mysql","root:root@/urlshort?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database " )
	}
	//Migrate the Schema
	db.AutoMigrate(&urlModel{})
}

func createUrlShorter(c *gin.Context)  {

}
type (
	urlModel struct{
		gorm.Model
		UrlHashId string `gorm:"primary_key"`
		Url string
		Shorten string
		ShortenHashId string `gorm:"primary_key"`

	}

)
