package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	fmt.Println(UrlHashId)

	Db.Where("url_hash_id = ?", UrlHashId).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No URL found!"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://"+todo.Url)
}
