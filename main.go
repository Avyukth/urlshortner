package main

import (
	//"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
	"net/http"
	"time"

	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB
 var url="www.google.com"

func main()  {
	router:=gin.Default()
	//router.GET("/", init)

	router.LoadHTMLGlob("views/*")
	//router.GET("/" ,createUrlShorter)
	v1:=router.Group("/")
	{
		v1.GET("/{id}", baseUrl)
		v1.POST("/long", getLongUrl)
		v1.GET("/short", createUrlShorter)

	//	//	//v1.GET("/", fetchAllUrlShorter)	//v1.PUT("/:id", updateUrlShorter)
	////	//v1.DELETE("/:id", deleteUrlShorter)
	}

	router.Run()

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

func getLongUrl(c *gin.Context)  {
	fmt.Println("Hello ")
	urlData := c.PostForm("url")
	hd := hashids.NewData()
	h,_ := hashids.NewWithData(hd)
	now := time.Now()
	urlHashId, _ := h.Encode([]int{int(now.Unix())})
	shortenUrl:= "http://rzp:3030/" + urlHashId


	//urlH := sha256.New()
	//urlH.Write([]byte(urlData))
	//urlHashId:=urlH.Sum(nil)
	//shortenUrl:="www.bing.com"
	//shortenUrlH:= sha256.New()
	//shortenUrlH.Write([]byte(shortenUrl))
	//shortenUrlHId:=urlH.Sum(nil)
	////fmt.Println(shortenUrlHId)
	data:=urlModel{UrlHashId:urlHashId,Url:urlData,Shorten:shortenUrl}
	//fmt.Println(data)
	//db.Exec("select * from url_models")
	db.Debug().Create(&data)

}
func baseUrl(c *gin.Context)  {

}
func createUrlShorter(c *gin.Context)  {
	fmt.Println("here")

	////tmpl := template.Must(template.ParseFiles("forms.html"))
	//
	//message, _ := c.GetQuery("m")
	//c.String(http.StatusOK, "Get works! you sent: "+message)
	//c.HTML(http.StatusOK,"forms.html",gin.H{"title": "Page file title!!"})
	//db,err=gorm.Open("mysql","root:root@/urlshort?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Println(err)
	//	panic("failed to connect database " )
	//}
	//Also we can use save that will return primary key
	//db.Debug().Save(&data)
	c.HTML(http.StatusOK, "forms.html", gin.H{
		"title": "Users",
	})



}

type (
	urlModel struct{
		gorm.Model
		UrlHashId string `gorm:"primary_key"`
		Url string
		Shorten string
		//ShortenHashId []byte
	}

)

var charToIndex map[string]int
func idConverter()  {
	initCharToIndex()
	initIndexToChar()

}
func initCharToIndex()  {
	for i := 0; i < 26; i++ {

	}
}
func initIndexToChar()  {

}

//b := [2]string{"Penn", "Teller"}