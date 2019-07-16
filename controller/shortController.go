package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
	"strings"
	. "urlshorner/migration"
	. "urlshorner/model"
)

func GetShortUrl(c *gin.Context) {

	c.HTML(http.StatusOK, "shortForms.html", gin.H{
		"title": "Users", "type": "Short",
	})
}

func RedirectShortUrl(c *gin.Context) {
	urlData := c.PostForm("url")
	hashID := strings.Split(urlData, "http://rzp.com/")
	var todo UrlModel
	UrlHashId := hashID[1]
	//fmt.Println(UrlHashId)

	Db.Where("url_hash_id = ?", UrlHashId).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No URL found!"})
		return
	}
	//Db.Where("hits = ?", UrlHashId).First(&todo)
	Db.Model(&todo).UpdateColumn("hits", gorm.Expr("hits + ?", 1))
	c.Redirect(http.StatusMovedPermanently, "http://"+todo.Url)
}

func Stats(c *gin.Context)  {
	//var todo UrlModel
	fmt.Println("stats")
	var urls []UrlModel
	Db.Find(&urls)
	fmt.Println(urls)
	//const templ = `<TABLE class= "myTable" >
 //       <tr class="headingTr">
 //           <td>Name</td>
 //       </tr>
 //       //{{range .urls}}
 //       <td>"Hello"</td>
 //       <td>"Hello again "</td>
 //       {{end}}
 //</TABLE>`

	fmt.Println(reflect.TypeOf(urls[0].UrlHashId))
	c.HTML(http.StatusOK, "stats.tmpl.html", urls)

}
